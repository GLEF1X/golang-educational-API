package routes

import (
	"fasthttp_restful/serializers"
	"fasthttp_restful/service"
	"log"
	"net/http"
)

type CustomerHandlers struct {
	customerService *service.CustomerService
	serializer      serializers.Serializer
}

func NewCustomerHandler(customerService *service.CustomerService, serializer serializers.Serializer) *CustomerHandlers {
	return &CustomerHandlers{customerService: customerService, serializer: serializer}
}

func (h *CustomerHandlers) GetCustomers(writer http.ResponseWriter, request *http.Request) {
	customers := h.customerService.GetAllCustomers()
	serialized, err := h.serializer.Serialize(customers)
	if err != nil {
		log.Fatal(err)
	}
	writer.Write(serialized)
}
