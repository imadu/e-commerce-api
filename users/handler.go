package users

import (
	"log"
	"strconv"

	"github.com/imadu/e-commerce-api/util"
	"github.com/labstack/echo"
)

//GetUserID returns the user by id
func GetUserID(context echo.Context) error {
	id := context.Param("id")

	user, err := GetUser(id)
	if err != nil {
		log.Fatal(err)
		return util.SendError(context, "failed", "something went wrong", "400")
	}
	return util.SendData(context, user)

}

// GetAllUsers returns all the users
func GetAllUsers(context echo.Context) error {
	limit, _ := strconv.ParseInt(context.QueryParam("limit"), 10, 64)
	size, _ := strconv.ParseInt(context.QueryParam("size"), 10, 64)

	users, err := GetUsers(limit, size)
	if err != nil {
		log.Fatal(err)
		return util.SendError(context, "failed", "something went wrong", "500")
	}
	return util.SendData(context, users)
}

// CreateNewUser creates a new user
func CreateNewUser(context echo.Context) error {
	user := new(User)
	if err := context.Bind(user); err != nil {
		log.Fatalf("Could not bind request to struct: %+v", err)
		return util.SendError(context, "", "", "")
	}

	NewUser, err := CreateUser(*user)

	if err != nil {
		log.Fatalf("Could not bind request to struct: %+v", err)
		return util.SendError(context, "", "", "")
	}

	return util.SendData(context, NewUser)

}

//UpdateUserDetails update a user based on id
func UpdateUserDetails(context echo.Context) error {
	id := context.Param("id")
	user := new(User)
	if err := context.Bind(user); err != nil {
		log.Fatalf("Could not bind request to struct: %+v", err)
		return util.SendError(context, "", "", "")
	}

	UpdatedUser, err := UpdateUser(id, *user)

	if err != nil {
		log.Fatalf("Could not bind request to struct: %+v", err)
		return util.SendError(context, "", "", "")
	}

	return util.SendData(context, UpdatedUser)

}

// DeleteUserAccount deletes a user
func DeleteUserAccount(context echo.Context) error {
	id := context.Param("id")

	DeletedID, err := DeleteUser(id)
	if err != nil {
		log.Fatalf("Could not bind request to struct: %+v", err)
		return util.SendError(context, "", "", "")
	}

	return util.SendSuccess(context, DeletedID)

}
