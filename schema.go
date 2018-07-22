package mapCreator

type Schema struct {
	y      int
	x      int
	schema [][]int
}

func NewSchema(y int, x int, def int) *Schema {
	s := new(Schema)
	s.y = y
	s.x = x
	s.schema = make([][]int, 0, 0)
	for y := 0; y < s.y; y++ {
		temp := make([]int, 0, 0)
		for x := 0; x < s.x; x++ {
			temp = append(temp, def)
		}
		s.schema = append(s.schema, temp)
	}
	return s
}

func (s *Schema) GetY() int {
	return s.y
}
func (s *Schema) GetX() int {
	return s.x
}
func (s *Schema) SetSchema(y int, x int, val int) {
	s.schema[y][x] = val
}

func (s *Schema) GetSchema(y int, x int) int {
	return s.schema[y][x]
}
