package definitions

type BoardState struct {
	Pieces      [BoardSize]int
	Pawns       [3]uint64
	KingSquares [2]uint64
	SideToMove  int
	EnPassantSq uint64
	FiftyMove   int
	Ply         int
	HisPly      int
	PosKey      uint64
	PieceNum    [13]int
	BigPieces   [3]int
	MajorPieces [3]int
	MinorPieces [3]int
	CastlePerm  int
	History     [MaxMoves]Undo

	PieceList [13][10]int
}

type Undo struct {
	Move        int
	CastlePerm  int
	EnPassantSq uint64
	FiftyMove   int
	PosKey      uint64
}

func InitialiseSq120To64() {

	for i := 0; i < BoardSize; i++ {
		Square120To64[i] = 65
	}

	sq64 := 0
	for rank := Rank1; rank <= Rank8; rank++ {
		for file := FileA; file <= FileH; file++ {
			sq120 := FileAndRank2Square(file, rank)
			Square120To64[sq120] = sq64
			Square64To120[sq64] = sq120
			sq64++
		}
	}
}

//---------------------------------- private functions --------------------------------------------------

func FileAndRank2Square(file, rank int) int {
	return 21 + file + rank*10
}
