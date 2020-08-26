package main

import (
	"fmt"
)

func main() {
	hp := Init([]int{234, 10, 58, 2, 47, 38, 22, 100, 22, 13, 44, 25, 16, 47, 28, 19})
	fmt.Printf("%#v\r\n", hp)
	fmt.Println(hp.Pop())
	fmt.Printf("%#v\r\n", hp)
}

// 二叉树是链式存储,而二叉堆是线性存储.
// 父节点和孩子节点的索引关系
// parent(i) = (i-1)/2
// left_child(i) = 2*i + 1
// right_child(i) = 2*i + 2

func parent(i int) int {
	if i == 0 {
		return 0
	}
	return (i - 1) / 2
}

func child(i int, left bool) int {
	v := 2*i + 1
	if !left {
		v += 1
	}
	return v
}

type minHeap []int

func (sf *minHeap) add(e int) {
	*sf = append(*sf, e)
	sf.up(len(*sf) - 1)
}

// 在成型的堆上,进行上浮,只需与父节点比较即可.
func (sf minHeap) up(idx int) {
	for idx > 0 {
		if pidx := parent(idx); sf[pidx] > sf[idx] {
			sf[pidx], sf[idx] = sf[idx], sf[pidx]
			idx = pidx
		}
	}
}

// 在成型的堆上,进行下沉,要与子节点的最大(或最小)比较
// 注意边界条件,有可能无子节点,或只有一个左节点
func (sf minHeap) down(idx int) {
	for child(idx, true) < len(sf) {
		pos := child(idx, true)
		min := sf[pos]
		if ridx := child(idx, false); ridx < len(sf) && sf[ridx] < sf[pos] {
			min = sf[ridx]
			pos = ridx
		}

		if sf[idx] < min {
			break
		}
		sf[pos], sf[idx] = sf[idx], sf[pos]
		idx = pos
	}
}

func (sf *minHeap) Pop() int {
	mh := *sf
	if len(mh) == 0 {
		return 0
	}

	v := mh[0]
	mh[len(mh)-1], mh[0] = mh[0], mh[len(mh)-1]
	*sf = mh[:len(mh)-1]
	sf.down(0)

	return v
}

// 构建无序的二叉堆调整为有序的二插堆,本质上是让所有非叶子节点依次下沉
func Init(b []int) minHeap {
	mh := minHeap(b)
	for i := parent(len(b) - 1); i >= 0; i-- {
		mh.down(i)
	}
	return mh
}
