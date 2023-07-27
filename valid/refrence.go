
cookie := fiber.Cookie{
	Name:     "token",
	Value:    tokenString,
	Expires:  expirationTime,
	HTTPOnly: true,
}

package main

import (
	"fmt"
	"net/http"
)

//func main() {
	// Define the endpoint URL
	url := "https://api.example.com/authenticate"

	// Create an HTTP client
	client := &http.Client{}

	// Create a request object
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the authentication header
	req.Header.Set("Authorization", "Bearer <access_token>")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Process the response
	// ...
}

//package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Replace these with your actual credentials
	username := "your_username"
	password := "your_password"

	// Encode the credentials in Base64
	credentials := username + ":" + password
	encodedCredentials := base64.StdEncoding.EncodeToString([]byte(credentials))

	// Set up the request URL
	url := "https://api.example.com/authenticate" // Replace this with the authentication endpoint URL

	// Prepare the request body (if needed)
	// requestBody := []byte("your_request_body_here")

	// Create the HTTP client and request
	client := &http.Client{}
	// req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	req, err := http.NewRequest("GET", url, nil) // Use GET if there's no request body, otherwise use POST or other appropriate methods

	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the authentication header
	req.Header.Set("Authorization", "Basic "+encodedCredentials)

	// Make the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Check the status code to ensure the request was successful
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Request failed with status code:", resp.StatusCode)
		fmt.Println("Response body:", string(body))
		return
	}

	// The token should be in the response body (if applicable)
	fmt.Println("Token:", string(body))
}

if err != nil {
	fmt.Println("error");
}
req.Header.Set("")

app := fiber.new()

auth := basicauth.New

type Config struct {
	Next func(c *fiber.Ctx) bool

	Users map[string]string
	Realm string

}

var ConfigDefault = Config{
    Next:            nil,
    Users:           map[string]string{},
    Realm:           "Restricted",
    Authorizer:      nil,
    Unauthorized:    nil,
    ContextUsername: "username",
    ContextPassword: "password",
}


type AuthLoginOut struct {
	Username string `json:"username" validate:"required,min=3,max=12"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Token1   string `json:"token"` 


}

request.Header.Add

type Config struct {
	Next func(c *fiber.Ctx) bool
	Users map[string]string
	Username 
	Password
	Email
	Token1
	Realm string 
	Authorizer func (string, string)bool{ 
	Username string `json:"username" validate:"required,min=3,max=12"`
	Password string `json:"password"`
	Email    string `json:"email"`
	}
	Unauthorized fiber.Handler 
	ContextUsername
	ContextPassword
	ContextEmail
}

type config struct {
	Realm  string 
	Unauthorized fiber.Handler
}

func TokenAuth(config Config) fiber.Handler {
	return func(c *fiber.Ctx)error {
		auth.header := C.Get("Authorization")
		if auth.Header  = " " || !strings.HasPrefix(auth.Header, "Bearer") {
			return config.Unauthorized(c)
		}
		Token := strings
	}
}


package main

import (
	"github.com/gofiber/fiber/v2"
	"strings"
)

type Config struct {
	Realm        string
	Unauthorized fiber.Handler
}


yo 



import (
	"github.com/gofiber/fiber/v2"
	"strings"
)

type Config struct {
	Realm        string
	Unauthorized fiber.Handler
}

func TokenAuth(config Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return config.Unauthorized(c)
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		c.Locals("token", token)

		return c.Next()
	}
}

func main() {
	app := fiber.New()

	tokenAuth := TokenAuth(Config{
		Realm:        "Restricted",
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
	})

	app.Use(tokenAuth)

	
	app.Get("/protected", func(c *fiber.Ctx) error {
	
		token := c.Locals("token").(string)

		return c.JSON(fiber.Map{
			"message": "Protected route",
			"token":   token,
		})
	})
	app.Listen(":3000")
}


func userLogin(response http.responseWriter, request *http.Request){
	response.Header().Set("content type", "application/json")
	json.newDecoder(request.Body).Decode(&user)
	err := collection.findOne egaaga       
}