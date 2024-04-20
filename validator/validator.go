package validator

import (
	"errors"
	"stringinator-go/model"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

// This method validates the request query param passed to stringinate GET method URI.
func ValidateQueryParam(c echo.Context) (string, error) {
	request_input := c.QueryParam("input")

	if request_input == "" {
		return "", errors.New("invalid query param \"input\" and cannot process further:" + request_input)
	}
	return request_input, nil

}

// This method validates the request body passed to stringinate POST method.
func ValidateRequestBody(c echo.Context) (*model.StringData, error) {
	requestData := new(model.StringData)
	if err := c.Bind(requestData); err != nil {
		return nil, err
	}

	// Create a new validator instance
	validate := validator.New()
	err := validate.Struct(requestData)
	if err != nil {
		return nil, err
	}
	return requestData, nil
}
