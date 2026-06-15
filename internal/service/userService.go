package service

import (
	"errors"
	"fmt"
	"user/internal/model"
	"user/internal/repository"
	"user/internal/request"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) Register(req *request.RegisterReq) (*model.User, error) {
	if req == nil {
		return nil, fmt.Errorf("请求不能为空")
	}

	// 查询用户（包含已软删除的记录），处理唯一索引冲突
	oldUser, err := s.userRepo.GetByUserNameUnscoped(req.Username)
	if err == nil && oldUser.ID > 0 {
		// 如果用户已被软删除，先物理删除再重新创建
		if oldUser.DeletedAt.Valid {
			if err := s.userRepo.HardDeleteByID(oldUser.ID); err != nil {
				return nil, fmt.Errorf("清理旧用户失败")
			}
		} else {
			return nil, fmt.Errorf("当前用户已存在")
		}
	}

	hashBytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, nil
	}
	user := &model.User{
		Username:     req.Username,
		PasswordHash: string(hashBytes),
		Nickname:     req.Nickname,
		Email:        req.Email,
		Phone:        req.Phone,
		Status:       1,
	}
	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) Login(req *request.LoginReq) (*model.User, error) {
	user, err := s.userRepo.GetByUserName(req.Username)
	if err != nil {
		return nil, fmt.Errorf("当前用户不存在")
	}
	if user.Status != 1 {
		return nil, fmt.Errorf("当前用户已禁用")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, errors.New("密码错误")
	}
	return user, nil
}

func (s *UserService) GetByID(id int64) (*model.User, error) {
	return s.userRepo.GetByID(id)
}

func (s *UserService) List(page, pageSize int) ([]model.User, int64, error) {
	return s.userRepo.List(page, pageSize)
}

func (s *UserService) DeleteByID(id int64) error {
	return s.userRepo.DeleteByID(id)
}
