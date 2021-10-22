package api

import (
	"github.com/GLEF1X/golang-educational-API/core"
	"github.com/GLEF1X/golang-educational-API/core/middlewares"
)

const (
	v1APIPathPrefix = "/api/v1"
)

func SetupRoutes(s *core.Server) {
	customerController := s.InitializeCustomerController()
	metricUtil := middlewares.NewMeasureTimeUtil()
	s.Router.GET(
		v1APIPathPrefix+"/users/getAll",
		metricUtil.MeasureResponseTimeMiddleware(middlewares.EnforceJSONMiddleware(customerController.GetCustomers)),
	)
	s.Router.PUT(
		v1APIPathPrefix+"/users/add",
		metricUtil.MeasureResponseTimeMiddleware(middlewares.EnforceJSONMiddleware(customerController.AddCustomer)),
	)
}
