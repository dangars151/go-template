package helpers

import "github.com/sirupsen/logrus"

func Log(err error, message string, data ...interface{}) {
	logrus.WithField("data", data).Errorf("%s: %+v", message, err)
}
