package jwt

import (
	"mt-scale/common"
	"mt-scale/entitys"
	"mt-scale/exception"
	"mt-scale/syslog"
	"mt-scale/utils"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	ignoreURL        []string
	defaultErrorFunc = func(code int) {
		exception.ThrowBusinessError(code)
	}
	tokenExp int = 48
)

// SecretKey
const (
	SecretKey = "my secret key"
)

// Options stores configurations for a JWT middleware.
type Options struct {
	ignoreURL []string
}

func init() {
	ignoreURL = common.GetConfArr("service.ignoreurl")
}

// Middleware validates JWT token.
func Middleware(options Options) gin.HandlerFunc {
	urls := options.ignoreURL
	if urls != nil {
		for _, v := range urls {
			ignoreURL = append(ignoreURL, v)
		}
	}

	return func(c *gin.Context) {
		path := c.Request.URL.Path
		if index := utils.MatchPath(ignoreURL, path); index != -1 {
			return
		}
		if !strings.HasPrefix(path, "/api/") {
			return
		}
		claims := getClaim(c)
		if claims == nil {
			defaultErrorFunc(common.JWTAuthFailedCode)
		}
		exp := claims["exp"]
		if exp == nil || exp.(float64) < float64(time.Now().Unix()) {
			defaultErrorFunc(common.TokenExpiredCode)
		}

		username := ""
		if val, ok := claims["username"]; ok {
			username = val.(string)
		}

		user := entitys.User{
			Username: username,
		}

		c.Set("loginuser", user)

		tokenString := GetTokenString(user)
		c.Header("jwt", tokenString)
	}
}

// GetTokenString Get token string
func GetTokenString(user entitys.User) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	if exp := common.GetConfInt("jwt.token.exp"); exp > 0 {
		tokenExp = exp
	}
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenExp)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["username"] = user.Username

	token.Claims = claims

	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		// log.Debug(err)
	}
	return tokenString
}

func getClaim(c *gin.Context) jwt.MapClaims {
	tokenString := c.GetHeader("Authorization")

	if len(tokenString) < 15 {
		return nil
	}

	syslog.Info(tokenString[7:])
	tokenString = tokenString[7:]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		return nil
	}

	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil
	}

	return claim
}
