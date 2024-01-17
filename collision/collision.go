package collision

import (
	raylib "github.com/gen2brain/raylib-go/raylib"
	"math"
)

func CheckPolys(polyA, polyB []raylib.Vector2) (raylib.Vector2, float32, bool) {
	var (
		normal         = raylib.NewVector2(0, 0)
		depth  float32 = math.MaxFloat32
	)

	for i := 0; i < len(polyA); i++ {
		edge := raylib.Vector2Subtract(polyA[i], polyA[(i+1)%len(polyA)])
		axis := raylib.NewVector2(-edge.Y, edge.X)
		axis = raylib.Vector2Normalize(axis)

		minA, maxA := project(axis, polyA)
		minB, maxB := project(axis, polyB)

		if maxA < minB || maxB < minA {
			return normal, depth, false
		}

		axisDepth := float32(math.Min(float64(maxA-minB), float64(maxB-minA)))

		if axisDepth < depth {
			depth = axisDepth
			normal = axis
		}
	}

	for i := 0; i < len(polyB); i++ {
		edge := raylib.Vector2Subtract(polyB[(i+1)%len(polyB)], polyB[i])
		axis := raylib.NewVector2(edge.Y, -edge.X)
		axis = raylib.Vector2Normalize(axis)

		minA, maxA := project(axis, polyA)
		minB, maxB := project(axis, polyB)

		if maxA < minB || maxB < minA {
			return normal, depth, false
		}

		axisDepth := float32(math.Min(float64(maxA-minB), float64(maxB-minA)))

		if axisDepth < depth {
			depth = axisDepth
			normal = axis
		}
	}

	depth /= raylib.Vector2Length(normal)
	normal = raylib.Vector2Normalize(normal)

	centerA := findArithmeticMean(polyA)
	centerB := findArithmeticMean(polyB)
	direction := raylib.Vector2Subtract(centerB, centerA)

	if raylib.Vector2DotProduct(direction, normal) > 0 {
		normal = raylib.NewVector2(-normal.X, -normal.Y)
	}

	return normal, depth, true
}

func findArithmeticMean(vertices []raylib.Vector2) raylib.Vector2 {
	var (
		sumX float32
		sumY float32
	)

	for i := 0; i < len(vertices); i++ {
		v := vertices[i]
		sumX += v.X
		sumY += v.Y
	}

	return raylib.NewVector2(sumX/float32(len(vertices)), sumY/float32(len(vertices)))
}

func project(axis raylib.Vector2, poly []raylib.Vector2) (float32, float32) {
	minPoint := raylib.Vector2DotProduct(axis, poly[0])
	maxPoint := minPoint

	for i := 1; i < len(poly); i++ {
		p := raylib.Vector2DotProduct(axis, poly[i])
		if p < minPoint {
			minPoint = p
		}
		if p > maxPoint {
			maxPoint = p
		}
	}

	return minPoint, maxPoint
}
