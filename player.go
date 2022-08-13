package main

type Player struct {
	playerCharacter Tile
}

func (p *Player) create(character Tile) {
	p.playerCharacter = character
}
