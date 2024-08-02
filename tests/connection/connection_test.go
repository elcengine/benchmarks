package e_tests

import (
	"benchmarks/mocks"
	. "benchmarks/tests"
	"context"
	"testing"

	e_connection "github.com/elcengine/elemental/connection"
	. "github.com/smartystreets/goconvey/convey"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectDriver() {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mocks.DB_URI))
	if err != nil {
		panic(err)
	}
	So(client, ShouldNotBeNil)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	client.Disconnect(ctx)
}

func ConnectElemental() {
	client := e_connection.ConnectURI(mocks.DB_URI)
	So(client, ShouldNotBeNil)
	client.Disconnect(context.Background())
}

func TestConnection(t *testing.T) {
	Convey("Connect to a database on Atlas", t, func() {
		Benchmark(ConnectDriver, ConnectElemental)
	})
}
