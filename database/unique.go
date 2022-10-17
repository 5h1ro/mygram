package database

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

type UniqueRule struct {
	db       *sql.DB
	ruleName string
}

func (r *UniqueRule) Rule(field string, rule string, message string, value interface{}) error {
	var queryRow *sql.Row
	var total int

	query := `SELECT COUNT(*) as total FROM %s WHERE %s = $1`
	params := strings.Split(strings.TrimPrefix(rule, fmt.Sprintf("%s:", r.ruleName)), ",")

	if len(params) == 2 {
		query = fmt.Sprintf(query, params[0], params[1])
		queryRow = r.db.QueryRow(query, value)
	} else if len(params) == 4 {
		query += ` AND %s != $`
		query = fmt.Sprintf(query, params[0], params[1], params[2])
		queryRow = r.db.QueryRow(query, value, params[3])
	} else {
		return fmt.Errorf("Arguments not enough")
	}
	err := queryRow.Scan(&total)
	if err != nil {
		return err
	}
	if total > 0 {
		if message != "" {
			return errors.New(message)
		}

		return fmt.Errorf("The %s has already been taken", field)
	}

	return nil
}

func NewUniqueRule(db *sql.DB, ruleName string) *UniqueRule {
	return &UniqueRule{db, ruleName}
}
