package db

import (
	"database/sql"
	"errors"
	"fmt"
)

type QueryPrepareError struct {
	Err error
}

func (e QueryPrepareError) Error() string {
	return fmt.Sprintf("prepare query failed: %v", e.Err)
}

func IsErrNotFound(err error) bool {
	return err != nil && errors.Is(err, sql.ErrNoRows)
}
