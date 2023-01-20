package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type M map[string]interface{}

type User struct {
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
}

type UserValidator struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age" validate:"gte=0,lte=80"`
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

var ActionIndex = func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("from action index"))
}

var ActionHome = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("from action home"))
	})

var ActionAbout = echo.WrapHandler(
	http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("from action about"))
		},
	))

func main() {
	r := echo.New()

	r.Validator = &CustomValidator{validator: validator.New()}

	r.GET("/index", func(ctx echo.Context) error {
		data := "Hello from index"
		return ctx.String(http.StatusOK, data)
	})

	r.GET("/html", func(ctx echo.Context) error {
		data := "<html><head>Test Golang</head><body><h1>My First Heading</h1><p>My first paragraph.</p>		</body>		</html>"
		return ctx.HTML(http.StatusOK, data)
	})

	r.GET("/page1", func(ctx echo.Context) error {
		name := ctx.QueryParam("name")
		data := fmt.Sprintf("Hello  %s", name)
		return ctx.HTML(http.StatusOK, data)
	})

	r.GET("/page2/:name", func(ctx echo.Context) error {
		name := ctx.Param("name")
		data := fmt.Sprintf("Hello from route %s", name)
		return ctx.HTML(http.StatusOK, data)
	})

	r.GET("/page3/:name/*", func(ctx echo.Context) error {
		name := ctx.Param("name")
		message := ctx.Param("*")

		data := fmt.Sprintf("Hello from route %s, i have message for you: %s", name, message)
		return ctx.HTML(http.StatusOK, data)
	})

	r.POST("/page4", func(ctx echo.Context) error {
		name := ctx.FormValue("name")
		message := ctx.FormValue("message")

		data := fmt.Sprintf("Hello from route %s, i have message for you: %s",
			name,
			strings.Replace(message, "/", "", 1))

		return ctx.HTML(http.StatusOK, data)
	})

	r.POST("/validatoruser", func(ctx echo.Context) error {
		u := new(UserValidator)
		if err := ctx.Bind(u); err != nil {
			return err
		}

		if err := ctx.Validate(u); err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, true)
	})

	r.Any("/user", func(c echo.Context) (err error) {
		u := new(User)
		if err = c.Bind(u); err != nil {
			return
		}

		return c.JSON(http.StatusOK, u)
	})

	r.GET("/echoindex", echo.WrapHandler(http.HandlerFunc(ActionIndex)))
	r.GET("/echohome", echo.WrapHandler(ActionHome))
	r.GET("/echoabout", ActionAbout)

	r.HTTPErrorHandler = func(err error, c echo.Context) {
		report, ok := err.(*echo.HTTPError)

		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		if castedObject, ok := err.(validator.ValidationErrors); ok {
			for _, err := range castedObject {
				switch err.Tag() {
				case "required":
					report.Message = fmt.Sprintf("%s is required", err.Field())
				case "email":
					report.Message = fmt.Sprintf("%s is not valid email", err.Field())
				case "gte":
					report.Message = fmt.Sprintf("%s value must be greater than %s", err.Field(), err.Param())
				case "lte":
					report.Message = fmt.Sprintf("%s value must be greater than %s", err.Field(), err.Param())
				}
				break
			}
		}

		errPage := fmt.Sprintf("assets/%d.html", report.Code)
		if err := c.File(errPage); err != nil {
			c.HTML(report.Code, "Errrroooooorrrrrrrrrrr ")
		}

		c.Logger().Error(report)
		c.JSON(report.Code, report)
	}

	r.Static("/static", "assets")

	fmt.Println("Server started at :9000")
	r.Start(":9000")
}
