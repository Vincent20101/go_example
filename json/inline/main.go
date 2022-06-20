package main

import (
	"encoding/json"
	"fmt"
)

type Project struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type JiraHttpReqField struct {
	Project     `json:",inline"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
}

func main() {
	dataProject := Project{
		Key:   "key",
		Value: "value",
	}
	dataJiraHttpReqField := &JiraHttpReqField{
		Project:     dataProject,
		Summary:     "Summary",
		Description: "Description",
	}
	data, _ := json.Marshal(dataJiraHttpReqField)
	fmt.Println(string(data))

}
