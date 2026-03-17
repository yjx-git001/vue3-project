package models

import (
	"database/sql/driver"
	"encoding/json"
)

type Array[T any] []T

func (c Array[T]) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *Array[T]) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}
func (c *Array[T]) Slice() []T {
	return *c
}
