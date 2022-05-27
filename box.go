package golang_united_school_homework

import "fmt"

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) >= b.shapesCapacity {
		return fmt.Errorf("box is full")
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i < 0 || i >= len(b.shapes) {
		return nil, fmt.Errorf("index out of range")
	}
	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if shape, err := b.GetByIndex(i); err != nil {
		return nil, err
	} else {
		b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
		return shape, nil
	}
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if oldShape, err := b.GetByIndex(i); err != nil {
		return nil, err
	} else {
		b.shapes[i] = shape
		return oldShape, nil
	}
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() (perimeter float64) {
	for _, shape := range b.shapes {
		perimeter = perimeter + shape.CalcPerimeter()
	}
	return
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() (area float64) {
	for _, shape := range b.shapes {
		area = area + shape.CalcArea()
	}
	return
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	exist := false
	for i := 0; i < len(b.shapes); i++ {
		if _, ok := b.shapes[i].(*Circle); ok {
			exist = true
			b.ExtractByIndex(i)
			i--
		}
	}
	if !exist {
		return fmt.Errorf("no circles in the box")
	}
	return nil
}
