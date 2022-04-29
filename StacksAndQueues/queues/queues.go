package queues

type Queue []int

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Enqueue(num int) {
	*q = append(*q, num)
}

func (q *Queue) Dequeue() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	} else {
		element := (*q)[0]
		*q = (*q)[1:]
		return element, true
	}
}
