// sha-diff-count counts the number of bits that are different in two SHA256 hashes.
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	hash1, hash2, diff := getHashDiff("x", "X")

	fmt.Printf("Hash 1: %x\nHash 2: %x\nDiff: %d\n", hash1, hash2, diff)
}

func getHashDiff(c1, c2 string) ([32]byte, [32]byte, int) {
	hash1 := sha256.Sum256([]byte(c1))
	hash2 := sha256.Sum256([]byte(c2))

	var diff int
	for i := 0; i < 32; i++ {
		if hash1[i] != hash2[i] {
			diff++
		}
	}

	return hash1, hash2, diff
}
