package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/vmware/harbor/git"
)

func InitClient() *git.Client {
	client, err := git.NewClient(os.Getenv("HARBOR_CATA_WORKSPACE"), os.Getenv("HARBOR_CATA_GITURL"), os.Getenv("HARBOR_CATA_BRANCH"))
	if err != nil {
		log.Error("the creation of git client failed")
	}
	err = client.Init()
	if err != nil {
		log.Error(fmt.Sprintf("%s :%s", "git clone failed,git uri is", os.Getenv("HARBOR_CATA_GITURL")))
	}
	client.RemoteAdd()
	return client
}
func PullTimer(client *git.Client) {
	timeinterval, _ := strconv.Atoi(os.Getenv("HARBOR_CATA_TIMEINTERVAL"))
	timer := time.NewTicker(time.Second * time.Duration(timeinterval))
	for {
		<-timer.C
		if err := client.Pull(); err != nil {
			log.Error(fmt.Sprintf("%s:%s", "git pull failed,the uri is:", client.URI))
		}
	}
}
