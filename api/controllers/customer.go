package routes

import (
	"github.com/GLEF1X/golang-educational-API/api/apiHelper"
	"github.com/GLEF1X/golang-educational-API/serializers"
	"github.com/GLEF1X/golang-educational-API/service"
	"net/http"
)

type CustomerController struct {
	customerService *service.CustomerService
	serializer      serializers.Serializer
}

func NewCustomerController(customerService *service.CustomerService, serializer serializers.Serializer) *CustomerController {
	return &CustomerController{customerService: customerService, serializer: serializer}
}

func (c *CustomerController) GetCustomers(writer http.ResponseWriter, request *http.Request) {
	customers := c.customerService.GetAllCustomers()
	if apiHelper.IsHeadersInvalid(request) {
		apiHelper.AnswerBadRequest(writer, "Invalid content type header", c.serializer)
		return
	}
	apiHelper.AnswerJson(writer, customers, c.serializer)
}

func (c *CustomerController) AddCustomer(writer http.ResponseWriter, request *http.Request) {
	apiHelper.IsHeadersInvalid(request)
	// err := c.customerService.AddCustomer()
	//if err != nil{
	//	apiHelper.AnswerBadRequest(writer, err.Error(), c.serializer)
	//	return
	//}
	apiHelper.Answer201(writer, c.serializer)
}
