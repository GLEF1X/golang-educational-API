package api

import (
	"fasthttp_restful/api/routes"
	"fasthttp_restful/configuration"
	"fasthttp_restful/domain"
	"fasthttp_restful/serializers"
	"fasthttp_restful/service"
	"github.com/gorilla/mux"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"log"
)

type App struct {
	router *mux.Router
	cfg    *configuration.APIConfiguration
}

func (app *App) Init() {
	router := mux.NewRouter()
	app.setupRoutes(router)
	app.cfg = configuration.GetAPIConfiguration()
}

func (app *App) getCustomerHandlers() *routes.CustomerHandlers {
	bunchOfCustomers := []domain.Customer{domain.Customer{Id: 1, Name: "Glib", City: "Konotop", Zipcode: "fake"}}
	customerRepository := domain.NewCustomerRepositoryStub(bunchOfCustomers)
	customerService := service.NewCustomerService(customerRepository)
	jsonSerializer := serializers.JsonSerializer{}
	return routes.NewCustomerHandler(customerService, jsonSerializer)
}

func (app *App) setupRoutes(router *mux.Router) {
	customerHandlers := app.getCustomerHandlers()
	router.HandleFunc("/getCustomers", customerHandlers.GetCustomers).Methods(fasthttp.MethodGet)
	app.router = router
}

func (app *App) Run() {
	if err := fasthttp.ListenAndServe(app.cfg.GetHTTPAddr(), fasthttpadaptor.NewFastHTTPHandler(app.router)); err != nil {
		log.Fatal(err.Error())
	}
}
