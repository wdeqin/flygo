package main

import "github.com/wdeqin/flyit/dispatch"
import "math/rand"

func main() {
	d := dispatch.NewDefaultDispatchee(7)
	dd := dispatch.NewThresholdDispatcher(4, &d)
	data := make([]interface{}, 10)
	for c := 0; c < 100; c++ {
		for i := range data {
			data[i] = rand.Int() % 100
		}
		dd.Dispatch(data)
	}
	dd.CleanUp()
	dd.Wait()
}
