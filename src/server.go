package main

import (
    "github.com/gofiber/fiber/v2"
)

func main() {
    fiberConf := fiber.Config {
        StrictRouting: true,
        CaseSensitive: true,
    }

    app := fiber.New(fiberConf)

    app.Listen(":3000")
}
