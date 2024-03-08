package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

// Transformation function for the input JSON
func transform(input map[string]interface{}) map[string]interface{} {
	output := make(map[string]interface{})

	for key, value := range input {
		// Sanitize key
		key = strings.TrimSpace(key)

		// Skip fields with empty keys
		if key == "" {
			continue
		}

		// Determine the data type and apply transformation
		switch val := value.(type) {
		case map[string]interface{}:
			// Handle map type
			output[key] = transformMap(val)
		case string:
			// Handle string type
			output[key] = transformString(val)
		case float64:
			// Handle number type
			output[key] = val
		case bool:
			// Handle boolean type
			output[key] = val
		case nil:
			// Handle null type
			continue
		}
	}

	return output
}

// Transformation function for map type
func transformMap(input map[string]interface{}) map[string]interface{} {
	output := make(map[string]interface{})

	for key, value := range input {
		// Sanitize key
		key = strings.TrimSpace(key)

		// Skip fields with empty keys
		if key == "" {
			continue
		}

		// Recursively transform nested maps
		switch val := value.(type) {
		case map[string]interface{}:
			output[key] = transformMap(val)
		case []interface{}:
			output[key] = transformList(val)
		case string:
			output[key] = transformString(val)
		case float64:
			output[key] = val
		case bool:
			output[key] = val
		case nil:
			continue
		}
	}

	return output
}

// Transformation function for list type
func transformList(input []interface{}) []interface{} {
	var output []interface{}

	for _, item := range input {
		switch val := item.(type) {
		case map[string]interface{}:
			output = append(output, transformMap(val))
		case string:
			output = append(output, transformString(val))
		case float64:
			output = append(output, val)
		case bool:
			output = append(output, val)
		case nil:
			continue
		}
	}

	return output
}

// Transformation function for string type
func transformString(input string) interface{} {
	// Attempt to parse string as time
	t, err := time.Parse(time.RFC3339, input)
	if err == nil {
		// If successful, return Unix Epoch time
		return t.Unix()
	}

	// Attempt to parse string as number
	num, err := strconv.ParseFloat(input, 64)
	if err == nil {
		// If successful, return stripped leading zeros
		return num
	}

	// Attempt to parse string as boolean
	switch input {
	case "1", "t", "T", "TRUE", "true", "True":
		return true
	case "0", "f", "F", "FALSE", "false", "False":
		return false
	default:
		// Return original string if no transformation applied
		return strings.TrimSpace(input)
	}
}

func main() {
	// Read input JSON file
	inputFile := "example.json"
	inputData, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input JSON file:", err)
		return
	}

	// Parse input JSON data
	var input map[string]interface{}
	if err := json.Unmarshal(inputData, &input); err != nil {
		fmt.Println("Error parsing input JSON:", err)
		return
	}

	// Transform input JSON
	output := transform(input)

	// Generate transformed JSON output
	transformedJSON, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling transformed JSON:", err)
		return
	}

	// Print output to stdout
	fmt.Println(string(transformedJSON))
}
