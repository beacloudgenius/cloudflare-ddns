package main

import (
	"os"

	"github.com/sirupsen/logrus"
	cloudflare "github.com/beacloudgenius/cloudflare-ddns"
)

func main() {
	c, err := cloudflare.New(os.Getenv("CF_API_KEY"), os.Getenv("CF_API_EMAIL"))
	if err != nil {
		logrus.WithError(err).Fatal("unable to instantiate ddns")
	}

	c.KeepUpdated(os.Getenv("CF_HOST"))
}
