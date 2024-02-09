package views

import (
	"encoding/json"
	"github.com/joegasewicz/gomek"
	"github.com/joegasewicz/noticeboard/identity_api/models"
	"github.com/joegasewicz/noticeboard/identity_api/schemas"
	"github.com/joegasewicz/noticeboard/identity_api/utilities"
	"io"
	"log"
	"net/http"
	"strings"
)

type Users struct{}

func (u *Users) Delete(w http.ResponseWriter, r *http.Request, d *gomek.Data) {

}

func (u *Users) Get(w http.ResponseWriter, r *http.Request, d *gomek.Data) {
	token := r.Header.Get("Authorization")
	if token == "" {
		log.Println("No token in headers")
		gomek.JSON(w, nil, http.StatusUnauthorized)
		return
	}
	tokenSplit := strings.Split(token, " ")[1]
	payload, err := utilities.ParseToken(tokenSplit)
	if err != nil {
		log.Println("Error parsing token")
		gomek.JSON(w, nil, http.StatusUnauthorized)
		return
	}
	var user models.User
	result := utilities.DB.First(&user, payload["user_id"])
	if result.Error != nil {
		log.Println(err.Error())
		gomek.JSON(w, nil, http.StatusUnauthorized)
		return
	}
	userDump := schemas.UserOmitPassword{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		RoleID:   user.RoleID,
		Avatar:   user.Avatar,
	}
	gomek.JSON(w, userDump, http.StatusOK)
}

func (u *Users) Post(w http.ResponseWriter, r *http.Request, d *gomek.Data) {
	// Get roles
	var user models.User
	var userSchema schemas.User

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read request body")
		return
	}
	err = json.Unmarshal(body, &userSchema)
	if err != nil {
		log.Println("Couldn't unmarshal JSON")
		gomek.JSON(w, nil, http.StatusInternalServerError)
		return
	}

	var role models.Role
	roleResult := utilities.DB.First(&role, "name = ?", "author")
	if roleResult.Error != nil {
		log.Println("Couldn't fetch user role form database")
		gomek.JSON(w, nil, http.StatusInternalServerError)
		return
	}
	hashedPassword, err := utilities.Hash(userSchema.Password)
	if err != nil {
		log.Println(err.Error())
		gomek.JSON(w, nil, http.StatusInternalServerError)
		return
	}
	user = models.User{
		Username: userSchema.Email,
		Email:    userSchema.Email,
		Password: hashedPassword,
		RoleID:   role.ID,
	}
	result := utilities.DB.Create(&user)
	if result.Error != nil {
		log.Println("Couldn't create new user")
		gomek.JSON(w, nil, http.StatusInternalServerError)
		return
	}
	// token
	userID := user.ID
	token, err := utilities.CreateNewJWT(userID)
	if err != nil {
		log.Println("error creating jwt token")
		gomek.JSON(w, nil, http.StatusInternalServerError)
		return
	}
	resData := schemas.ResponseData{
		Token: token,
	}
	gomek.JSON(w, resData, http.StatusAccepted)
}

func (u *Users) Put(w http.ResponseWriter, r *http.Request, d *gomek.Data) {

}
