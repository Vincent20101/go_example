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
	Project
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
	fmt.Println(dataJiraHttpReqField.Value)
	data, _ := json.Marshal(dataJiraHttpReqField)
	fmt.Println(string(data))

}
