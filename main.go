package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/benebobaa/harisenin-mini-project/helper"
	"github.com/benebobaa/harisenin-mini-project/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/lib/pq"
	"time"
)

var db *sql.DB

func initDB(dbDriver string, dbSource string) (*sql.DB, error) {

	db, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

type PostData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	fmt.Println("Hello World")
	config, err := utils.LoadConfig(".")
	helper.PanicIfError(err)

	db, err := initDB(config.DBDriver, config.DBSource)
	helper.PanicIfError(err)

	app := fiber.New(
		fiber.Config{
			IdleTimeout: time.Second * 5,
		})

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Post("/ping", func(c fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Post("/api/user", func(c fiber.Ctx) error {
		// Parse JSON body
		var postData PostData
		if err := json.Unmarshal(c.Body(), &postData); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
		}

		// Insert data into the database
		SQL := `INSERT INTO "user" (username, password) VALUES ($1, $2)`
		_, err := db.Exec(SQL, postData.Username, postData.Password)
		if err != nil {
			if pqErr, ok := err.(*pq.Error); ok {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": pqErr.Message})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Data inserted successfully"})
	})

	err = app.Listen(config.ServerAddress)
	helper.PanicIfError(err)
}
