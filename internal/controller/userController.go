package controller

import (
	"strconv"
	"user/internal/request"
	"user/internal/service"
	"user/pkg/ecode"
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

	response.Success(ctx, gin.H{
		"user":  user,
		"token": "这里后面替换成真实 JWT",
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
