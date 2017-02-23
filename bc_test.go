package bc_test

import (
	"testing"

	"github.com/blainsmith/baconian-cipher"
)

var text = "the quick brown fox jumps over the lazy dog"
var message = "bacon's cipher is a method of steganography created by francis bacon. this task is to implement a program for encryption and decryption of plaintext using the simple alphabet of the baconian cipher or some other kind of representation of this alphabet (make anything signify anything). the baconian alphabet may optionally be extended to encode all lower case characters individually and/or adding a few punctuation characters such as the space."
var encoded = "BacON's cIPHer Is a METhoD of stEgAnogRaphy crEatEd By FRAncis baCOn. thIs TASk Is TO imPLeMENT a proGrAm FOR eNcRYPTIOn anD deCRyPtioN Of plAINTExt UsING the SIMpLe AlPhaBet Of thE BAConIan CIphER Or sOme OTHer kInD Of reprESenTATion OF This alPHaBET (makE An"

func TestEncrypt(t *testing.T) {
	if bc.Encrypt(text, message) != encoded {
		t.Error("Encoded message is incorrect")
	}
}

func BenchmarkEncrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bc.Encrypt(text, message)
	}
}

func TestDecrypt(t *testing.T) {
	if bc.Decrypt(encoded) != text {
		t.Error("Decoded message is incorrect")
	}
}

func BenchmarkDecrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bc.Decrypt(encoded)
	}
}
