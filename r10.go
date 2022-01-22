package main

import (
	"bytes"
	"embed"
	"image"
	_ "image/png"
	"log"
	"pr10/ann"
	"pr10/sb"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

//go:embed asset
var f embed.FS

var desireed = sb.NewVec2(320, 240)

type Game struct {
	Gem     *sb.Sprite
	Vehcile *ann.Vehicle
	targets []sb.Vec2
}

func (g *Game) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.genTargets()
	}
	g.Vehcile.Steer(g.targets, desireed)
	g.Vehcile.Update(640, 480)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i := 0; i < 8; i++ {
		g.Gem.Pos = g.targets[i].Clone()
		g.Gem.Draw(screen)
	}
	g.Vehcile.Rocket.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
func (g *Game) genTargets() {
	g.targets = nil
	for i := 0; i < 8; i++ {
		g.targets = append(g.targets, sb.NewVec2(sb.RandIntn(640), sb.RandIntn(480)))
	}
}
func NewGame() *Game {
	game := new(Game)
	game.Gem = sb.NewSprite(320, 240, true)
	game.Vehcile = ann.NewVehicle(8, float64(sb.RandIntn(640)), float64(sb.RandIntn(480)))
	game.targets = make([]sb.Vec2, 0)
	data, _ := f.ReadFile("asset/gem2.png")
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
	data1, _ := f.ReadFile("asset/car.png")
	img1, _, err := image.Decode(bytes.NewReader(data1))
	if err != nil {
		log.Fatal(err)
	}
	game.Gem.AddAnimFrame("default", img, 120, 6, true)
	game.Vehcile.Rocket.AddAnimFrame("default", img1, 60, 1, true)
	return game
}
func main() {
	gg := NewGame()
	gg.genTargets()
	gg.Gem.Start()
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Animation (Ebiten Demo)")
	if err := ebiten.RunGame(gg); err != nil {
		log.Fatal(err)
	}
}
