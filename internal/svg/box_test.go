package svg

import "testing"

func TestBox_Combine(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		var box Box
		other := Box{
			Size: Point{X: 3, Y: 4},
		}
		if box.Combine(other) != other {
			t.Errorf("Expected %v, but got %v", other, box)
		}
	})
	t.Run("gap", func(t *testing.T) {
		box := Box{Position: Point{X: 1, Y: 2}, Size: Point{X: 3, Y: 4}}
		other := Box{Position: Point{X: 5, Y: 6}, Size: Point{X: 7, Y: 8}}
		expected := Box{Position: Point{X: 1, Y: 2}, Size: Point{X: 11, Y: 12}}
		if result := box.Combine(other); result != expected {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})
}
