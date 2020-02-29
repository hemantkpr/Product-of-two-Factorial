package main

func Factorial (n int, c chan int){
	fact := 1
	for i:=1;i<=n; i++ {
		fact = fact*i
	}
	c <- fact
}