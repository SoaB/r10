package sb

import (
	"fmt"
	"math"
)

// Vec2 represents a 2D Vec2 or point.
type Vec2 struct {
	X, Y float64
}

// NewVec2 : create a new vec2
func NewVec2(x, y int) Vec2 {
	return Vec2{float64(x), float64(y)}
}

// Clone : Clone a vector
func (a Vec2) Clone() Vec2 {
	return Vec2{a.X, a.Y}
}

// String returns an string representation of this Vec2.
func (a Vec2) String() string {
	return fmt.Sprintf("Vec2(X=%f, Y=%f)", a.X, a.Y)
}

// Equals tells if a == b using the default EPSILON value.
func (a Vec2) Equals(b Vec2) bool {
	return math.Abs(a.X-b.X) < EPSILON &&
		math.Abs(a.Y-b.Y) < EPSILON
}

// Add performs a componentwise addition of the two Vec2s, returning a + b.
func (a Vec2) Add(b Vec2) Vec2 {
	return Vec2{a.X + b.X, a.Y + b.Y}
}

// AddScalar performs a componentwise scalar addition of a + b.
func (a Vec2) AddScalar(b float64) Vec2 {
	return Vec2{a.X + b, a.Y + b}
}

// Sub performs a componentwise subtraction of the two Vec2s, returning
// a - b.
func (a Vec2) Sub(b Vec2) Vec2 {
	return Vec2{a.X - b.X, a.Y - b.Y}
}

// SubScalar performs a componentwise scalar subtraction of a - b.
func (a Vec2) SubScalar(b float64) Vec2 {
	return Vec2{a.X - b, a.Y - b}
}

// Mul performs a componentwise multiplication of the two Vec2s, returning
// a * b.
func (a Vec2) Mul(b Vec2) Vec2 {
	return Vec2{a.X * b.X, a.Y * b.Y}
}

// MulScalar performs a componentwise scalar multiplication of a * b.
func (a Vec2) MulScalar(b float64) Vec2 {
	return Vec2{a.X * b, a.Y * b}
}

// Div performs a componentwise division of the two Vec2s, returning a * b.
func (a Vec2) Div(b Vec2) Vec2 {
	return Vec2{a.X / b.X, a.Y / b.Y}
}

// DivScalar performs a componentwise scalar division of a * b.
func (a Vec2) DivScalar(b float64) Vec2 {
	return Vec2{a.X / b, a.Y / b}
}

// IsNaN tells if any components of this Vec2 are not an number.
func (a Vec2) IsNaN() bool {
	return math.IsNaN(a.X) || math.IsNaN(a.Y)
}

// Less tells if a is componentwise less than b:
//  return a.X < b.X && a.Y < b.Y
func (a Vec2) Less(b Vec2) bool {
	return a.X < b.X && a.Y < b.Y
}

// Greater tells if a is componentwise greater than b:
//  return a.X > b.X && a.Y > b.Y
func (a Vec2) Greater(b Vec2) bool {
	return a.X > b.X && a.Y > b.Y
}

// AnyLess tells if a is componentwise any less than b:
//  return a.X < b.X || a.Y < b.Y
func (a Vec2) AnyLess(b Vec2) bool {
	return a.X < b.X || a.Y < b.Y
}

// AnyGreater tells if a is componentwise any greater than b:
//  return a.X > b.X || a.Y > b.Y
func (a Vec2) AnyGreater(b Vec2) bool {
	return a.X > b.X || a.Y > b.Y
}

// Clamp clamps each value in the Vec2 to the range of [min, max] and returns
// it.
func (a Vec2) Clamp(min, max float64) Vec2 {
	return Vec2{
		Clamp(a.X, min, max),
		Clamp(a.Y, min, max),
	}
}

// Radians converts each value in the Vec2 from degrees to radians and
// returns it.
func (a Vec2) Radians() Vec2 {
	return Vec2{
		Radians(a.X),
		Radians(a.Y),
	}
}

// Degrees converts each value in the Vec2 from radians to degrees and
// returns it.
func (a Vec2) Degrees() Vec2 {
	return Vec2{
		Degrees(a.X),
		Degrees(a.Y),
	}
}

// Rounded rounds each value in the Vec2 to the nearest whole number and
// returns it.
func (a Vec2) Rounded() Vec2 {
	return Vec2{
		Rounded(a.X),
		Rounded(a.Y),
	}
}

// Dot returns the dot product of a and b.
func (a Vec2) Dot(b Vec2) float64 {
	return a.X*b.X + a.Y*b.Y
}

// Inverse returns the inverse (negated) Vec2 -a.
func (a Vec2) Inverse() Vec2 {
	return Vec2{-a.X, -a.Y}
}

// LengthSq returns the magnitude squared of this Vec2, useful for comparing
// distances.
func (a Vec2) LengthSq() float64 {
	return a.X*a.X + a.Y*a.Y
}

// Length returns the magnitude of this Vec2. To avoid a sqrt call when
// strictly comparing distances, LengthSq can be used instead.
func (a Vec2) Length() float64 {
	return math.Sqrt(a.X*a.X + a.Y*a.Y)
}

// Normalized returns the normalized (i.e. length/magnitude == 1) Vec2 of a.
// If the Vec2's length is zero (and division by zero would occur) then
// [Vec2Zero, false] is returned.
func (a Vec2) Normalized() (v Vec2, ok bool) {
	length := math.Sqrt(a.X*a.X + a.Y*a.Y)
	if Equal(length, 0) {
		return Vec2{0, 0}, false
	}
	return Vec2{
		a.X / length,
		a.Y / length,
	}, true
}

// Normalize :
func (a *Vec2) Normalize() {
	length := math.Sqrt(a.X*a.X + a.Y*a.Y)
	if length != 0 {
		a.X = a.X / length
		a.Y = a.Y / length
	}
}

// Proj returns a Vec2 representing the projection of Vec2 a onto b.
func (a Vec2) Proj(b Vec2) Vec2 {
	return b.MulScalar(a.Dot(b) / b.LengthSq())
}

// Min returns a Vec2 representing the smallest components of both the
// Vec2s.
func (a Vec2) Min(b Vec2) Vec2 {
	var r Vec2
	if a.X < b.X {
		r.X = a.X
	} else {
		r.X = b.X
	}
	if a.Y < b.Y {
		r.Y = a.Y
	} else {
		r.Y = b.Y
	}
	return r
}

// Max returns a Vec2 representing the largest components of both the
// Vec2s.
func (a Vec2) Max(b Vec2) Vec2 {
	var r Vec2
	if a.X > b.X {
		r.X = a.X
	} else {
		r.X = b.X
	}
	if a.Y > b.Y {
		r.Y = a.Y
	} else {
		r.Y = b.Y
	}
	return r
}

// Lerp returns a Vec2 representing the linear interpolation between the
// Vec2s a and b. The t parameter is the amount to interpolate (0.0 - 1.0)
// between the Vec2s.
func (a Vec2) Lerp(b Vec2, t float64) Vec2 {
	return a.Mul(b.MulScalar(t))
}

// Angle returns the angle in deg between the two Vec2s.
func (a Vec2) Angle(b Vec2) float64 {
	agl := math.Atan2(b.Y-a.Y, b.X-a.X) * (180.0 / math.Pi)
	if agl < 0 {
		agl += 360.0
	}
	return agl
}

// AngleRad returns the angle in rad between the two Vec2s.
func (a Vec2) AngleRad(b Vec2) float64 {
	angle := math.Atan2(b.Y-a.Y, b.X-a.X)
	if angle < 0 {
		angle += 2 * math.Pi
	}
	return angle + math.Pi
}

// Limit limit to val
func (a *Vec2) Limit(n float64) {
	l := a.Length()
	if n <= l {
		ra := n / l
		a.X *= ra
		a.Y *= ra
	}
}

/*
Perpendicular : return a Perpendicular vector
垂直是一個幾何術語。
在平面幾何中，如果一條直線與另一條直線相交，
且它們構成的任意相鄰兩個角相等，
那麼這兩條直線相互垂直。
*/
func (a Vec2) Perpendicular() Vec2 {
	return Vec2{a.Y, -a.X}
}

// Rotate : rotate vector with angle
func (a Vec2) Rotate(angle float64) Vec2 {
	result := Vec2{0, 0}
	ra := angle * math.Pi / 180.0
	result.X = a.X*math.Cos(ra) - a.Y*math.Sin(ra)
	result.Y = a.X*math.Sin(ra) + a.Y*math.Cos(ra)

	return result
}

// Distance : return distance between two points
func (a Vec2) Distance(b Vec2) float64 {
	xx := a.X - b.X
	yy := a.Y - b.Y
	return math.Sqrt(xx*xx + yy*yy)
}
