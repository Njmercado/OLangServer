package mongodb

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//MongoDB structure to handle mongodb connections
type MongoDB struct {
	Client *mongo.Client
	URL    string
	DB     *mongo.Database
}

func getMongoClient(url string) (*mongo.Client, error) {

	ctx, cancel := context.WithCancel(context.Background())

	//Client options, db, pass   "mongodb://localhost:27017"
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		log.Print("The given database do not exist")
		log.Fatal(err)
		cancel()
		return nil, err
	}

	// defer client.Disconnect(context.TODO())
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Print("DataBase do not exist")
		log.Fatal(err)
		cancel()
		return nil, err
	}

	return client, nil
}

//NewMongoDB permite poder establecer una nueva conexión con la base de datos a usar
func NewMongoDB(url, db string) *MongoDB {

	client, err := getMongoClient(url)
	if err != nil {
		log.Print("Couldn't be able to connect to given database")
		log.Fatal(err)
	}

	fmt.Println("DB connection has been succesfull")

	return &MongoDB{
		Client: client,
		URL:    url,
		DB:     client.Database(db),
	}
}
