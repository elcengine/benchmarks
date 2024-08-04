package e_tests

import (
	"benchmarks/mocks"
	"benchmarks/setup"
	. "benchmarks/tests"
	"context"
	"fmt"
	"testing"

	"github.com/elcengine/elemental/connection"
	"github.com/go-faker/faker/v4"
	"github.com/samber/lo"
	. "github.com/smartystreets/goconvey/convey"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	documentCount = 500
	data          = make([]User, documentCount)
)

func Driver() {
	ctx := context.Background()
	client := e_connection.GetConnection()
	for i := range data {
		_, err := client.Database(mocks.DEFAULT_DB).Collection("users").InsertOne(ctx, data[i])
		if err != nil {
			panic(err)
		}
	}
	count, err := client.Database(mocks.DEFAULT_DB).Collection("users").CountDocuments(ctx, map[string]interface{}{})
	if err != nil {
		panic(err)
	}
	So(count, ShouldEqual, documentCount)
	client.Database(mocks.DEFAULT_DB).Drop(ctx)
}

func Elemental() {
	for i := range data {
		UserModel.Create(data[i]).Exec()
	}
	count := UserModel.CountDocuments().Exec()
	So(count, ShouldEqual, documentCount)
	UserModel.Drop()
}

func TestInsert(t *testing.T) {
	setup.Connection()
	defer setup.Teardown()
	for i := range data {
		data[i] = User{
			ID:         primitive.NewObjectID(),
			Name:       faker.UUIDHyphenated(),
			Age:        lo.Must(faker.RandomInt(1, 100))[0],
			Occupation: faker.UUIDDigit(),
			Weapons:    []string{faker.UUIDHyphenated(), faker.UUIDHyphenated(), faker.UUIDHyphenated()},
			School:     lo.ToPtr(faker.UUIDHyphenated()),
		}
	}
	Convey(fmt.Sprintf("Insert %d records", documentCount), t, func() {
		Benchmark(Driver, Elemental, 1)
	})
}