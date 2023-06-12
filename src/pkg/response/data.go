package response

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

/*
   Process get the json file and parse the content
   Params:
       - data: name of json string
   Return:
       - map[string]interface{} with the json data to send
	- error if precess fail
*/
func Process(data string) (map[string]interface{}, error){
	if !isAFile(data){
		// Body with default response
		jsonString := `{"msg":"Hello from MOCKIT 🧉"}`
		var data map[string]interface{}
		json.Unmarshal([]byte(jsonString), &data)
		return data, nil
	}

	return dataInfo(data)
}

// Check if the requestData param make reference to a json file
func isAFile(data string) bool {
	return strings.Contains(data,".json")
}

func dataInfo(data string) (map[string]interface{}, error){
	jsonFile, err := os.Open(data)
	if err != nil {
		return nil, err
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	return result, nil
}
