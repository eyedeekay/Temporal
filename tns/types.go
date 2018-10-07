package tns

import (
	ci "github.com/libp2p/go-libp2p-crypto"
	host "github.com/libp2p/go-libp2p-host"
)

const (
	// CommandEcho is a test command used to test if we have successfully connected to a tns daemon
	CommandEcho = "/echo/1.0.0"
)

var (
	commands = []string{CommandEcho}
)

// RecordRequest is a message sent when requeting a record form TNS, the response is simply Record
type RecordRequest struct {
	RecordName string `json:"record_name"`
}

// ZoneRequest is a message sent when requesting a reccord from TNS.
type ZoneRequest struct {
	ZoneName           string `json:"zone_name"`
	ZoneManagerKeyName string `json:"zone_manager_key_name"`
}

// Zone is a mapping of human readable names, mapped to a public key. In order to retrieve the latest
type Zone struct {
	Manager   *ZoneManager `json:"zone_manager"`
	PublicKey ci.PubKey    `json:"zone_public_key"`
	// A human readable name for this zone
	Name string `json:"name"`
	// A map of records managed by this zone
	Records                 map[string]*Record   `json:"records"`
	RecordNamesToPublicKeys map[string]ci.PubKey `json:"record_names_to_public_keys"`
}

// Record is a particular name entry managed by a zone
type Record struct {
	PublicKey *ci.PubKey `json:"public_key"`
	// A human readable name for this record
	Name string `json:"name"`
	// User configurable meta data for this record
	MetaData map[string]interface{} `json:"meta_data"`
}

// ZoneManager is the authorized manager of a zone
type ZoneManager struct {
	PublicKey ci.PubKey `json:"public_key"`
}

// HostOpts is our options for when we create our libp2p host
type HostOpts struct {
	IPAddress string `json:"ip_address"`
	Port      string `json:"port"`
	IPVersion string `json:"ip_version"`
	Protocol  string `json:"protocol"`
}

// Manager is used to manipulate a zone on TNS and run a daemon
type Manager struct {
	PrivateKey        ci.PrivKey
	ZonePrivateKey    ci.PrivKey
	RecordPrivateKeys map[string]ci.PrivKey
	Zone              *Zone
	Host              host.Host
}

// Client is used to query a TNS daemon
type Client struct {
	PrivateKey ci.PrivKey
	Host       host.Host
}

// Host is an interface used by a TNS client or daemon
type Host interface {
	MakeHost(pk ci.PrivKey, opts *HostOpts) error
}