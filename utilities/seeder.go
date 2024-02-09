package utilities

import (
	"github.com/joegasewicz/locksmith/models"
	"log"
)

type Seeder struct{}

func (s *Seeder) CreateUser() {
	var superRole models.Role
	DB.First(&superRole, "name = ?", "super")
	var authRole models.Role
	DB.First(&authRole, "name = ?", "author")
	hashPassword, err := Hash("wizard")
	if err != nil {
		log.Panic(err.Error())
	}
	users := []models.User{
		{
			Email:    "joegasewicz@gmail.com",
			Username: "TestSuper",
			Password: hashPassword,
			RoleID:   superRole.ID,
		},
		{
			Email:    "pymailio@gmail.com",
			Username: "TestAuthor",
			Password: hashPassword,
			RoleID:   authRole.ID,
		},
	}

	userResult := DB.Create(&users)
	if userResult == nil {
		log.Println("successfully seeded db with users")
	}

}

func (s *Seeder) CreateRoles() {
	var superRole models.Role
	superRoleResult := DB.First(&superRole, "name = ?", "super")
	if superRoleResult.Error != nil {
		// This means we have no roles on the db
		roles := []models.Role{
			{Name: "super"},
			{Name: "author"},
			{Name: "editor"},
		}
		rolesResult := DB.Create(&roles)
		if rolesResult.Error == nil {
			log.Println("successfully seeded db with roles")
		}
	}
}
