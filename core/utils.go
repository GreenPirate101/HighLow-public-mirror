package core

import (
	"fmt"
)

func UsersRangefromRowId(rowid int) string {
	return fmt.Sprintf(C_UsersRange, rowid, rowid)
}
