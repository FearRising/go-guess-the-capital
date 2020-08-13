package main

import (
	"fmt"
	//"math"
	"time"
	"strings"
)
func sleep(secs time.Duration){
	time.Sleep(secs * time.Second)
}
func clear(){
	fmt.Println("\033[H\033[2J")
}
func topBar(q int){
	fmt.Printf("##########  %d  ##########\n\n", q)
}
func displayScore(name string, score int, total int){
	clear()
	fmt.Printf("\n\n\n\033[1m%s\033[0m scored \033[1m%d\033[0m  point(s) out of \033[1m%d\033[0m!\n\n\n \033[1mThank you for playing\n\nMade By Simon Kellet on 13/08/20\033[0m\n", name, score, total)
}

func main(){
	var answer string
	var name string
	var question int
	var score int
	score = 0
		//create a map
	countryCapitalsMap := map[string] string{
		"France":"Paris",
		"Italy":"Rome",
		"Japan":"Tokyo",
		"United Kingdom":"London",
		"United States":"WashingtonDC",
		"Turkey":"Ankara",
		"Greece":"Athens",
		"Netherlands":"Amsterdam",
		"Iraq":"Baghdad",
		"Switzerland":"Bern",
		"Christmas Island":"FlyingFishCove",
		"Gibraltar":"Gibraltar",
		"Palestine":"Jerusalem",
		"Portugal":"Lisbon",
		"Russia":"Moscow",
		"Spain":"Madrid",
		"Germany":"Berlin",
		"Australia":"Sydney",
		"Scotland":"Edinburgh",
	}


	clear()
	fmt.Println("Please Enter your first name:\n")
	fmt.Scanln(&name)
	clear()
	sleep(2)
	fmt.Printf("Nice to meet you \033[1m%s\033[0m!\n", name)
	sleep(2)
	fmt.Printf("Lets do a quiz together! (%d Qs)\n", len(countryCapitalsMap))
	sleep(3)
	fmt.Println("Please type your answer with NO SPACES e.g. NewYork, RebublicOfIreland")
	sleep(4)
	fmt.Println("OK! Let's Go!")
	sleep(3)

	question = 1
	if question > len(countryCapitalsMap){
		displayScore(name, score, len(countryCapitalsMap))
	} else {
		for country, capital := range countryCapitalsMap{
			clear()
			topBar(question)
			fmt.Printf("Question %d:\n", question,)
			fmt.Printf("What is the Capital of %s?\n", country)
			fmt.Scanln(&answer)
			if strings.ToLower(answer) == strings.ToLower(countryCapitalsMap[country]){
				fmt.Printf("%s\n", answer)
				fmt.Printf("You got it correct!\n\n",)
				score++
				question++
				sleep(1)
			} else {
				fmt.Printf("%s\n", answer)
				fmt.Printf("\nIncorrect!\nIt was: \033[1m%s\033[0m\n\n", capital)
				question++
				sleep(3)
			}
		}
		displayScore(name, score, len(countryCapitalsMap))
	}
}
