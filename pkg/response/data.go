package response

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Process(data string) (map[string]interface{}, error){
	if !isAFile(data){
		// Body with default response
		jsonString := `{"msg":"hello from MOCKIT"}`
		var data map[string]interface{}
		json.Unmarshal([]byte(jsonString), &data)
		return data, nil
	}

	return dataInfo(data), nil
}

// Check if the requestData param make reference to a json file
func isAFile(data string) bool {
	return strings.Contains(data,".json")
}

func dataInfo(data string) map[string]interface{}{
	jsonFile, err := os.Open(data)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	return result
}
