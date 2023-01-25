package main

type Player struct {
	Speed         int
	Might         int
	Knowledge     int
	Sanity        int
	PlayerName    string
	CharacterName string
	Objects       []Object
}

func getAllCharacters() []Player {
	Rob := Player{
		Speed:         4,
		Might:         4,
		Knowledge:     4,
		Sanity:        4,
		CharacterName: "Rob",
		Objects:       make([]Object, 0),
	}
	Bree := Player{
		Speed:         4,
		Might:         4,
		Knowledge:     4,
		Sanity:        4,
		CharacterName: "Bree",
		Objects:       make([]Object, 0),
	}
	Kyle := Player{
		Speed:         4,
		Might:         4,
		Knowledge:     4,
		Sanity:        4,
		CharacterName: "Kyle",
		Objects:       make([]Object, 0),
	}
	Nate := Player{
		Speed:         4,
		Might:         4,
		Knowledge:     4,
		Sanity:        4,
		CharacterName: "Nate",
		Objects:       make([]Object, 0),
	}
	Ellie := Player{
		Speed:         4,
		Might:         4,
		Knowledge:     4,
		Sanity:        4,
		CharacterName: "Ellie",
		Objects:       make([]Object, 0),
	}

	return []Player{Nate, Ellie, Rob, Bree, Kyle}
}
