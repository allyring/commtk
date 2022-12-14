package cfg

import (
	jsonschema "github.com/xeipuuv/gojsonschema"
)

var Config = &Cfg{}

type Cfg struct {
	Layout	bool
}


func LoadLayoutFile(filepath string) (error, []jsonschema.ResultError) {
	// Load JSON, then verify it against a schema
	// Parse it recursively into a node tree.
	// Return node tree

	//var documentLoader jsonschema.JSONLoader
	//
	//if filepath == "" {
	//	documentLoader = jsonschema.NewStringLoader(defaultLayoutConfig)
	//} else {
	//	documentLoader = jsonschema.NewReferenceLoader(filepath)
	//}
	//
	//
	//
	//schemaLoader := jsonschema.NewStringLoader(layoutJsonSchema)
	//
	//result, err := jsonschema.Validate(schemaLoader, documentLoader)
	//if err != nil {
	//	// Error while reading schema
	//	return err, nil
	//}
	//
	//
	//if result.Valid() {
	//	// Parse data into variable here
	//	var fileData []byte
	//
	//	if filepath != "" {
	//		fdCache, err := os.ReadFile(filepath)
	//
	//		if err != nil {
	//			return err, nil
	//		}
	//
	//		fileData = fdCache
	//	} else {
	//		fileData = []byte(defaultLayoutConfig)
	//	}
	//
	//
	//	var cells [][]GridCell
	//	if err := json.Unmarshal(fileData, &cells); err != nil {
	//		return err, nil
	//	}
	//
	//	Config.Layout = cells
	//
	//
	//	return nil, nil
	//} else {
	//	// Error in config file
	//	return err, result.Errors()
	//	}


	return nil, nil
}
