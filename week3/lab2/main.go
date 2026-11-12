package main

import (

   "fmt"
   "log"
   "jeffbezo/utils/crack"
)


func main(){

  sha1_hash := "aa1c7d931cf140bb35a5a16adeb83a551649c3b9"
  wordlist := "../../passwords/nord_vpn.txt"

  fmt.Printf("Attempting to crack hash %s...\n", sha1_hash)

  err := crack.CrackSHA1(sha1_hash, wordlist)
  if err != nil {
     log.Fatal(err)
  }

}
