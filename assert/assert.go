package assert

import (
	"fmt"
)

func That(msg ...any) {
	panic(fmt.Sprint(msg...))
}
