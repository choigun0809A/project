package main

import (
	"image/color"
	"log"

	Font_loader "github.com/choigun0809A/project/Font_Loader"
	"github.com/choigun0809A/project/Image"
	"github.com/choigun0809A/project/Node"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Nodes map[string]*Node.Node
}

func (g *Game) Update() error {
	return nil
}

func draw_screen_grid(screen *ebiten.Image) {

}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)

	if node, ok := g.Nodes["card_background"]; ok {
		node.Draw_node(screen)
	}

	if node, ok := g.Nodes["text"]; ok {
		node.Draw_node(screen)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func load_images_and_nodes(game *Game) {
	Image.New_Image("assets/card_background.png", "card_background")

	Font_loader.Load_Font("arial.ttf", "arial-16", 16)
	Font_loader.Load_Font("arial.ttf", "arial-20", 20)

	game.Nodes["text"] = Node.New_Text("text", "arial-16", color.Black)
	game.Nodes["text"].Change_position(100, 100)

	game.Nodes["card_background"] = Node.New_Node("card_background")
	game.Nodes["card_background"].Change_scale(0.5, 0.5)

}

func main() {
	game := Game{
		Nodes: make(map[string]*Node.Node),
	}

	ebiten.SetWindowSize(414, 896)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetTPS(60)

	load_images_and_nodes(&game)

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
