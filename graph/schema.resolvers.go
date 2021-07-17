package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"minefit_auth/graph/generated"
	"minefit_auth/graph/model"
	"minefit_auth/internal/linksignout"
	"minefit_auth/internal/users"
	"minefit_auth/mongo/tstlog"
	"minefit_auth/pkg/jwt"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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
	inlog, err := tstlog.Getcol().InsertOne(ctx, bson.M{"From": "CreateUser", "data_newuser_in": input, "timestamp": time.Now()})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(inlog)
	//=================================================
	t1 := time.Now().Unix()
	t2 := time.Now().UnixNano()
	var user users.User
	user.ID = strconv.FormatInt(int64(input.Roleid), 10) + "-" + strconv.FormatInt(t1, 16) + strconv.FormatInt(t2, 16)
	user.Username = input.Username
	user.Password = input.Password
	user.Clinicid = input.Createtime
	user.Clinicid = input.Clinicid
	user.Roleid = int64(input.Roleid)
	user.Accountstatus = int64(input.Accountstatus)
	user.Sessionlifetime = int64(input.Sessionlifetime)

	//----------------------------------------------

	var useronly users.Useronly
	useronly.Username = input.Username
	msg := users.Checkuser(useronly)
	//----------------------------------------------

	//=================================================
	outlog, err := tstlog.Getcol().InsertOne(ctx, bson.M{"From": "CreateUser", "data_newuser_out": msg, "timestamp": time.Now()})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(outlog)
	//=================================================

	if len(msg) == 0 {
		user.Create()
		token, err := jwt.GenerateToken(user.Username)
		if err != nil {
			return "", err
		}
		return token, nil
	} else {
		return "", nil
	}
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (*model.ReturnSignin, error) {
	//=================================================
	inlog, err := tstlog.Getcol().InsertOne(ctx, bson.M{"From": "Login", "data_ligin_in": input, "timestamp": time.Now()})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(inlog)
	//=================================================
	var output *model.ReturnSignin

	var user users.User
	user.Username = input.Username
	user.Password = input.Password
	correct := user.Authenticate()
	if !correct {
		//
		dummyReturnSignin := model.ReturnSignin{
			Status: "incorrect login",
		}
		output = &dummyReturnSignin
		return output, nil
	}
	//------------------------------ query data from auth
	var useronly users.Useronly
	var userdata model.ReturnSignin
	useronly.Username = input.Username
	userdata = users.Finduser(useronly)
	fmt.Println(userdata)

	//------------------------------

	// token, err := jwt.GenerateToken(user.Username)
	// if err != nil {
	// 	return "", err
	// }

	//------------------------------

	dummyReturnSignin := model.ReturnSignin{
		Status:          "OK",
		Sessionid:       userdata.Sessionid,
		Clinicid:        userdata.Clinicid,
		Userid:          userdata.Userid,
		Roleid:          userdata.Roleid,
		Accountstatus:   userdata.Accountstatus,
		Sessionlifetime: userdata.Sessionlifetime,
	}
	output = &dummyReturnSignin

	//=================================================
	outlog, err := tstlog.Getcol().InsertOne(ctx, bson.M{"From": "Login", "data_login_out": "data_out", "timestamp": time.Now()})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(outlog)
	//=================================================

	return output, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	username, err := jwt.ParseToken(input.Token)
	if err != nil {
		return "", fmt.Errorf("access denied")
	}
	token, err := jwt.GenerateToken(username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *queryResolver) Links(ctx context.Context) ([]*model.LinkSignout, error) {
	var resultLinks []*model.LinkSignout
	var dbLinks []linksignout.Link = linksignout.GetAll()
	for _, link := range dbLinks {
		resultLinks = append(resultLinks, &model.LinkSignout{Username: link.Username, Clinicid: link.Clinicid, Signouttime: link.Signouttime})
	}
	return resultLinks, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
