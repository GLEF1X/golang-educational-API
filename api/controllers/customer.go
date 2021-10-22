package routes

import (
	"github.com/GLEF1X/golang-educational-API/api/apiHelper"
	"github.com/GLEF1X/golang-educational-API/dto"
	"github.com/GLEF1X/golang-educational-API/serializers"
	"github.com/GLEF1X/golang-educational-API/service"
	"github.com/valyala/fasthttp"
)

type CustomerController struct {
	customerService *service.UserService
	serializer      serializers.Serializer
}

func NewCustomerController(customerService *service.UserService, serializer serializers.Serializer) *CustomerController {
	return &CustomerController{customerService: customerService, serializer: serializer}
}

func (c *CustomerController) GetCustomers(ctx *fasthttp.RequestCtx) {
	customers := c.customerService.GetAllUsers()
	apiHelper.AnswerJson(ctx, customers, c.serializer)
}

func (c *CustomerController) AddCustomer(ctx *fasthttp.RequestCtx) {
	user := &dto.User{}
	err := c.serializer.Deserialize(ctx.Request.Body(), user)
	if err != nil {
		apiHelper.AnswerBadRequest(ctx, "Wrong type of payload", c.serializer)
		return
	}
	c.customerService.AddUser(user)
	apiHelper.Answer201(ctx, c.serializer, &dto.User{})
}
