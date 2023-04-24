package diffsquares

func SquareOfSum(n int) int {
	//panic("Please implement the SquareOfSum function")
    return (n*(n+1)/2)*(n*(n+1)/2)
}

func SumOfSquares(n int) int {
	//panic("Please implement the SumOfSquares function")
    return (n*(n+1)*(2*n+1))/6
}

func Difference(n int) int {
	//panic("Please implement the Difference function")
    return SquareOfSum(n) - SumOfSquares(n)
}
