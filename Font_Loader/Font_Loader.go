package Font_loader

import (
	"embed"
	"fmt"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//go:embed *.ttf
var fontsFS embed.FS

var Font_Faces = make(map[string]font.Face)

func Load_Font(font_path string, font_name string, font_size float64) {
	font_data, err := fontsFS.ReadFile(font_path)
	if err != nil {
		fmt.Println("failed when loading...")
		panic(err)
	}
	tt, err := opentype.Parse(font_data)
	if err != nil {
		fmt.Println("failed when parsing...")
		panic(err)
	}
	ff, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    font_size,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		fmt.Println("failed when creating face...")
		panic(err)
	}
	Font_Faces[font_name] = ff
}
