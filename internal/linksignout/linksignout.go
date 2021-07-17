package linksignout

import (
	"context"
	"fmt"
	"log"
	"minefit_auth/internal/users"
	"minefit_auth/mongo/singoutdb"

	"go.mongodb.org/mongo-driver/bson"
)

type Link struct {
	Username    string
	Clinicid    string
	Signouttime string
	User        *users.User
}

var ctx = context.TODO()

func (link Link) Save() int64 {
	res, insertErr := singoutdb.Getcol().InsertOne(ctx, link)
	if insertErr != nil {
		log.Fatal(insertErr)
	}
	fmt.Println(res)

	return 0
}

func GetAll() []Link {

	//opts := options.Find()
	//opts.SetSort(bson.D{{"_id", -1}})

	res, insertErr := singoutdb.Getcol().Find(ctx, bson.D{{}})
	if insertErr != nil {
		log.Fatal(insertErr)
	}
	fmt.Println(res)

	var msg []Link
	if insertErr = res.All(ctx, &msg); insertErr != nil {
		panic(insertErr)
	}

	return msg

}
