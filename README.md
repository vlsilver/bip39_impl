## [About BIP39 - Bitcoin Improvement proposal](https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki#Wordlist)
```go
//ENT: length of entropy
//CS: checksum
//MS: length of the generated mnemonic sentence.

var ENT int = 128 //256
CS := ENT/32
MS := (ENT+CS)/11
```