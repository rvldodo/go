// package main

// import "fmt"

// type shape interface {
// 	getArea() float64
// }

// type square struct{
// 	sideLength float64
// }

// type triangle struct {
// 	height float64
// 	base float64
// }

// func main() {
// 	sq := square{5.0}
// 	tr := triangle{9.0, 5.0}

// 	printArea(sq)
// 	printArea(tr)
// }

// func (s square) getArea() float64 {
// 	return s.sideLength * s.sideLength
// }

// func(t triangle) getArea() float64 {
// 	return 0.5 * t.base * t.height
// }

// func printArea(s shape) {
// 	fmt.Printf("Total area: %v\n", s.getArea())
// }

package main

import "fmt" 

type shape interface {     
	getArea() float64 
} 
type triangle struct {     
	height float64     
	base   float64 
} 
type square struct {     
	sideLength float64 
} 

func main() {     
	triangle := triangle{height: 12.0, base: 17.0}     
	square := square{sideLength: 4.0}    
	printArea(triangle)     
	printArea(square) 
} 

func (t triangle) getArea() float64 {     
	return 0.5 * t.base * t.height 
} 
func (s square) getArea() float64 {     
	return s.sideLength * s.sideLength 
} 
func printArea(s shape) {     
	fmt.Println(s.getArea()) 
}