package apiHelper

import (
	"github.com/GLEF1X/golang-educational-API/core/logging"
	"github.com/GLEF1X/golang-educational-API/serializers"
	"github.com/valyala/fasthttp"
)

const (
	JsonContentType = "application/json"
)

type APIResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message,omitempty"`
}

func answer500OnError(ctx *fasthttp.RequestCtx, err error) {
	ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	logging.GetLogger().Fatal(err)
}

func AnswerJson(ctx *fasthttp.RequestCtx, unserializedData interface{}, serializer serializers.Serializer) {
	ctx.Response.Header.Add("Content-Type", JsonContentType)
	serializedData, sErr := serializer.Serialize(unserializedData)
	if sErr != nil {
		logging.GetLogger().Fatal("Something went wrong, cannot serialize data")
		return
	}
	ctx.SetBody(serializedData)
}

func AnswerBadRequest(ctx *fasthttp.RequestCtx, message string, serializer serializers.Serializer) {
	response := &APIResponse{Ok: false, Message: message}
	serializedAPIResponse, sErr := serializer.Serialize(response)
	if sErr != nil {
		answer500OnError(ctx, sErr)
		return
	}
	ctx.SetBody(serializedAPIResponse)
}

func Answer201(ctx *fasthttp.RequestCtx, serializer serializers.Serializer, serializedStruct interface{}) {
	response := &APIResponse{Ok: true}
	serializedAPIResponse, sErr := serializer.Serialize(response)
	if sErr != nil {
		answer500OnError(ctx, sErr)
		return
	}
	ctx.SetBody(serializedAPIResponse)
}
