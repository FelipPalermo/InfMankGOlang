package main

import (
 "fmt"
 "math/rand"
 "time"
)

// 1 : criar dicionario
// 2 : selecao do caractere
// 3 : cagar
// 4 : concatenar caracteres
// 5 : juntar e exibir resultado 
// 6 : ficar de pau duro para fury

func RN() int{
  number := rand.Intn(25)
  return number
}


func main(){
  
  rand.Seed(time.Now().UnixNano())  

  var letra [1]string
  var holder string
  var palavra string

  alfabeto := [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
  "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

  for{
  for i := 0; i < 3; i++ {
    
    holder = ""
    letra[0] = alfabeto[RN()]
    holder = holder + letra[0] 

  } 
  
    palavra = holder

  fmt.Print(palavra)
  time.Sleep(time.Second * 2)

}


}
