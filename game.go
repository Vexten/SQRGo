package sqrgame

import (
	"math/rand"
	"time"
	"container/list"
	obj "github.com/Vexten/SQRGo/objects"
)

var sharedRand rand.Rand = *rand.New(rand.NewSource(time.Now().UnixNano()))

type BoardSize byte

const (
	Small BoardSize = 30
	Medium BoardSize = 50
	Large BoardSize = 70
	ExtraLarge BoardSize = 90
)

type GameState byte

const (
	NextMove GameState = iota
	WrongMove
	GameEnd
	Cheating
)

type Move struct {
	Player byte
	Dim1 int
	Dim2 int
}

type Stats struct {
	Winner byte
	Scores *[]int
}

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

func (this *GameInstance) generateMove() {
	this.currDim1 = this.random.Intn(7)
	this.currDim2 = this.random.Intn(7)
}

func (this *GameInstance) nextPlayer() {
	this.currPlayer++
	if (this.currPlayer == this.players) {
		this.currPlayer = 0
	}
}

/*
TODO:
first turn logic
*/
func (this *GameInstance) addRect(x int, y int, width int, height int, player byte) bool {
	if (x + width > int(this.edge) || y + height > int(this.edge)) {
		return false
	}
	near := false
	newRect := obj.NewRect(x, y, width, height, player)
	iter := this.rects.Front()
	for iter != nil {
		rect := iter.Value.(obj.Rect)
		diff := rect.Start().Diff(*newRect.Start())
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
	if (!near) {
		return false
	}
	this.rects.PushBack(newRect)
	return true
}

func (this *GameInstance) boardComplete() bool {
	currArea := 0
	iter := this.rects.Front()
	for iter != nil {
		currRect := iter.Value.(obj.Rect)
		currArea += currRect.Area()
	}
	return (currArea > this.boardArea);
}

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

func NewGameInstance(size BoardSize, players byte, endPercentage float32) *GameInstance {
	return NewGameInstanceSeeded(size, players, endPercentage, sharedRand.Int63())
}

func (this *GameInstance) CurrentMove() Move {
	return Move{this.currPlayer, this.currDim1, this.currDim2}
}

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

func (this *GameInstance) SkipMove() {
	this.generateMove()
	this.nextPlayer()
}

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

func (this *GameInstance) Reset() {
	this.rects.Init()
	this.currPlayer = 0
	this.generateMove()
}