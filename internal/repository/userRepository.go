package repository

import (
	"user/internal/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) GetByUserName(username string) (*model.User, error) {
	var user model.User
	err := r.db.Where("username = ?", username).First(&user).Error
	return &user, err
}

// GetByUserNameUnscoped 查询用户（包含已软删除的记录）
func (r *UserRepository) GetByUserNameUnscoped(username string) (*model.User, error) {
	var user model.User
	err := r.db.Unscoped().Where("username = ?", username).First(&user).Error
	return &user, err
}

func (r *UserRepository) GetByID(id int64) (*model.User, error) {
	var user model.User
	err := r.db.Where("id = ?", id).First(&user).Error
	return &user, err
}

func (r *UserRepository) List(page, pagesize int) ([]model.User, int64, error) {
	var users []model.User
	var total int64

	db := r.db.Model(&model.User{})
	db.Count(&total)
	offset := (page - 1) * pagesize
	if err := db.Offset(offset).Limit(pagesize).Find(&users).Error; err != nil {
		return nil, 0, err
	}
	return users, total, nil
}

func (r *UserRepository) UpdateProfile(id int64, nickname, email, phone, avatar string) error {
	updates := map[string]any{
		"nickname": nickname,
		"email":    email,
		"phone":    phone,
		"avatar":   avatar,
	}
	return r.db.Model(&model.User{}).Where("id = ?", id).Updates(updates).Error
}

func (r *UserRepository) DeleteByID(id int64) error {
	return r.db.Delete(&model.User{}, id).Error
}

// HardDeleteByID 物理删除（忽略软删除，直接从数据库删除记录）
func (r *UserRepository) HardDeleteByID(id int64) error {
	return r.db.Unscoped().Delete(&model.User{}, id).Error
}
