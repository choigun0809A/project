package Image

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var Images = make(map[string]*ebiten.Image)

func New_Image(image_path string, desired_name string) {
	_, image, _ := ebitenutil.NewImageFromFile(image_path)
	Images[desired_name] = ebiten.NewImageFromImage(image)
}
