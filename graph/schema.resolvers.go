package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/distrotion/gql-no-db/graph/generated"
	"github.com/distrotion/gql-no-db/graph/model"
	"github.com/distrotion/gql-no-db/internal/linksignout"
	"github.com/distrotion/gql-no-db/internal/users"
)

func (r *mutationResolver) CreateSignout(ctx context.Context, input model.NewSignout) (*model.LinkSignout, error) {
	var link linksignout.Link
	link.Username = input.Username
	link.Clinicid = input.Clinicid
	link.Signouttime = input.Signouttime
	linkID := link.Save()
	fmt.Println(linkID)
	return &model.LinkSignout{Username: link.Username, Clinicid: link.Clinicid, Signouttime: link.Signouttime}, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	fmt.Print(input)
	//=================================================
	// inlog, err := tstlog.Getcol().InsertOne(ctx, bson.M{"From": "CreateUser", "data_newuser_in": input,"timestamp":time.Now()})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(inlog)
	//=================================================
	// t1 := time.Now().Unix()
	// t2 := time.Now().UnixNano()
	// var user users.User
	// user.ID = strconv.FormatInt(int64(input.Roleid), 10) + "-" + strconv.FormatInt(t1, 16) + strconv.FormatInt(t2, 16)
	// user.Username = input.Username
	// user.Password = input.Password
	// user.Clinicid = input.Createtime
	// user.Clinicid = input.Clinicid
	// user.Roleid = int64(input.Roleid)
	// user.Accountstatus = int64(input.Accountstatus)
	// user.Sessionlifetime = int64(input.Sessionlifetime)

	//----------------------------------------------

	// var useronly users.Useronly
	// useronly.Username = input.Username
	// msg := users.Checkuser(useronly)
	//----------------------------------------------

	//=================================================
	// outlog, err := tstlog.Getcol().InsertOne(ctx, bson.M{"From": "CreateUser", "data_newuser_out": msg,"timestamp":time.Now()})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(outlog)
	//=================================================

	// if len(msg) == 0 {
	// 	user.Create()
	// 	token, err := jwt.GenerateToken(user.Username)
	// 	if err != nil {
	// 		return "", err
	// 	}
	// 	return token, nil
	// } else {
	// 	return "", nil
	// }
	return "", nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (*model.ReturnSignin, error) {
	//=================================================
	// inlog, err := tstlog.Getcol().InsertOne(ctx, bson.M{"From": "Login", "data_ligin_in": input,"timestamp":time.Now()})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(inlog)
	//=================================================
	var output *model.ReturnSignin
	//fmt.Println(input.Username)
	//=================================================

	var user users.User
	user.Username = input.Username
	user.Password = input.Password
	//correct := user.Authenticate()
	correct := false
	if user.Username == "parin@mail.com" && user.Password == "123456789" {
		correct = true
	} else {
		correct = false
	}

	if !correct {
		//
		// output.Status = "incorrect login"
		// output.Sessionid = ""
		// output.Userid = ""
		// output.Roleid = 0
		// output.Accountstatus = 0
		// output.Sessionlifetime = 0
		dummyReturnSignin := model.ReturnSignin{
			Status: "incorrect login",
			// Sessionid:       "",
			// Userid:          "",
			// Clinicid:        "",
			// Roleid:          0,
			// Accountstatus:   0,
			// Sessionlifetime: 0,
		}
		output = &dummyReturnSignin
		return output, nil
	}
	//------------------------------

	//------------------------------

	// token, err := jwt.GenerateToken(user.Username)
	// if err != nil {
	// 	return "", err
	// }
	//=================================================
	// outlog, err := tstlog.Getcol().InsertOne(ctx, bson.M{"From": "Login", "data_login_out": token, "timestamp": time.Now()})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(outlog)
	//=================================================
	// output.Status = "OK"
	// output.Sessionid = "A-11"
	// output.Userid = "0-12ab34cd56ef"
	// output.Roleid = 1
	// output.Accountstatus = 1
	// output.Sessionlifetime = 1
	dummyReturnSignin := model.ReturnSignin{
		Status:          "OK",
		Sessionid:       "A-11",
		Clinicid:        "C-A",
		Userid:          "0-12ab34cd56ef",
		Roleid:          1,
		Accountstatus:   1,
		Sessionlifetime: 1,
	}
	output = &dummyReturnSignin

	return output, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	// username, err := jwt.ParseToken(input.Token)
	// if err != nil {
	// 	return "", fmt.Errorf("access denied")
	// }
	// token, err := jwt.GenerateToken(username)
	// if err != nil {
	// 	return "", err
	// }
	// return token, nil
	return "", nil
}

func (r *queryResolver) Links(ctx context.Context) ([]*model.LinkSignout, error) {
	var resultLinks []*model.LinkSignout
	// var dbLinks []linksignout.Link = linksignout.GetAll()
	// for _, link := range dbLinks {
	// 	resultLinks = append(resultLinks, &model.LinkSignout{Username: link.Username, Clinicid: link.Clinicid, Signouttime: link.Signouttime})
	// }
	return resultLinks, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
// type ReturnSignin struct {
// 	Status           string `json:"status"`
// 	Session_id       string `json:"session_id"`
// 	Clinic_id        string `json:"clinic_id"`
// 	User_id          string `json:"user_id"`
// 	Role_id          int    `json:"role_id"`
// 	Account_status   int    `json:"account_status"`
// 	Session_lifetime int    `json:"session_lifetime"`
// }
