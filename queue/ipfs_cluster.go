package queue

import (
	"encoding/json"

	"github.com/RTradeLtd/Temporal/models"
	"github.com/RTradeLtd/Temporal/rtfs_cluster"
	"github.com/RTradeLtd/config"
	"github.com/jinzhu/gorm"
	"github.com/streadway/amqp"
)

// ProcessIPFSClusterPins is used to process messages sent to rabbitmq requesting be pinned to our cluster
// TODO: add in email notification and metric strategies
func (qm *QueueManager) ProcessIPFSClusterPins(msgs <-chan amqp.Delivery, cfg *config.TemporalConfig, db *gorm.DB) error {
	clusterManager, err := rtfs_cluster.Initialize(cfg.IPFSCluster.APIConnection.Host, cfg.IPFSCluster.APIConnection.Port)
	if err != nil {
		return err
	}
	uploadManager := models.NewUploadManager(db)
	qm.LogInfo("processing ipfs cluster pins")
	for d := range msgs {
		qm.LogInfo("new message detected")
		clusterAdd := IPFSClusterPin{}
		err = json.Unmarshal(d.Body, &clusterAdd)
		if err != nil {
			qm.LogError(err, "failed to unmarshal message")
			d.Ack(false)
			continue
		}
		if clusterAdd.NetworkName != "public" {
			qm.refundCredits(clusterAdd.UserName, "pin", clusterAdd.CreditCost, db)
			qm.LogError(err, "private networks not supported for ipfs cluster")
			d.Ack(false)
			continue
		}
		qm.LogInfo("successfully unmarshaled message, decoding hash string")
		encodedCid, err := clusterManager.DecodeHashString(clusterAdd.CID)
		if err != nil {
			qm.refundCredits(clusterAdd.UserName, "pin", clusterAdd.CreditCost, db)
			qm.LogError(err, "failed to decode hash string")
			d.Ack(false)
			continue
		}
		qm.LogInfo("pinning hash to cluster")
		err = clusterManager.Pin(encodedCid)
		if err != nil {
			qm.refundCredits(clusterAdd.UserName, "pin", clusterAdd.CreditCost, db)
			qm.LogError(err, "failed to pin hash to cluster")
			d.Ack(false)
			continue
		}
		_, err = uploadManager.FindUploadByHashAndNetwork(clusterAdd.CID, clusterAdd.NetworkName)
		if err != nil && err != gorm.ErrRecordNotFound {
			qm.LogError(err, "failed to search database for upload")
			d.Ack(false)
			continue
		}

		if err == gorm.ErrRecordNotFound {
			_, err = uploadManager.NewUpload(clusterAdd.CID, "pin-cluster", models.UploadOptions{
				NetworkName:      clusterAdd.NetworkName,
				Username:         clusterAdd.UserName,
				HoldTimeInMonths: clusterAdd.HoldTimeInMonths})
			if err != nil {
				qm.LogError(err, "failed to create upload in database")
				d.Ack(false)
				continue
			}
		} else {
			_, err = uploadManager.UpdateUpload(clusterAdd.HoldTimeInMonths, clusterAdd.UserName, clusterAdd.CID, clusterAdd.NetworkName)
			if err != nil {
				qm.LogError(err, "failed to update upload in database")
				d.Ack(false)
				continue
			}
		}
		qm.LogInfo("successfully pinned hash to cluster")
		d.Ack(false)
	}
	return nil
}
