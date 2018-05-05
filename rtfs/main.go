package rtfs

import (
	ipfsapi "github.com/ipfs/go-ipfs-api"
)

type IpfsManager struct {
	Shell *ipfsapi.Shell
}

func Initialize() *IpfsManager {
	url := ""
	manager := IpfsManager{}
	manager.Shell = establishShell(url)
	return &manager
}

func establishShell(url string) *ipfsapi.Shell {
	if len(url) < 10 {
		shell := ipfsapi.NewShell("/ip4/127.0.0.1/tcp/9095/ipfs/QmV7gaSDfUsTRMMRq7QgUBWbkYcrwjoNMriQJrXahjp7NJ")
		return shell
	}
	shell := ipfsapi.NewShell(url)
	return shell
}