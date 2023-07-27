package routes

import (
	"errors"

	"github.com/cosmovid21/fiber/database"
	"github.com/cosmovid21/fiber/models"
	// "github.com/go-playground/validator/v10/translations/id"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2"
)

type product struct {
	Id           uint   `json:"id"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

func CreateResponseProduct(product *models.Product) (product) {
	// var &product = product{
	return product{Id: product.ID, Name: product.name, SerialNUmber: product.SerialNumber }
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.error())
	}

	database.Database.Db.Create(&product)
	reponseProduct := CreateResponseProduct(product)

	return c.Status(200).JSON(responseProduct)
}


func GetProducts(c *fiber.Ctx) error {
	Products := []models.Product{}

	database.Database.Db.Find(&products)
	responseProducts := []Product{}

	for _,Product := range Products{
		responseProducts := CreateResponseProduct(&Products)
		responseProducts = append(responseProducts responseProduct)
	}

	return c.status(200).JSON(responseProducts)

}

func findProduct(id = int,products *models.Product) error {
	database.Database.Db.find(&Product, "id =?", id)
	if product.id == 0 {
		errors.New("not found")
	}
	return nil;
}

func GetProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var product models.Product

	if err != nil {
		c.status(400).JSON(err.Error())
	}
	if err := findProduct(id, &product); err != nil {
		c.status(400).JSON(err.Error())
	}

	responeProduct := CreateProduct(product)

	return c.Status(200).JSON(responeProduct)
}

func UpdateProduct(c *fiber.Ctx) error {
	id, err = c.ParamsInt("id")

	var product models.Product

	if err != nil{
		c.Status(400).JSON(err.Error())
	}

	if err := findProduct(id, &product); err != nil{
		c.Status(400).JSON(err.Error())
	}

	Type UpdateProduct Struct{
		Name string `json : "name"`
		SerialNumber uint `json: "serial_number"`
	}

	var updateData UpdateProduct 

	

}