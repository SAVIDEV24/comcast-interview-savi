package service

import (
	"net/http"
	"stringinator-go/model"
	"stringinator-go/validator"

	"github.com/labstack/echo/v4"
)

type StringinatorService struct {
	Seen_strings map[string]int
}

var newStringinatorService = func(Seen_strings map[string]int) StringinatorService {
	return StringinatorService{
		Seen_strings: Seen_strings,
	}

}

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
			return c.JSON(http.StatusBadRequest, nil)
		}
		request_input = request_data.Input

	}

	s.remember(request_input)
	result, err := findMostOccurredChar(request_input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, result)
}

func (S *StringinatorService) Stats(c echo.Context) (err error) {
	max := 0
	longest := 0
	longestStr := ""
	statistics := model.Statistics{}
	for k, v := range S.Seen_strings {
		if v > max {
			max = v
		}

		if len(k) > longest {
			longest = len(k)
			longestStr = k
		}
	}

	for k, v := range S.Seen_strings {
		if v == max {
			statistics.Mostoccurred = append(statistics.Mostoccurred, k)
		}
	}

	statistics.LongestInput = longestStr
	return c.JSON(http.StatusOK, statistics)
}

func (s *StringinatorService) remember(input string) {
	if s.Seen_strings[input] == 0 {
		s.Seen_strings[input] = 1
	} else {
		s.Seen_strings[input] += 1
	}
}

func findMostOccurredChar(str string) ([]model.CharCount, error) {
	occurmap := make(map[rune]int, len(str))
	var result []model.CharCount
	max := 0

	for _, r := range str {
		if ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z') {
			occurmap[r]++
		}
	}

	for _, v := range occurmap {
		if v > max {
			max = v
		}
	}

	for k, v := range occurmap {
		if v == max {
			result = append(result, model.CharCount{Char: string(k), Occurance: v})
		}
	}

	return result, nil

}
