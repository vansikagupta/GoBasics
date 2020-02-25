package main

import ("fmt"
		"math")
		
type Vector struct{
	X, Y float64
}
//methods associated with type Vector
//value receiver
func (v Vector) abs() float64{
	v.X = 3// changes made are local; they do not reflect on the actual Vector variable
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}
//pointer receiver
func (v *Vector) Scale() {
	v.X *= 10
	v.Y *= 10 // changes reflect on original copy
}

//this is a regular function : methods are same as functions
func absfunc(v Vector) float64{
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func main(){
	v := Vector{4,5}
	fmt.Println(v)
	fmt.Println(v.abs())
	v.Scale()
	fmt.Println(v)
	fmt.Println(absfunc(v))
}