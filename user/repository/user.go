package repository

import (
	"context"
	"errors"
	"log"
	"strconv"

	"gorm.io/gorm"

	"github.com/mokhlesur-rahman/golang-basic-crud-api-server/domain"
	"github.com/mokhlesur-rahman/golang-basic-crud-api-server/domain/dto"
	"github.com/mokhlesur-rahman/golang-basic-crud-api-server/internal/config"
	"github.com/mokhlesur-rahman/golang-basic-crud-api-server/internal/utils"
)

func New(db *gorm.DB) domain.AuthRepository {
	return &AuthSqlStorage{
		db: db,
	}
}

type AuthSqlStorage struct {
	db *gorm.DB
}

func (a *AuthSqlStorage) User(ctx context.Context, ctr *domain.User) (*domain.User, error) {
	hashedPassword := ""

	if ctr.Password != "" && ctr.PasswordConfirm != "" {
		if ctr.Password != ctr.PasswordConfirm {
			return nil, errors.New("password doesn't match")
		} else {
			hash, err := utils.HashPassword(ctr.Password)
			if err != nil {
				log.Println("Password hashing failed")
			}
			hashedPassword = hash
		}
	} else {
		return nil, errors.New("please provide both password")
	}

	ctr.Password = hashedPassword
	ctr.PasswordConfirm = hashedPassword

	user := domain.User{}

	if ctr.Email != "" {
		//Check if email already exists
		mail := a.db.First(&user, "email=?", ctr.Email)
		cred := &domain.User{}
		if err := mail.WithContext(ctx).Take(cred).Error; err != nil {
			log.Println(err)
		}
		if cred.Email != "" {
			return nil, errors.New("email already exists")
		}
	}

	if err := a.db.Create(ctr).Error; err != nil {
		return nil, err
	}

	// Retrieve the ID of the newly created record
	createdId := ctr.ID

	userResp := domain.User{
		ID:              createdId,
		Name:            ctr.Name,
		Email:           ctr.Email,
		CreatedAt:       ctr.CreatedAt,
		Password:        "",
		PasswordConfirm: "",
	}

	return &userResp, nil
}

func (a *AuthSqlStorage) SignIn(ctx context.Context, ctr *dto.SignIn) (*domain.JWTToken, error) {
	jwt := config.JWT()
	user := domain.User{}

	if ctr.Email != "" && ctr.Password != "" {
		qry := a.db.Find(&user, "email=?", ctr.Email)
		cred := &domain.User{}
		if err := qry.WithContext(ctx).Take(cred).Error; err != nil {
			log.Println(err)
		}
		if cred.Email == "" {
			err := errors.New("Invalid Data")
			return nil, err
		}

		if err := utils.VerifyPassword(cred.Password, ctr.Password); err != nil {
			err := errors.New("Invalid Password")
			return nil, err
		}

		token, err := utils.GenerateToken(jwt.ExpiredIn, strconv.Itoa(int(cred.ID)), jwt.Secret)
		if err != nil {
			log.Println(err)
		}

		loggedInData := &domain.LoggerInUserData{}
		loggedInData.Name = user.Name
		loggedInData.Email = user.Email
		loggedInData.ID = user.ID

		reqJwt := &domain.JWTToken{User: loggedInData, Secret: token, MaxAge: jwt.MaxAge, ExpiredIn: jwt.ExpiredIn, Message: "success"}
		return reqJwt, nil
	}

	err := errors.New("Invalid Data")
	return nil, err

}
