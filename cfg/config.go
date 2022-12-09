package cfg

import (
	"encoding/json"
	jsonschema "github.com/xeipuuv/gojsonschema"
	"os"
)
type GridCell struct {
	Tool	string	`json:"tool"`
	RatioX	int		`json:"ratioX"`
	RatioY	int		`json:"ratioY"`
}

var Config = &Cfg{}

const defaultLayoutConfig = "[  [ {\"ratioX\": 1, \"ratioY\":1, \"tool\": \"ifconfig\" },  {\"ratioX\": 1, \"ratioY\": 1, \"tool\": \"ping\" }, {\"ratioX\": 1, \"ratioY\": 1, \"tool\": \"traceroute\" }]]"
const layoutJsonSchema = "{\"type\": \"array\", \"minItems\": 1, \"items\":    {\"type\": \"array\", \"items\":    {\"type\": \"object\", \"properties\": {\"ratioX\": {\"type\": \"integer\", \"minimum\": 1}, \"ratioY\": {\"type\": \"integer\", \"minimum\": 1}, \"tool\": {\"type\": \"string\", \"enum\": [\"ifconfig\",\"ping\",\"traceroute\"]}}}, \"additionalItems\": false}}"

type Cfg struct {
	Layout	[][]GridCell
}


func LoadLayoutFile(filepath string) (error, []jsonschema.ResultError) {
	var documentLoader jsonschema.JSONLoader

	if filepath == "" {
		documentLoader = jsonschema.NewStringLoader(defaultLayoutConfig)
	} else {
		documentLoader = jsonschema.NewReferenceLoader(filepath)
	}



	schemaLoader := jsonschema.NewStringLoader(layoutJsonSchema)

	result, err := jsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		// Error while reading schema
		return err, nil
	}


	if result.Valid() {
		// Parse data into variable here
		var fileData []byte

		if filepath != "" {
			fdCache, err := os.ReadFile(filepath)

			if err != nil {
				return err, nil
			}

			fileData = fdCache
		} else {
			fileData = []byte(defaultLayoutConfig)
		}


		var cells [][]GridCell
		if err := json.Unmarshal(fileData, &cells); err != nil {
			return err, nil
		}

		Config.Layout = cells


		return nil, nil
	} else {
		// Error in config file
		return err, result.Errors()
		}
}
