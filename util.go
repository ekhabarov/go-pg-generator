package main

import (
	"io/ioutil"
	"log"
	"strings"
)

const id = "ID"

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func snake2Camel(word string) string {
	//ID should be always uppercase
	words := strings.Split(word, "_")
	for i, w := range words {
		if strings.ToLower(w) == strings.ToLower(id) {
			words[i] = strings.ToUpper(id)
		}
	}

	return strings.Replace(strings.Title(strings.Join(words, " ")), " ", "", -1)
}

func saveToFile(f string, b []byte) {
	err := ioutil.WriteFile(f, b, 0644)
	check(err)
}
