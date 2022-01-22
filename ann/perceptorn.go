package ann

// Simple Perceptron
import (
	"pr10/sb"
)

// Perceptron :
type Perceptron struct {
	Weights []float64 // Slice of weight for inputs
	C       float64   // learning constant
}

/*
NewPerceptron :
    create with n weights and learning constant
*/
func NewPerceptron(n int, c float64) *Perceptron {
	weights := make([]float64, n)
	// Start with random weight
	for i := 0; i < n; i++ {
		weights[i] = sb.RandFloat64Area(0, 1)
	}
	return &Perceptron{weights, c}
}

/*
FeedForward :
    Give me a steering result
*/
func (p *Perceptron) FeedForward(forces []sb.Vec2) sb.Vec2 {
	sum := sb.NewVec2(0, 0)
	for i := 0; i < len(p.Weights); i++ {
		forces[i] = forces[i].MulScalar(p.Weights[i])
		sum = sum.Add(forces[i])
	}
	return sum
}

/*
Train :
    train the perceptron
    Weights are adjusted based on vehicle's error
*/
func (p *Perceptron) Train(forces []sb.Vec2, err sb.Vec2) {
	for i := 0; i < len(p.Weights); i++ {
		p.Weights[i] += p.C * err.X * forces[i].X
		p.Weights[i] += p.C * err.Y * forces[i].Y
		p.Weights[i] = sb.Constrain(p.Weights[i], 0, 1)
	}
}
