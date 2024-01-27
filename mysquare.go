package main

import "fmt"

func primeChecker(num int) {
	isPrime := true
	if num == 2{
		isPrime = true
	}else if num < 2 || num % 2 == 0 {
		isPrime = false
	} else {
		for i := 2; i < int(num/2)+1; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}
	}
	if isPrime {
		fmt.Printf(">>>>>>\t%d is a prime number\n",num)
	} else {
		fmt.Printf(">>>>>>>\t%d is not a prime number\n",num)
	}

}

func main() {
	fmt.Println("\t\t.........Press Ctrl-C to quit .........\n\n")
	k:= 0
	for k < 1{
	         var number int
	         fmt.Print("Enter number to check if its a prime number : ")
	         fmt.Scanln(&number)
			 primeChecker(number)
	}
}
