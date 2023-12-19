package Services

import (
	"akasia/Config"
	"akasia/Routes"
	"flag"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

var AppEnv = flag.String("env", "", "define environment")

func init() {
	flag.Parse()
	if *AppEnv == "" {
		*AppEnv = Config.Localhost
	}
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func AppInitialization() {
	//Config DB SQL
	Config.GetEnvironment(*AppEnv).InitDB()

	var validate = echo.New()
	validate.Validator = &CustomValidator{validator: validator.New()}

	//Collect Routes
	var routes Routes.Routes
	routes.CollectRoutes(validate)
}
