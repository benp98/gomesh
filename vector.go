package gomesh

import (
	"fmt"
	"math"
)

type Vector2D [2]float64

// Float64AlmostEqual checks whether the difference between the floats a and b is smaller than or equal to the given tolerance or not
func Float64AlmostEqual(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

// Length returns the length of the Vector2D
func (v Vector2D) Length() float64 {
	return math.Sqrt(math.Pow(math.Abs(v[0]), 2) + math.Pow(math.Abs(v[1]), 2))
}

// AlmostEqual checks whether the difference between this vector and another vector is smaller than or equal to the given tolerance or not
func (v Vector2D) AlmostEqual(other Vector2D, tolerance float64) bool {
	return Float64AlmostEqual(v[0], other[0], tolerance) && Float64AlmostEqual(v[1], other[1], tolerance)
}

// Normalized returns a normalized version of the Vector2D
func (v Vector2D) Normalized() Vector2D {
	nv := Vector2D{v[0], v[1]}

	l := v.Length()
	if l != 0 {
		nv[0] /= l
		nv[1] /= l
	}

	return nv
}

// AddVector2D returns a new Vector2D which is the sum of this vector and the given other vector
func (v Vector2D) AddVector2D(other Vector2D) Vector2D {
	nv := Vector2D{}

	nv[0] = v[0] + other[0]
	nv[1] = v[1] + other[1]

	return nv
}

// SubtractVector2D returns a new Vector2D which is the difference of this vector and the given other vector
func (v Vector2D) SubtractVector2D(other Vector2D) Vector2D {
	nv := Vector2D{}

	nv[0] = v[0] - other[0]
	nv[1] = v[1] - other[1]

	return nv
}

// MulVector2D returns a new Vector2D which is the product of this vector and the given other vector
func (v Vector2D) MulVector2D(other Vector2D) Vector2D {
	nv := Vector2D{}

	nv[0] = v[0] * other[0]
	nv[1] = v[1] * other[1]

	return nv
}

// MulScalar returns a new Vector2D which is the sum of this vector and the given scalar value
func (v Vector2D) MulScalar(scalar float64) Vector2D {
	nv := Vector2D{}

	nv[0] = v[0] * scalar
	nv[1] = v[1] * scalar

	return nv
}

// String returns a string representation of the Vector2D
func (v Vector2D) String() string {
	return fmt.Sprintf("(%f, %f)", v[0], v[1])
}

type Vector3D [3]float64

// Length returns the length of the Vector3D
func (v Vector3D) Length() float64 {
	return math.Sqrt(math.Pow(math.Abs(v[0]), 2) + math.Pow(math.Abs(v[1]), 2) + math.Pow(math.Abs(v[2]), 2))
}

// AlmostEqual checks whether the difference between this vector and another vector is smaller than or equal to the given tolerance or not
func (v Vector3D) AlmostEqual(other Vector3D, tolerance float64) bool {
	return Float64AlmostEqual(v[0], other[0], tolerance) && Float64AlmostEqual(v[1], other[1], tolerance) && Float64AlmostEqual(v[2], other[2], tolerance)
}

// Normalized returns a normalized version of the Vector3D
func (v Vector3D) Normalized() Vector3D {
	nv := Vector3D{v[0], v[1], v[2]}

	l := v.Length()
	if l != 0 {
		nv[0] /= l
		nv[1] /= l
		nv[2] /= l
	}

	return nv
}

// AddVector3D returns a new Vector3D which is the sum of this vector and the given other vector
func (v Vector3D) AddVector3D(other Vector3D) Vector3D {
	nv := Vector3D{}

	nv[0] = v[0] + other[0]
	nv[1] = v[1] + other[1]
	nv[2] = v[2] + other[2]

	return nv
}

// SubtractVector3D returns a new Vector3D which is the difference of this vector and the given other vector
func (v Vector3D) SubtractVector3D(other Vector3D) Vector3D {
	nv := Vector3D{}

	nv[0] = v[0] - other[0]
	nv[1] = v[1] - other[1]
	nv[2] = v[2] - other[2]

	return nv
}

// MulVector3D returns a new Vector3D which is the product of this vector and the given other vector
func (v Vector3D) MulVector3D(other Vector3D) Vector3D {
	nv := Vector3D{}

	nv[0] = v[0] * other[0]
	nv[1] = v[1] * other[1]
	nv[2] = v[2] * other[2]

	return nv
}

// MulScalar returns a new Vector3D which is the sum of this vector and the given scalar value
func (v Vector3D) MulScalar(scalar float64) Vector3D {
	nv := Vector3D{}

	nv[0] = v[0] * scalar
	nv[1] = v[1] * scalar
	nv[2] = v[2] * scalar

	return nv
}

// String returns a string representation of the Vector3D
func (v Vector3D) String() string {
	return fmt.Sprintf("(%f, %f, %f)", v[0], v[1], v[2])
}
