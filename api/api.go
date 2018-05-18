package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/RTradeLtd/Temporal/api/rtfs_cluster"
	"github.com/RTradeLtd/Temporal/database"
	"github.com/RTradeLtd/Temporal/models"
	gocid "github.com/ipfs/go-cid"

	"github.com/RTradeLtd/Temporal/api/rtfs"
	"github.com/RTradeLtd/Temporal/queue"
	"github.com/gin-contrib/rollbar"
	"github.com/gin-gonic/gin"
	"github.com/stvp/roll"
)

// Setup is used to initialize our api.
// it invokes all  non exported function to setup the api.
func Setup() *gin.Engine {
	// we use rollbar for logging errors
	token := os.Getenv("ROLLBAR_TOKEN")
	if token == "" {
		log.Fatal("invalid token")
	}
	roll.Token = token
	roll.Environment = "development"
	r := gin.Default()
	r.Use(rollbar.Recovery(false))
	setupRoutes(r)
	return r
}

// setupRoutes is used to setup all of our api routes
func setupRoutes(g *gin.Engine) {

	g.POST("/api/v1/ipfs/pin/:hash", pinHashLocally)
	g.POST("/api/v1/ipfs/add-file", addFileLocally)
	g.POST("/api/v1/ipfs-cluster/pin/:hash", pinHashToCluster)
	g.POST("/api/v1/ipfs-cluster/sync-errors-local", syncClusterErrorsLocally)
	g.DELETE("/api/v1/ipfs/remove-pin/:hash", removePinFromLocalHost)
	g.DELETE("/api/v1/ipfs-cluster/remove-pin/:hash", removePinFromCluster)
	g.GET("/api/v1/ipfs-cluster/status-local-pin/:hash", getLocalStatusForClusterPin)
	g.GET("/api/v1/ipfs-cluster/status-global-pin/:hash", getGlobalStatusForClusterPin)
	g.GET("/api/v1/ipfs-cluster/status-local", fetchLocalClusterStatus)
	g.GET("/api/v1/ipfs/pins", getLocalPins)
	g.GET("/api/v1/database/uploads", getUploadsFromDatabase)
	g.GET("/api/v1/database/uploads/:address", getUploadsForAddress)
}

// removePinFromLocalHost is used to remove a pin from the ipfs instance
// TODO: fully implement
func removePinFromLocalHost(c *gin.Context) {
	// fetch hash param
	hash := c.Param("hash")
	// initialise a connetion to the local ipfs node
	manager := rtfs.Initialize()
	// remove the file from the local ipfs state
	err := manager.Shell.Unpin(hash)
	if err != nil {
		c.Error(err)
		return
	}
	// TODO:
	// change to send a message to the cluster to depin
	qm, err := queue.Initialize(queue.IpfsQueue)
	if err != nil {
		c.Error(err)
		return
	}
	qm.PublishMessage(hash)
	c.JSON(http.StatusOK, gin.H{"deleted": hash})
}

// removePinFromCluster is used to remove a pin from the cluster global state
// this will mean that all nodes in the cluster will no longer track the pin
// TODO: fully implement
func removePinFromCluster(c *gin.Context) {
	hash := c.Param("hash")
	manager := rtfs_cluster.Initialize()
	err := manager.RemovePinFromCluster(hash)
	if err != nil {
		c.Error(err)
		return
	}
	qm, err := queue.Initialize(queue.IpfsQueue)
	if err != nil {
		c.Error(err)
		return
	}
	qm.PublishMessage(hash)
	c.JSON(http.StatusOK, gin.H{"deleted": hash})
}

// fetchLocalClusterStatus is used to fetch the status of the localhost's
// cluster state, and not the rest of the cluster
// TODO: cleanup
func fetchLocalClusterStatus(c *gin.Context) {
	// this will hold all the retrieved content hashes
	var cids []*gocid.Cid
	// this will hold all the statuses of the content hashes
	var statuses []string
	// initialize a connection to the cluster
	manager := rtfs_cluster.Initialize()
	// fetch a map of all the statuses
	maps, err := manager.FetchLocalStatus()
	if err != nil {
		c.Error(err)
		return
	}
	// parse the maps
	for k, v := range maps {
		cids = append(cids, k)
		statuses = append(statuses, v)
	}
	c.JSON(http.StatusOK, gin.H{"cids": cids, "statuses": statuses})
}

// syncCluserErrorsLocally is used to parse through the local cluster state
// and sync any errors that are detected.
func syncClusterErrorsLocally(c *gin.Context) {
	// initialize a conection to the cluster
	manager := rtfs_cluster.Initialize()
	// parse the local cluster status, and sync any errors, retunring the cids that were in an error state
	syncedCids, err := manager.ParseLocalStatusAllAndSync()
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"synced-cids": syncedCids})
}

// getUploadsFromDatabase is used to read a list of uploads from our database
// TODO: cleanup
func getUploadsFromDatabase(c *gin.Context) {
	// open a connection to the database
	db := database.OpenDBConnection()
	// create an upload manager interface
	um := models.NewUploadManager(db)
	// fetch the uplaods
	uploads := um.GetUploads()
	if uploads == nil {
		um.DB.Close()
		c.JSON(http.StatusNotFound, nil)
		return
	}
	um.DB.Close()
	c.JSON(http.StatusFound, gin.H{"uploads": uploads})
}

// getUploadsForAddress is used to read a list of uploads from a particular eth address
// TODO: cleanup
func getUploadsForAddress(c *gin.Context) {
	// open connection to the database
	db := database.OpenDBConnection()
	// establish a new upload manager
	um := models.NewUploadManager(db)
	// fetch all uploads for that address
	uploads := um.GetUploadsForAddress(c.Param("address"))
	if uploads == nil {
		um.DB.Close()
		c.JSON(http.StatusNotFound, nil)
		return
	}
	um.DB.Close()
	c.JSON(http.StatusFound, gin.H{"uploads": uploads})
}

// getLocalPins is used to get the pins tracked by the local ipfs node
func getLocalPins(c *gin.Context) {
	manager := rtfs.Initialize()
	pinInfo, err := manager.Shell.Pins()
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"pins": pinInfo})
}

// getGlobalStatusForClusterPin is used to get the global cluster status for a particular pin
func getGlobalStatusForClusterPin(c *gin.Context) {
	hash := c.Param("hash")
	manager := rtfs_cluster.Initialize()
	status, err := manager.GetStatusForCidGlobally(hash)
	if err != nil {
		c.Error(err)
		return
	}
	qm, err := queue.Initialize(queue.IpfsQueue)
	if err != nil {
		c.Error(err)
		return
	}
	qm.PublishMessage(fmt.Sprintf("%v", status))
	c.JSON(http.StatusFound, gin.H{"status": status})
}

// getLocalStatusForClusterPin is used to get teh localnode's cluster status for a particular pin
func getLocalStatusForClusterPin(c *gin.Context) {
	hash := c.Param("hash")
	manager := rtfs_cluster.Initialize()
	status, err := manager.GetStatusForCidLocally(hash)
	if err != nil {
		c.Error(err)
		return
	}
	qm, err := queue.Initialize(queue.IpfsQueue)
	if err != nil {
		c.Error(err)
		return
	}
	qm.PublishMessage(fmt.Sprintf("%v", status))
	c.JSON(http.StatusFound, gin.H{"status": status})
}

// pinHashToCluster is used to pin a hash to the global cluster state
func pinHashToCluster(c *gin.Context) {
	hash := c.Param("hash")
	manager := rtfs_cluster.Initialize()
	contentIdentifier := manager.DecodeHashString(hash)
	manager.Client.Pin(contentIdentifier, -1, -1, hash)
	qm, err := queue.Initialize(queue.IpfsQueue)
	if err != nil {
		c.Error(err)
		return
	}
	qm.PublishMessage(hash)
	c.JSON(http.StatusOK, gin.H{"hash": hash})
}

// pinHashLocally is used to pin a hash to the local ipfs node
func pinHashLocally(c *gin.Context) {
	hash := c.Param("hash")
	uploadAddress := c.PostForm("uploadAddress")
	holdTimeInMonths := c.PostForm("holdTime")
	holdTimeInt, err := strconv.ParseInt(holdTimeInMonths, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}
	dpa := queue.DatabasePinAdd{
		Hash:             hash,
		UploaderAddress:  uploadAddress,
		HoldTimeInMonths: holdTimeInt,
	}

	qm, err := queue.Initialize(queue.DatabasePinAddQueue)
	if err != nil {
		c.Error(err)
		return
	}

	err = qm.PublishMessage(dpa)
	if err != nil {
		c.Error(err)
		return
	}
	manager := rtfs.Initialize()
	err = manager.Shell.Pin(hash)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"upload": dpa})
}

// addFileLocally is used to add a file to our local ipfs node
// this will have to be done first before pushing any file's to the cluster
// this needs to be optimized so that the process doesn't "hang" while uploading
func addFileLocally(c *gin.Context) {
	fileHandler, err := c.FormFile("file")
	if err != nil {
		c.Error(err)
		return
	}
	uploaderAddress := c.PostForm("uploaderAddress")
	fmt.Println(uploaderAddress)
	holdTimeinMonths := c.PostForm("holdTime")
	holdTimeinMonthsInt, err := strconv.ParseInt(holdTimeinMonths, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}
	openFile, err := fileHandler.Open()
	if err != nil {
		c.Error(err)
		return
	}
	manager := rtfs.Initialize()
	resp, err := manager.Shell.Add(openFile)
	if err != nil {
		c.Error(err)
		return
	}

	dfa := queue.DatabaseFileAdd{
		Hash:             resp,
		HoldTimeInMonths: holdTimeinMonthsInt,
		UploaderAddress:  uploaderAddress,
	}

	qm, err := queue.Initialize(queue.DatabaseFileAddQueue)
	if err != nil {
		c.Error(err)
		return
	}
	err = qm.PublishMessage(dfa)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": resp})
}
