package running

import (
	"database/sql/driver"
	"fmt"
)

// PgBoolAsInt 把 PostgreSQL 的 boolean 读成 1/0，写回去时再把 1/0 转成 true/false
type PgBoolAsInt struct {
	Int *int
}

// Scan 读库：t/f → 1/0
func (b *PgBoolAsInt) Scan(src interface{}) error {
	if src == nil {
		b.Int = nil
		return nil
	}

	var v bool
	switch s := src.(type) {
	case bool:
		v = s
	case []byte:
		v = len(s) > 0 && (s[0] == 't' || s[0] == 'T' || s[0] == '1')
	case string:
		v = s == "t" || s == "true" || s == "T" || s == "1"
	default:
		return fmt.Errorf("unsupported type: %T", src)
	}

	if v {
		one := 1
		b.Int = &one
	} else {
		zero := 0
		b.Int = &zero
	}
	return nil
}

// Value 写库：1 → true，0 或 nil → false
func (b PgBoolAsInt) Value() (driver.Value, error) {
	if b.Int == nil {
		return false, nil
	}
	return *b.Int != 0, nil
}
