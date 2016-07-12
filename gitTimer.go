package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/vmware/harbor/git"
)

func InitClient() (*git.Client, error) {
	client, err := git.NewClient("/go/bin/harborCatalog", os.Getenv("HARBOR_CATA_GITURL"), "master")
	if err != nil {
		log.Error("the creation of git client failed")
		return nil, err
	}
	err = client.Init()
	if err != nil {
		log.Error(fmt.Sprintf("%s :%s", "git clone failed:", err))
		return nil, err
	}
	client.RemoteAdd()
	return client, nil
}

func PullTimer(client *git.Client) {
	timeinterval, _ := strconv.Atoi(os.Getenv("REPO_FETCH_INTERVAL"))
	timer := time.NewTicker(time.Second * time.Duration(timeinterval))
	for {
		<-timer.C
		if err := client.Reset(); err != nil {
			log.Error(fmt.Sprintf("%s:%s", "git reset failed:", err))
		}
		if err := client.Pull(); err != nil {
			log.Error(fmt.Sprintf("%s:%s", "git pull failed:", err))
		}
	}
}
