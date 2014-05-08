package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
#########
##################
####################################

Guess the number.         ##################

####################################
###################
#########
*/






// GENERATES A RANDOM NUMBER
func xrand(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}




// INITIATES THE GAME
func init_GuessNumber(from, to int){
	myrand := xrand(1, 10);
	guessListener(myrand);
}



// LISTEN FOR GUESSES
func guessListener(rNumber int){
	var myrand int
	myrand = rNumber
	var myname string
	myname = "SomeUser"
   	var guess int
	fmt.Println("#### GUESS THE NUMBER INITIATED")
	// TODO: LISTEN FOR ChAT MESSAGES INSTEAD of console prompt
	for  {
		fmt.Println("#### Take a guess...")		
		fmt.Scanf("%d", &guess)
		
		
		if guess < myrand {
			fmt.Printf("#### The number is larger than %d ", guess)
		}
		if guess > myrand {
			fmt.Printf("#### The number is smaller than %d ", guess)
		}
		if guess == myrand {
			break
		}
	}
	// CORRECT! 
	if guess == myrand {correct(myname)}
}


// CORRECT MSG
func correct(name string){
	fmt.Printf("Good job %s! You guessed my number\n", name)
}



func main() {
	init_GuessNumber(1,10);
}



















