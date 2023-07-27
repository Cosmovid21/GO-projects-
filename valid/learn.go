type static struct {
	ByteRange bool `json :"byte_range"`
	browse bool `json:"browse"`
	download bool `json:"downlaod"`
}

app.Get("/api/list", func(c *fiber Ctx))error {
	return c.SendString("get request")
}

app.Get("/api/login", func(c *fiber.Ctx))error {
	return c.SendString("POST REQUEST")
}

app.Use(func(c *fiber Ctx))error {
	return c.Next
}

app.Use(func(c *fiber Ctx))error {
	return c.Next
}

app.Use(func(c *fiber Ctx))error {
	return c.Next
}



app.Use(requestid.New())

app.Use(requestid.New(requestid.Config{
	Header: "X-Custom-Header",
	Generator: func() string{
		return "static-id"
	},
}))

type config struct {
	Next func (c *fiber Ctx) bool
	Header string
	Generator func() string
	contextKey interface{}
}


func main() {
	app := fiber.New()

	app.Get("/login", func(c *fiber.Ctx)error {
		return c.Sendstring("empty")
	}
	)
	app.Post("/login", func (c *fiber.Ctx) error{
		username := c.FormValue("username")
		password := c.FormValue("password")
		email := c.FormValue("email")
	}
)





func main() {
	app := fiber.New()

	app.Get("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var user *User
		for i := range users {
			if users[i].ID == id {
				user = &users[i]
				break
			}
		}

		if user == nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
		}

		return c.JSON(user)
	})

	app.Put("/users/:id", func(c *fiber.Ctx) error {
		
		id := c.Params("id")

		var user *User
		for i := range users {
			if users[i].ID == id {
				user = &users[i]
				break
			}
		}

		if user == nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
		}

		var updatedUser User
		if err := c.BodyParser(&updatedUser); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
		}

		user.Username = updatedUser.Username
		user.Email = updatedUser.Email

		return c.JSON(user)
	})

	app.Listen(":3000")
}

//

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

var users = []User{}

func main() {
	app := fiber.New()

	app.Post("/register", func(c *fiber.Ctx) error {
		var user User
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
		}
		for _, u := range users {
			if u.Username == user.Username {
				return c.Status(fiber.StatusConflict).JSON(fiber.Map{"message": "Username already taken"})
			}
		}

		user.ID = len(users) + 1

		users = append(users, user)
		return c.JSON(fiber.Map{"message": "Registration successful", "user": user})
	})

	app.Get("/users", func(c *fiber.Ctx) error {
		return c.JSON(users)
	})

	app.Listen(":3000")
}

