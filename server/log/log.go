package log

import (
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetReportCaller(true)

	logrus.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})
}
