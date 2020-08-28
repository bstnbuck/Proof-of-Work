package main

import (
	"fmt"
)

func main(){
	starti := false

	for starti == false{
		var startinput int
		fmt.Println("Start Proof of Work? 1= Yes, 2= Yes, but simple with text 3= Proof generated pow hash 9= No: ")
		_, err:= fmt.Scan(&startinput)
		if err != nil{
			return
		}
		switch startinput{
		case 1:
			pown()

		case 2:
			var text string
			fmt.Print("Enter the text you will use for proof-of-work: ")
			_, err:= fmt.Scan(&text)
			if err != nil{
				return
			}
			pow(text)

		case 3:
			var hasht, text string
			fmt.Print("Enter the hash: ")
			_, err:= fmt.Scan(&hasht)
			fmt.Print("Enter the text with the nonce: ")
			_, err1:= fmt.Scan(&text)
			if err != nil || err1 != nil{
				return
			}
			proof(hasht, text)

		case 9:
			fmt.Println("Stopped")
			starti = true
		}

	}
}