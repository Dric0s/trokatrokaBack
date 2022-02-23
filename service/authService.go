package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
	"os"
	"time"
	"trokatrokaBack/model"
)

func CreateToken(idUser uint64) (*model.TokenDetails, error) {

	tokenDetails := &model.TokenDetails{}
	tokenDetails.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	tokenDetails.AccessUuid = uuid.NewV4().String()

	tokenDetails.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	tokenDetails.RefreshUuid = uuid.NewV4().String()

	var err error

	os.Setenv("ACCESS_SECRET", "TCH0LA5") //Colocar em um env file
	authorizationTokenClaims := jwt.MapClaims{}
	authorizationTokenClaims["authorized"] = true
	authorizationTokenClaims["access_uuid"] = tokenDetails.AccessUuid
	authorizationTokenClaims["user_id"] = idUser
	authorizationTokenClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	authorizationToken := jwt.NewWithClaims(jwt.SigningMethodHS256, authorizationTokenClaims)

	tokenDetails.AccessToken, err = authorizationToken.SignedString([]byte(os.Getenv("ACCESS_SECRET")))

	if err != nil {
		return nil, err
	}

	os.Setenv("REFRESH_SECRET", "N01AS") //Colocar em um env file
	refreshTokenClaims := jwt.MapClaims{}
	refreshTokenClaims["authorized"] = true
	refreshTokenClaims["access_uuid"] = tokenDetails.RefreshUuid
	refreshTokenClaims["user_id"] = idUser
	refreshTokenClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	tokenDetails.RefreshToken, err = refreshToken.SignedString([]byte(os.Getenv("ACCESS_SECRET")))

	if err != nil {
		return nil, err
	}

	return tokenDetails, nil
}
