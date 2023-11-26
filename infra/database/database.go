package database

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/IcaroSilvaFK/free-code-source-back/infra/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB_CONTEXT = context.Background()

func NewDBConnection() *mongo.Database {

	fmt.Println("entro aq")

	ctx, cancel := context.WithTimeout(DB_CONTEXT, time.Second*10)

	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv(utils.DATABASE_URL)))

	if !errors.Is(err, nil) {
		panic(err)
	}

	// defer func() {
	// 	if err = client.Disconnect(ctx); !errors.Is(err, nil) {
	// 		panic(err)
	// 	}
	// }()

	fmt.Println(client.Ping(ctx, nil))

	db := client.Database("monorail")

	return db

}
