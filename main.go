package main

import "fmt"

users :=make(map[string]int){
Altoha:= 48
temik:=87
nurzhan:=63
}
users, ok := users["temik"]
if ok {
fmt.Println("найдено число:" users)
}else{
fmt.Println("в мапе нет")

delete(users, temik)
users[aruzhan] = 51