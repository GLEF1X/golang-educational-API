package api

import (
	"github.com/GLEF1X/golang-educational-API/adapters/orm"
	"github.com/GLEF1X/golang-educational-API/adapters/repositories"
	"github.com/GLEF1X/golang-educational-API/api/controllers"
	"github.com/GLEF1X/golang-educational-API/core"
	"github.com/GLEF1X/golang-educational-API/serializers"
	"github.com/GLEF1X/golang-educational-API/service"
	"github.com/gorilla/mux"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"log"
)

type App struct {
	router *mux.Router
	cfg    *core.APIConfiguration
}

func (app *App) Init() {
	router := mux.NewRouter()
	app.cfg = core.NewAPIConfiguration()
	app.setupRoutes(router)
}

func (app *App) initializeCustomerController() *routes.CustomerController {
	db := orm.PrepareDatabase(app.cfg.GetDsn())
	customerRepository := repositories.NewCustomerRepository(db)
	customerService := service.NewCustomerService(customerRepository)
	jsonSerializer := serializers.JsonSerializer{}
	return routes.NewCustomerController(customerService, jsonSerializer)
}

func (app *App) setupRoutes(router *mux.Router) {
	customerHandlers := app.initializeCustomerController()
	router.HandleFunc("/getCustomers", customerHandlers.GetCustomers).Methods(fasthttp.MethodGet)
	app.router = router
}

func (app *App) Run() {
	if err := fasthttp.ListenAndServe(app.cfg.GetHTTPAddr(), fasthttpadaptor.NewFastHTTPHandler(app.router)); err != nil {
		log.Fatal(err.Error())
	}
}
