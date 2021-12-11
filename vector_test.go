package gomesh_test

import (
	"testing"

	"github.com/benp98/gomesh"
)

const maximumTolerance = 1e-8

func TestVector2DLength(t *testing.T) {
	testCases := []struct {
		v gomesh.Vector2D
		l float64
	}{
		{
			v: gomesh.Vector2D{3, 4},
			l: 5,
		},
		{
			v: gomesh.Vector2D{-3, 4},
			l: 5,
		},
		{
			v: gomesh.Vector2D{-3, -4},
			l: 5,
		},
		{
			v: gomesh.Vector2D{3, -4},
			l: 5,
		},
		{
			v: gomesh.Vector2D{-5, 2},
			l: 5.38516480,
		},
	}

	for _, testCase := range testCases {
		normalized := testCase.v.Normalized()

		// Check whether the length of the vector is correct regardless of orientation
		l := testCase.v.Length()
		if !gomesh.Float64AlmostEqual(l, testCase.l, maximumTolerance) {
			t.Errorf("Expected vector %s length of %f, got %f", testCase.v.String(), testCase.l, l)
		}

		// Check whether the normalized vector is of length 1
		normalizedLength := normalized.Length()
		if !gomesh.Float64AlmostEqual(1, normalizedLength, maximumTolerance) {
			t.Errorf("Expected length of the normalized vector %s to be 1, got %f", normalized.String(), normalizedLength)
		}
	}
}

func TestVector2DMath(t *testing.T) {
	a := gomesh.Vector2D{1, 2}
	b := gomesh.Vector2D{3, 4}

	expectedC := gomesh.Vector2D{4, 6}
	c := a.AddVector2D(b)
	if !c.AlmostEqual(expectedC, maximumTolerance) {
		t.Errorf("Expected vector %s to be %s as a result of %s+%s", c.String(), expectedC.String(), a.String(), b.String())
	}

	expectedD := gomesh.Vector2D{3, 8}
	d := a.MulVector2D(b)
	if !d.AlmostEqual(expectedD, maximumTolerance) {
		t.Errorf("Expected vector %s to be %s as a result of %s*%s", d.String(), expectedD.String(), a.String(), b.String())
	}

	expectedE := gomesh.Vector2D{2.5, 5}
	e := a.MulScalar(2.5)
	if !e.AlmostEqual(expectedE, maximumTolerance) {
		t.Errorf("Expected vector %s to be %s as a result of %s*2.5", e.String(), expectedE.String(), a.String())
	}

	expectedF := gomesh.Vector2D{-2, -2}
	f := a.SubtractVector2D(b)
	if !f.AlmostEqual(expectedF, maximumTolerance) {
		t.Errorf("Expected vector %s to be %s as a result of %s-%s", f.String(), expectedF.String(), a.String(), b.String())
	}
}

func TestVector3DLength(t *testing.T) {
	testCases := []struct {
		v gomesh.Vector3D
		l float64
	}{
		{
			v: gomesh.Vector3D{3, 4, 5},
			l: 7.07106781,
		},
		{
			v: gomesh.Vector3D{6, 3, 7},
			l: 9.69535971,
		},
		{
			v: gomesh.Vector3D{-5, 3, -9},
			l: 10.7238053,
		},
	}

	for _, testCase := range testCases {
		normalized := testCase.v.Normalized()

		// Check whether the length of the vector is correct regardless of orientation
		l := testCase.v.Length()
		if !gomesh.Float64AlmostEqual(l, testCase.l, maximumTolerance) {
			t.Errorf("Expected vector %s length of %f, got %f", testCase.v.String(), testCase.l, l)
		}

		// Check whether the normalized vector is of length 1
		normalizedLength := normalized.Length()
		if !gomesh.Float64AlmostEqual(1, normalizedLength, maximumTolerance) {
			t.Errorf("Expected length of the normalized vector %s to be 1, got %f", normalized.String(), normalizedLength)
		}
	}
}

func TestVector3DMath(t *testing.T) {
	a := gomesh.Vector3D{1, 2, 3}
	b := gomesh.Vector3D{4, 5, 6}

	expectedC := gomesh.Vector3D{5, 7, 9}
	c := a.AddVector3D(b)
	if !c.AlmostEqual(expectedC, maximumTolerance) {
		t.Errorf("Expected vector %s to be %s as a result of %s+%s", c.String(), expectedC.String(), a.String(), b.String())
	}

	expectedD := gomesh.Vector3D{4, 10, 18}
	d := a.MulVector3D(b)
	if !d.AlmostEqual(expectedD, maximumTolerance) {
		t.Errorf("Expected vector %s to be %s as a result of %s*%s", d.String(), expectedD.String(), a.String(), b.String())
	}

	expectedE := gomesh.Vector3D{2.5, 5, 7.5}
	e := a.MulScalar(2.5)
	if !e.AlmostEqual(expectedE, maximumTolerance) {
		t.Errorf("Expected vector %s to be %s as a result of %s*2.5", e.String(), expectedE.String(), a.String())
	}

	expectedF := gomesh.Vector3D{-3, -3, -3}
	f := a.SubtractVector3D(b)
	if !f.AlmostEqual(expectedF, maximumTolerance) {
		t.Errorf("Expected vector %s to be %s as a result of %s-%s", f.String(), expectedF.String(), a.String(), b.String())
	}
}
