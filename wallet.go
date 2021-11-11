package main

import (
	"math/rand"
)

type Wallet struct {
	Entropy []byte
}

func (w *Wallet) NewEntropy(bitSize int) ([]byte, error) {
	entropy := make([]byte, bitSize/8)
	_, _ = rand.Read(entropy) //err is always nil
	w.Entropy = entropy
	return entropy, nil
}




