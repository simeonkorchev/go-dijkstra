package go_dijkstra

type keyVal struct {
	key int
	val int
}

type indexedMinQueue struct {
	data        []keyVal
	dataToIndex map[keyVal]int
	keyMap      map[int]struct{}
}

func (i *indexedMinQueue) insert(item keyVal) {
	i.keyMap[item.key] = struct{}{}
	i.data = append(i.data, item)

	if len(i.data) == 1 {
		i.dataToIndex[i.data[0]] = 0
		return
	}
	i.swim(len(i.data) - 1)
}

func (i *indexedMinQueue) swim(k int) {
	for k > 0 {
		if i.data[k/2].val < i.data[k].val {
			return
		}
		i.dataToIndex[i.data[k]] = k / 2
		i.dataToIndex[i.data[k/2]] = k
		i.data[k], i.data[k/2] = i.data[k/2], i.data[k]
		k /= 2
	}
}

func (i *indexedMinQueue) deleteMin() keyVal {
	min := i.data[0]
	delete(i.keyMap, min.key)
	delete(i.dataToIndex, min)
	i.data[0] = i.data[len(i.data)-1]
	i.data = i.data[:len(i.data)-1]
	if len(i.data) == 0 {
		return min
	}
	i.dataToIndex[i.data[0]] = 0
	if len(i.data) > 1 {
		i.sink(0)
	}
	return min
}

func (i *indexedMinQueue) sink(k int) {
	for 2*k < len(i.data) {
		index := i.lesserChildIndex(2*k, 2*k+1)
		if i.data[k].val < i.data[index].val {
			return
		}
		i.dataToIndex[i.data[k]] = index
		i.dataToIndex[i.data[index]] = k
		i.data[index], i.data[k] = i.data[k], i.data[index]
		k = index
	}
}

func (i *indexedMinQueue) lesserChildIndex(j, k int) int {
	if j >= len(i.data)-1 {
		return len(i.data) - 1
	}
	if i.data[j].val < i.data[k].val {
		return j
	}
	return k
}

func (i *indexedMinQueue) decreaseKey(kv keyVal, newVal int) {
	if !i.contains(kv) {
		return
	}
	index := i.dataToIndex[kv]
	delete(i.dataToIndex, kv)
	kv.val = newVal
	i.data[index] = kv
	i.dataToIndex[kv] = index
	if len(i.data) > 1 {
		i.sink(index)
		i.swim(index)
	}
}

func (i *indexedMinQueue) contains(e keyVal) bool {
	_, ok := i.keyMap[e.key]
	//if !ok {
	//	return false
	//}
	//_, ok = i.dataToIndex[e]
	return ok
}

func (i *indexedMinQueue) isEmpty() bool {
	return len(i.data) == 0
}
