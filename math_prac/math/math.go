package math

import (
	"math"
)

type NotAnInteger struct{
	F float64
}

func (f *NotAnInteger) Subtract(x float64, y float64){
	f.F = x - y
}

func (f *NotAnInteger) Add(x float64, y float64){
	f.F = x + y
}

func (f *NotAnInteger) Multiply(x float64, y float64){
	f.F = x * y
}

func (f *NotAnInteger) Pow(x float64, y float64){
	f.F = math.Pow(x, y)
}

func (f *NotAnInteger) Divide(x float64, y float64){
	f.F = x / y
}
