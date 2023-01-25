package main

type Tile struct {
	RoomName string

	Top    WallAttribute
	Right  WallAttribute
	Left   WallAttribute
	Bottom WallAttribute

	SpecialCard SpecialCard
}

type WallAttribute string

const (
	Wall   = WallAttribute("wall")
	Window = WallAttribute("window")
	Door   = WallAttribute("door")
)

type SpecialCard string

const (
	Omen  = SpecialCard("omen")
	Item  = SpecialCard("item")
	Event = SpecialCard("event")
)

func (w WallAttribute) toSideDisplay() string {
	switch w {
	case Window:
		return "#"
	case Door:
		return " "
	default:
		return "|"
	}
}

func (w WallAttribute) toTopBottomDisplay() string {
	switch w {
	case Window:
		return "##"
	case Door:
		return "  "
	default:
		return "--"
	}
}
