package main

import "fmt"

type Game struct {
	Gameboard [][]Tile
	Tiles     []Tile
}

func getRandomWallAttribute() WallAttribute {
	num := getRandomNumber(0, 9)
	switch {
	case num <= 3:
		return Window
	case num >= 6:
		return Door
	default:
		return Wall
	}
}

func newGame() Game {
	tiles := make([]Tile, 0)
	for i := 0; i < 44; i++ {

		tiles = append(tiles,
			Tile{RoomName: fmt.Sprintf("Room %d", i),
				Top:    getRandomWallAttribute(),
				Bottom: getRandomWallAttribute(),
				Left:   getRandomWallAttribute(),
				Right:  getRandomWallAttribute()})
	}

	return Game{Tiles: tiles,
		Gameboard: [][]Tile{
			{tiles[0], tiles[1], tiles[2]},
			{tiles[3], tiles[4], tiles[5]},
			{tiles[6], tiles[7], tiles[8]}}}
}

type TilePlacement string

const (
	Top    = TilePlacement("top")
	Bottom = TilePlacement("bottom")
	Right  = TilePlacement("right")
	Left   = TilePlacement("left")
)

func (g Game) showBoard() {
	// number of rows
	n := len(g.Gameboard)

	// number of columns
	m := len(g.Gameboard[0])
	for i := 0; i < n; i++ {
		//blank space
		for j := 0; j < m; j++ {
			fmt.Printf("-------")
		}

		for j := 0; j < m; j++ {
			fmt.Printf("----%s----", g.Gameboard[i][j].Top.toTopBottomDisplay())
		}
		fmt.Printf("\n")
		for j := 0; j < m+2; j++ {
			fmt.Printf("|        |")
		}

		fmt.Printf("\n")
		fmt.Printf("|        |")
		for j := 0; j < m; j++ {
			fmt.Printf("%s %s %s", g.Gameboard[i][j].Left.toSideDisplay(), g.Gameboard[i][j].RoomName, g.Gameboard[i][j].Right.toSideDisplay())
		}
		fmt.Printf("|        |")
		fmt.Printf("\n")
		for j := 0; j < m+2; j++ {
			fmt.Printf("|        |")
		}

		fmt.Printf("\n")
		fmt.Printf("----------")
		for j := 0; j < m; j++ {
			fmt.Printf("----%s----", g.Gameboard[i][j].Bottom.toTopBottomDisplay())
		}
		fmt.Printf("----------")

		fmt.Printf("\n")
	}
}

//func (g Game) placeTile(placement TilePlacement) Game {
//
//}
