package handler

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/rs/zerolog/log"
)

type HTTPResponse struct {
	StatusCode int
	Headers    map[string]string
	Message    string
	Data       interface{}
}

func (h HTTPResponse) body() string {
	serialised, err := json.Marshal(struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{
		Message: h.Message,
		Data:    h.Data,
	})

	if err != nil {
		log.Error().
			Str("message", h.Message).
			Interface("data", h.Data).
			Err(err).
			Msg("failed to serialise HTTPResponse")

		return `{"message":"failed to serialise response","data":null}`
	}

	return string(serialised)
}

func (h HTTPResponse) Response() (events.APIGatewayV2HTTPResponse, error) {
	log.Debug().
		Int("status_code", h.StatusCode).
		Str("message", h.Message).
		Interface("data", h.Data).
		Send()

	result := events.APIGatewayV2HTTPResponse{
		StatusCode: h.StatusCode,
		Headers:    h.Headers,
		Body:       h.body(),
	}

	var err error
	if h.StatusCode >= 500 {
		err = fmt.Errorf("request failed: %s", h.Message)
	}

	return result, err
}

func (h HTTPResponse) Header(key, value string) HTTPResponse {
	h.Headers[key] = value
	return h
}

func (h HTTPResponse) Ok(message string, data interface{}) HTTPResponse {
	h.StatusCode = 200
	h.Message = message
	h.Data = data
	return h
}

func (h HTTPResponse) BadRequest(err error) HTTPResponse {
	h.StatusCode = 400
	h.Message = err.Error()
	return h
}

func (h HTTPResponse) BadRequestf(format string, a ...any) HTTPResponse {
	return h.BadRequest(fmt.Errorf(format, a...))
}

func (h HTTPResponse) NotFound() HTTPResponse {
	h.StatusCode = 404
	h.Message = "not found"
	return h
}

func (h HTTPResponse) Error(err error) HTTPResponse {
	h.StatusCode = 500
	h.Message = err.Error()
	return h
}

func (h HTTPResponse) Errorf(format string, a ...any) HTTPResponse {
	return h.Error(fmt.Errorf(format, a...))
}
