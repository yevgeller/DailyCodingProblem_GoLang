package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	row, exists := cb[file]
	if exists == false {
		return 0
	}

	occupied := 0
	for _, x := range row {
		if x == true {
			occupied += 1
		}
	}
	return occupied
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank <= 0 || rank > 8 {
		return 0
	}
	var occupied int
	for x := range cb {
		if cb[x][rank-1] {
			occupied += 1
		}
	}
	return occupied
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var all int
	for x := range cb {
		for range cb[x] {
			all++
		}
	}
	return all
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	panic("Please implement CountOccupied()")

	var occupied int
	for x := range cb {
		for _, y := range cb[x] {
			if y {
				occupied++
			}
		}
	}
	return occupied

}
