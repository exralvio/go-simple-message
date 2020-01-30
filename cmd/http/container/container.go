package container

import "github.com/labstack/echo"

type AppContainer struct {
	HTTPService *echo.Echo
}

func NewAppContainer() *AppContainer {
	app := &AppContainer{
		HTTPService: echo.New(),
	}

	return app
}

func (app *AppContainer) Run() {
	app.HTTPService.Logger.Error(
		app.HTTPService.Start(":3000"),
	)
}
