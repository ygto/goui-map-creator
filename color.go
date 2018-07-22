package mapCreator

import "github.com/ygto/goui/color"

func GetColor(val int) *color.Color {
	var c *color.Color
	switch val {
	case TILE_WALL:
		c = color.CreateColor(0, 0, 0)
	case TILE_ROAD:
		c = color.CreateColor(255, 255, 255)
	case TILE_GOAL:
		c = color.CreateColor(0, 255, 0)
	case TILE_DANGER:
		c = color.CreateColor(255, 0, 0)
	}
	return c
}
