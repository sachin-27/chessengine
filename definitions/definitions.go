package definitions

const (
	MaxMoves  = 2048
	BoardSize = 120
)

type Piece int

const (
	Empty Piece = iota
	wP          = iota
	wN          = iota
	wB          = iota
	wR          = iota
	wQ          = iota
	wK          = iota
	bP          = iota
	bN          = iota
	bB          = iota
	bR          = iota
	bQ          = iota
	bK          = iota
)

type File int

const (
	FileA    File = iota
	FileB         = iota
	FileC         = iota
	FileD         = iota
	FileE         = iota
	FileF         = iota
	FileG         = iota
	FileH         = iota
	FileNone      = iota
)

type Rank int

const (
	Rank1    Rank = iota
	Rank2         = iota
	Rank3         = iota
	Rank4         = iota
	Rank5         = iota
	Rank6         = iota
	Rank7         = iota
	Rank8         = iota
	RankNone      = iota
)

type Color int

const (
	White Color = iota
	Black       = iota
	Both        = iota
)

type Square int

const (
	A1, B1, C1, D1, E1, F1, G1, H1 = iota + 21, iota + 22, iota + 23, iota + 24, iota + 25, iota + 26, iota + 27, iota + 28
	A2, B2, C2, D2, E2, F2, G2, H2 = iota + 30, iota + 31, iota + 32, iota + 33, iota + 34, iota + 35, iota + 36, iota + 37
	A3, B3, C3, D3, E3, F3, G3, H3 = iota + 39, iota + 40, iota + 41, iota + 42, iota + 43, iota + 44, iota + 45, iota + 46
	A4, B4, C4, D4, E4, F4, G4, H4 = iota + 48, iota + 49, iota + 50, iota + 51, iota + 52, iota + 53, iota + 54, iota + 55
	A5, B5, C5, D5, E5, F5, G5, H5 = iota + 57, iota + 58, iota + 59, iota + 60, iota + 61, iota + 62, iota + 63, iota + 64
	A6, B6, C6, D6, E6, F6, G6, H6 = iota + 66, iota + 67, iota + 68, iota + 69, iota + 70, iota + 71, iota + 72, iota + 73
	A7, B7, C7, D7, E7, F7, G7, H7 = iota + 75, iota + 76, iota + 77, iota + 78, iota + 79, iota + 80, iota + 81, iota + 82
	A8, B8, C8, D8, E8, F8, G8, H8 = iota + 84, iota + 85, iota + 86, iota + 87, iota + 88, iota + 89, iota + 90, iota + 91
	NoSquare                       = iota + 91
)

type Castling int

const (
	WhiteKingCastle  Castling = iota + 1
	WhiteQueenCastle          = 2
	BlackKingCastle           = 4
	BlackQueenCastle          = 8
)

var Square120To64 = [BoardSize]int{}
var Square64To120 = [64]int{}
