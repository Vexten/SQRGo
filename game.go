/*
Package sqrgame contains structs and functions that
define a game of "Squares"
*/
package sqrgame

import (
	"container/list"
	"math/rand"
	"time"
	obj "github.com/Vexten/SQRGo/objects"
)

var sharedRand rand.Rand = *rand.New(rand.NewSource(time.Now().UnixNano()))

//Avaliable board sizes
type BoardSize byte

const (
	Small BoardSize = 30
	Medium BoardSize = 50
	Large BoardSize = 70
	ExtraLarge BoardSize = 90
)

//go:generate stringer -type=BoardSize

//State of the game after an attemped Move
type GameState byte

const (
	NextMove GameState = iota
	WrongMove
	GameEnd
)

//go:generate stringer -type=GameState

//Current move.
//	Player - player who must make a move
//	Width, Height - dimensions of current rect
type Move struct {
	Player byte
	Width int
	Height int
}

//Game results
type Stats struct {
	Winner byte
	Scores *[]int
}

/*
A single game instance.
Only one routine shoud interact with a single instance at a time
*/
type GameInstance struct {
	edge BoardSize
	boardArea int
	rects *list.List
	players byte
	currPlayer byte
	currDim1 int
	currDim2 int
	random *rand.Rand
}

//Set current rect dims as two 6-sided die rolls
func (game *GameInstance) generateMove() {
	game.currDim1 = game.random.Intn(6) + 1
	game.currDim2 = game.random.Intn(6) + 1
}

//Set current player num to next player in order
func (game *GameInstance) nextPlayer() {
	game.currPlayer++
	if (game.currPlayer == game.players) {
		game.currPlayer = 0
	}
}

//Perform move validity checks and add a rect to collection
func (game *GameInstance) addRect(x int, y int, width int, height int, player byte) bool {
	//check field bounds
	if (x + width > int(game.edge) || y + height > int(game.edge)) {
		return false
	}
	if (x < 0 || y < 0) {
		return false
	}
	newRect := obj.NewRect(x, y, width, height, player)
	near := false
	iter := game.rects.Front()
	for iter != nil {
		rect := iter.Value.(obj.Rect)
		diff := rect.Start().Diff(*newRect.Start())
		//only check rects that are no more than 6 units (max side len) away
		if (diff.X() > -7 && diff.Y() > -7 && diff.X() < 7 && diff.Y() < 7) {
			if (rect.CollidesWith(&newRect) || newRect.CollidesWith(&rect)) {
				return false
			}
			if (rect.Player() == player && !near)	{
				near = rect.Near(&newRect)
			}
		}
		iter = iter.Next()
	}
	//also check for first turns
	if (!near && !(game.rects.Len() < int(game.players))) {
		return false
	}
	game.rects.PushBack(newRect)
	return true
}

//Check if borad is filled
func (game *GameInstance) boardComplete() bool {
	currArea := 0
	iter := game.rects.Front()
	for iter != nil {
		currRect := iter.Value.(obj.Rect)
		currArea += currRect.Area()
		iter = iter.Next()
	}
	return (currArea > game.boardArea);
}

//Create a GameInstance with seeded generator.
//endPercentage defines how full should the board be to count as filled.
func NewGameInstanceSeeded(size BoardSize, players byte, endPercentage float32, seed int64) *GameInstance {
	a := GameInstance{}
	a.edge = size
	a.boardArea = int(float32(int(size) * int(size)) * endPercentage)
	a.rects = list.New()
	a.players = players
	a.currPlayer = 0
	a.random = rand.New(rand.NewSource(seed))
	a.generateMove()
	return &a
}

//Create a GameInstance with a random seed.
//endPercentage defines how full should the board be to count as filled.
func NewGameInstance(size BoardSize, players byte, endPercentage float32) *GameInstance {
	return NewGameInstanceSeeded(size, players, endPercentage, sharedRand.Int63())
}

//Returns current move
func (game *GameInstance) CurrentMove() Move {
	return Move{game.currPlayer, game.currDim1, game.currDim2}
}

//Try to make a move as the current player
//	returns GameEnd if board is filled
//	returns Cheating if provided rect dimensions differ from stored inside instance
func (game *GameInstance) MakeMove(x int, y int, rotation byte) GameState {
	width := game.currDim1
	height := game.currDim2
	if (rotation == 1) {
		width = game.currDim2
		height = game.currDim1
	}
	if (game.addRect(x, y, width, height, game.currPlayer)) {
		if (game.boardComplete()) {
			return GameEnd
		}
		game.generateMove()
		game.nextPlayer()
		return NextMove
	}
	return WrongMove
}

//Generates new move and increments current player
func (game *GameInstance) SkipMove() {
	game.generateMove()
	game.nextPlayer()
}

//Calculates points and determines a winner
func (game *GameInstance) End() Stats {
	stat := Stats{}
	results := make([]int, game.players)
	max := -1
	iter := game.rects.Front()
	for iter != nil {
		currRect := iter.Value.(obj.Rect)
		results[currRect.Player()] += currRect.Area()
	}
	for i := 0; i < len(results); i++ {
		if (results[i] > max) {
			max = results[i];
			stat.Winner = byte(i)
		}
	}
	stat.Scores = &results
	return stat
}

//Clears the board, resets current player and generates a new move
func (game *GameInstance) Reset() {
	game.rects.Init()
	game.currPlayer = 0
	game.generateMove()
}