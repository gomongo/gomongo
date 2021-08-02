package repl

import (
	"context"
	
	"github.com/stretchr/testify/assert"
	
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"testing"
)

type MockClient struct{
	RunCommandArg interface{}
}

func (m *MockClient) RunCommand (ctx context.Context, runCommand interface{}, options ...*options.RunCmdOptions) *mongo.SingleResult {
	m.RunCommandArg = runCommand
	return &mongo.SingleResult{}
}

func TestInitiateReplicaSet(t *testing.T) {
	mockClient := MockClient{}
	
	replicaSetConfig := ReplSetInitiateConfig{
		ID: "myReplicaSet",
		Members: []ReplSetInitiateMember{
			{
				ID: 0,
				Host: "myHost",
			},
			{
				ID: 1,
				Host: "myHost2",
			},
			{
				ID: 2,
				Host: "myHost3",
			},
		},
	}
	expectedCommand := bson.M{
		"replSetInitiate": replicaSetConfig,
	}

	_ = InitiateReplicaSet(context.TODO(), &mockClient, replicaSetConfig)

	assert.EqualValues(t, expectedCommand, mockClient.RunCommandArg)

}