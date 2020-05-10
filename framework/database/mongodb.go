package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

func ConnectMongoDB() *mongo.Database {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	conn, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("DB_CONNECTION")))
	if err != nil {
		return nil
	}

	db := conn.Database(os.Getenv("DB_NAME"))

	return db
}
