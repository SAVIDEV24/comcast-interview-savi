package model

type (
	StringData struct {
		Input  string `param:"input" query:"input" form:"input" json:"input" xml:"input" validate:"required"`
		Length int    `json:"length"`
	}

	CharCount struct {
		Char      string `json:"character"`
		Occurance int    `json:"No of occurances"`
	}

	StatsData struct {
		Inputs map[string]int `json:"inputs"`
	}

	Statistics struct {
		Mostoccurred []string `json:"most_popular"`
		LongestInput string   `json:"longest_input_received"`
	}
)
