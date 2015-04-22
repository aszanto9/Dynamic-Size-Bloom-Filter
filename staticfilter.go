/* 

 This file implements the basis for our CS51 final project. 

 Joseph Kahn    josephkahn@college.harvard.edu
 Grace Lin      glin@college.harvard.edu
 Aron Szanto    aszanto@college.harvard.edu

 */

package main 

// Including library packages referenced in this file 
import (
    //"fmt"
    "math"
)


func main() {

}

// type definition for standard bloom filter
type Bloom struct{
    // bset BitSet // will use external library for this
    //hf hash_function // to generate list of bites values h(0)…h(k). type: 
    m uint64 // size of bit array
    k uint64 // number of hash functions
    n uint64 // predicted number of elements in filter
    eps float64 // false positive error bound
}

// initialize key values using math!!
func new (num uint64, err_bound float64) Bloom {
    new_bloom Bloom
    eps := err_bound
    n := num
    m := math.Ceil(-1 * (float64(n) * math.Log(eps)) / ((math.Log(2) * math.Log(2))))
    //hf := Hash.hash_fun // will use external hash function
    k := math.Ceil((m / float64(n)) * math.Log(2))
    //bset = BitSet.new(m) // will use external bites implementation
    return new_bloom
}

// basic methods. may later include methods for changing epsilon, estimating how many elements are in filter, checking how saturated filter is, etc.

func insert (key string) Bloom
func check (key string) bool
func reset()