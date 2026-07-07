package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Move struct {
	Color int
	Col   int
}

type Board struct {
	BoardString string
}

type place struct {
	occupied bool
	Color    int
}

type ConnectGame int

var gameBoard [][]int

func (t *ConnectGame) Move(args *Move, reply *int) error {

	for i := range gameBoard {
		if gameBoard[7-i][args.Col] == -1 {
			gameBoard[7-i][args.Col] = args.Color
		}
	}
	return nil
}

func (t *ConnectGame) Get(args *int, reply *Board) error {
	reply.BoardString = "Hello World"
	return nil
}

// 6 rows 7 columns
func main() {
	gameBoard = make([][]int, 6) // Allocates the outer slice

	for i := range gameBoard {
		gameBoard[i] = make([]int, 7) // Allocates each inner row
	}

	for i := range gameBoard {
		for j := range gameBoard[i] {
			gameBoard[i][j] = -1
		}
	}

	cg := new(ConnectGame)
	rpc.Register(cg)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	log.Println("Serving on PORT 1234")
	http.Serve(l, nil)
}
