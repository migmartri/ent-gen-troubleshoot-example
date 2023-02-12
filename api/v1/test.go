package v1

import (
	"database/sql/driver"
)

type Test struct {
}

// Required to be stored in a byte[] database column with ent
// See https://entgo.io/docs/faq/#how-to-store-protobuf-objects-in-a-blob-column
// Implementation of scan interface
func (x *Test) Value() (driver.Value, error) {
	return "string", nil
}

func (x *Test) Scan(src any) error {
	return nil
}
