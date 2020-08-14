package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main(){
   pass := "Admin"
   newpass,_ := bcrypt.GenerateFromPassword([]byte(pass), 10)
   fmt.Println(newpass)

   decrypt := bcrypt.CompareHashAndPassword(newpass, []byte("admin"))
   if decrypt != nil{
	   fmt.Println("Invali Password")
   }else{
  
	fmt.Println("valid Password")
   }
}