package main

import (
	"fmt"
	"net/http"
	"github.com/yosssi/ace"
)

func main(){


	fmt.Println("hello")

	http.HandleFunc("/",f)

	http.ListenAndServe(":8000",nil)
}

func f(w http.ResponseWriter, r *http.Request) {

	tpl, err := ace.Load("base", "inner", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tpl.Execute(w, map[string]string{"Msg": "Hello Ace"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}