package sb

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type AnimFrame struct {
	Image     *ebiten.Image
	Duration  *TimeCounter
	Frect     []image.Rectangle
	Center    Vec2
	MaxFrame  int
	CurrFrame int
	Loop      bool
}

func NewAnimFrame(img image.Image, duration int, frames int, loop bool) *AnimFrame {
	animFrame := new(AnimFrame)
	animFrame.Image = ebiten.NewImageFromImage(img)
	animFrame.MaxFrame = frames
	animFrame.Frect = make([]image.Rectangle, animFrame.MaxFrame)
	w, h := animFrame.Image.Size()
	w = w / animFrame.MaxFrame
	animFrame.Center = NewVec2(w/2, h/2)
	for i := 0; i < animFrame.MaxFrame; i++ {
		animFrame.Frect[i] = image.Rect(i*w, 0, i*w+w, h)
	}
	animFrame.Duration = NewCounter(duration)
	animFrame.Loop = loop
	return animFrame
}

func (af *AnimFrame) SetDuration(d int) {
	af.Duration.ResetCounter(d)
}

func (af *AnimFrame) SetLoop(b bool) {
	af.Loop = b
}

func (af *AnimFrame) SetFrameRect(rect []image.Rectangle) {
	af.MaxFrame = len(rect)
	af.Frect = nil
	af.Frect = append(af.Frect, rect...)
}

type Sprite struct {
	CurrAnimName       string
	AnimFrames         map[string]*AnimFrame
	Pos                Vec2
	Angle              float64
	Alpha              float64
	Visible            bool
	Animated           bool
	CenterCoordonnates bool
}

func NewSprite(x, y int, center bool) *Sprite {
	sprite := new(Sprite)
	sprite.CurrAnimName = "default"
	sprite.AnimFrames = make(map[string]*AnimFrame)
	sprite.Pos = NewVec2(x, y)
	sprite.Angle = 0
	sprite.Alpha = 1
	sprite.Visible = true
	sprite.Animated = false
	sprite.CenterCoordonnates = center
	return sprite
}

func (spr *Sprite) AddAnimFrame(label string, img image.Image,
	duration int, frames int, loop bool) {
	spr.AnimFrames[label] = NewAnimFrame(img, duration, frames, loop)
}
func (spr *Sprite) nextFrame() {
	currAnim := spr.AnimFrames[spr.CurrAnimName]
	if currAnim.MaxFrame > 1 {
		if spr.Animated {
			if currAnim.Duration.TimeUp() {
				currAnim.CurrFrame++
				if currAnim.CurrFrame+1 > currAnim.MaxFrame {
					if currAnim.Loop {
						currAnim.CurrFrame = 0
					} else {
						spr.Animated = false
						spr.Visible = false
					}
				}
			}
		}
	}
}
func (spr *Sprite) Draw(surface *ebiten.Image) {
	if spr.Visible {
		currFram := spr.AnimFrames[spr.CurrAnimName]
		opt := &ebiten.DrawImageOptions{}
		if spr.CenterCoordonnates {
			opt.GeoM.Translate(-currFram.Center.X, -currFram.Center.Y)
		}
		opt.GeoM.Rotate(spr.Angle)
		opt.GeoM.Translate(spr.Pos.X, spr.Pos.Y)
		opt.ColorM.Scale(1, 1, 1, spr.Alpha)
		surface.DrawImage(currFram.Image.SubImage(currFram.Frect[currFram.CurrFrame]).(*ebiten.Image), opt)
		spr.nextFrame()
	}
}

func (spr *Sprite) Start() {
	spr.Animated = true
	spr.Visible = true
}

func (spr *Sprite) Stop() {
	spr.Animated = false
	spr.Visible = false
}

func (spr *Sprite) Reset() {
	spr.Animated = false
	spr.Visible = false
	spr.AnimFrames[spr.CurrAnimName].CurrFrame = 0
}
