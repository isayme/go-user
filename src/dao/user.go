package dao

import (
	"github.com/drexedam/gravatar"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/isayme/go-user/src/httperror"
	"github.com/isayme/go-user/src/mongo"
	passwordUtil "github.com/isayme/go-user/src/password"
	"github.com/isayme/go-user/src/schema"
)

// User user dao
type User struct {
	session *mongo.Session
}

// NewUser ...
func NewUser(session *mongo.Session) *User {
	return &User{
		session: session,
	}
}

// Signup ...
func (u *User) Signup(username, email, password string) (*schema.User, error) {
	s, c := u.session.GetCollection("users")
	defer s.Close()

	now := bson.Now()
	user := &schema.User{
		ID:       bson.NewObjectId(),
		Username: username,
		Email:    email,
		Avatar:   gravatar.New(email).AvatarURL(),
		Created:  now,
		Updated:  now,
	}

	password, err := passwordUtil.Generate(password)
	if err != nil {
		return nil, err
	}
	user.Password = password

	err = c.Insert(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Login ...
func (u *User) Login(email, password string) (*schema.User, error) {
	s, c := u.session.GetCollection("users")
	defer s.Close()

	var user schema.User

	selector := bson.M{
		"email": email,
	}
	err := c.Find(selector).One(&user)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, httperror.EmailPasswordNotMatch
		}
		return nil, err
	}

	if passwordUtil.Compare(user.Password, password) == false {
		return nil, httperror.EmailPasswordNotMatch
	}

	return &user, nil
}

// Me ...
func (u *User) Me(ID bson.ObjectId) (*schema.User, error) {
	s, c := u.session.GetCollection("users")
	defer s.Close()

	var user schema.User

	selector := bson.M{
		"_id": ID,
	}
	err := c.Find(selector).One(&user)
	return &user, err
}
