package service

import (
	"fmt"
	"net/http"
	"stringinator-go/interfaces"
	"stringinator-go/model"
	"stringinator-go/validator"

	"github.com/labstack/echo/v4"
)

type StringinatorService struct {
	store interfaces.Store
}

var NewStringinatorService = func(store interfaces.Store) *StringinatorService {
	return &StringinatorService{
		store: store,
	}

}

// This method accepts string input and return back the most occurred character from the given input and its number of occurances.
func (s *StringinatorService) Stringinate(c echo.Context) (err error) {
	var requestInput string
	if c.Request().Method == "GET" {
		queryparam, err := validator.ValidateQueryParam(c)
		if err != nil {
			c.Logger().Error("invalid Query param", err)
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		requestInput = queryparam
	} else {
		requestData, err := validator.ValidateRequestBody(c)
		if err != nil {
			c.Logger().Error("invalid Request body: ", err)
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		requestInput = requestData.Input
	}

	// Add input to data store
	err = s.store.SaveStrings(requestInput)
	if err != nil {
		c.Logger().Error("error while saving input", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	result := findMostOccurredChar(requestInput)
	return c.JSON(http.StatusOK, result)
}

// This method return statistics from temporary/persistent ims about all strings the server has seen,
// including the number of times each input has been received along with the longest and most popular strings etc.
func (S *StringinatorService) Stats(c echo.Context) (err error) {

	seenStrings, err := S.store.GetStrings()

	if err != nil {
		c.Logger().Error("error while retrieving stored inputs", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	max := 0
	longest := 0
	longestStr := ""
	statistics := model.Statistics{}

	for k, v := range seenStrings {
		if v > max {
			max = v
		}

		if len(k) > longest {
			longest = len(k)
			longestStr = k
		}
	}

	for k, v := range seenStrings {
		if v == max {
			statistics.Mostoccurred = append(statistics.Mostoccurred, k)
		}
	}

	statistics.LongestInput = longestStr
	return c.JSON(http.StatusOK, statistics)
}

// This method returns most occurred character and its count and error
func findMostOccurredChar(str string) []model.CharCount {
	fmt.Println("inside it")
	occurmap := make(map[rune]int)
	var result []model.CharCount
	max := 0

	// Count the character occurance
	for _, r := range str {
		if ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z') {
			occurmap[r]++
		}
	}

	//Find the maximum number of occurance.
	for _, v := range occurmap {
		if v > max {
			max = v
		}
	}

	//Append most frequented character and its length to the result slice
	for k, v := range occurmap {
		if v == max {
			result = append(result, model.CharCount{Char: string(k), Occurance: v})
		}
	}

	return result

}
