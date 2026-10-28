package main

import (
  
	"fmt"
	"encoding/hex"

)


// source code: https://wgallagher86.medium.com/implementing-and-breaking-repeating-key-xor-in-go-acb561958430

func xorEncrypt(text, key []byte) []byte {

	eb := make([]byte, len(text))

	for i := 0; i< len(text); i++ {

     eb[i] = text[i] ^ key[i%len(key)]

	}

	return eb

}


func main(){

 var plaintext string

 fmt.Print("Enter string: ")
 fmt.Scan(&plaintext)

 res := xorEncrypt([]byte(plaintext), []byte("ICE"))

 ret := make([]byte, hex.EncodedLen(len(res)))
 hex.Encode(ret, res)

 fmt.Printf("%s \n", ret)

}
