package RSA

import (
    "fmt"
    "math"
    "math/Big"
)

type PublicKey struct {
    modulus  *big.Int
    exponent int
}

type PrivateKey struct {
    PublicKey
    private_exp *big.Int
    primes      *[]big.Int
    
}

// yes I'm aware Go has a package for this already
func main() {
	
}
