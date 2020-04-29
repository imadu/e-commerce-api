package auth

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/imadu/e-commerce-api/users"
	"github.com/labstack/echo"
)

//Login function returns jwt token for authorization
func Login(context echo.Context) error {
	username := context.FormValue("username")
	password := context.FormValue("password")

	usernameDetails, err := users.GetUsername(username)
	if err != nil {
		return err
	}

	correctPassword := users.CheckPasswordHash(password, usernameDetails.Password)

	if correctPassword == false {
		return echo.ErrUnauthorized
	}

	token := jwt.New(jwt.SigningMethodHS256)

	//set the claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = usernameDetails.Username
	claims["email"] = usernameDetails.Email
	claims["phone"] = usernameDetails.Phonenumber
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Generate encoded token and send it as response.
	generatedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, map[string]string{
		"token": generatedToken,
	})

}
