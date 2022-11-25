package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)



func timeAcc(w http.ResponseWriter, response *http.Request) {
	// p := fmt.Println
	P := fmt.Fprintln
	res := response.Method

	switch res {
	case "GET":
		time := time.Now()
		timeHour := time.Hour()
		timeMin := time.Minute()
		P(w, ">", timeHour, "h", timeMin)
	}
}

func dThousand(w http.ResponseWriter, response *http.Request) {
	// p := fmt.Println
	P := fmt.Fprintln
	res := response.Method
	rand.Seed(time.Now().UnixNano())
	rand := rand.Intn(9999)
	switch res {
	case "GET":
		P(w, ">",rand)
	}

}

func dRandomFace(w http.ResponseWriter, response *http.Request) {
	P := fmt.Fprintln
	rand.Seed(time.Now().UnixNano())
	dArray := []int{}
	res := response.Method

	two := rand.Intn(2)
	sqr := rand.Intn(4)
	dsix := rand.Intn(6)
	two_sqr := rand.Intn(8)
	dten := rand.Intn(10)
	dtwelve := rand.Intn(12)
	doubleTen := rand.Intn(20)
	hundr := rand.Intn(100)

	dArray = append(dArray, two,sqr,dsix,two_sqr,dten,dtwelve,doubleTen,hundr)
	rand.Shuffle(len(dArray),func(i, j int) { dArray[i], dArray[j] = dArray[j], dArray[i] })
	switch res {
	case "GET":
		P(w,">",dArray[0],dArray[1],dArray[2],dArray[3],dArray[4],dArray[5],dArray[6],dArray[7])
	}

}

func randWord (w http.ResponseWriter, response *http.Request){
	
}

func main() {
	http.HandleFunc("/", timeAcc)       // affichant l’heure qu’il est.
	http.HandleFunc("/dice", dThousand) // affichant le résultat d’un dé à 1000 faces
	http.HandleFunc("/dices", dRandomFace) // affichant quinze dés aux nombres de faces aléatoires
	http.HandleFunc("/randomize-words", randWord)
	// http.HandleFunc("/semi-capitalize-sentence", randWord)
	http.ListenAndServe(":4567", nil)
}
