package main

import(
	"p2b/math_prac/math"
	"fmt"
	"os"
	"strconv"
	"flag"
)

func main(){
	args := os.Args[1:]
	rA, _ := strconv.ParseFloat(args[1], 64)
	rB, _ := strconv.ParseFloat(args[1], 64)
	rC, _ := strconv.ParseFloat(args[2], 64)
	rD, _ := strconv.ParseFloat(args[3], 64)
	var a = flag.Float64("a", rA, "a")
	var b = flag.Float64("b", rB, "b")
	var c = flag.Float64("c", rC, "c")
	var d = flag.Float64("d", rD, "d")
	flag.Parse()
	var f, x, y, z math.NotAnInteger
	x.Pow(*a,2)
	y.Multiply(2, *b)
	z.Multiply(*c,*d)
	f.Add(x.F,y.F)
	f.Add(f.F, z.F)
	fmt.Println(f.F)
}