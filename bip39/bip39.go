package bip39

import "bip39_impl/bip39/worldlists"

const (
	English string = "english"
)

var (
	worldlist []string
)

func init() {
	// worldList is the set of words to use.
	worldlist = worldlists.English
}

// GetWorldlist gets the list of words to use for mnemonics.
func GetWorldlist(language string) []string {
	switch language {
	case English:
		return worldlists.English
	default:
		return worldlist
	}
}
