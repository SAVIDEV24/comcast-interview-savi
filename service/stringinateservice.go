package service

import (
	"net/http"
	"stringinator-go/datastore"
	"stringinator-go/model"
	"stringinator-go/validator"

	"github.com/labstack/echo/v4"
)

type StringinatorService struct {
	SeenStrings map[string]int
	Ims         datastore.InMemoryStore
}

var NewStringinatorService = func(SeenStrings map[string]int, ims datastore.InMemoryStore) *StringinatorService {
	return &StringinatorService{
		SeenStrings: SeenStrings,
		Ims:         ims,
	}

}

// This method accepts string input and return back the most occurred character from the given input and its number of occurances.
func (s *StringinatorService) Stringinate(c echo.Context) (err error) {

	var request_input string
	if c.Request().Method == "GET" {
		query, err := validator.ValidateQueryParam(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		request_input = query
	} else {
		request_data, err := validator.ValidateRequestBody(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		request_input = request_data.Input

	}

	// Add input to temporary data store
	s.remember(request_input)

	// Add input to persistent data store
	s.Ims.AddInput(request_input)
	result := findMostOccurredChar(request_input)
	return c.JSON(http.StatusOK, result)
}

// This method return statistics from temporary ims about all strings the server has seen,
// including the number of times each input has been received along with the longest and most popular strings etc.
func (S *StringinatorService) Stats(c echo.Context) (err error) {
	max := 0
	longest := 0
	longestStr := ""
	statistics := model.Statistics{}
	for k, v := range S.SeenStrings {
		if v > max {
			max = v
		}

		if len(k) > longest {
			longest = len(k)
			longestStr = k
		}
	}

	for k, v := range S.SeenStrings {
		if v == max {
			statistics.Mostoccurred = append(statistics.Mostoccurred, k)
		}
	}

	statistics.LongestInput = longestStr
	return c.JSON(http.StatusOK, statistics)
}

// This method return statistics from temporary ims about all strings the server has seen,
// including the number of times each input has been received along with the longest and most popular strings etc.
func (S *StringinatorService) StatsFromIms(c echo.Context) (err error) {
	SeenStringMap := S.Ims.GetSeenStrings()
	max := 0
	longest := 0
	longestStr := ""
	statistics := model.Statistics{}

	//Find the maximum number of occurance and the longest input
	for k, v := range SeenStringMap {
		if v > max {
			max = v
		}

		if len(k) > longest {
			longest = len(k)
			longestStr = k
		}
	}

	for k, v := range SeenStringMap {
		if v == max {
			statistics.Mostoccurred = append(statistics.Mostoccurred, k)
		}
	}

	statistics.LongestInput = longestStr
	return c.JSON(http.StatusOK, statistics)
}

// This method used to store inputs to temporary in memory data store.
func (s *StringinatorService) remember(input string) {
	if s.SeenStrings[input] == 0 {
		s.SeenStrings[input] = 1
	} else {
		s.SeenStrings[input] += 1
	}
}

// This method returns most occurred character and its count and error
func findMostOccurredChar(str string) []model.CharCount {
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
