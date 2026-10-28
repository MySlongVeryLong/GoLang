package main

import (
 
  "encoding/base64"
  "encoding/hex"
  "fmt"
  "os"
  "bufio"
)

func main() {
  
  fmt.Print("Enter your string: ")
  reader := bufio.NewReader(os.Stdin)
  
  input, _ := reader.ReadString('\n')
 
  fmt.Printf("String: %v", input)
  
  myString := []byte(input)

  // binary encoding
  
  binaryStr := fmt.Sprintf("%08b", myString) 
  fmt.Println("Binary =", binaryStr)
  
  // hexadecimal encoding

  base16Str := hex.EncodeToString(myString)
  fmt.Println("Hexadecimal =", base16Str)

  // base64 encoding
  encoded := base64.StdEncoding.EncodeToString(myString)

  fmt.Println("Base64 =", encoded)


}
