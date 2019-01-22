package mongo

import (
	"context"
	"github.com/batazor/go-smpp/pkg/utils"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/sirupsen/logrus"
)

var (
	log = logrus.New()

	// ENV
	MONGO_URL             = utils.Getenv("MONGO_URL", "mongodb://localhost:27017")
	MONGO_DATABASE_SMPP   = utils.Getenv("MONGO_DATABASE_SMPP", "go-smpp-message")
	MONGO_DATABASE_WORKER = utils.Getenv("MONGO_DATABASE_WORKER", "go-smpp-worker")

	// Channel
	// message
	AddMessageCh    = make(chan Document)
	UpdateMessageCh = make(chan Document)
	DeleteMessageCh = make(chan Document)
	StatusMessageCh = make(chan Document)
	// worker
	AddWorkerCh    = make(chan Document)
	UpdateWorkerCh = make(chan Document)
	DeleteWorkerCh = make(chan Document)
)

func init() {
	// Logging =================================================================
	// Setup the logger backend using Sirupsen/logrus and configure
	// it to use a custom JSONFormatter. See the logrus docs for how to
	// configure the backend at github.com/Sirupsen/logrus
	log.Formatter = new(logrus.JSONFormatter)
}

func Listen() session {
	var c session
	var err error

	c.client, err = mongo.NewClient(MONGO_URL)
	if err != nil {
		log.Fatal("Error uri for MongoDB: ", err)
	}

	err = c.client.Connect(context.TODO())
	if err != nil {
		log.Fatal("Error connect to MongoDB: ", err)
	}

	log.Info("Run MongoDB")

	go func() {
		for {
			select {
			case <-StatusMessageCh:
				//c.statusMessage(cmd)
			case <-AddMessageCh:
				//c.addMessage(cmd)
			case <-UpdateMessageCh:
				//c.updateMessage(cmd)
			case <-DeleteMessageCh:
				//c.deleteMessage(cmd)
			case <-AddWorkerCh:
				//c.addWorker(cmd)
			case <-UpdateWorkerCh:
				//c.updateWorker(cmd)
			case <-DeleteWorkerCh:
				//c.deleteWorker(cmd)
			}
		}
	}()

	return c
}
