# Audio

### Preliminary
- We often write $\cos(\omega t)$, with $\omega=2\pi f$, with $f$ the frequency of the signal, as a shortcut for $\cos(2\pi f t)$
- Sine and cosine factorisation: we can factor $A\cos(\omega t) + B\sin(\omega t)$ like this:
$$\begin{equation}
\begin{split}   A\cos(\omega t) + B\sin(\omega t)
&=\sqrt{A^2+B^2}(\frac {A} {\sqrt{A^2+B^2}}\cos(\omega t) + \frac {B} {\sqrt{A^2+B^2}}\sin(\omega t))\\
&=\sqrt{A^2+B^2}(\cos(\phi) \cos(\omega t) - \sin(\phi)\sin(\omega t))\\
&=\sqrt{A^2+B^2}\cos(\omega t + \phi)\\
\end{split}
\end{equation}$$
$${\text with}\begin{cases}
   \cos(\phi) = \frac {A} {\sqrt{A^2+B^2}}\\
   \sin(\phi) = \frac {B} {\sqrt{A^2+B^2}}
\end{cases}$$
$$\text {so }\tan(\phi) = \frac { \frac {B} {\sqrt{A^2+B^2}}} {\frac {A} {\sqrt{A^2+B^2}}} = \frac B A  $$
$$\text {and therefore }\phi ≡ \arctan(\frac B A) + \frac \pi 2 [\pi]$$
>*note: the signals must have the same frequency*


### Discrete fourier transform (DFT)
Let $N$ the amount of samples.

For a given vector $(s_k)_{k\in{ \llbracket0, N-1 \rrbracket}}$ of samples, its fourier transform $(c_n)_{n\in{ \llbracket0, N-1 \rrbracket}}$ is defined as:  

$$\forall n \in{\llbracket0, N-1 \rrbracket},  c_n = \sum_{k=0}^{N-1} s_ke^{-i2\pi k \frac n N}$$  

> I suggest you to watch [This 3Blue1Brown video](https://youtu.be/spUNpyF58BY), which gives a good idea of the intuition behind it.

>Here, $c_n$ is a complex numbers whose real part is "the amount of cosine of the signal" and whose complex part is "the amount of sine of the signal", so if we where to "extract" the nth frequency of the signal, it would be like :
>$$s_f(t) = Re(c_n)\cos(n\omega t) +Im(c_n)\sin(n\omega t)$$*

If we were to factor the sine and cosine, we would get :
$$\begin{equation}
\begin{split}  Re(c_n)\cos(n\omega t) +Im(c_n)\sin(n\omega t)
&=\sqrt{Re(c_n)^2+Im(c_n)^2}\cos(n\omega t + \phi) \\
&=|c_n|\cos(n\omega t + \arg(c_n))
\end{split}
\end{equation}$$

>In other words, $|c_n|$ is the magnitude of the frequency and $\arg(c_n)$ is the phase

So to compute naively all the coefficients, we would have two nested loops of N repetitions and the complexity would be $O(N^2)$.

>you can look the [implementation in Go](https://github.com/RugiSerl/audio/blob/main/math/fourier.go#L50). 

The inverse fourier transform is defined as:

$$\forall k \in{\llbracket0, N-1 \rrbracket},  s_k = \frac 1 N\sum_{n=0}^{N-1} c_ne^{i2\pi n \frac k N}$$  
*replace $c_n$ by its definition above and everything should cancel*.

So the inverse fourier transform is almost identical to the fourier transform.

### Fast fourier transform (FFT) - Cooley–Tukey algorithm
For the rest of the document, we define $\omega = e^{\frac {i2\pi} N}$, a Nth root of unity.

For a given vector $(s_k)_{k\in{ \llbracket0, N-1 \rrbracket}}$ of data and $P = \sum_{k=0}^{N-1}s_kX^k$, computing the fourier transform for $c_n$ is the same thing as evaluating $P(\omega^n)$

So instead of naively evalutating this polynomial at $1, \omega, ..., \omega^{n-1}$, which would result in a complexity of $O(N^2)$ we will first factor our polynomial like this :
$$P = \underbrace{\sum_{k=0}^{\frac{N-1} 2}s_{2k}X^{2k}}_{P_{even}} + X\underbrace{\sum_{k=0}^{\frac{N-1} 2}s_{2k+1}X^{2k}}_{P_{odd}}$$

>*please note that we impose that $N = 2^k, k \in \Bbb N$ for this algorithm*

So we have :

$$P(X) = P_{even}(X^2) + XP_{odd}(X^2)$$

And since $P_{even}$ and $P_{odd}$ are even functions, we can deduce that :

$$P(-X) = P_{even}(X^2) - XP_{odd}(X^2)$$

and in our context: 

$$P(-\omega^k) = P(e^{i\pi}\omega^k) = P(\omega^{k+\frac N 2}) = P_{even}(\omega^{2k}) - \omega^kP_{odd}(\omega^{2k})$$

>This will allow us to reduce by half the amount of evaluations of $P$, since we will only have to compute $P(\omega^k)$ for $k \in \llbracket 0, \frac N 2-1 \rrbracket$, and deduce the other values

So the point of the algorithm will be to calculate recursively $P_{even}(1), P_{even}(\omega^2), ..., P_{even}(\omega^{2(N-1)})$ and $P_{odd}(1), P_{odd}(\omega^2), ..., P_{odd}(\omega^{2(N-1)})$

and now for $k \in \llbracket 0, \frac N 2-1 \rrbracket$ :
$$\begin{cases}
   P(\omega^k) = P_{even}(\omega^{2k}) + \omega^kP_{odd}(\omega^{2k}) \\
   P(\omega^{k+\frac N 2}) = P_{even}(\omega^{2k}) - \omega^kP_{odd}(\omega^{2k})
\end{cases}$$

here's a pseudo-code of the fft


```
function fft(P, ω) -> ComplexArray {
    let N = |P|;
    if N = 1 then return P(1);
    let Podd = P.odd();
    let Peven = P.even();
    
    let oddResult = fft(Podd, ω²);
    let evenResult = fft(Peven, ω²);
    let result = new(ComplexArray);

    let ωk = 1; // ω to the power of k
 
    for k=0 to N/2-1 do {
        result[k] <- evenResult[k] + ωk*oddResults[k];
        result[k+N/2] <- evenResult[k] - ωk*oddResults[k];
        ωk <- ωk*ω;
    }
    return result;
}
```

>here's the [implementation in Go](https://github.com/RugiSerl/audio/blob/main/math/fourier.go#L83). 




