package dispatch

import "fmt"
import "bytes"

type Dispatcher interface {
	Dispatch([]interface{}) bool
	CleanUp() bool
}

type SorDataList struct {
	SorNum int
	Data   []interface{}
}

type Dispatchee interface {
	GetNumOfSor() int
	GetSorNum(interface{}) int
	ProcessSor(int, []interface{}) int
	ProcessSors([]SorDataList) int
}

type thresholdDispatcher struct {
	threshold  int
	dataList   [][]interface{}
	dataCount  []int
	dispatchee Dispatchee
}

func NewThresholdDispatcher(threshold int, dispatchee Dispatchee) thresholdDispatcher {
	var d thresholdDispatcher
	if threshold <= 0 {
		panic(fmt.Sprintf("threshold must greater than zero, got %d", threshold))
	}

	if dispatchee == nil {
		panic("dispatchee must not be nil")
	}

	d.threshold = threshold
	d.dispatchee = dispatchee

	d.dataList = make([][]interface{}, dispatchee.GetNumOfSor())
	for i, _ := range d.dataList {
		d.dataList[i] = make([]interface{}, d.threshold)
	}
	d.dataCount = make([]int, dispatchee.GetNumOfSor())

	return d
}

func (d *thresholdDispatcher) Dispatch(data []interface{}) bool {
	for _, e := range data {
		sor := d.dispatchee.GetSorNum(e)
		d.dataList[sor][d.dataCount[sor]] = e
		d.dataCount[sor]++
		if d.dataCount[sor] >= d.threshold {
			d.dispatchee.ProcessSor(sor, d.dataList[sor])
			d.dataList[sor] = make([]interface{}, d.threshold)
			d.dataCount[sor] = 0
		}
	}
	return true
}

func (d *thresholdDispatcher) CleanUp() bool {
	for i, l := range d.dataList {
		if d.dataCount[i] > 0 {
			d.dispatchee.ProcessSor(i, l)
			d.dataList[i] = make([]interface{}, d.threshold)
		}
	}
	return true
}

type defaultDispatchee struct {
	numOfSor int
}

func NewDefaultDispatchee(numOfSor int) defaultDispatchee {
	if numOfSor <= 0 {
		panic(fmt.Sprintf("numOfSor must grater than zero, got %d", numOfSor))
	}
	var dd defaultDispatchee
	dd.numOfSor = numOfSor
	return dd
}

func (dd *defaultDispatchee) GetNumOfSor() int {
	return dd.numOfSor
}

func (dd *defaultDispatchee) GetSorNum(e interface{}) int {
	ee, ok := e.(int)
	if ok {
		return ee % dd.numOfSor
	} else {
		panic("e must be int")
	}
}

func (dd *defaultDispatchee) ProcessSor(sorNum int, data []interface{}) int {
	go func() {
		var buf bytes.Buffer
		buf.WriteString(fmt.Sprintf("$%d [", sorNum))
		for _, e := range data {
			if e == nil {
				break
			}
			buf.WriteString(fmt.Sprintf("%2v, ", e))
		}
		buf.WriteString(fmt.Sprintf("\b\b]\n"))
		fmt.Print(buf.String())
	}()
	return 0
}

func (dd *defaultDispatchee) ProcessSors([]SorDataList) int {
	panic("not implemented")
}
