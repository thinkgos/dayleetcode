package main

func main() {
	if isPowerOfTwo(100) {
		panic("not expect")
	}

	if !isPowerOfTwo(1024) {
		panic("not expect")
	}
}

func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}
