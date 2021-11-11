package main

import "fmt"

func main() {
	wallet := Wallet{}
	wallet.NewEntropy(16)
	fmt.Println("This is Wallet: ", wallet)
}
