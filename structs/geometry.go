package geometry

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (t Triangle) Perimeter() float64 {
	return 0
}

func (t Triangle) Area() float64 {
	return t.Width * t.Height / 2
}
