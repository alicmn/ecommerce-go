package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type Book struct {
	name string
	isbn string
}

func main() {

	book := Book{
		name: "Valley of the Wolfs",
		isbn: "1564165bf651651651",
	}
	var p *Book = &book
	printBookInfo(p)

	app := fiber.New()
	hello := "<h1>Hello</h1>"

	app.Get("/hello", func(ctx *fiber.Ctx) error {
		return ctx.SendString(hello)
	})

	app.Listen(":3000")

}
func printBookInfo(book *Book) {
	fmt.Println(book.name, " ", book.isbn)
}
