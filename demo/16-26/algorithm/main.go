package main

import (
	"algorithm/queue"
	"log"
)

type DataSource struct {
	Value int

	//数组编号
	ArrayNum int

	//数组中的index
	Index int
}

func NewDataSource(val, arrayNum, index int) *DataSource {
	return &DataSource{
		Value:    val,
		ArrayNum: arrayNum,
		Index:    index,
	}
}

func (d DataSource) compare(dataSource *DataSource) int {
	if d.Value > dataSource.Value {
		return -1
	}

	if d.Value < dataSource.Value {
		return 1
	}
	return 0
}
func getTop(data [][]int) []int {
	if len(data) == 0 {
		return nil
	}
	rowSize := len(data)
	columnSize := len(data[0])
	ret := make([]int, columnSize)
	//保持一个最小堆存放来自20个数组的最小值
	heap := queue.NewSliceQueue()
	for i := 0; i < rowSize; i++ {
		//记录源数组的编号和index
		heap.EnQueue(NewDataSource(data[i][0], i, 0))
	}
	num := 0
	for num < columnSize {
		//删除定点元素
		d := heap.DeQueue().(*DataSource)
		ret[num] = d.Value
		num++
		if num > columnSize {
			break
		}
		d.Value = data[d.ArrayNum][d.Index+1]
		d.Index++
		heap.EnQueue(d)
	}
	return ret
}

func main() {
	data := [][]int{{9, 7, 6, 3, 1}, {100, 67, 23, 14, 5}, {137, 121, 111, 104, 101}, {77, 67, 54, 36, 29}}
	log.Println(getTop(data))
}
