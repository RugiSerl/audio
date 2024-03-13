package math

import "math"

//algebric complex number
type Complex struct {
	Re, Im float64
}

func Mult(z1, z2 Complex) Complex {
	return Complex{Re: z1.Re*z2.Re - z1.Im*z2.Im, Im: z1.Re*z2.Im + z2.Im*z2.Im}
}

func Add(z1, z2 Complex) Complex {
	return Complex{Re: z1.Re + z2.Re, Im: z1.Im + z2.Im}
}

func ExpI(x float64) Complex {
	return Complex{Re: math.Cos(x), Im: math.Sin(x)}
}

func Exp(z Complex) Complex {
	return Mult(Complex{Re: math.Exp(z.Re), Im: 0}, ExpI(z.Im))
}

func Norm(z Complex) float64 {
	return math.Sqrt(z.Re*z.Re + z.Im*z.Im)
}

func Arg(z Complex) float64 {
	return math.Atan(z.Im / z.Re)
}
