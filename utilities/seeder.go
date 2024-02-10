package utilities

import (
	"github.com/joegasewicz/locksmith/models"
	"log"
)

type Seeder struct{}

func (s *Seeder) CreateUser(y *yamlScheme) {
	var users []models.User
	for _, user := range y.Users {
		hashPassword, err := Hash(user.Password)
		if err != nil {
			log.Panic(err.Error())
		}
		var role models.Role
		result := DB.First(&role, "name = ?", user.Role)
		if result.Error != nil {
			log.Fatalf("Error: can not assign a role with name %s\n", user.Role)
		}
		u := models.User{
			Username: user.Name,
			Email:    user.Email,
			Password: hashPassword,
			RoleID:   role.ID,
		}
		users = append(users, u)
	}

	userResult := DB.Create(&users)
	if userResult == nil {
		log.Println("successfully seeded db with users")
	}

}

func (s *Seeder) CreateRoles(y *yamlScheme) {
	for _, role := range y.Roles {
		r := models.Role{
			Name: role,
		}
		rolesResult := DB.Create(&r)
		if rolesResult.Error == nil {
			log.Println("successfully seeded db with roles")
		}
	}
}
