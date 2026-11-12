package main

import (

   "fmt"
   "log"
   "jeffbezo/utils/crack"
)


func main(){

  sha1_hash := "485f5c36c6f8474f53a3b0e361369ee3e32ee0de2f368b87b847dd23cb284b316bb0f026ada27df76c31ae8fa8696708d14b4d8fa352dbd8a31991b90ca5dd38"
  wordlist := "../../passwords/nord_vpn.txt"

  fmt.Printf("Attempting to crack hash %s...\n", sha1_hash)

  err := crack.CrackSHA512(sha1_hash, wordlist)
  if err != nil {
     log.Fatal(err)
  }

}
