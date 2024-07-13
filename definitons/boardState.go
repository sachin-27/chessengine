package definitions

type BoardState struct {
	Pieces      [120]int
	Pawns       [3]int64
	KingSquares [2]int64
	SideToMove  int
	EnPassantSq int64
	FiftyMove   int
	Ply         int
	HisPly      int
	PosKey      int64
	PieceNum    [13]int
	BigPieces   [3]int
	MajorPieces [3]int
	MinorPieces [3]int
	CastlePerm  int
	History     [MaxMoves]Undo
}

type Undo struct {
	Move        int
	CastlePerm  int
	EnPassantSq int64
	FiftyMove   int
	PosKey      int64
}

func InitialiseSq120To64() {

	sq120 := BoardSize
	for i := 0; i < BoardSize; i++ {
		Square120To64[i] = sq120
	}

	sq64 := 0
	for rank := Rank1; rank <= Rank8; rank++ {
		for file := FileA; file <= FileH; file++ {
			sq120 = FileAndRank2Square(file, rank)
			Square120To64[sq120] = sq64
			Square64To120[sq64] = sq120
			sq64++
		}
	}
}

//---------------------------------- private functions --------------------------------------------------

func FileAndRank2Square(file File, rank Rank) int {
	return 21 + int(file) + int(rank)*10
}
