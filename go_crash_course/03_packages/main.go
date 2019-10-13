package main

import (
	"fmt"

	. "github.com/keriwisnu/go_crash_course/03_packages/strutil"
)

func main() {
	fmt.Println("helloworld")
	fmt.Println(strutil.reverse("helloworld"))
}
