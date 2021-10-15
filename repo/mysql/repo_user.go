package mysql

import (
	"context"

	"github.com/huyhoang1990/beego-investigation/entity"
	"github.com/huyhoang1990/beego-investigation/repo"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func NewUserRepo(db *gorm.DB) repo.UserRepo {
	return &userRepo{
		db: db,
	}
}

type userRepo struct {
	db *gorm.DB
}

func (repo *userRepo) FindUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	var row UserDao
	q := repo.db.WithContext(ctx).
		Where("username = ?", username).
		First(&row)
	if err := q.Error; err != nil {
		return nil, err
	}
	return row.toStruct(), nil
}

func (repo *userRepo) InsertOne(ctx context.Context, newUser *entity.User) error {
	var row = new(UserDao).fromStruct(newUser)
	q := repo.db.WithContext(ctx).
		Omit(clause.Associations).
		Create(&row)
	return q.Error
}

func (repo *userRepo) FindById(ctx context.Context, id string) (*entity.User, error) {
	var row UserDao
	q := repo.db.WithContext(ctx).
		Where("id = ?", id).
		First(&row)
	if err := q.Error; err != nil {
		return nil, err
	}
	return row.toStruct(), nil
}

type UserDao struct {
	ID        string `gorm:"column:id;primaryKey"`
	Username  string `gorm:"column:username;not null;unique"`
	Password  string `gorm:"column:password;not null"`
	CreatedAt int64  `gorm:"column:created_at;not null"`
	ExpiredAt int64  `gorm:"column:expired_at;not null"`
	Status    string `gorm:"column:status;not null"`
}

func (dao *UserDao) TableName() string {
	return "users"
}

func (dao *UserDao) fromStruct(user *entity.User) *UserDao {
	dao.ID = user.ID
	dao.Username = user.Username
	dao.Password = user.Password
	dao.CreatedAt = user.CreatedAt
	dao.ExpiredAt = user.ExpiredAt
	dao.Status = string(user.Status)
	return dao
}

func (dao *UserDao) toStruct() *entity.User {
	return &entity.User{
		ID:        dao.ID,
		Username:  dao.Username,
		Password:  dao.Password,
		CreatedAt: dao.CreatedAt,
		ExpiredAt: dao.ExpiredAt,
		Status:    entity.UserStatus(dao.Status),
	}
}
