package controller

import (
	"strconv"
	"user/internal/middleware"
	"user/internal/request"
	"user/internal/service"
	"user/pkg/ecode"
	auth "user/pkg/jwt"
	"user/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (c *UserController) Register(ctx *gin.Context) {
	var req request.RegisterReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Fail(ctx, ecode.ParamError, "参数错误")
		return
	}
	user, err := c.userService.Register(&req)
	if err != nil {
		response.Fail(ctx, ecode.UserExists, "用户已存在")
		return
	}
	response.Success(ctx, user)
}

func (c *UserController) Login(ctx *gin.Context) {
	var req request.LoginReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Fail(ctx, ecode.ParamError, "参数错误")
		return
	}

	user, err := c.userService.Login(&req)
	if err != nil {
		response.Fail(ctx, ecode.PasswordError, err.Error())
		return
	}

	token, err := auth.GenerateToken(user.ID, user.Username)
	if err != nil {
		response.Fail(ctx, ecode.InternalError, "Token 生成失败")
		return
	}
	response.Success(ctx, gin.H{
		"user":  user,
		"token": token,
	})
}

func (c *UserController) List(ctx *gin.Context) {
	var req request.ListUsersReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.Fail(ctx, ecode.ParamError, "参数错误")
		return
	}
	users, total, err := c.userService.List(req.Page, req.PageSize)
	if err != nil {
		response.Fail(ctx, ecode.InternalError, "查询失败")
		return
	}

	response.Success(ctx, gin.H{
		"list":      users,
		"total":     total,
		"page":      req.Page,
		"page_size": req.PageSize,
	})

}

func (c *UserController) UpdateProfile(ctx *gin.Context) {
	userIDValue, exists := ctx.Get(middleware.CurrentUserIDKey)
	if !exists {
		response.Fail(ctx, ecode.TokenInvalid, "未获取到登录用户")
		return
	}

	userID, ok := userIDValue.(int64)
	if !ok {
		response.Fail(ctx, ecode.TokenInvalid, "用户 ID 类型错误")
		return
	}

	var req request.UpdateProfileReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Fail(ctx, ecode.ParamError, "参数错误")
		return
	}
	if err := c.userService.UpdateProfile(userID, &req); err != nil {
		response.Fail(ctx, ecode.UserNotFound, "更新用户信息失败")
		return
	}
	user, err := c.userService.GetByID(userID)
	if err != nil {
		response.Fail(ctx, ecode.UserNotFound, "用户不存在")
	}
	
	response.Success(ctx, user)

}

func (c *UserController) GetProfile(ctx *gin.Context) {
	userIDVal, exit := ctx.Get(middleware.CurrentUserIDKey)
	if !exit {
		response.Fail(ctx, ecode.ParamError, "未获取到登陆用户")
		return
	}
	userID, ok := userIDVal.(int64)
	if !ok {
		response.Fail(ctx, ecode.TokenInvalid, "用户ID类型错误")
		return
	}
	user, err := c.userService.GetByID(userID)
	if err != nil {
		response.Fail(ctx, ecode.UserNotFound, "用户不存在")
		return
	}

	response.Success(ctx, user)
}

func (c *UserController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.Fail(ctx, ecode.ParamError, "用户 ID 错误")
		return
	}
	if err := c.userService.DeleteByID(id); err != nil {
		response.Fail(ctx, ecode.InternalError, "删除失败")
		return
	}
	response.Success(ctx, nil)
}
