package main

import (
	"context"
	"time"

	"github.com/gomongo/gomongo/pkg/repl"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	clientOpts := options.Client().ApplyURI("mongodb://localHost:27017/?connect=direct")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx,clientOpts)
	if err != nil {
		panic(err)
	}

	replicaSetConfig := repl.ReplSetInitiateConfig{
		ID: "replset1",
		Members: []repl.ReplSetInitiateMember{
			{
				ID: 0,
				Host: "mongo_1",
			},
			{
				ID: 1,
				Host: "mongo_2",
			},
			{
				ID: 2,
				Host: "mongo_3",
			},
		},
	}
	
	dbClient := client.Database("admin")

	err = repl.InitiateReplicaSet(ctx,dbClient,replicaSetConfig)
	if err != nil {
		panic(err)
	}

}