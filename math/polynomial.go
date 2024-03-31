package math

type Polynomial struct {
	coefs []Complex // Representation of polynomials with coefficients
}

func NewPolynomialFromReal(coefs []float64) Polynomial {
	p := Polynomial{make([]Complex, len(coefs))}
	for i := 0; i < len(coefs); i++ {
		p.coefs[i] = Real(coefs[i])
	}
	return p
}

// Returns the odd part of a polynomial
func (p Polynomial) Odd() Polynomial {
	r := Polynomial{make([]Complex, len(p.coefs)/2)}
	for i := 0; i < len(p.coefs)/2; i++ {
		r.coefs[i] = p.coefs[2*i]
	}
	return r
}

// Returns the even part of a polynomial
func (p Polynomial) Even() Polynomial {
	r := Polynomial{make([]Complex, len(p.coefs)/2)}
	for i := 0; i < len(p.coefs)/2; i++ {
		r.coefs[i] = p.coefs[2*i+1]
	}
	return r
}

// Evaluate the polynomial at a certain point
func (p Polynomial) Evaluate(z Complex) Complex {
	s := Complex{Re: 0, Im: 0}
	for i := 0; i < len(p.coefs); i++ {
		s = Add(s, Mult(p.coefs[i], z.Pow(i)))
	}
	return s
}

// Fast exponentiation algorithm in O(log(n))
func (z Complex) Pow(exponent int) Complex {
	if exponent == 0 {
		return Real(1)
	} else if exponent == 1 {
		return z
	} else if exponent%2 == 0 {
		r := z.Pow(exponent / 2)
		return Mult(r, r)
	} else if exponent%2 == 1 {
		r := z.Pow(exponent / 2)
		return Mult(z, Mult(r, r))
	} else {
		return Complex{Re: 1, Im: 0}
	}
}
