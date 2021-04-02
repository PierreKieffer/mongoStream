package services

import (
	"context"
	"github.com/PierreKieffer/mongoStream/dataModel"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func SetOptions(resumability bool, startTime uint32) *options.ChangeStreamOptions {

	var cso *options.ChangeStreamOptions
	if resumability == true {
		cso = options.ChangeStream().SetStartAtOperationTime(&primitive.Timestamp{T: startTime})
	}
	return cso

}

func InitBuffer(bufferSize int) chan dataModel.Oplog {
	var buffer = make(chan dataModel.Oplog, bufferSize)
	return buffer
}

func ListenerWorker(mongoUri string, database string, collection string, buffer chan dataModel.Oplog, csoOptional ...*options.ChangeStreamOptions) {

	log.Println("Listening on database : " + database + "| collection : " + collection)

	var cso *options.ChangeStreamOptions

	if len(csoOptional) > 0 {
		cso = csoOptional[0]
	}

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))

	coll := client.Database(database).Collection(collection)

	var pipeline mongo.Pipeline

	cur, err := coll.Watch(ctx, pipeline, cso)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var oplog dataModel.Oplog
		if err := cur.Decode(&oplog); err != nil {
			log.Fatal(err)
		}

		buffer <- oplog

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

}
