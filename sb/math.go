package sb

import (
	"math"
	"math/rand"
	"time"
)

var rnd *rand.Rand

func init() {
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RandFloat64Area : return a float64 between min and max
func RandFloat64Area(min, max float64) float64 {
	return (min + rnd.Float64()*(max-min))
}

// RandFloat64n : return a float64 with limit n
func RandFloat64n(n float64) float64 {
	return rnd.Float64() * n
}

// RandFloat64 : return a float64 (0.0-1.0)
func RandFloat64() float64 {
	return rnd.Float64()
}

// RandIntArea : return a Integer between min and max
func RandIntArea(min, max int) int {
	return min + rnd.Intn(max-min)
}

// RandIntn : return a Integer less n
func RandIntn(n int) int {
	return rnd.Intn(n)
}

// RandInt : return a Integer
func RandInt() int {
	return rnd.Int()
}

//Remap : Remap a num from a1-a2 to b1-b2
func Remap(val, a1, a2, b1, b2 float64) float64 {
	return b1 + (val-a1)*(b2-b1)/(a2-a1)
}

//Constrain :
func Constrain(a, min, max float64) float64 {
	if a < min {
		return min
	} else if a > max {
		return max
	}
	return a
}

// EPSILON :The default epsilon value used for floating point comparisons.
const EPSILON = 1.0e-8

// AlmostEqual tells if the two floating point values x and y are considered
// equal within the specified absolute==relative tolerence value.
//
// The method of comparison used is that described at:
//  http://realtimecollisiondetection.net/blog/?p=89
func AlmostEqual(x, y, absTol float64) bool {
	if x == y || (math.Abs(x-y) <= absTol*math.Max(1.0, math.Max(math.Abs(x), math.Abs(y)))) {
		return true
	}
	return false
}

// Equal tells if the two floating point values a and b are considered equal
// within the default EPSILON comparison value. It is short-handed for:
//  AlmostEqual(a, b, EPSILON)
func Equal(a, b float64) bool {
	return AlmostEqual(a, b, EPSILON)
}

// Clamp returns the value v clamped to the range of [min, max].
func Clamp(v, min, max float64) float64 {
	return math.Max(math.Min(v, max), min)
}

// Radians converts from degrees to radians.
func Radians(degrees float64) (radians float64) {
	return math.Pi * degrees / 180.0
}

// Degrees converts from radians to degrees.
func Degrees(radians float64) (degrees float64) {
	return radians * (180.0 / math.Pi)
}

// Rounded returns the value rounded to the nearest whole number.
func Rounded(v float64) float64 {
	if v < 0 {
		return math.Ceil(v - 0.5)
	}
	return math.Floor(v + 0.5)
}

// Lerp performs a linear interpolation between a and b. The t parameter is a
// number in the range 0.0-1.0. Some examples:
//
//  Lerp(0, 10, 0) == 0
//  Lerp(0, 10, 0.5) == 5
//  Lerp(0, 10, 1) == 10
//
// The interpolation method is precise, so it is guaranteed that:
//  Lerp(a, b, 1) == a
func Lerp(a, b, t float64) float64 {
	return (1-t)*a + t*b
}
