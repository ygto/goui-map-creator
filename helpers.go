package mapCreator

func ArrayPop(array []*Node) (*Node, []*Node) {
	if len(array) == 0 {
		return nil, nil
	}
	return array[len(array)-1], array[:len(array)-1]
}
