/// [date] 2022-04-28
/// [ref] https://stackoverflow.com/questions/49448302/convert-interface-to-struct

package main

/// @param raw: unstructured graphql response
/// @return: array of nodes
func GetNodes(raw interface{}) []interface{} {
	res, _ := raw.(map[string]interface{})["projects"].(map[string]interface{})["nodes"].([]interface{})
	return res
}

/// @param node: unstructured node element of graphql response
/// @return: name field of node
func GetName(node interface{}) string {
	res, ok := node.(map[string]interface{})["name"]
	if !ok {
		return "_"
	}
	return res.(string)
}

/// @param node: unstructured node element of graphql response
/// @return: forksCount field of node
func GetForksCount(node interface{}) float64 {
	res, ok := node.(map[string]interface{})["forksCount"]
	if !ok {
		return 0
	}
	return res.(float64)
}
