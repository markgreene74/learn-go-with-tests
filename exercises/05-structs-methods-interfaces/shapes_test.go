package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// v1
// func TestArea(t *testing.T) {
// 	checkArea := func(t testing.TB, shape Shape, want float64) {
// 		t.Helper()
// 		got := shape.Area()
// 		if got != want {
// 			t.Errorf("got %g want %g", got, want)
// 		}
// 	}

// 	t.Run("rectangles", func(t *testing.T) {
// 		rectangle := Rectangle{12.0, 6.0}
// 		checkArea(t, rectangle, 72.0)
// 	})

// 	t.Run("circles", func(t *testing.T) {
// 		circle := Circle{10}
// 		checkArea(t, circle, 314.1592653589793)
// 	})

// 	t.Run("triangles", func(t *testing.T) {
// 		triangle := Triangle{10.5, 5.5}
// 		checkArea(t, triangle, 28.875)
// 	})
// }

// v2
// func TestArea(t *testing.T) {

// 	testCases := []struct {
// 		shape Shape
// 		want  float64
// 	}{
// 		{Circle{10}, 314.1592653589793},
// 		{Rectangle{12.0, 6.0}, 72.0},
// 		{Triangle{10.5, 5.5}, 28.875},
// 	}

// 	for _, tt := range testCases {
// 		got := tt.shape.Area()
// 		if got != tt.want {
// 			t.Errorf("got %g want %g", got, tt.want)
// 		}
// 	}
// }

// v3
func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, hasArea float64) {
		t.Helper()
		got := shape.Area()
		if got != hasArea {
			t.Errorf("%#v got %g want %g", shape, got, hasArea)
		}
	}

	testCases := []struct {
		shape   Shape
		hasArea float64
	}{
		{shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{shape: Rectangle{Width: 12.0, Height: 6.0}, hasArea: 72.0},
		{shape: Triangle{Base: 10.5, Height: 5.5}, hasArea: 28.875},
	}

	for _, tt := range testCases {
		checkArea(t, tt.shape, tt.hasArea)
	}
}
