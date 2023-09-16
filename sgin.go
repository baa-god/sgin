package sgin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler any

type Router interface {
	Use(args ...Handler) Router
	Get(path string, handlers ...Handler) Router
	Post(path string, handlers ...Handler) Router
	Group(path string, handlers ...Handler) Router
	Add(method string, path string, handlers ...Handler) Router
}

type App struct {
	e *gin.Engine
}

func New() *App {
	app := &App{e: gin.New()}
	return app
}

func (app *App) Use(args ...Handler) Router {
	app.e.Use(HandlerFunc(args...)...)
	return app
}

func (app *App) Get(path string, handlers ...Handler) Router {
	return app.Add(http.MethodGet, path, handlers...)
}

func (app *App) Post(path string, handlers ...Handler) Router {
	return app.Add(http.MethodPost, path, handlers...)
}

func (app *App) Group(path string, handlers ...Handler) Router {
	r := app.e.Group(path, HandlerFunc(handlers...)...)
	return &Group{r: r}
}

func (app *App) Add(method string, path string, handlers ...Handler) Router {
	app.e.Handle(method, path, HandlerFunc(handlers...)...)
	return app
}
