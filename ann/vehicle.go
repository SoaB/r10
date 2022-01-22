package ann

import "pr10/sb"

type Vehicle struct {
	Brain        *Perceptron
	Rocket       *sb.Sprite
	Velocity     sb.Vec2
	Acceleration sb.Vec2
	R            float64
	MaxForce     float64
	MaxSpeed     float64
}

func NewVehicle(n int, x, y float64) *Vehicle {
	brain := NewPerceptron(n, 0.001)
	rocket := sb.NewSprite(int(x), int(y), true)
	acceleration := sb.NewVec2(0, 0)
	velocity := sb.NewVec2(0, 0)
	r := 3.0
	maxspeed := 3.0
	maxforce := 0.05
	return &Vehicle{
		Brain:        brain,
		Rocket:       rocket,
		Velocity:     velocity,
		Acceleration: acceleration,
		R:            r,
		MaxForce:     maxforce,
		MaxSpeed:     maxspeed}
}

func (v *Vehicle) Update(w, h int) {
	lastPos := v.Rocket.Pos.Clone()
	// Update velocity
	v.Velocity = v.Velocity.Add(v.Acceleration)
	// Limit speed
	v.Velocity.Limit(v.MaxSpeed)
	v.Rocket.Pos = v.Rocket.Pos.Add(v.Velocity)
	// Reset acceleration to 0 each cycle
	v.Acceleration = sb.Vec2{X: 0, Y: 0}
	v.Rocket.Pos.X = sb.Constrain(v.Rocket.Pos.X, 0, float64(w))
	v.Rocket.Pos.Y = sb.Constrain(v.Rocket.Pos.Y, 0, float64(h))
	v.Rocket.Angle = v.Rocket.Pos.AngleRad(lastPos)
}

func (v *Vehicle) ApplyForce(force sb.Vec2) {
	v.Acceleration = v.Acceleration.Add(force)
}

/*
Steer :
    Here is where the brain processes everything
*/
func (v *Vehicle) Steer(targets []sb.Vec2, desired sb.Vec2) {
	// Make a Slice of forces
	forces := make([]sb.Vec2, len(targets))
	// Steer towards all targets
	for i := 0; i < len(forces); i++ {
		forces[i] = v.Seek(targets[i])
	}
	// The array of forces is the input to the brain
	result := v.Brain.FeedForward(forces)
	// Use the result to steer the vehicle
	v.ApplyForce(result)
	// Train the brain according to the error
	err := desired.Sub(v.Rocket.Pos)
	v.Brain.Train(forces, err)
}

/*
Seek :
   A metod that calculates a steering force towards
   a target
   STEER = DESIRED MINUS VELOCITY
*/
func (v *Vehicle) Seek(target sb.Vec2) sb.Vec2 {
	desired := target.Sub(v.Rocket.Pos)
	// Normalize desired and scale to maxium speed
	desired.Normalize()
	desired = desired.MulScalar(v.MaxSpeed)
	// steering = Desired minus velocity
	steer := desired.Sub(v.Velocity)
	steer.Limit(v.MaxForce)
	return steer
}
