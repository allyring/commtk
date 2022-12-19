package cfg

import (
	"encoding/json"
	"errors"
	jsonschema "github.com/xeipuuv/gojsonschema"
	"strconv"
	"time"
)

var Config = &Cfg{}


type LayoutNode struct {
	childSizePercentages []int
	verticalStacked bool

	children []LayoutNode

	tool string
	address string
}

type JsonLayoutImport struct {
	Leaf bool `json:"Leaf"`
	Tool string `json:"ToolName"`

	ChildOrderedSizePercentages []int `json:"ChildOrderedSizePercentages"`
	VerticalStacked bool `json:"VerticalStacked"`

	ChildrenParsed []JsonLayoutImport `json:"Children"`
}


type Cfg struct {
	LayoutRoot LayoutNode
}

const (
	defaultConfig string = "{'ChildOrderedSizePercentages': [33,67],'Children': [{'Leaf': true,'ToolName': 'T1'},{'Leaf': false,'VerticalStacked': true,'Children': [{'Children': [{'leaf': true,'toolName': 'T3'},{'leaf': true, 'toolName': 'T4'}]},{'leaf': true, 'toolName': 'T2'}]}]}"

)

func ParseJSON(currentJson string) (JsonLayoutImport, error) {
	parsed := JsonLayoutImport{}
	err := json.Unmarshal([]byte(currentJson), &parsed)

	return parsed, err
}

func ParseToLayoutTree(parsedJSON JsonLayoutImport) (LayoutNode, error) {
	currentNode := LayoutNode{}

	currentNode.childSizePercentages = parsedJSON.ChildOrderedSizePercentages
	currentNode.verticalStacked = parsedJSON.VerticalStacked

	if (parsedJSON.Leaf) && (len(parsedJSON.ChildrenParsed) > 0) {
		// Leaf has children, so return an error
		return currentNode, errors.New("cannot have a tool with children")
	}


	if parsedJSON.Leaf {
		currentNode.tool = parsedJSON.Tool
		// Generate address suffix for currentNode based on unix time nanoseconds
		currentNode.address = parsedJSON.Tool + strconv.FormatInt(time.Now().UnixNano(), 10)
		return currentNode, nil
	}

	// Dealing with child nodes
	for _, child := range parsedJSON.ChildrenParsed {
		node, err := ParseToLayoutTree(child)
		if err != nil {
			return currentNode, err
		}
		currentNode.children = append(currentNode.children, node)
	}

	return currentNode, nil
}

func LoadLayoutFile(fileData string) (error, []jsonschema.ResultError) {
	// Load JSON, then verify it against the schema

	var documentLoader jsonschema.JSONLoader

	if fileData == "" {
		// No config file passed in, so use default layout config
		documentLoader = jsonschema.NewStringLoader(defaultConfig)
	} else {
		documentLoader = jsonschema.NewStringLoader(fileData)
	}

	// Load JSON schema from file
	schemaLoader := jsonschema.NewReferenceLoader("file:///mnt/c/Users/Ally%20R/Desktop/Programming/Go/commtk/schema.json")

	// Validate document
	result, err := jsonschema.Validate(documentLoader, schemaLoader)

	// Return any json validation errors
	// TODO: Update schema to allow single tool layouts
	if err != nil {
		return err, nil
	}

	if result.Errors() != nil {
		return nil, result.Errors()
	}

	// Parse it recursively into a layout tree
	var jsonData []byte
	if fileData != "" {
		jsonData = []byte(fileData)
	} else {
		jsonData = []byte(defaultConfig)
	}


	parsedJson, err := ParseJSON(string(jsonData))

	// Debug for JSON parsing
	// s, _ := json.MarshalIndent(final, "", "\t")
	// fmt.Println(string(s))

	// Perform some extra parsing on the parsed JSON

	final, err := ParseToLayoutTree(parsedJson)

	Config.LayoutRoot = final

	// Save configNode tree into variable Config
	return nil, nil
}
