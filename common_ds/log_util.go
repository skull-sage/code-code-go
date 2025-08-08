package common_ds

import (
	"fmt"
)

func Log(msg string) {
	fmt.Println("#", "", msg)
}

func Logf(format string, a ...any) {
	fmt.Printf("# %s\n", fmt.Sprintf(format, a...))
}

func Logln(a ...any) {
	fmt.Println(a...)
}
