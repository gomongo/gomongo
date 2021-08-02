package repl

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ReplClient is an abstraction of the MongoDB Golang driver's [Database](https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo?utm_source=godoc#Database) type.
type ReplClient interface{
	RunCommand (context.Context, interface{}, ...*options.RunCmdOptions) *mongo.SingleResult
}

// ReplSetInitiateMember represents a single member in the configuration for initializing a replica set.
type ReplSetInitiateMember struct{
	ID int `bson:"_id"`
	Host string `bson:"host"`
}

// ReplSetInitiateConfig represents the information required to initialize a set of members in a replica set.
type ReplSetInitiateConfig struct{
	ID string `bson:"_id"`
	Members []ReplSetInitiateMember `bson:"members"`
}

// ReplSetMember represents a single member of a replica set.
type ReplSetMember struct {
	ID int `bson:"_id"`
	Host string `bson:"host"`
	ArbiterOnly bool `bson:"arbiterOnly"`
	BuildIndexes bool `bson:"buildIndexes"`
	Hidden bool `bson:"hidden"`
	Priority int `bson:"priority"`
	Tags map[string]interface{} `bson:"tags"`
	SlaveDelay int `bson:"slaveDelay"`
	Votes int `bson:"votes"`
}

// ReplSetConfig represents the information needed to configure an existing replica set, 
// mirroring the document structure returned by the [`replSetGetConfig`](https://docs.mongodb.com/v4.2/reference/command/replSetGetConfig/) command.
type ReplSetConfig struct {
	ID string `bson:"_id"`
	Members []ReplSetMember `bson:"members"`
	Version int `bson:"version"`
	ProtocolVersion float32 `bson:"protocolVersion"`
	WriteConcernMajorityJournalDefault bool `bson:"writeConcernMajorityJournalDefault"`
}

// InitiateReplicaSet will initialize a new replica set on a set of MongoDB instances.
func InitiateReplicaSet(ctx context.Context, client ReplClient, config ReplSetInitiateConfig) error {
	
	cmd := bson.M{
		"replSetInitiate": config,
	}
	result := client.RunCommand(ctx,cmd)
	err := result.Err()
	
	return err
}

// ConfigureReplicaSet will configure an existing replica set within a set of MongoDB instances.
func ConfigureReplicaSet(ctx context.Context, client ReplClient, config ReplSetConfig, force bool, maxTimeMS int) error {
	cmd := bson.M{
		"replSetReconfig": config,
		"force": force,
		"maxTimeMS": maxTimeMS,
	}
	result := client.RunCommand(ctx,cmd)
	err := result.Err()
	
	return err
}