package math

import(
	"testing"
	"math"
)

func TestMath(t *testing.T){
	cases := []struct {
		a float64
		b float64
		c float64
		d float64
	}{
		{1.0, 2.0, 3.0, 4.0},
		{4.0, 3.0, 2.0, 1.0},
		{0.0, 1.0, 0.0,1.0},
		{3.45, 4.56, 5.67, 6.78},
	}
	var f, f1,f2,f3,f4,f5, a, b, c NotAnInteger
	for _, x := range cases{
		f1.Subtract(x.a, x.b)
		if f1.F != (x.a - x.b) {
			t.Errorf("Sub failed with %F - %F", x.a, x.b)
		}

		f2.Add(x.a, x.b)
		if f2.F != (x.a + x.b){
			t.Errorf("Add failed with %F + %F", x.a, x.b)
		}

		f3.Multiply(x.a, x.b)
		if f3.F != (x.a * x.b){
			t.Errorf("Multiply failed with %F * %F", x.a, x.b)
		}

		f4.Pow(x.a, x.b)
		if f4.F != (math.Pow(x.a,x.b)){
			t.Errorf("Pow failed with %F ** %F", x.a, x.b)
		}

		f5.Divide(x.a, x.b)
		if f5.F != (x.a / x.b){
			t.Errorf("Divide failed with %F / %F", x.a, x.b)
		}

		a.Pow(x.a, 2)
		b.Multiply(2,x.b)
		c.Multiply(x.c,x.d)
		f.Add(b.F, a.F)
		f.Add(f.F, c.F)
		if f.F != math.Pow(x.a, 2) + (2 * x.b) + (x.c * x.d){
			t.Errorf("Pythagorean failed with %F ^ 2 + 2 * %F + %F * %F", x.a, x.b, x.c, x.d)
		}
	}
}
