package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	wordcount := make(map[string]int)
	myString := strings.Fields(s)
	for _,word2:= range myString {
		word:=string(word2)
		//fmt.Println(word)
		i,y:= wordcount[word]
		if y == false {
			wordcount[word]=1
		} else {
			wordcount[word]=i+1
		}
		
	}
	return wordcount //map[string]int{"x": 1}
}

func main() {
	wc.Test(WordCount)
}
