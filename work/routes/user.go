package routes

import (
	"errors"

	"github.com/cosmovid21/fiber/database"
	"github.com/cosmovid21/fiber/models"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateResponseUser (user models.User) (User){
return User{ID: user.ID, FirstName: user.FirstName, LastName: user.LastName}
}

func CreateUser( c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(& user)
	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}

	database.Database.Db.Find(&users)
	responseUsers := []User{}
	for _, user := range users {
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
}

return c.Status(200).JSON(responseUsers)
}

func findUser(id int, user *models.User)error {
	database.Database.Db.Find(&user, "id = ?", id )
	if user.ID == 0 {
		return errors.New("non existent")
	}
	return nil
}
 

func GetUser(c *fiber.Ctx)error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).SendString("check you id again")
	}

	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

func Updateuser (c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := findUser(id, &user); err != nil {
		return c.Status(200).JSON(err.Error())
	}

	type Updateuser struct {
		FirstName string `json:"first_name"`
		LastName string `json:"last_name"`
	}

	var updateData Updateuser

	if err := c.BodyParser(&updateData); err != nil{
		return c.Status(590).JSON(err.Error())

	}

	user.FirstName = updateData.FirstName
	user.LastName = updateData.LastName

	database.Database.Db.Save(&user)
	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)
}

func DeleteUser (c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var user models.User

	if err != nil {
		return c.Status(400).JSON("please ensure that it is an id")
	}

	if err := findUser(id, &user); err != nil {
		return  c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&user).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.Status(200).SendString("successfully deleted")
}
