package main

func GetNodes(raw interface{}) []interface{} {
	res, _ := raw.(map[string]interface{})["projects"].(map[string]interface{})["nodes"].([]interface{})
	return res
}

func GetName(node interface{}) string {
	res, ok := node.(map[string]interface{})["name"]
	if !ok {
		return "_"
	}
	return res.(string)
}

func GetForksCount(node interface{}) float64 {
	res, ok := node.(map[string]interface{})["forksCount"]
	if !ok {
		return 0
	}
	return res.(float64)
}
