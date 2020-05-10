package oauth

import (
	"microservice-go/application/users"
	"microservice-go/domain/oauth"
	"microservice-go/framework/utils/bcrypt"
	"microservice-go/framework/utils/jwt"
)

type OAuthService struct {
	OAuthRepository users.UserRepository
}

func (service *OAuthService) DoLogin(email string, password string) (*oauth.OAuth, error) {
	var auth oauth.OAuth
	user, err := service.OAuthRepository.Fetch(email)
	if err != nil {
		return nil, err
	}

	isValid := bcrypt.ComparedPassword(user.Password, password)
	if isValid == false {
		return nil, err
	}

	generate, err := jwt.GenerateTokenJWT(user)
	if err != nil {
		return nil, err
	}
	auth.Token = generate

	return &auth, nil
}
