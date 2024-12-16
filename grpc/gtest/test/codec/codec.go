package codec

import (
	"encoding/json"
)

type JSONCoder struct{}

func (j *JSONCoder) String() string {
	//log.Println("JSONCoder String")
	return "JSONCoder"
}

func (j *JSONCoder) Marshal(v interface{}) ([]byte, error) {
	//log.Println("JSONCoder Marshal")
	return json.Marshal(v)
}

func (j *JSONCoder) Unmarshal(data []byte, v interface{}) error {
	//log.Println("JSONCoder UnMarshal")
	return json.Unmarshal(data, v)
}

func (j *JSONCoder) Name() string {
	//log.Println("JSONCoder String")
	return "JSONCoder"
}
