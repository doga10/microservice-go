package application

import (
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"microservice-go/application/oauth"
	"microservice-go/application/status"
	"microservice-go/application/users"
	"microservice-go/framework/database"
)

func Router(router *mux.Router) {
	status.StatusRouter(router)
	oauth.OAuthRouter(router)
	users.UserRouter(router)
}

func LoadModules() {
	db := database.ConnectMongoDB()
	bootstrapUser(db)
	bootstrapOAuth(db)
}

func bootstrapUser(db *mongo.Database) interface{} {
	repository := users.UserRepository{Db: db}
	service := users.UserService{UserRepository: repository}
	return service
}

func bootstrapOAuth(db *mongo.Database) interface{} {
	repository := users.UserRepository{Db: db}
	service := oauth.OAuthService{OAuthRepository: repository}
	return service
}
