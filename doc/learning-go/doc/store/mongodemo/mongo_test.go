package mongodemo

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func TestMongoPing(t *testing.T) {
	var ctx = context.Background()
	t.Run("Ping", func(t *testing.T) {
		if err := mg.Ping(ctx, &readpref.ReadPref{}); err != nil {
			t.Fatalf("Error mg.Ping() %s\n", err)
		}
		t.Log("mg.Ping() Pong")
	})
}
func TestMongoDatabase(t *testing.T) {
	t.Run("DataBase", func(t *testing.T) {
		db := mg.Database("admin", &options.DatabaseOptions{})
		t.Logf("Database Name: %s\n", db.Name())

		dbs, err := mg.ListDatabaseNames(ctx, bson.D{}, &options.ListDatabasesOptions{})
		if err != nil {
			t.Fatalf("Error mg.ListDatabaseNames() %s\n", err)
		}
		t.Logf("mg.ListDatabaseNames() %s\n", dbs)

		dbResult, err := mg.ListDatabases(ctx, bson.D{}, &options.ListDatabasesOptions{})
		if err != nil {
			t.Fatalf("Error mg.ListDatabases() %s\n", err)
		}
		t.Logf("dbResult.Databases %+v\n", dbResult.Databases)
		t.Logf("dbResult.TotalSize %d\n", dbResult.TotalSize)
	})
}
func TestMongoCollection(t *testing.T) {
	var db *mongo.Database
	t.Run("Collection", func(t *testing.T) {
		db = mg.Database("test", &options.DatabaseOptions{})
		t.Logf("Database Name: %s\n", db.Name())
		colls, err := db.ListCollectionNames(ctx, bson.D{}, &options.ListCollectionsOptions{})
		if err != nil {
			t.Fatalf("Error db.ListCollectionNames() %s\n", err)
		}
		t.Logf("db.db.ListCollectionNames() %s\n", colls)

		cols, err := db.ListCollectionSpecifications(ctx, bson.D{}, &options.ListCollectionsOptions{})
		if err != nil {
			t.Fatalf("Error db.ListCollectionSpecifications() %s\n", err)
		}
		t.Logf("cols %+v\n", cols)

		cursor, err := db.ListCollections(ctx, bson.D{}, &options.ListCollectionsOptions{})
		if err != nil {
			t.Fatalf("Error db.ListCollections() %s\n", err)
		}
		t.Logf("cursor %+v\n", cursor)
	})
	t.Run("Insert", func(t *testing.T) {
		result, err := db.Collection("inventory", &options.CollectionOptions{}).InsertMany(ctx, []any{bson.D{
			{Key: "item", Value: "journal"},
			{Key: "qty", Value: int32(25)},
			{Key: "tags", Value: bson.A{"blank", "red"}},
			{Key: "size", Value: bson.D{{Key: "h", Value: 14}, {Key: "w", Value: 21}, {Key: "uom", Value: "cm"}}},
			{Key: "status", Value: "A"},
		},
			bson.D{
				{Key: "item", Value: "mat"},
				{Key: "qty", Value: int32(25)},
				{Key: "tags", Value: bson.A{"gray"}},
				{Key: "size", Value: bson.D{{Key: "h", Value: 27.9}, {Key: "w", Value: 35.5}, {Key: "uom", Value: "cm"}}},
				{Key: "status", Value: "A"},
			},
			bson.D{
				{Key: "item", Value: "mousepad"},
				{Key: "qty", Value: 25},
				{Key: "tags", Value: bson.A{"gel", "blue"}},
				{Key: "size", Value: bson.D{{Key: "h", Value: 19}, {Key: "w", Value: 22.85}, {Key: "uom", Value: "cm"}}},
				{Key: "status", Value: "A"},
			}}, &options.InsertManyOptions{})
		if err != nil {
			t.Fatalf("Error db.Collection(\"inventory\").InsertMany() %s\n", err)
		}
		t.Log("result.InsertedIDs", result.InsertedIDs)
	})
	t.Run("UpdateOne", func(t *testing.T) {
		result, err := db.Collection("inventory", &options.CollectionOptions{}).UpdateOne(ctx, bson.D{{Key: "item", Value: "mat"}},
			bson.D{
				{Key: "$set", Value: bson.D{{Key: "size.uom", Value: "cm"}, {Key: "status", Value: "P"}}},
				{Key: "$currentDate", Value: bson.D{{Key: "lastModified", Value: true}}},
			}, &options.UpdateOptions{})
		if err != nil {
			t.Fatalf("Error db.Collection(\"inventory\").UpdateOne() %s\n", err)
		}
		t.Logf("MatchedCount %d ModifiedCount %d UpsertedCount %d UpsertedID %v\n", result.MatchedCount, result.ModifiedCount, result.UpsertedCount, result.UpsertedID)
	})
	t.Run("Find", func(t *testing.T) {
		cursor, err := db.Collection("inventory", &options.CollectionOptions{}).Find(ctx, bson.D{}, &options.FindOptions{})
		if err != nil {
			t.Fatalf("Error db.Collection(\"inventory\").Find() %s\n", err)
		}
		t.Logf("cursor %+v\n", cursor)
	})
	t.Run("Delete", func(t *testing.T) {
		result, err := db.Collection("inventory", &options.CollectionOptions{}).DeleteOne(ctx, bson.D{{Key: "item", Value: "mat"}}, &options.DeleteOptions{})
		if err != nil {
			t.Fatalf("Error db.Collection(\"inventory\").DeleteOne() %s\n", err)
		}
		t.Logf("DeletedCount %d\n", result.DeletedCount)
	})
	t.Run("Find", func(t *testing.T) {
		cursor, err := db.Collection("inventory", &options.CollectionOptions{}).Find(ctx, bson.D{}, &options.FindOptions{})
		if err != nil {
			t.Fatalf("Error db.Collection(\"inventory\").Find() %s\n", err)
		}
		t.Logf("cursor %+v\n", cursor)
	})
}
