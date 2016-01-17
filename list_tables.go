package dbutil

// ListTables returns a list of
func ListTables(db DBLike) ([]string, error) {
	rows, err := db.Query(".tables")
	result := []string{}

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var tableName string
		err := rows.Scan(&tableName)
		if err != nil {
			return nil, err
		}

		result = append(result, tableName)
	}

	return result, nil
}