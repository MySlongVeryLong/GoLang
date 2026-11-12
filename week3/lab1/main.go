package main

import (
  
   "fmt"
   "log"
   "jeffbezo/utils/crack"
)


func main(){

  md5_hash := "6a85dfd77d9cb35770c9dc6728d73d3f"
  wordlist := "../../passwords/nord_vpn.txt"

  fmt.Printf("Attempting to crack hash %s...\n", md5_hash)

  err := crack.CrackMD5(md5_hash, wordlist)
  if err != nil {
     log.Fatal(err)
  }
  
}
