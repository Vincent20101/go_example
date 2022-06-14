package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//{
//"test1" : {
//"aa" : 0,
//"bb" : 0,
//"cc" :  [ 1, 0 , 0, 0, 0, 0],
//"dd" : 1
//},
//"test2" : {
//"ee" : [ 1, 0 ],
//"tt" : 288,
//"gg" : 1,
//"nn" : 0
//}
//}

func main() {
	left, _ := readFileOfJson("E:\\jsonDataold.json")
	right, _ := readFileOfJson("E:\\jsonData.json")
	jsonDiffDict(left, right, 0)
	fmt.Println("left right change place")
	jsonDiffDict(right, left, 1)
	return
}

func jsonDiffDict(json1, json2 map[string]interface{}, sign int) {
	for k, value := range json1 {
		if k == "test1" || k == "test2" {
			json2_value := json2[k].(map[string]interface{})
			for k1, value1 := range value.(map[string]interface{}) {
				if _, ok := json2_value[k1]; ok {
					switch value1.(type) {
					case float64:
						if value1 == json2_value[k1] {
							continue
						} else {
							if sign == 0 {
								fmt.Printf("exist diff %16s %30s %9.0f %9.0f\n", k, k1, value1, json2_value[k1])
							}
						}
					case interface{}:
						value11 := value1.([]interface{})
						value22 := json2_value[k1].([]interface{})
						length := len(value11)
						if length > 4 { //只对比切片的前4个
							length = 4
						}
						for i := 0; i < length; i++ {
							if value11[i] == value22[i] {
								continue
							} else {
								if sign == 0 {
									fmt.Printf("exist diff %16s %30s %2d %9.0f %9.0f\n", k, k1, i, value11[i], value22[i])
								}
							}
						}
					default:
						fmt.Println("exist char type not sure ", k, k1)
					}
				} else {
					if sign == 0 {
						fmt.Println("right not exist", k, k1)
					} else {
						fmt.Println("left not exist", k, k1)
					}
				}
			}
		}
	}
}
func readFileOfJson(filepath string) (map[string]interface{}, error) {
	mapInfo := make(map[string]interface{})
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("ReadFile: err", err.Error())
		return nil, err
	}
	if err := json.Unmarshal(bytes, &mapInfo); err != nil {
		fmt.Println("Unmarshal: err", err.Error())
		return nil, err
	}
	return mapInfo, nil
}
