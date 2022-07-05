/*
Package sqrgame contains structs and functions that
define a game of "Squares"
*/
package sqrgame

import (
	"math/rand"
	"time"
	"container/list"
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

//State of the game after an attemped Move
type GameState byte

const (
	NextMove GameState = iota
	WrongMove
	GameEnd
	Cheating
)

//Current move.
//	Player - player who must make a move
//	Dim1, Dim2 - dimensions of current rect
type Move struct {
	Player byte
	Dim1 int
	Dim2 int
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
func (this *GameInstance) generateMove() {
	this.currDim1 = this.random.Intn(7)
	this.currDim2 = this.random.Intn(7)
}

//Set current player num to next player in order
func (this *GameInstance) nextPlayer() {
	this.currPlayer++
	if (this.currPlayer == this.players) {
		this.currPlayer = 0
	}
}

//Perform move validity checks and add a rect to collection
func (this *GameInstance) addRect(x int, y int, width int, height int, player byte) bool {
	//check field bounds
	if (x + width > int(this.edge) || y + height > int(this.edge)) {
		return false
	}
	newRect := obj.NewRect(x, y, width, height, player)
	near := false
	iter := this.rects.Front()
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
	if (!near && !(this.rects.Len() < int(this.players))) {
		return false
	}
	this.rects.PushBack(newRect)
	return true
}

//Check if borad is filled
func (this *GameInstance) boardComplete() bool {
	currArea := 0
	iter := this.rects.Front()
	for iter != nil {
		currRect := iter.Value.(obj.Rect)
		currArea += currRect.Area()
	}
	return (currArea > this.boardArea);
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
func (this *GameInstance) CurrentMove() Move {
	return Move{this.currPlayer, this.currDim1, this.currDim2}
}

//Try to make a move as the current player
//	returns GameEnd if board is filled
//	returns Cheating if provided rect dimensions differ from stored inside instance
func (this *GameInstance) MakeMove(x int, y int, width int, heignt int) GameState {
	if (width != this.currDim1) {
		if (width != this.currDim2 || heignt != this.currDim1) {
			return Cheating
		}
	} else {
		if (heignt != this.currDim2) {
			return Cheating
		}
	}
	if (this.addRect(x, y, width, heignt, this.currPlayer)) {
		if (this.boardComplete()) {
			return GameEnd
		}
		this.generateMove()
		this.nextPlayer()
		return NextMove
	}
	return WrongMove
}

//Generates new move and increments current player
func (this *GameInstance) SkipMove() {
	this.generateMove()
	this.nextPlayer()
}

//Calculates points and determines a winner
func (this *GameInstance) End() Stats {
	stat := Stats{}
	results := make([]int, this.players)
	max := -1
	iter := this.rects.Front()
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
func (this *GameInstance) Reset() {
	this.rects.Init()
	this.currPlayer = 0
	this.generateMove()
}