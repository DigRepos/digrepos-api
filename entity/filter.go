package entity

import (
	"fmt"
	"strings"
)

type Star struct {
	Low  string `query:"low"`
	High string `query:"high"`
}

type Filter struct {
	Keywords []string `query:"keywords"`
	Star     Star     `query:"star"`
	Language string   `query:"language"`
	License  string   `query:"license"`
}

func (f *Filter) BuildQuery() string {
	queryArr := []string{}
	if len(f.Keywords) > 0 {
		queryArr = append(queryArr, strings.Join(f.Keywords, " "))
	}
	if queryTrim(f.Star.Low) != "" {
		queryArr = append(queryArr, "stars:>="+f.Star.Low)
	}
	if queryTrim(f.Star.High) != "" {
		queryArr = append(queryArr, "stars:<="+f.Star.High)
	}
	if queryTrim(f.Language) != "" {
		queryArr = append(queryArr, "language="+f.Language)
	}
	if queryTrim(f.License) != "" {
		queryArr = append(queryArr, "license="+f.License)
	}
	fmt.Println("[Filter] BuildQuery", strings.Join(queryArr, " "))

	return strings.Join(queryArr, " ")
}

func queryTrim(query string) string {
	return strings.Trim(query, " ")
}
