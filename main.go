package main

import (
	"SeperatedAxisTheorem/collision"
	"SeperatedAxisTheorem/object"
	raylib "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	raylib.InitWindow(800, 450, "Raylib SAT Collision Detection/Resolution")
	raylib.SetTargetFPS(60)

	cube := object.NewObject([]raylib.Vector2{{0, 0}, {800, 0}, {800, 10}, {0, 10}}, raylib.Vector2{0, 440})
	cube2 := object.NewObject([]raylib.Vector2{{0, 0}, {800, 0}, {800, 10}, {0, 10}}, raylib.Vector2{500, 389})
	//poly := object.NewObject([]raylib.Vector2{{-50, -50}, {50, -50}, {80, 50}, {0, 70}, {-80, 50}}, raylib.Vector2{400, 200})
	poly := object.NewObject([]raylib.Vector2{{0, 0}, {100, -50}, {100, 0}}, raylib.Vector2{400, 440})
	player := object.NewObject([]raylib.Vector2{{-25, -25}, {25, -25}, {25, 25}, {-25, 25}}, raylib.Vector2{400, 225})

	velocity := raylib.NewVector2(0, 0)
	canJump := false
	for !raylib.WindowShouldClose() {
		if raylib.IsKeyDown(raylib.KeyD) {
			velocity.X += 2
		} else if raylib.IsKeyDown(raylib.KeyA) {
			velocity.X -= 2
		}

		if raylib.IsKeyDown(raylib.KeySpace) && canJump {
			velocity.Y -= 25
			canJump = false
		}

		velocity.Y += 2

		velocity = raylib.Vector2Scale(velocity, 0.8)
		player.Position = raylib.Vector2Add(player.Position, velocity)

		normal, depth, colliding := collision.CheckPolys(cube.ToWorldSpace(), player.ToWorldSpace())
		if colliding {
			player.Position = raylib.Vector2Subtract(player.Position, raylib.Vector2Scale(normal, depth))
			canJump = true
		}
		normal, depth, colliding = collision.CheckPolys(poly.ToWorldSpace(), player.ToWorldSpace())
		if colliding {
			player.Position = raylib.Vector2Subtract(player.Position, raylib.Vector2Scale(normal, depth))
			canJump = true
		}
		normal, depth, colliding = collision.CheckPolys(cube2.ToWorldSpace(), player.ToWorldSpace())
		if colliding {
			player.Position = raylib.Vector2Subtract(player.Position, raylib.Vector2Scale(normal, depth))
			canJump = true
		}

		raylib.BeginDrawing()
		raylib.ClearBackground(raylib.RayWhite)

		if colliding {
			raylib.DrawText("Collision!", 10, 10, 20, raylib.Red)
		}
		cube.Draw()
		cube2.Draw()
		poly.Draw()
		player.Draw()

		raylib.EndDrawing()
	}
}
