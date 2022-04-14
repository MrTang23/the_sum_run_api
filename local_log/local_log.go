package local_log

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

func Output_Info(x string) {
	writer, err := os.OpenFile("./local_log/log.txt", os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("Create/Open file log.txt failed: %v", err)
	}

	logrus.SetOutput(writer)
	logrus.Info(x)
}

func Output_Error(x string) {
	writer, err := os.OpenFile("./local_log/log.txt", os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("Create/Open file log.txt failed: %v", err)
	}

	logrus.SetOutput(writer)
	logrus.Error(x)
}

func Output_Warn(x string) {
	writer, err := os.OpenFile("./local_log/log.txt", os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("Create/Open file log.txt failed: %v", err)
	}

	logrus.SetOutput(writer)
	logrus.Warn(x)
}
