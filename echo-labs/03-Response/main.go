package main
import (	
	"net/http"	
	"github.com/labstack/echo"
)

type Person struct {
	FirstName string 
	LastName string 
}

func main()  {
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
		FirstName : "qwe",
		LastName : "asd",
	}

	ps := make([]Person, 0)
	ps = append(ps, p)
	ps = append(ps, p)
	ps = append(ps, p)
	return ps
}
