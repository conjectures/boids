package main

import (
    "math"
    "math/rand"
    "time"
)

type Boid struct {
    position Vector2D
    velocity Vector2D
    id       int
}

func (b *Boid) calcAcceleration() Vector2D {
    // get upper and lower radius of boid sight
    upper, lower := b.position.AddValue(viewRadius), b.position.AddValue(-viewRadius)
    avgPosition, avgVelocity := Vector2D{x: 0, y: 0}, Vector2D{x:0, y:0}
    count := 0.0
    // calculate acceleration

    lock.Lock()
    for i := math.Max(lower.x, 0); i <= math.Min(upper.x, screenWidth); i++ {
        for j := math.Max(lower.y, 0); j <= math.Min(upper.y, screenHeight); j++ {
            if otherBoidId := boidMap[int(i)][int(j)]; otherBoidId != -1 && otherBoidId != b.id {
                // boid might be in view box but not in view Radius
                if dist := boids[otherBoidId].position.Distance(b.position); dist < viewRadius {
                    count++
                    avgVelocity = avgVelocity.Add(boids[otherBoidId].velocity)
                    avgPosition = avgPosition.Add(boids[otherBoidId].position)
                }
            }
        }
    }
    lock.Unlock()

    accel := Vector2D{x: 0, y: 0}
    if count > 0 {
        avgPosition, avgVelocity = avgPosition.DivideValue(count), avgVelocity.DivideValue(count)

        accelAlignment := avgVelocity.Subtract(b.velocity).MultiplyValue(adjRate)
        accelCohesion := avgPosition.Subtract(b.position).MultiplyValue(adjRate)
        accel = accel.Add(accelAlignment).Add(accelCohesion)
    }

    return accel
}
func (b *Boid) moveOne() {
    acceleration := b.calcAcceleration()
    lock.Lock()
    b.velocity = b.velocity.Add(acceleration).limit(-1, 1)
    boidMap[int(b.position.x)][int(b.position.y)] = -1
    b.position = b.position.Add(b.velocity)
    boidMap[int(b.position.x)][int(b.position.y)] = b.id

    next := b.position.Add(b.velocity)

    if next.x >= screenWidth || next.x < 0 {
        b.velocity = Vector2D{x: -b.velocity.x, y: b.velocity.y}
    }
    if next.y >= screenHeight || next.y < 0 {
        b.velocity = Vector2D{x: b.velocity.x, y: -b.velocity.y}
    }
    lock.Unlock()
}
func (b *Boid) start() {
    for {
        b.moveOne()
        time.Sleep(5 * time.Millisecond)
    }
}

func createBoid(bid int) {
    b := Boid{
        position: Vector2D{x: rand.Float64() * screenWidth, y: rand.Float64() * screenHeight},
        // random number between -1 and 1 for velocity (so that boids move only 1 pixel per update)
        velocity: Vector2D{x: (rand.Float64() * 2) - 1.0, y: (rand.Float64() * 2) - 1.0},
        id:       bid,
    }
    // Reference created structure for each boid
    boids[bid] = &b
    boidMap[int(b.position.x)][int(b.position.y)] = b.id
    go b.start()
}
