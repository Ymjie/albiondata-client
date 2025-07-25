package client

import (
	"os"
	"path/filepath"

	"github.com/ao-data/albiondata-client/log"
)

func processOffline(path string) {
	log.Infof("Beginning offline process with %v", path)

	r := newRouter()
	go r.run()

	_, err := os.Stat(path)

	if err != nil {
		log.Error("Could not find {}: ", path, err)

		return
	}

	l := newListener(r)

	fileExtension := filepath.Ext(path)
	switch fileExtension {
	case ".pcap":
		l.startOfflinePcap(path)
	case ".gob":
		l.startOfflineCommandGob(path)
	default:
		log.Error("Only .pcap and .gob files supported at this time.")
	}
}
