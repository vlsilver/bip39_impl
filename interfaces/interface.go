package interfaces

type WalletInterface interface {
	// NewEntropy will create random entropy bytes.
	NewEntropy(bitSize int) ([]byte, error)
}

