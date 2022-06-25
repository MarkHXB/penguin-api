package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	result := getSum(r)
	fmt.Fprintf(w, "Hello Gyula, az eredményed: %v", result)
}

func getSum(r *http.Request) int{
	params := r.URL.Query()

	a := params.Get("a")
	b := params.Get("b")

	if a == ""{
		a = "0"
	}
	if b == ""{
		b = "0"
	}

	aNumber, err := strconv.Atoi(a)
	if err != nil{
		fmt.Println(err)
	}
	bNumber, err := strconv.Atoi(b)
	if err != nil{
		fmt.Println(err)
	}
	return aNumber + bNumber
}
