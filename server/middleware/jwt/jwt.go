package jwt

import (
	"errors"
	"mt-scale/utils"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	ignoreURL        []string
	defaultErrorFunc = func(c *gin.Context) {
		panic(errors.New(">>>>>> :jwt token mismatch"))
	}
	tokenExp int = 48
)

// SecretKey
const (
	SecretKey = "my secret key"
)

// User User entity
type User struct {
	Email      string
	Password   string
	Firstname  string
	Middlename string
	Lastname   string
	Role       string
	Address    string
	Group      string
	User       string
	Time       time.Time
}

// Options stores configurations for a JWT middleware.
type Options struct {
	ignoreURL []string
	ErrorFunc gin.HandlerFunc
}

func init() {
	ignoreURL = utils.GetConfArr("service.ignoreurl")
}

// Middleware validates JWT token.
func Middleware(options Options) gin.HandlerFunc {
	urls := options.ignoreURL
	if urls != nil {
		for _, v := range urls {
			ignoreURL = append(ignoreURL, v)
		}
	}

	errorFunc := options.ErrorFunc
	if errorFunc == nil {
		errorFunc = defaultErrorFunc
	}

	return func(c *gin.Context) {
		path := c.Request.URL.Path
		if _, f := utils.Find(ignoreURL, path); f {
			return
		}
		claim := getClaim(c)
		if claim == nil {
			return
		}
		exp := claim["exp"]
		if exp == nil || exp.(float64) < float64(time.Now().Unix()) {
			// log.Info("token expired, redirect to login")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Login expired, login again."})
			return
		}

		email := ""
		if val, ok := claim["email"]; ok {
			email = val.(string)
		}
		role := ""
		if val, ok := claim["role"]; ok {
			role = val.(string)
		}
		group := ""
		if val, ok := claim["group"]; ok {
			group = val.(string)
		}
		firstname := ""
		if val, ok := claim["firstname"]; ok {
			firstname = val.(string)
		}
		lastname := ""
		if val, ok := claim["lastname"]; ok {
			lastname = val.(string)
		}

		user := User{
			Email:     email,
			Role:      role,
			Group:     group,
			Firstname: firstname,
			Lastname:  lastname,
		}

		c.Set("loginuser", user)

		// log.Debug(user)

		tokenString := getTokenString(user)

		// log.Debug("New TOKEN:   ", tokenString)
		c.Header("jwt", tokenString)
	}
}

func getTokenString(user User) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	if exp := utils.GetConfInt("jwt.token.exp"); exp > 0 {
		tokenExp = exp
	}
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenExp)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["email"] = user.Email
	claims["role"] = user.Role
	claims["group"] = user.Group
	claims["firstname"] = user.Firstname
	claims["lastname"] = user.Lastname

	token.Claims = claims

	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		// log.Debug(err)
	}
	return tokenString
}

func getClaim(c *gin.Context) jwt.MapClaims {
	tokenString := c.GetHeader("Authorization")

	// log.Debug(tokenString)

	if len(tokenString) < 15 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Unauthorized"})
		return nil
	}

	// log.Debug(tokenString[7:])
	tokenString = tokenString[7:]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		// log.Debug(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Unauthorized"})
		return nil
	}

	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		// log.Error(ok)
	}
	return claim
}
