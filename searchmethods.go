func LinearSeach(v [] int, e int) int {
	for i:=0; i<len(v); i++{
		if v[i] == e{
			return v[i]
		}
	}
	return -1
}

func BinarySearch(v [] int, e int, ini int, fim int) int {
	if ini > fim {
		return -1
	}
	mid = (ini+fim)/2
	if v[mid] == e{
		return mid
	} else {
		if e < v[mid]{
			return BinarySearch(v, e, ini, mid-1)
		}
		return BinarySearch(v, e, mid+1, fim)
	}
} 