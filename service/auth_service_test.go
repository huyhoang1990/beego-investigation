package service

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/huyhoang1990/beego-investigation/entity"
	"github.com/huyhoang1990/beego-investigation/infras"
	"github.com/huyhoang1990/beego-investigation/repo/mysql"
	"gorm.io/gorm"
)

func GetDatabaseConnection() *gorm.DB {
	db, err := infras.NewMysqlSession()
	if err != nil {
		panic(err)
	}
	return db
}

func getUserRepo() AuthService {
	userRepo := mysql.NewUserRepo(GetDatabaseConnection())
	passwordService := NewPasswordService()
	svc := NewAuthService(userRepo, passwordService)
	return svc

}

func TestAddNewUser(t *testing.T) {
	svc := getUserRepo()

	_, err := svc.AddNewUser(context.Background(), &entity.User{
		Username: "hoang",
		Password: "123456",
	})
	if err.Error() == "this username is registered" {
		fmt.Println("success -- this username is registered")
	}

	_, err = svc.AddNewUser(context.Background(), &entity.User{
		Username: "hoang1",
		Password: "123456",
	})

	if strings.Contains(err.Error(), "password must be larger") {
		fmt.Println("success -- password must be larger or equal to 8 character")
	}

	_, err = svc.AddNewUser(context.Background(), &entity.User{
		Username: "hoang1",
		Password: "sdf123456!435345@",
	})

	if strings.Contains(err.Error(), "password is not strong enough") {
		fmt.Println("success -- password is not strong enough")
	}

	user, err := svc.AddNewUser(context.Background(), &entity.User{
		Username: "hoang2",
		Password: "Hoang@1234",
	})

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	fmt.Println(user)
}
