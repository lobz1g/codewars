package src

func DirReduc(arr []string) []string {
	if len(arr) == 0 {
		return []string{}
	}

	tmpArr := []string{}

	for i := len(arr) - 1; i > 0; i-- {
		if (arr[i] == "NORTH" && arr[i-1] == "SOUTH") ||
			(arr[i] == "SOUTH" && arr[i-1] == "NORTH") ||
			(arr[i] == "EAST" && arr[i-1] == "WEST") ||
			(arr[i] == "WEST" && arr[i-1] == "EAST") {
			tmpArr = append(tmpArr, arr[:i-1]...)
			tmpArr = append(tmpArr, arr[i+1:]...)
			return DirReduc(tmpArr)
		}
	}

	return arr
}
