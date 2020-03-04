package common

import (
	"log"
	"encoding/json"
)

//自定义一个JSON的编解码器

type JSONCoder struct {}

func (j *JSONCoder) Marshal(v interface{}) ([]byte, error) {
	log.Println("JSONCoder Marshal")
	return json.Marshal(v)
}

func (j *JSONCoder) Unmarshal(data []byte, v interface{}) error{
	log.Println("JSONCoder UnMarshal")
	return json.Unmarshal(data, v)
}

func (j *JSONCoder) String() string {
	log.Println("JSONCoder String")
	return "JSONCoder"
}
