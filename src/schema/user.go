package schema

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

// User 账号
type User struct {
	ID       bson.ObjectId `json:"_id" bson:"_id"`
	Username string        `json:"username" bson:"username"`
	Email    string        `json:"email" bson:"email"`
	Avatar   string        `json:"avatar" bson:"avatar"`
	Password string        `json:"-" bson:"password"`
	Created  time.Time     `json:"created" bson:"created"`
	Updated  time.Time     `json:"updated" bson:"updated"`
}
