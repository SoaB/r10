package sb

type Canvas struct {
	Buf    []byte
	W, H   int
	Stride int
}

type Rgba struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

var (
	ColorBlack   = Rgba{0, 0, 0, 255}
	ColorWhite   = Rgba{255, 255, 255, 255}
	ColorRed     = Rgba{255, 0, 0, 255}
	ColorLime    = Rgba{0, 255, 0, 255}
	ColorBlue    = Rgba{0, 0, 255, 255}
	ColorYellow  = Rgba{255, 255, 0, 255}
	ColorCyan    = Rgba{0, 255, 255, 255}
	ColorMagenta = Rgba{255, 0, 255, 255}
	ColorSilver  = Rgba{192, 192, 192, 255}
	ColorGray    = Rgba{128, 128, 128, 255}
	ColorMaroon  = Rgba{128, 0, 0, 255}
	ColorOlive   = Rgba{128, 128, 0, 255}
	ColorGreen   = Rgba{0, 128, 0, 255}
	ColorPurple  = Rgba{128, 0, 128, 255}
	ColorTeal    = Rgba{0, 128, 128, 255}
	ColorNavy    = Rgba{0, 0, 128, 255}
)

func NewRgba(r, g, b, a int) Rgba {
	return Rgba{uint8(r), uint8(g), uint8(b), uint8(a)}
}

func NewCanvas(width, height int) *Canvas {
	cn := new(Canvas)
	cn.W = width
	cn.H = height
	cn.Stride = width * 4
	cn.Buf = make([]byte, width*height*4)
	return cn
}

func (c *Canvas) At(x, y int) Rgba {
	i := c.BufOffset(x, y)
	return Rgba{c.Buf[i], c.Buf[i+1], c.Buf[i+2], c.Buf[i+3]}
}

func (c *Canvas) Set(x, y int, color Rgba) {
	if x < 0 || x >= c.W || y < 0 || y >= c.H {
		return
	}
	i := c.BufOffset(x, y)
	s := c.Buf[i : i+4 : i+4] // Small cap improves performance, see https://golang.org/issue/27857
	s[0] = color.R
	s[1] = color.G
	s[2] = color.B
	s[3] = color.A
}

func (c *Canvas) BufOffset(x, y int) int {
	return (y*c.Stride + x*4)
}

func (c *Canvas) Clear(color Rgba) {
	var cnt int
	var i, k int
	for i = 0; i < c.H; i++ {
		for k = 0; k < c.W; k++ {
			c.Buf[cnt] = color.R
			c.Buf[cnt+1] = color.G
			c.Buf[cnt+2] = color.B
			c.Buf[cnt+3] = color.A
			cnt += 4
		}
	}
}

func intAbs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
func (c *Canvas) DrawRect(x, y, w, h int, col Rgba) {
	for i := y; i < y+h; i++ {
		c.Bresenham(x, i, x+w, i, col)
	}
}

// Bresenham draws a line between (x0, y0) and (x1, y1)
func (c *Canvas) Line(x0, y0, x1, y1, thick int, color Rgba) {
	if thick <= 1 {
		c.Bresenham(x0, y0, x1, y1, color)
		return
	}
	dx := intAbs(x1 - x0)
	dy := intAbs(y1 - y0)
	var sx, sy int
	if x0 < x1 {
		sx = 1
	} else {
		sx = -1
	}
	if y0 < y1 {
		sy = 1
	} else {
		sy = -1
	}
	err := dx - dy

	var e2 int
	for {
		c.DrawDisk(x0, y0, thick, color)
		if x0 == x1 && y0 == y1 {
			return
		}
		e2 = 2 * err
		if e2 > -dy {
			err = err - dy
			x0 = x0 + sx
		}
		if e2 < dx {
			err = err + dx
			y0 = y0 + sy
		}
	}
}

// PolylineBresenham draws a polyline to an image
func (c *Canvas) PolylineBresenham(color Rgba, s ...float64) {
	for i := 2; i < len(s); i += 2 {
		c.Bresenham(int(s[i-2]+0.5), int(s[i-1]+0.5), int(s[i]+0.5), int(s[i+1]+0.5), color)
	}
}

// Bresenham draws a line between (x0, y0) and (x1, y1)
func (c *Canvas) Bresenham(x0, y0, x1, y1 int, color Rgba) {
	dx := intAbs(x1 - x0)
	dy := intAbs(y1 - y0)
	var sx, sy int
	if x0 < x1 {
		sx = 1
	} else {
		sx = -1
	}
	if y0 < y1 {
		sy = 1
	} else {
		sy = -1
	}
	err := dx - dy

	var e2 int
	for {
		c.Set(x0, y0, color)
		if x0 == x1 && y0 == y1 {
			return
		}
		e2 = 2 * err
		if e2 > -dy {
			err = err - dy
			x0 = x0 + sx
		}
		if e2 < dx {
			err = err + dx
			y0 = y0 + sy
		}
	}
}

//Bresenham circle with center at (xc,yc) with radius and red green blue color
func (cn *Canvas) DrawCircle(xc, yc, radius int, color Rgba) {
	if xc-radius < 0 || xc+radius >= cn.W || yc-radius < 0 || yc+radius >= cn.H {
		return
	}
	x := 0
	y := radius
	p := 3 - (radius << 1)
	for x <= y {
		a := xc + x
		b := yc + y
		c := xc - x
		d := yc - y
		e := xc + y
		f := yc + x
		g := xc - y
		h := yc - x
		cn.Set(a, b, color)
		cn.Set(c, d, color)
		cn.Set(e, f, color)
		cn.Set(g, f, color)
		if x > 0 { //avoid drawing pixels at same position as the other ones
			cn.Set(a, d, color)
			cn.Set(c, b, color)
			cn.Set(e, h, color)
			cn.Set(g, h, color)
		}
		if p < 0 {
			p += (x << 2) + 6
			x++
		} else {
			p += ((x - y) << 2) + 10
			x++
			y--
		}
	}
}

//Filled bresenham circle with center at (xc,yc) with radius and red green blue color
func (cn *Canvas) DrawDisk(xc, yc, radius int, color Rgba) {
	if xc+radius < 0 || xc-radius >= cn.W || yc+radius < 0 || yc-radius >= cn.H {
		return
	}
	x := 0
	y := radius
	p := 3 - (radius << 1)
	pb := yc + radius + 1
	pd := yc + radius + 1
	for x <= y {
		a := xc + x
		b := yc + y
		c := xc - x
		d := yc - y
		e := xc + y
		f := yc + x
		g := xc - y
		h := yc - x
		if b != pb {
			cn.Bresenham(a, b, c, b, color)
		}
		if d != pd {
			cn.Bresenham(a, d, c, d, color)
		}
		if f != b {
			cn.Bresenham(e, f, g, f, color)
		}
		if h != d && h != f {
			cn.Bresenham(e, h, g, h, color)
		}
		pb = b
		pd = d
		if p < 0 {
			p += (x << 2) + 6
			x++
		} else {
			p += ((x - y) << 2) + 10
			x++
			y--
		}
	}
}
