package go_dijkstra

import "testing"

func TestIndexedMinQueue(t *testing.T) {
	imq := indexedMinQueue{
		data:        make([]keyVal, 0),
		dataToIndex: make(map[keyVal]int),
		keyMap:      make(map[int]struct{}),
	}
	data := []int{5, 4, 3, 2, 1}

	for _, item := range data {
		imq.insert(keyVal{key: item, val: item})
		if !imq.contains(keyVal{key: item, val: item}) {
			t.Fatalf("expected queue to contain %d, but it doesn't", item)
		}
		if imq.dataToIndex[keyVal{key: item, val: item}] != 0 {
			t.Fatalf("expected queue to have the minimum element %d at index 0, but it has %d", item, imq.data[0])
		}
		if len(imq.dataToIndex) != len(imq.data) && len(imq.data) != len(imq.keyMap) {
			t.Fatalf("expected lengths for dataToIndex, data and keyMap to be equal")
		}
	}

	min := imq.deleteMin()
	if min.val != 1 {
		t.Fatalf("expected min element to be 1, but it is %d", min)
	}
	if min.key != 1 {
		t.Fatalf("expected min key to be %d, but it is %d", 1, min.key)
	}

	if len(imq.dataToIndex) != len(imq.data) && len(imq.data) != len(imq.keyMap) {
		t.Fatalf("expected lengths for dataToIndex, data and keyMap to be equal")
	}
	if imq.contains(keyVal{key: 1, val: 1}) {
		t.Fatal("expected queue not to contain 1")
	}
	if _, ok := imq.dataToIndex[min]; ok {
		t.Fatalf("expected index for %d to be not present, but it is", min)
	}

	imq.decreaseKey(keyVal{key: 5, val: 5}, 1)
	if len(imq.dataToIndex) != len(imq.data) && len(imq.data) != len(imq.keyMap) {
		t.Fatalf("expected lengths for dataToIndex, data and keyMap to be equal")
	}
	if !imq.contains(keyVal{key: 5, val: 1}) {
		t.Fatal("expected queue to contain 5")
	}
	min = imq.deleteMin()
	if len(imq.dataToIndex) != len(imq.data) && len(imq.data) != len(imq.keyMap) {
		t.Fatalf("expected lengths for dataToIndex, data and keyMap to be equal")
	}
	if min.key != 5 {
		t.Fatalf("expected min key to be %d, but it is %d", 5, min.key)
	}
	if min.val != 1 {
		t.Fatalf("expected min value to be 1, but it is %d", min)
	}
	if imq.contains(keyVal{key: 5, val: 1}) {
		t.Fatal("expected queue not to contain 5")
	}
	if _, ok := imq.dataToIndex[min]; ok {
		t.Fatalf("expected index for %d to be not present, but it is", min)
	}
}
