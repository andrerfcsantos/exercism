package chessboard

// File stores if a square is occupied by a piece
type File []bool

// Chessboard contains eight Files, accessed with values from 'A' to 'H'
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given rank
func CountInFile(cb Chessboard, file string) int {
	ret := 0
	for _, occupied := range cb[file] {
		if occupied {
			ret++
		}
	}

	return ret
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given file
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}

	ret := 0

	for _, file := range cb {
		if rank <= len(file) && file[rank-1] {
			ret++
		}
	}
	return ret
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	ret := 0
	for _, rank := range cb {
		ret += len(rank)
	}
	return ret
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	ret := 0
	for file := range cb {
		ret += CountInFile(cb, file)
	}
	return ret
}
