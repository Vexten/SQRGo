package sqrgame

import (
	"fmt"
	"testing"
	obj "github.com/Vexten/SQRGo/objects"
)

func TestFirstMoves(t *testing.T) {
	inst := NewGameInstance(Small, 2, .8)
	move := inst.CurrentMove()
	state := inst.MakeMove(0,0,0)
	if (state != NextMove) {
		t.Errorf("First move failed, got: %s",state)
	}
	move = inst.CurrentMove()
	state = inst.MakeMove(int(inst.edge)-move.Width,int(inst.edge)-move.Height,0)
	if (state != NextMove) {
		t.Errorf("Second move failed, got: %s",state)
	}
}

type rectOperationsTest struct {
	Rect1 obj.Rect
	Rect2 obj.Rect
	result bool
}

var near = []rectOperationsTest {
	{obj.NewRect(0,0,3,3,0),obj.NewRect(0,3,3,3,0),true},
	{obj.NewRect(0,0,3,3,0),obj.NewRect(3,0,3,3,0),true},
	{obj.NewRect(0,0,3,3,0),obj.NewRect(3,3,3,3,0),true},
	{obj.NewRect(0,0,3,3,0),obj.NewRect(4,4,3,3,0),false},
	{obj.NewRect(0,0,3,3,0),obj.NewRect(2,2,3,3,0),false},
	{obj.NewRect(0,0,3,5,0),obj.NewRect(3,1,2,2,0),true},
	{obj.NewRect(0,0,5,3,0),obj.NewRect(1,3,1,1,0),true},
	{obj.NewRect(0,0,6,6,0),obj.NewRect(2,2,2,2,0),false},
}

func TestNear(t *testing.T) {
	var res bool
	for num, test := range near {
		res = test.Rect1.Near(&test.Rect2)
		if (res != test.result) {
			t.Errorf("Test %d failed on 1->2: expected %t, got %t.", num, test.result, res)
		}
		res = test.Rect2.Near(&test.Rect1)
		if (res != test.result) {
			t.Errorf("Test %d failed on 2->1: expected %t, got %t.", num, test.result, res)
		}
	}
}

var collision = []rectOperationsTest {
	//point collisions
	{obj.NewRect(5,5,3,3,0),obj.NewRect(3,3,3,3,0),true},
	{obj.NewRect(5,5,3,3,0),obj.NewRect(7,3,3,3,0),true},
	{obj.NewRect(5,5,3,3,0),obj.NewRect(3,7,3,3,0),true},
	{obj.NewRect(5,5,3,3,0),obj.NewRect(7,7,3,3,0),true},
	//line collisions
	{obj.NewRect(5,5,3,3,0),obj.NewRect(5,7,3,3,0),true},
	{obj.NewRect(5,5,3,3,0),obj.NewRect(7,5,3,3,0),true},
	{obj.NewRect(5,5,3,3,0),obj.NewRect(3,5,3,3,0),true},
	{obj.NewRect(5,5,3,3,0),obj.NewRect(5,3,3,3,0),true},
	//contains part
	{obj.NewRect(5,5,4,4,0),obj.NewRect(6,6,2,2,0),true},
	{obj.NewRect(5,5,4,4,0),obj.NewRect(8,6,2,2,0),true},
	{obj.NewRect(5,5,4,4,0),obj.NewRect(6,8,2,2,0),true},
	{obj.NewRect(5,5,4,4,0),obj.NewRect(4,6,2,2,0),true},
	{obj.NewRect(5,5,4,4,0),obj.NewRect(6,4,2,2,0),true},
	//touching
	{obj.NewRect(5,5,3,3,0),obj.NewRect(2,2,3,3,0),false},
	{obj.NewRect(5,5,3,3,0),obj.NewRect(8,2,3,3,0),false},
	{obj.NewRect(5,5,3,3,0),obj.NewRect(5,2,3,3,0),false},
	{obj.NewRect(5,5,3,3,0),obj.NewRect(2,8,3,3,0),false},
	{obj.NewRect(5,5,3,3,0),obj.NewRect(2,5,3,3,0),false},
	{obj.NewRect(5,5,3,3,0),obj.NewRect(8,8,3,3,0),false},
	{obj.NewRect(5,5,3,3,0),obj.NewRect(8,5,3,3,0),false},
	{obj.NewRect(5,5,3,3,0),obj.NewRect(5,8,3,3,0),false},
	//not even touching
	{obj.NewRect(6,6,1,1,0),obj.NewRect(2,2,3,3,0),false},
	{obj.NewRect(6,6,1,1,0),obj.NewRect(8,2,3,3,0),false},
	{obj.NewRect(6,6,1,1,0),obj.NewRect(5,2,3,3,0),false},
	{obj.NewRect(6,6,1,1,0),obj.NewRect(2,8,3,3,0),false},
	{obj.NewRect(6,6,1,1,0),obj.NewRect(2,5,3,3,0),false},
	{obj.NewRect(6,6,1,1,0),obj.NewRect(8,8,3,3,0),false},
	{obj.NewRect(6,6,1,1,0),obj.NewRect(8,5,3,3,0),false},
	{obj.NewRect(6,6,1,1,0),obj.NewRect(5,8,3,3,0),false},
}

func TestCollisions(t *testing.T) {
	var res bool
	for num, test := range collision {
		res = test.Rect1.CollidesWith(&test.Rect2)
		if (res != test.result) {
			t.Errorf("Test %d failed on 1->2: expected %t, got %t.", num + 1, test.result, res)
		}
		res = test.Rect2.CollidesWith(&test.Rect1)
		if (res != test.result) {
			t.Errorf("Test %d failed on 2->1: expected %t, got %t.", num + 1, test.result, res)
		}
	}
}

func TestWrong(t *testing.T) {
	inst := NewGameInstance(Small,2,.8)
	move := inst.CurrentMove()
	t.Log("Move: (" + fmt.Sprint(move.Width) + ";" + fmt.Sprint(move.Height) + ")")
	state := inst.MakeMove(-2,-2,0)
	if (state != WrongMove) {
		t.Errorf("Negative failed, got %s", state)
	}
	state = inst.MakeMove(int(inst.edge),int(inst.edge),0)
	if (state != WrongMove) {
		t.Errorf("(size;size) failed, got %s", state)
	}
	state = inst.MakeMove(int(inst.edge),0,0)
	if (state != WrongMove) {
		t.Errorf("(size;0) failed, got %s", state)
	}
	state = inst.MakeMove(0,int(inst.edge),0)
	if (state != WrongMove) {
		t.Errorf("(0;size) failed, got %s", state)
	}
	inst.rects.PushBack(obj.NewRect(0,0,6,6,0))
	inst.rects.PushBack(obj.NewRect(24,24,6,6,1))
	state = inst.MakeMove(15,15,0)
	if (state != WrongMove) {
		t.Errorf("Middle failed, got %s", state)
	}
}

func TestWrongMulti(t *testing.T) {
	for i := 0; i < 100; i++ {
		TestWrong(t)
	}
}