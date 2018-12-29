package heaps

type MaxPriorityQueue struct {
	array []int
}

func New(a []int) *MaxPriorityQueue {
	maxHeapify(a, 0, len(a))
	return &MaxPriorityQueue{array: a}
}

func (m *MaxPriorityQueue) Max() (int, bool) {

	if len(m.array) > 0 {
		return m.array[0], true
	} else {
		return 0, false
	}
}

func (m *MaxPriorityQueue) ExtractMax() (int, bool) {

	if len(m.array) == 0 {
		return 0, false
	}

	max := m.array[0]
	m.array[0] = m.array[len(m.array)-1]
	m.array = append([]int(nil), m.array[0:len(m.array)-1]...)
	maxHeapify(m.array, 0, len(m.array))
	return max, true

}

func (m *MaxPriorityQueue) Insert(v int) {

	m.array = append(m.array, v)
	for index := len(m.array) - 1; index > 0; index = parentIndex(index) {
		if m.array[parentIndex(index)] < m.array[index] {
			m.array[parentIndex(index)], m.array[index] = m.array[index], m.array[parentIndex(index)]
		} else {
			break
		}
	}

}
