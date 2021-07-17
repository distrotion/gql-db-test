package users

import (
	"context"
	"fmt"
	"log"
	"minefit_auth/graph/model"
	"minefit_auth/mongo/maindb"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID              string `json:"id"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	Createtime      string `json:"createtime"`
	Clinicid        string `json:"clinicid"`
	Roleid          int64  `json:"roleid"`
	Accountstatus   int64  `json:"accountstatus"`
	Sessionlifetime int64  `json:"sessionlifetime"`
}

type Userlogin struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

type Useronly struct {
	Username string `json:"name"`
}

var ctx = context.TODO()

func (user *User) Create() {
	statement, err := maindb.Getcol().InsertOne(ctx, user)
	print(statement)
	if err != nil {
		log.Fatal(err)
	}

	//-------------------------------------------------------------- hash password

	// hashedPassword, err := HashPassword(user.Password)
	// // _, err = statement.Exec(user.Username, hashedPassword)
	// _, err = maindb.Getcol().UpdateOne(ctx, bson.M{"username": user.Username}, bson.M{"$set": bson.M{"password": hashedPassword}})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//--------------------------------------------------------------
}

func GetUserIdByUsername(username string) (int, error) {
	res, err := maindb.Getcol().Find(ctx, username)
	if err != nil {
		log.Fatal(err)
	}

	var msg User
	if err = res.All(ctx, &msg); err != nil {
		panic(err)
	}

	var Id int
	Id, err = strconv.Atoi(msg.ID)
	if err != nil {
		log.Fatal(err)
	}

	return Id, nil
}

func Checkuser(useronly Useronly) []bson.M {
	res, err := maindb.Getcol().Find(ctx, useronly)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(res)

	var msg []bson.M
	if err = res.All(ctx, &msg); err != nil {
		panic(err)
	}
	fmt.Println(len(msg))
	return msg
}

func Finduser(useronly Useronly) model.ReturnSignin {
	res, err := maindb.Getcol().Find(ctx, useronly)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(res)

	var msg []User
	if err = res.All(ctx, &msg); err != nil {
		panic(err)
	}
	fmt.Println(len(msg))

	var out model.ReturnSignin
	out.Sessionid = msg[0].Username
	out.Clinicid = msg[0].Clinicid
	out.Userid = msg[0].ID
	out.Roleid = int(msg[0].Roleid)
	out.Accountstatus = int(msg[0].Accountstatus)
	out.Sessionlifetime = int(msg[0].Sessionlifetime)
	return out
}

func (user *User) Authenticate() bool {
	// statement, err := database.Db.Prepare("select Password from Users WHERE Username = ?")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// row := statement.QueryRow(user.Username)
	//fmt.Println(user)
	res, err := maindb.Getcol().Find(ctx, bson.M{"username": user.Username})
	if err != nil {
		log.Fatal(err)
	}

	var msg []User
	if err = res.All(ctx, &msg); err != nil {
		panic(err)
	}

	fmt.Println(msg)

	//-------------------------------------------------------------- hash password

	// var hashedPassword string
	// hashedPassword = msg[0].Password
	// //fmt.Print(msg.Password)
	// return CheckPasswordHash(user.Password, hashedPassword)

	//--------------------------------------------------------------
	var result bool
	if len(msg) == 0 {
		result = false

	} else {
		if msg[0].Password == user.Password {
			result = true
		} else {
			result = false
		}
	}

	return result

}

//-------------------------------------------------------------------------------function
//HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
