package khasher

import "hash/fnv"

// KHasher is an implementation of the hashing functions mentioned in the paper
// https://www.eecs.harvard.edu/~michaelm/postscripts/rsa2008.pdf
// A similar algorithm was mentioned in The Art of Computer Programing Vol.3 Sorting and Searching
// In the section of double hashing open adderssing.
type KHasher struct {
	k int
}

// NewKHasher creates a new instance of the hasher
func NewKHasher(k int) *KHasher {
	if k < 1 {
		return &KHasher{1}
	}
	return &KHasher{k}
}

// Hash returns K number of hash values of the given key
// m is the size of the hash table
func (k *KHasher) Hash(key []byte, m uint32) ([]uint32, error) {
	h := fnv.New32()
	_, err := h.Write(key)
	if err != nil {
		return nil, err
	}
	ha := h.Sum32()
	hb := murmur3(key)
	hashes := []uint32{}
	for i := 0; i < k.k; i++ {
		//  (a + b * i ) % m
		hashes = append(hashes, (ha+hb*uint32(i))%m)
	}
	return hashes, nil
}
