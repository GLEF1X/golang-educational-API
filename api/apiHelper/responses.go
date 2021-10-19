package apiHelper

import (
	"github.com/GLEF1X/golang-educational-API/serializers"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
)

type APIResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message,omitempty"`
}

const (
	JsonContentType = "application/json"
)

func answer500OnError(writer http.ResponseWriter, err error) {
	writer.WriteHeader(fasthttp.StatusInternalServerError)
	log.Println(err.Error())
}

func AnswerJson(rw http.ResponseWriter, unserializedData interface{}, serializer serializers.Serializer) {
	rw.Header().Add("Content-Type", JsonContentType)
	serializedData, sErr := serializer.Serialize(unserializedData)
	if sErr != nil {
		log.Println("Something went wrong, cannot serialize data")
		return
	}
	_, wErr := rw.Write(serializedData)
	if wErr != nil {
		log.Println("Cannot write serialized data to apiHelper")
		return
	}
}

func IsHeadersInvalid(request *http.Request) bool {
	return request.Header.Get("content-type") != JsonContentType
}

func AnswerBadRequest(writer http.ResponseWriter, message string, serializer serializers.Serializer) {
	response := &APIResponse{Ok: false, Message: message}
	serializedResponse, sErr := serializer.Serialize(response)
	if sErr != nil {
		answer500OnError(writer, sErr)
		return
	}
	_, wErr := writer.Write(serializedResponse)
	if wErr != nil {
		answer500OnError(writer, wErr)
		return
	}
}

func Answer201(writer http.ResponseWriter, serializer serializers.Serializer) {

}
