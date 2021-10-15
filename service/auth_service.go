package service

import (
	"context"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/huyhoang1990/beego-investigation/entity"
	"github.com/huyhoang1990/beego-investigation/repo"
	"github.com/huyhoang1990/beego-investigation/utils"
	"golang.org/x/crypto/argon2"
)

func NewAuthService(
	userRepo repo.UserRepo,
	passwordSvc PasswordService,
) AuthService {
	return &authService{
		userRepo: userRepo,
		passSrv:  passwordSvc,
		passwordConfig: &PasswordConfig{
			time:    1,
			memory:  64 * 1024,
			threads: 4,
			keyLen:  24,
		},
	}
}

type AuthService interface {
	AddNewUser(ctx context.Context, user *entity.User) (*entity.User, error)
	ValidateUser(ctx context.Context, username string, password string) (*entity.User, error)
}

type authService struct {
	userRepo       repo.UserRepo
	passSrv        PasswordService
	passwordConfig *PasswordConfig
}

func (svc *authService) AddNewUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	existed, err := svc.userRepo.FindUserByUsername(ctx, user.Username)
	if err == nil && existed != nil {
		return nil, fmt.Errorf("this username is registered")
	} else if err = svc.passSrv.ValidateStrength(user.Password); err != nil {
		return nil, err
	}

	hashed, err := svc.generatePassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.ID = uuid.New().String()
	user.Password = hashed
	fmt.Println(hashed)
	user.CreatedAt = utils.GetUnixTimeNow()
	user.ExpiredAt = utils.AddUnixTimeNow(1 * 24 * 60 * 60)
	user.Status = entity.UserStatus_ACTIVE

	if err := svc.userRepo.InsertOne(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (svc *authService) ValidateUser(ctx context.Context, username string, password string) (*entity.User, error) {
	fmt.Println("==========================vaiiiasdfasdfasdfasdfasgdfhgaskjfhgaskdhjfgkajshdfglllllllll======")
	currentUser, err := svc.userRepo.FindUserByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("incorrect email or password")
	}

	valid, err := svc.comparePassword(password, currentUser.Password)

	if err != nil {
		return nil, err
	}

	if !valid {
		return nil, fmt.Errorf("incorrect email or password")
	}

	if currentUser.IsInActive() {
		return nil, fmt.Errorf("your account is deactivated")
	}

	return currentUser, nil
}

func (svc *authService) comparePassword(password string, hash string) (bool, error) {
	parts := strings.Split(hash, "$")

	c := &PasswordConfig{}
	_, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &c.memory, &c.time, &c.threads)
	if err != nil {
		return false, err
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, err
	}

	decodedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, err
	}
	c.keyLen = uint32(len(decodedHash))

	comparisonHash := argon2.IDKey([]byte(password), salt, c.time, c.memory, c.threads, c.keyLen)

	return subtle.ConstantTimeCompare(decodedHash, comparisonHash) == 1, nil
}

func (svc *authService) generatePassword(password string) (string, error) {
	// Generate a Salt
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, svc.passwordConfig.time, svc.passwordConfig.memory, svc.passwordConfig.threads, svc.passwordConfig.keyLen)

	// Base64 encode the salt and hashed password.
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	format := "$funfun$v=%d$m=%d,t=%d,p=%d$%s$%s"
	full := fmt.Sprintf(format, argon2.Version, svc.passwordConfig.memory, svc.passwordConfig.time, svc.passwordConfig.threads, b64Salt, b64Hash)
	return full, nil
}

type PasswordConfig struct {
	time    uint32
	memory  uint32
	threads uint8
	keyLen  uint32
}
