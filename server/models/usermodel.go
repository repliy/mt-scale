package models

import (
	"mt-scale/common"
	"mt-scale/entitys"
	"mt-scale/exception"
	"mt-scale/middleware/jwt"
	"mt-scale/models/dto"
	"mt-scale/models/vo"
	"mt-scale/syslog"
	"mt-scale/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddUser Add a new user
func AddUser(user entitys.User) primitive.ObjectID {
	col, ctx := Collection("user")
	filter := bson.D{
		primitive.E{
			Key:   "username",
			Value: user.Username,
		},
	}
	cur, err := col.Find(ctx, filter)
	if err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.DatabaseErrorCode)
	}
	if cur.Next(ctx) {
		exception.ThrowBusinessErrorMsg("用户信息已存在，请更改...")
	}
	user.Password = utils.Md5(user.Password)
	result, err := col.InsertOne(ctx, user)
	if err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.DatabaseErrorCode)
	}
	return result.InsertedID.(primitive.ObjectID)
}

// Login User login
func Login(dto dto.LoginDto) vo.LoginVo {
	col, ctx := Collection("user")
	filter := bson.D{
		primitive.E{
			Key:   "username",
			Value: dto.Username,
		},
	}
	cur, err := col.Find(ctx, filter)
	if err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.DatabaseErrorCode)
	}
	var result *vo.LoginVo = new(vo.LoginVo)
	var user entitys.User
	if cur.Next(ctx) {
		if err := cur.Decode(&user); err != nil {
			syslog.Error(err)
			exception.ThrowBusinessError(common.DatabaseErrorCode)
		}
		if utils.Md5(dto.Password) != user.Password {
			exception.ThrowBusinessErrorMsg("用户名密码不相符")
		}
		result.ID = user.ID
		result.Username = user.Username
		var token *vo.AuthTokenVo = new(vo.AuthTokenVo)
		token.AccessToken = jwt.GetTokenString(user)
		result.Token = *token
		return *result
	}
	exception.ThrowBusinessErrorMsg("用户名密码不相符")
	return *result
}
