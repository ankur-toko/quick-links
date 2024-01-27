package models

import (
	"encoding/json"
	"fmt"
)

type QuickLink struct {
	Key string
	URL string
}

func (q *QuickLink) ToJSON() string {
	str, err := json.Marshal(q)
	if err != nil {
		fmt.Print("error marshalling quickLink to json", err)
	}
	return string(str)
}
