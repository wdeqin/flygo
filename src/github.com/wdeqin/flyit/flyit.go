package main

import "github.com/wdeqin/flyit/dispatch"
import "math/rand"
import "time"

func main() {
	d := dispatch.NewDefaultDispatchee(7)
	dd := dispatch.NewThresholdDispatcher(4, &d)
	data := make([]interface{}, 10)
	for c := 0; c < 10; c++ {
		for i, _ := range data {
			data[i] = rand.Int() % 100
		}
		dd.Dispatch(data)
	}
	dd.CleanUp()

	<-time.After(5 * time.Second)
}
