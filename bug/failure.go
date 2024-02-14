package bug

import (
	"fmt"
)

func Exit(msg ...any) {
	panic(fmt.Sprint(msg...))
}
