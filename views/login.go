package views

import (
	"encoding/json"
	"github.com/joegasewicz/gomek"
	"github.com/joegasewicz/locksmith/models"
	"github.com/joegasewicz/locksmith/schemas"
	"github.com/joegasewicz/locksmith/utilities"
	"io"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request, d *gomek.Data) {
	var credentials schemas.Login
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}
	err = json.Unmarshal(body, &credentials)
	if err != nil {
		log.Println(err.Error())
		gomek.JSON(w, nil, http.StatusBadRequest)
	}
	var user models.User
	result := utilities.DB.First(&user, "email = ?", credentials.Email)
	if result.Error != nil {
		log.Println(err.Error())
		gomek.JSON(w, nil, http.StatusInternalServerError)
		return
	}
	if result.RowsAffected == 0 {
		log.Println("User not found")
		gomek.JSON(w, nil, http.StatusNotFound)
		return
	}
	passwordsMatch := utilities.Compare(user.Password, credentials.Password)
	if !passwordsMatch {
		log.Println("Passwords don't match")
		gomek.JSON(w, nil, http.StatusUnauthorized)
		return
	}
	token, err := utilities.CreateNewJWT(user.ID)
	userWithToken := &schemas.UserWithToken{
		User: schemas.UserOmitPassword{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			RoleID:   user.RoleID,
			Avatar:   user.Avatar,
		},
		Token: token,
	}
	gomek.JSON(w, userWithToken, http.StatusOK)
}
