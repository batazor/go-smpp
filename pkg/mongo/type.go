package mongo

import (
	"github.com/mongodb/mongo-go-driver/mongo"
)

type session struct {
	client *mongo.Client
}

type Response struct {
	error   error
	payload chan *interface{}
}

type Document struct {
	ID       *string
	Item     interface{}
	Response *Response
}
