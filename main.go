package main

type MyCircularQueue struct {
	slice []int
	size  int
}

func Constructor(k int) MyCircularQueue {

	return MyCircularQueue{
		slice: []int{},
		size:  k,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.size > len(this.slice) {
		this.slice = append(this.slice, value)
		return true
	} else {
		return false
	}
}

func (this *MyCircularQueue) DeQueue() bool {
	if len(this.slice) > 0 {
		this.slice = this.slice[1:len(this.slice)]
		return true
	} else {
		return false
	}
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	} else {
		return this.slice[0]
	}
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	} else {
		return this.slice[len(this.slice)-1]
	}
}

func (this *MyCircularQueue) IsEmpty() bool {
	return len(this.slice) == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return len(this.slice) == this.size
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */

func main() {

}
