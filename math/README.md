## math module
- "complex.go" defines the complex number in the for $a + ib$, with $a$ and $b$ real numbers and elementary operations associated
- "polynomial.go" implements polynomials with complex coefficients, defined above, and useful operations
- "fourier.go" implements the fourier transform, in two ways :
    - Discrete fourier transform: 
    this algorithm naively compute each fourier coefficient one by one, with a complexity of $O(n^2)$.  
    Discrete fourier transform definition :
    $$ c_n = \sum_{k=0}^Ns_ne^{i2\pi k \frac n N}$$
    - Fast fourier transform (FFT), using Cooley–Tukey algorithm