package main

import u "github.com/wdeqin/flyit/util"
import "fmt"

func main() {
	u.MyPrint("1 + 1 = ", u.Add(1, 1), "\n")
	fmt.Printf("%v\n", u.MyPrint)
	fmt.Printf("%+v\n", u.MyPrint)
	fmt.Printf("%#v\n", u.MyPrint)
}
