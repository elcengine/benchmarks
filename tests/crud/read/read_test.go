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
	documentCount = 10000
	data          = make([]User, documentCount)
)

func Driver() {
	ctx := context.Background()
	client := e_connection.GetConnection()
	docs := make([]interface{}, documentCount)
	cursor, err := client.Database(mocks.DEFAULT_DB).Collection("users").Find(ctx, map[string]interface{}{})
	if err != nil {
		panic(err)
	}
	cursor.All(ctx, &docs)
	So(len(docs), ShouldEqual, documentCount)
}

func Elemental() {
	docs := UserModel.Find().Exec().([]User)
	So(len(docs), ShouldEqual, documentCount)
}

func TestRead(t *testing.T) {
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
	UserModel.InsertMany(data).Exec()
	Convey(fmt.Sprintf("Read %d records", documentCount), t, func() {
		Benchmark(Driver, Elemental, 10)
	})
}
