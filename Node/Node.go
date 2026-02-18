package Node

import (
	"image/color"
	common_tools "project/Common_tools"
	font_loader "project/Font_Loader"
	image "project/Image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Status int

const (
	Text_type Status = iota
	Image_type
)

type Node struct {
	Op    ebiten.DrawImageOptions
	Image *ebiten.Image

	Text_Op    text.DrawOptions
	Text       string
	font_color color.Color

	Font_name string
	Scale     common_tools.Vector
	Position  common_tools.Vector

	Children []*Node
	Parent   *Node

	Type_ Status
}

func (n *Node) Draw_node(screen *ebiten.Image) {
	switch n.Type_ {
	case Text_type:
		face, ok := font_loader.Font_Faces[n.Font_name]
		if !ok || face == nil {
			return
		}

		// Wrap old font.Face into v2 text.Face
		go_face := text.NewGoXFace(face)

		op := &text.DrawOptions{}

		// Position
		op.GeoM.Translate(float64(n.Position.X), float64(n.Position.Y))

		text.Draw(screen, n.Text, go_face, &n.Text_Op)
	case Image_type:
		screen.DrawImage(n.Image, &n.Op)
	}

}

func (n *Node) Draw_Child_On_Parent() {
	for _, child := range n.Children {
		child.Draw_node(n.Image)
	}
}

func (n *Node) Update_Op() {
	switch n.Type_ {
	case Text_type:
		n.Text_Op.GeoM.Reset()
		n.Text_Op.GeoM.Scale(float64(n.Scale.X), float64(n.Scale.Y))
		n.Text_Op.GeoM.Translate(float64(n.Position.X), float64(n.Position.Y))
	case Image_type:
		n.Op.GeoM.Reset()
		n.Op.GeoM.Scale(float64(n.Scale.X), float64(n.Scale.Y))
		n.Op.GeoM.Translate(float64(n.Position.X), float64(n.Position.Y))
	}

}

func (n *Node) Change_scale(x, y float64) {
	n.Scale = common_tools.Vector{X: float32(x), Y: float32(y)}
	n.Update_Op()
}

func (n *Node) Change_position(x, y float64) {
	n.Position = common_tools.Vector{X: float32(x), Y: float32(y)}
	n.Update_Op()
}

func (n *Node) Get_changed_size() (float32, float32) {
	return float32(n.Image.Bounds().Size().X) * n.Scale.X, float32(n.Image.Bounds().Size().Y) * n.Scale.Y
}

func (n *Node) Add_Child(child *Node) {
	n.Children = append(n.Children, child)
	child.Parent = n
}

func (n *Node) Remove_Child(child *Node) {
	for i := 0; i < len(n.Children); i++ {
		if n.Children[i] == child {
			n.Children = append(n.Children[:i], n.Children[i+1:]...)
		}
	}
}

func New_Node(image_name string) *Node {
	n := Node{
		Op:    ebiten.DrawImageOptions{},
		Image: image.Images[image_name],

		Scale:    common_tools.Vector{X: 1, Y: 1},
		Position: common_tools.Vector{X: 0, Y: 0},
		Children: []*Node{nil},
		Parent:   nil,

		Type_: Image_type,
	}
	return &n
}

func New_Text(text_ string, font_name string, font_color color.Color) *Node {
	n := Node{
		Text_Op:   text.DrawOptions{},
		Text:      text_,
		Font_name: font_name,
		Scale:     common_tools.Vector{X: 1, Y: 1},
		Position:  common_tools.Vector{X: 0, Y: 0},
		Children:  []*Node{nil},
		Parent:    nil,

		Type_: Text_type,
	}
	return &n
}
