package math

import "math"

//Bruh I did all this without realizing that the complex type is already available natively in go.

// algebric complex number
type Complex struct {
	Re, Im float64
}

func Real(x float64) Complex {
	return Complex{x, 0}
}

// Multiply two complex numbers
func Mult(z1, z2 Complex) Complex {
	return Complex{Re: z1.Re*z2.Re - z1.Im*z2.Im, Im: z1.Re*z2.Im + z1.Im*z2.Re}
}

// Add two complex numbers
func Add(z1, z2 Complex) Complex {
	return Complex{Re: z1.Re + z2.Re, Im: z1.Im + z2.Im}
}

// Substract two complex numbers
func Substract(z1, z2 Complex) Complex {
	return Complex{Re: z1.Re - z2.Re, Im: z1.Im - z2.Im}
}

// Imaginary exponential
func ExpI(x float64) Complex {
	return Complex{Re: math.Cos(x), Im: math.Sin(x)}
}

// Complex exponential
func Exp(z Complex) Complex {
	return Mult(Real(math.Exp(z.Re)), ExpI(z.Im))
}

// just a shortcut
func Omega(power float64) Complex {
	return ExpI(2 * math.Pi * float64(power))
}

// Returns |z|
func Norm(z Complex) float64 {
	return math.Sqrt(z.Re*z.Re + z.Im*z.Im)
}

func Arg(z Complex) float64 {
	return math.Atan(z.Im / z.Re)
}
