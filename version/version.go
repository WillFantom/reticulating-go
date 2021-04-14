package version

import (
	semver "github.com/hashicorp/go-version"
	"github.com/sirupsen/logrus"
)

const (
	unversioned string = "unversioned"
	rolling     string = "rolling"
)

var (
	appVersion string = unversioned
)

func init() {
	if appVersion == unversioned || appVersion == rolling {
		return
	}
	_, err := semver.NewVersion(appVersion)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"given version": appVersion,
			"error message": err.Error(),
		}).Fatalln("🆘	app does not have a valid version")
	}
}

func GetVersion() string {
	return appVersion
}

func UpdateAvaliable() bool {
	// add this in at some point.... maybe....
	return false
}
