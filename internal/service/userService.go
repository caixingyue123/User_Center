package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"
	"user/internal/model"
	"user/internal/repository"
	"user/internal/request"

	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

const userInfoCacheTTL = 30 * time.Minute

type UserService struct {
	userRepo    *repository.UserRepository
	redisClient *redis.Client
}

func NewUserService(userRepo *repository.UserRepository, redisClient *redis.Client) *UserService {
	return &UserService{
		userRepo:    userRepo,
		redisClient: redisClient,
	}
}

func userInfoCacheKey(id int64) string {
	return fmt.Sprintf("user:info:%d", id)
}

func parseCachedUser(cacheValue string) (*model.User, bool) {
	var user model.User
	if err := json.Unmarshal([]byte(cacheValue), &user); err != nil {
		return nil, false
	}
	return &user, true
}

func marshalCachedUser(user *model.User) ([]byte, error) {
	return json.Marshal(user)
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

func (s *UserService) UpdateProfile(id int64, req *request.UpdateProfileReq) error {
	_, err := s.userRepo.GetByID(id)
	if err != nil {
		return err
	}
	if err := s.userRepo.UpdateProfile(id, req.Nickname, req.Email, req.Phone, req.Avatar); err != nil {
		return err
	}
	s.deleteUserInfoCache(id)
	return nil
}

func (s *UserService) GetByID(id int64) (*model.User, error) {
	ctx := context.Background()
	cacheKey := userInfoCacheKey(id)

	if s.redisClient != nil {
		cacheValue, err := s.redisClient.Get(ctx, cacheKey).Result()
		if err == nil {
			if user, ok := parseCachedUser(cacheValue); ok {
				return user, nil
			}
		}
	}
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if s.redisClient != nil {
		cacheBytes, err := marshalCachedUser(user)
		if err == nil {
			_ = s.redisClient.Set(ctx, cacheKey, cacheBytes, userInfoCacheTTL).Err()
		}
	}
	return user, nil
}

func (s *UserService) List(page, pageSize int) ([]model.User, int64, error) {
	return s.userRepo.List(page, pageSize)
}

func (s *UserService) DeleteByID(id int64) error {
	if err := s.userRepo.DeleteByID(id); err != nil {
		return err
	}
	s.deleteUserInfoCache(id)
	return nil
}

func (s *UserService) deleteUserInfoCache(id int64) {
	if s.redisClient == nil {
		return
	}
	_ = s.redisClient.Del(context.Background(), userInfoCacheKey(id)).Err()
}
