package main

import (
    "fmt"
    "math"
)

//TODO: test rigorously
//      implement bloom filter to facilitate multipattern searching
func rabinKarpSearch(pattern, text []string, radix, prime int) {
    textLength := len(text)
    patternLength := len(pattern)

    var patternHash, textHash float64 = 0, 0//hash values for pattern and text

    //high order digit position of pattern length digit window
    highOrder := int(math.Pow(float64(radix), float64(patternLength -1))) % prime
	
    //pattern and text hash values computed
    for ii := 0; ii <= textLength; ii++ {
        patternHash = float64(radix) * patternHash + int(pattern[ii]) % float64(prime)
        textHash = (float64(radix) * textHash + text[ii]) % prime
    }
        
    for ii := 0; ii < textLength - patternLength; ii++ {
        if patternHash == textHash {//hash patterns match, check char by char
            for jj := 0; jj < patternLength; jj++ {
                if pattern[ii] == text[ii+jj] {
                    //check each pattern char against sequential text segment of same length
                    fmt.Print("Pattern occurs with shift %d", ii)
                }
            }

            //not a match, move along
            if ii < textLength - patternLength {
                text[ii+1] = (radix * (textHash - int(text[ii]) * highOrder) + text[ii + patternLength]) % prime
            }
        }
    }
}

func main() {
    text := []string{"Dummy text to test pattern matching algorithm against"}
    pattern := []string{"te"}

    //radix == |Σ| i.e. the radix notation currently used
    var radix, prime int = 256, 3 
    
    rabinKarpSearch(pattern, text, radix, prime)
}
