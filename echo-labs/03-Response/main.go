package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type Person struct {
	FirstName string
	LastName  string
}

type User struct {
	Name  string   `json:"name" xml:"name" form:"name" query:"name"`
	Email string   `json:"email" xml:"email" form:"email" query:"email"`
	Tags  []string `json:"tags" xml:"tags" form:"tags" query:"tags"`
}

func main() {
	e := echo.New()

	//Text/Plain
	e.GET("/plain", func(c echo.Context) error {
		return c.String(http.StatusOK, "Prueba de texto plano")
	})

	//HTML
	e.GET("/html", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<h1>Prueba de HTML</h1><script>alert('Evento HTML')</script>")
	})

	//No-content
	e.GET("/nocontent", func(c echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return err
		}

		return c.NoContent(http.StatusOK)
	})

	people := GetPerson()
	//JSON
	e.GET("/json", func(c echo.Context) error {
		return c.JSON(http.StatusOK, people)
	})

	//XML
	e.GET("/xml", func(c echo.Context) error {
		return c.XML(http.StatusOK, people)
	})

	e.Start(":8080")
}

func GetPerson() []Person {
	p := Person{
		FirstName: "qwe",
		LastName:  "asd",
	}

	ps := make([]Person, 0)
	ps = append(ps, p)
	ps = append(ps, p)
	ps = append(ps, p)
	return ps
}
