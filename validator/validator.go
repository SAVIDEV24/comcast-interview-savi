package validator

import (
	"errors"
	"fmt"
	"stringinator-go/model"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func ValidateQueryParam(c echo.Context) (string, error) {
	request_input := c.QueryParam("input")

	if request_input == "" {
		return "", errors.New("invalid query param and cannot process further")
	}
	return request_input, nil

}

func ValidateRequestBody(c echo.Context) (*model.StringData, error) {

	// Create a new validator instance
	validate := validator.New()
	request_data := new(model.StringData)
	if err := c.Bind(request_data); err != nil {
		return nil, err
	}

	// Validate the request_data struct
	err := validate.Struct(request_data)
	if err != nil {
		// Handle validation errors
		fmt.Println("Validation errors:")
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf("- %s is required\n", err.Field())
		}
		return nil, err
	}
	return request_data, nil
}
