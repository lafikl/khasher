# khasher
KHasher is a simple utility to create k number of hashes out of two hash functions only.

Which is an implementation of the hashing functions mentioned in [this paper](https://www.eecs.harvard.edu/~michaelm/postscripts/rsa2008.pdf).

A similar algorithm was also mentioned in The Art of Computer Programing Vol.3 Sorting and Searching
In the section of double hashing open adderssing.


# Example
```go

import (
    "fmt"

    "github.com/lafikl/khasher"
)
func main() {
    kh := khasher.NewKHasher(5)
    size_of_table := uint32(67)
    hashes, err := kh.Hash([]byte("Donald E.K"), size_of_table)
    if err != nil {
        t.Fatal(err)
    }
    fmt.Println(hashes)
}
```
