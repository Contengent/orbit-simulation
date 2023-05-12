/*
* <2023-05-11>
* hihi it's edelstine back at it again.
*
 */

package main

import (
	// go packages

	"image/color"
	_ "image/png"
	"log"
	"math"
	"strconv"

	// ebiten packages

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct{}

var img *ebiten.Image

var planetInformation string = ""

var colors map[string]color.RGBA = map[string]color.RGBA{
	"black": {0, 0, 0, 0xff},
	"white": {255, 255, 255, 0xff},
	"red":   {255, 0, 0, 0xff},
	"blue":  {0, 0, 255, 0xff},
}

var planet ballObject = ballObject{
	shape:    shape{10, 10, colors["white"]},
	position: position{80, 180, 10, 0},
	creation: creation{nil, nil},
}

var blackHole ballObject = ballObject{
	shape:    shape{20, 20, colors["blue"]},
	position: position{180, 140, 0, 0},
	creation: creation{nil, nil},
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 400, 320
}

func main() {
	ebiten.SetWindowSize(800, 640)
	ebiten.SetWindowTitle("Gravity demo")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

/* render the guy */
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colors["black"])
	planet.ballCreation()
	blackHole.ballCreation()
	planet.gravityCalculation(blackHole)
	planet.velocityCalculation()

	x, y := ebiten.CursorPosition()

	planetInformation =
		"Position: (" + strconv.FormatFloat(planet.position.xpos, 'f', 0, 64) + ", " + strconv.FormatFloat(planet.position.ypos, 'f', 0, 64) + ")" + "\n" +
			"Velocity: (" + strconv.FormatFloat(planet.position.xvelocity, 'f', 2, 64) + ", " + strconv.FormatFloat(planet.position.yvelocity, 'f', 2, 64) + ")" + "\n" +
			"MousePos: (" + strconv.Itoa(x) + ", " + strconv.Itoa(y) + ")"

	planet.movePlanetToMouse()

	screen.DrawImage(planet.creation.image, planet.creation.options)
	screen.DrawImage(blackHole.creation.image, blackHole.creation.options)

	ebitenutil.DebugPrint(screen, planetInformation)

}

/* balls */
type ballObject struct {
	shape    shape
	position position
	creation creation
}

type shape struct {
	width  int
	height int
	color  color.RGBA
}

type position struct {
	xpos      float64
	ypos      float64
	xvelocity float64
	yvelocity float64
}

type creation struct {
	options *ebiten.DrawImageOptions
	image   *ebiten.Image
}

func updateBalls() *ebiten.DrawImageOptions {
	e := &ebiten.DrawImageOptions{}

	e.GeoM.Translate(0, 0)

	return e

}

func (ball *ballObject) velocityCalculation() {
	ball.position.xpos += ball.position.xvelocity
	ball.position.ypos += ball.position.yvelocity
}

func (ball *ballObject) ballCreation( /*screen *ebiten.Image*/ ) {

	ball.creation.image = ebiten.NewImage(ball.shape.width, ball.shape.height)
	ball.creation.image.Fill(ball.shape.color)

	ball.creation.options = &ebiten.DrawImageOptions{}
	ball.creation.options.GeoM.Translate(ball.position.xpos, ball.position.ypos)

	//screen.DrawImage(ball.creation.image, ball.creation.options)
}

func (planet *ballObject) gravityCalculation(blackhole ballObject) {
	xcomponent := blackhole.position.xpos - planet.position.xpos
	ycomponent := blackhole.position.ypos - planet.position.ypos

	vectorDistance := math.Sqrt(math.Pow(xcomponent, 2) + math.Pow(ycomponent, 2))

	xcomponent /= vectorDistance
	ycomponent /= vectorDistance

	planet.position.xvelocity += xcomponent
	planet.position.yvelocity += ycomponent
}

func (planet *ballObject) movePlanetToMouse() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		x, y := ebiten.CursorPosition()
		planet.position.xpos, planet.position.ypos = float64(x), float64(y)
		planet.position.xvelocity, planet.position.yvelocity = 0, 0
	}
}
