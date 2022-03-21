package main

import "github.com/hugoalmeidahh/go-restful/api"

func main() {
	api.Run()
}

// app := fiber.New()

// app.Prefork = true // enable prefork

// app.Get("/", func(c *fiber.Ctx) {
// 	c.Send(fmt.Sprintf("Hi, I'm worker #%v", os.Getpid()))
// 	// => Hi, I'm worker #16858
// 	// => Hi, I'm worker #16877
// 	// => Hi, I'm worker #16895
// })

// app.Listen(8080)
