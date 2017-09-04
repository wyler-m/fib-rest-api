package main

import "math/big"

func fib(r *big.Int) *big.Int{
	// defines objects for performing bigInt arithmatic 
	k := new(big.Int).Set(r)
	out := new(big.Int).SetUint64(0)
	zero := new(big.Int).SetUint64(0)
	one := new(big.Int).SetUint64(1)
	two := new(big.Int).SetUint64(2)
	
	// cases for if k <= 1 or k == 2 
	if k.Cmp(one) == -1 || k.Cmp(one) == 0 { 
		return k
	} else if (k.Cmp(two) == 0) { 
		return one
	} 

	//even case
	// f(2k) = f(k) * (2*f(k+1)-f(k))
	//odd case
	// f(2k+1) = f(k+1)**2 +f(k)**2
	// f(r) = f(r)**2 +f(k)**2

	out.Mod(k,two) 
	if out.Cmp(zero) == 0 {
		
		// r := k/2
		k.Quo(k,two)
		// out = (2 * fib(r+1) )
		out.Mul(fib(k.Add(k,one)),two)
		// out = out - fib(r)
		out.Sub(out,fib(k.Sub(k,one)))
		// out = out * fib(r) 
		out.Mul(out,fib(k))

		return out
	} else {

		// r := (k-1)/2
		k.Sub(k,one)
		k.Quo(k,two)
		// out = fib(r+1)
		out.Set(fib(k.Add(k,one)))
		// out = fib(r+1)**2
		out.Mul(out,out)
		// using zero as temporary var to square second term
		// out = (fib(r+1))**2 + fib(r)**2
		zero.Set(fib(k.Sub(k,one)))
		zero.Mul(zero,zero)
		out.Add(out, zero)

		return out
	}
}

func fib_processer(nString string) string {
	n := new(big.Int)
	n.SetString(nString,10)
	f_n := fib(n)
	fnstring:= f_n.String()
	return fnstring
}