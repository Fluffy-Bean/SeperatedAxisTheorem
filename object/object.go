package object

import (
	raylib "github.com/gen2brain/raylib-go/raylib"
)

type Object struct {
	Vertices []raylib.Vector2
	Position raylib.Vector2
	Velocity raylib.Vector2
}

// ToWorldSpace converts the vertices of the object to world space based on the object's position
func (o *Object) ToWorldSpace() []raylib.Vector2 {
	worldSpace := make([]raylib.Vector2, len(o.Vertices))
	for i, v := range o.Vertices {
		worldSpace[i] = raylib.Vector2Add(v, o.Position)
	}
	return worldSpace
}

func (o *Object) Draw() {
	worldSpace := o.ToWorldSpace()
	for i := 0; i < len(worldSpace); i++ {
		raylib.DrawLineEx(worldSpace[i], worldSpace[(i+1)%len(worldSpace)], 1, raylib.Black)
	}
}

func NewObject(vertices []raylib.Vector2, position raylib.Vector2) *Object {
	return &Object{
		Vertices: vertices,
		Position: position,
	}
}
