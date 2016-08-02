package experiment

import "os"
import "strconv"
import "reflect"

func Add(lhs, rhs int) int {
	return lhs + rhs
}

func MyPrint(args ...interface{}) {
	for _, arg := range args {
		switch v := reflect.ValueOf(arg); v.Kind() {
		case reflect.String:
			os.Stdout.WriteString(v.String())
		case reflect.Int:
			os.Stdout.WriteString(strconv.FormatInt(v.Int(), 10))
		}
	}
}
