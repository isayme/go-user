package router

import (
	"net/http"

	"github.com/globalsign/mgo/bson"

	"github.com/gin-gonic/gin"
	"github.com/isayme/go-user/src/constant"
	"github.com/isayme/go-user/src/dao"
	"github.com/isayme/go-user/src/httperror"
	"github.com/isayme/go-user/src/jwt"
	"github.com/isayme/go-user/src/schema"
)

// User ...
type User struct {
	db *dao.User
}

// NewUser ...
func NewUser(db *dao.User) *User {
	return &User{
		db: db,
	}
}

// SignupRequest signup request body
type SignupRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Signup singup router
func (u *User) Signup(c *gin.Context) {
	var body SignupRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		panic(httperror.InvalidParams.WithErr(err))
	}

	user, err := u.db.Signup(body.Username, body.Password)
	if err != nil {
		panic(err)
	}

	token, err := generateAccessToken(user)
	if err != nil {
		panic(err)
	}

	c.Header(constant.HTTPHeaderAuthorization, token)

	c.JSON(http.StatusOK, user)
}

// LoginRequest login request body
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login login router
func (u *User) Login(c *gin.Context) {
	var body LoginRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		panic(httperror.InvalidParams.WithErr(err))
	}

	user, err := u.db.Login(body.Username, body.Password)
	if err != nil {
		panic(err)
	}

	token, err := generateAccessToken(user)
	if err != nil {
		panic(err)
	}

	c.Header(constant.HTTPHeaderAuthorization, token)

	c.JSON(http.StatusOK, user)
}

// Me get current user info
func (u *User) Me(c *gin.Context) {
	userID, _ := c.Get(constant.UserIDKey)
	user, err := u.db.Me(bson.ObjectIdHex(userID.(string)))
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, user)
}

func generateAccessToken(user *schema.User) (string, error) {
	return jwt.Sign(map[string]interface{}{
		"uid":      user.ID,
		"username": user.Username,
	})
}
