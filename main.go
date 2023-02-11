package main

import ( 
	"github.com/hajimehoshi/ebiten/v2"
    "image/color"
    "log"
)

const (
    screenWidth, screenHeight = 640, 360
    boidCount = 500
)

var (
    green = color.RGBA{R:10, G:255, B:50, A:255}
    // list of `boidCount` elements of (pointer) Boid type
    boids [boidCount] *Boid
)

type Game struct{}

func (g *Game) Update() error {
    return nil
}

func (g *Game) Draw(screen *ebiten.Image){
    for _, boid := range boids {
        // iterate over a list of boids and draw a with a diamond shape (4 pixels)
        screen.Set(int(boid.position.x+1), int(boid.position.y), green)
        screen.Set(int(boid.position.x-1), int(boid.position.y), green)
        screen.Set(int(boid.position.x), int(boid.position.y-1), green)
        screen.Set(int(boid.position.x), int(boid.position.y+1), green)
    }
}

func (g *Game) Layout(_,_ int) (w, h int) {
    return screenWidth, screenHeight
}

func main() {
    for i := 0; i < boidCount; i++ {
        createBoid(i)
    }
    ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
    ebiten.SetWindowTitle("Boids in a box")
    // Run graphics library by giving a number of parameters (screen size, etc)
    if err := ebiten.RunGame(&Game{}); err != nil {
        log.Fatal(err)
    }
}
