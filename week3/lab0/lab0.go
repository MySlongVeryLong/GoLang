package main

import (

   "fmt"
   "crypto/md5"
   //"encoding/hex"
   "crypto/sha1"
   "crypto/sha256"
   "crypto/sha512"
   "crypto/sha3"
)

func md5_compare(a string, b string){

  sum1 := md5.Sum([]byte(a))
  //fmt.Println(hex.EncodeToString(sum1[:]))
  sum2 := md5.Sum([]byte(b))
  //fmt.Println(hex.EncodeToString(sum2[:]))
  if sum1 == sum2 {
     fmt.Println("Hash (MD5): Output A / output B => Match!")
  } else {
     fmt.Println("Hash (MD5): Output A / output B => No Match!")
  }

}

func sha1_compare(a string, b string){

  sum1 := sha1.Sum([]byte(a))
  //fmt.Println(hex.EncodeToString(sum1[:]))
  sum2 := sha1.Sum([]byte(b))
  //fmt.Println(hex.EncodeToString(sum2[:]))
  if sum1 == sum2 {
     fmt.Println("Hash (SHA1): Output A / output B => Match!")
  } else {
     fmt.Println("Hash (SHA1): Output A / output B => No Match!")
  }

}

func sha256_compare(a string, b string){

  sum1 := sha256.Sum256([]byte(a))
  //fmt.Println(hex.EncodeToString(sum1[:]))
  sum2 := sha256.Sum256([]byte(b))
  //fmt.Println(hex.EncodeToString(sum2[:]))
  if sum1 == sum2 {
     fmt.Println("Hash (SHA256): Output A / output B => Match!")
  } else {
     fmt.Println("Hash (SHA256): Output A / output B => No Match!")
  }

}

func sha512_compare(a string, b string){

  sum1 := sha512.Sum512([]byte(a))
  //fmt.Println(hex.EncodeToString(sum1[:]))
  sum2 := sha512.Sum512([]byte(b))
  //fmt.Println(hex.EncodeToString(sum2[:]))
  if sum1 == sum2 {
     fmt.Println("Hash (SHA512): Output A / output B => Match!")
  } else {
     fmt.Println("Hash (SHA512): Output A / output B => No Match!")
  }

}

func sha3_compare(a string, b string){

  sum1 := sha3.Sum384([]byte(a))
  //fmt.Println(hex.EncodeToString(sum1[:]))
  sum2 := sha3.Sum384([]byte(b))
  //fmt.Println(hex.EncodeToString(sum2[:]))
  if sum1 == sum2 {
     fmt.Println("Hash (SHA3): Output A / output B => Match!")
  } else {
     fmt.Println("Hash (SHA3): Output A / output B => No Match!")
  }

}

func main(){

 fmt.Printf("======Name + Hashing Program======\n")

 var str1,str2 string

 fmt.Print("Enter the first string: ") 
 fmt.Scan(&str1)
 fmt.Print("Enter the second string: ")
 fmt.Scan(&str2)
 
 md5_compare(str1,str2)
 sha1_compare(str1,str2)
 sha256_compare(str1,str2)
 sha512_compare(str1,str2)
 sha3_compare(str1,str2)

}
