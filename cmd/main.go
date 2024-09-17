package main

import (
	"context"
	"log"

	"github.com/sbaitman-service/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://admin:SBaitman123!@localhost:27018"))
	if err != nil {
		log.Fatal("Unable to connect to mongoDB")
		panic(err)
	}
	log.Println("Successfully connected to mongoDB")

	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()
	user1 := model.User{UserId: "ssb001", FirstName: "Shyam", LastName: "Baitmangalkar", Email: "sbaitman@hotmail.com"}
	user2 := model.User{UserId: "jw002", FirstName: "Johnny", LastName: "Walker", Email: "jwalker@hotmail.com"}
	user3 := model.User{UserId: "om003", FirstName: "Old", LastName: "Monk", Email: "omonk@hotmail.com"}
	user4 := model.User{UserId: "jb004", FirstName: "Jim", LastName: "Beam", Email: "jbeam@hotmail.com"}
	user5 := model.User{UserId: "jd005", FirstName: "Jack", LastName: "Daniel", Email: "jdaniel@hotmail.com"}
	user6 := model.User{UserId: "wl006", FirstName: "William", LastName: "Lawson", Email: "wlawson@hotmail.com"}
	user7 := model.User{UserId: "jp007", FirstName: "John", LastName: "Paul", Email: "jpaul@hotmail.com"}
	user8 := model.User{UserId: "os008", FirstName: "Oaks", LastName: "Smith", Email: "osmith@hotmail.com"}
	user9 := model.User{UserId: "li009", FirstName: "Long", LastName: "Island", Email: "lisland@hotmail.com"}
	users := make([]interface{}, 0)
	users = append(users, user1, user2, user3, user4, user5, user6, user7, user8, user9)
	coll := client.Database("user-account").Collection("user")
	res, insErr := coll.InsertMany(context.Background(), users)
	if insErr != nil {
		log.Printf("Unable to insert user %+v in database", users)
		panic(insErr)
	}
	log.Printf("Successfully inserted one record: %+v", res.InsertedIDs...)

}
