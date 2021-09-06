package errors

import "fmt"

type InvalidAvroTableName struct {
	ErrorCode string
	TableName interface{}
}

func (e InvalidAvroTableName)Error() string {
	return fmt.Sprintf("%v is not a valid Avro table name", e.TableName)
}