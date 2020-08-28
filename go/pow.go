package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func pow(text string){
	found := false
	nonce := rand.Intn(1000000)

	for found == false{
		hash := sha256.New()
		hash.Write([]byte(text+strconv.Itoa(nonce)))
		hasht := hex.EncodeToString(hash.Sum(nil))
		fmt.Println(hasht)
		if strings.HasPrefix(hasht, "002") || strings.HasPrefix(hasht, "003"){
			found = true
			fmt.Println("Hash found! ",hasht)
			fmt.Println("Text and Nonce: ",text," + ",nonce)
		}
		hash.Reset()
		nonce += 1
	}
}

func pown(){
	nonce := rand.Intn(1000000)
	count := 0

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the text you will use for proof-of-work: ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	var mode int; var nullsvar int
	fmt.Print("Enter the leading nulls the hash should have: ")
	_, err:= fmt.Scan(&nullsvar)
	nulls := makeStringNulls(nullsvar)
	fmt.Println("Which mode? 1= SHA256, 2= SHA512 3= MD5: ")
	_, err1:= fmt.Scan(&mode)
	if err != nil || err1 != nil{
		return
	}


	switch mode{
	case 1:
		for {
			hash := sha256.New()
			hash.Write([]byte(text+strconv.Itoa(nonce)))
			hasht := hex.EncodeToString(hash.Sum(nil))
			//fmt.Println(hasht)

			if strings.HasPrefix(hasht, nulls){
				fmt.Println("Hash found! ",hasht)
				fmt.Println("Text and Nonce: ",text," + ",nonce," runs: ",count)
				return
			}
			nonce++
			count++
		}

	case 2:

		for {
			hash := sha512.New()
			hash.Write([]byte(text+strconv.Itoa(nonce)))
			hasht := hex.EncodeToString(hash.Sum(nil))
			//fmt.Println(hasht)

			if strings.HasPrefix(hasht, nulls){
				fmt.Println("Hash found! ",hasht)
				fmt.Println("Text and Nonce: ",text," + ",nonce," runs: ",count)
				return
			}
			nonce++
			count++
		}
	case 3:
		for {
			hash := md5.New()
			hash.Write([]byte(text + strconv.Itoa(nonce)))
			hasht := hex.EncodeToString(hash.Sum(nil))
			//fmt.Println(hasht)
			if strings.HasPrefix(hasht, nulls) {
				fmt.Println("Hash found! ", hasht)
				fmt.Println("Text and Nonce: ", text, " + ", nonce," runs: ",count)
				return
			}
			nonce++
			count++
		}
	default:
		fmt.Println("ERROR! Wrong mode!")
		return
	}
}

func proof(hasht string, text string){
	hash := sha256.New()
	hash.Write([]byte(text))
	hashg := hex.EncodeToString(hash.Sum(nil))
	fmt.Println(hashg)
	if hasht == hashg{
		fmt.Println("Hash correct! ",hasht)
	}else{
		fmt.Println("Hash incorrect! ",hasht)
	}
}

func makeStringNulls(nulls int) (strnulls string) {
	for i := 0; i < nulls; i++ {
		strnulls += "0"
	}
	return strnulls
}
