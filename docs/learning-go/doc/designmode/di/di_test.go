package di_test

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/luoleiself/learning-go/designmode/di"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initMySQLDriver(dataSourceName string) (*sql.DB, error) {
	return sql.Open("mysql", dataSourceName)
}

func initMongoDriver(dataSourceName string, dbName string, collectionName string) (*mongo.Collection, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(dataSourceName))
	if err != nil {
		return nil, err
	}
	return client.Database(dbName).Collection(collectionName), nil
}

func TestDI(t *testing.T) {
	t.Run("TestDI", func(t *testing.T) {
		mysqlDB, err := initMySQLDriver("user:password@tcp(localhost:3306)/mydb")
		if err != nil {
			log.Fatal(err)
		}
		mongoDB, err := initMongoDriver("mongodb://localhost:27017", "mydb", "users")
		if err != nil {
			log.Fatal(err)
		}

		userMysql := di.NewUserMySQLRepository(mysqlDB)
		userMongo := di.NewUserMongoRepository(mongoDB)
		userCase := di.NewUserCase(userMysql)
		if os.Getenv("USER_MONGO_DB") == "true" {
			userCase = di.NewUserCase(userMongo)
		}

		userCase.Create(context.Background(), &di.User{
			Name:     "John Doe",
			Email:    "john.doe@example.com",
			Password: "password",
		})
	})
}
