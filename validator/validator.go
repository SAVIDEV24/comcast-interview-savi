package validator

import (
	"errors"
	"stringinator-go/model"

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
	request_data := new(model.StringData)
	if err := c.Bind(request_data); err != nil {
		return nil, err
	}

	// Validate the request_data struct
	err := c.Validate(request_data)
	if err != nil {
		return nil, err
	}
	return request_data, nil
}
