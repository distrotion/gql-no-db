package linksignout

import (
	"context"
	_ "fmt"
	_ "log"

	"github.com/distrotion/gql-no-db/internal/users"
	_ "github.com/distrotion/gql-no-db/mongo/singoutdb"

	_ "go.mongodb.org/mongo-driver/bson"
)

type Link struct {
	Username    string
	Clinicid    string
	Signouttime string
	User        *users.User
}

var ctx = context.TODO()

func (link Link) Save() int64 {
	// res, insertErr := singoutdb.Getcol().InsertOne(ctx, link)
	// if insertErr != nil {
	// 	log.Fatal(insertErr)
	// }
	// fmt.Println(res)

	return 0
}

func GetAll() []Link {

	//opts := options.Find()
	//opts.SetSort(bson.D{{"_id", -1}})

	// res, insertErr := singoutdb.Getcol().Find(ctx, bson.D{{}})
	// if insertErr != nil {
	// 	log.Fatal(insertErr)
	// }
	// fmt.Println(res)

	var msg []Link
	// if insertErr = res.All(ctx, &msg); insertErr != nil {
	// 	panic(insertErr)
	// }

	return msg

}
