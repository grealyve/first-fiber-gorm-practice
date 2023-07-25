package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/grealyve/first-fiber-gorm-practice/models"
	"github.com/grealyve/first-fiber-gorm-practice/storage"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

func (r *Repository) NewBook(c *fiber.Ctx) error {
	book := Book{}

	err := c.BodyParser(&book)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "request failed"})
		return err
	}

	err = r.DB.Create(&book).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "couldn't create a new book"})
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{"message": "book has been added"})
	return nil
}

func (r *Repository) GetBooks(c *fiber.Ctx) error {
	bookModels := &[]models.Books{}

	err := r.DB.Find(bookModels).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "couldn't get books"})
		return err
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{"message": "books fetched successfully", "data": bookModels})
	return nil
}

func (r *Repository) DeleteBook(c *fiber.Ctx) error {
	bookModel := &models.Books{}
	id := c.Params("id")

	if id == "" {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "id cannot be empty"})
		return nil
	}

	err := r.DB.Delete(bookModel, id)
	if err.Error != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "couldn't delete the book"})
		return err.Error
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{"message": "the book deleted successfully"})
	return nil
}

func (r *Repository) GetBookByID(c *fiber.Ctx) error {
	id := c.Params("id")
	bookModel := &models.Books{}

	if id == "" {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "id cannot be empty"})
		return nil
	}

	fmt.Println("the ID is", id)

	err := r.DB.Where("id = ?", id).First(bookModel).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "couldn't get the book"})
		return err
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{"message": "Here it is", "data": bookModel})
	return nil
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/new_books", r.NewBook)
	api.Delete("delete_book/:id", r.DeleteBook)
	api.Get("/get_book/:id", r.GetBookByID)
	api.Get("/books", r.GetBooks)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("could not load the database")
	}

	err = models.MigrateBooks(db)
	if err != nil {
		log.Fatal("could not migrate DB")
	}

	r := Repository{
		DB: db,
	}

	app := fiber.New()
	r.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
