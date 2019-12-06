package main

import (
	"fmt"
	"math"
	"time"
)

// public class MollerTrumbore {

//     private static double EPSILON = 0.0000001;

type point struct {
	x float64
	y float64
	z float64
}

func (v1 point) sub(v2 point) vector {
	return vector{v1.x - v2.x, v1.y - v2.y, v1.z - v2.z}
}

func (v1 point) sum(v2 vector) vector {
	return vector{v1.x + v2.x, v1.y + v2.y, v1.z + v2.z}
}

func (v1 vector) dotV(v2 vector) float64 {
	return ((v1).x*(v2).x + (v1).y*(v2).y + (v1).z*(v2).z)
}

type plane struct {
	p1     point
	normal vector
}

type vector struct {
	x float64
	y float64
	z float64
}

func (v1 vector) cross(v2 vector) vector {
	return vector{v1.y*v2.z - v1.z*v2.y, v1.z*v2.x - v1.x*v2.z, v1.x*v2.y - v1.y*v2.x}
}

func (v1 vector) scale(k float64) vector {
	return vector{k * v1.x, k * v1.y, k * v1.z}
}

func (v1 vector) crossF(f float64) vector {
	return vector{v1.x * f, v1.y * f, v1.z * f}
}

func (v1 vector) minus(v2 vector) vector {
	return vector{v1.x - v2.x, v1.y - v2.y, v1.z - v2.z}
}

func (v1 vector) magnitude() float64 {
	return math.Sqrt(v1.x*v1.x + v1.y*v1.y + v1.z*v1.z)
}

func (v1 vector) sub(v2 vector) vector {
	return vector{v1.x - v2.x, v1.y - v2.y, v1.z - v2.z}
}

func (v1 vector) subP(v2 point) vector {
	return vector{v1.x - v2.x, v1.y - v2.y, v1.z - v2.z}
}

func (v1 vector) dot(v2 vector) float64 {
	return ((v1).x*(v2).x + (v1).y*(v2).y + (v1).z*(v2).z)
}

func (v1 vector) dotP(v2 point) float64 {
	return ((v1).x*(v2).x + (v1).y*(v2).y + (v1).z*(v2).z)
}

//dot(u,v)   ((u).x * (v).x + (u).y * (v).y + (u).z * (v).z)

type ray struct {
	start     point
	direction vector
	end       point
}

type triangle struct {
	p0 point
	p1 point
	p2 point
}

func (tri triangle) collisionPoint(r ray) *vector {
	a := tri.p0
	b := tri.p1
	c := tri.p2

	ba := b.sub(a)
	ca := c.sub(a)

	normal := ba.cross(ca)
	// d := -normal.dotP(c)
	aToStart := a.sub(r.start)
	endToStart := r.end.sub(r.start)

	intervalPointNum := aToStart.dotV(normal)
	intervalPointDen := endToStart.dotV(normal)
	interval := intervalPointNum / intervalPointDen
	intersectionPoint := r.start.sum(endToStart.crossF(interval))

	u := b.sub(a)
	v := c.sub(a)
	w := intersectionPoint.subP(a)
	uu := u.dotV(u)
	uv := u.dotV(v)
	uw := u.dotV(w)
	vv := v.dotV(v)
	vw := v.dotV(w)

	s := (uv*vw - vv*uw) / (uv*uv - uu*vv)
	t := (uv*uw - uu*vw) / (uv*uv - uu*vv)
	if 0 <= s+t && s+t <= 1 && 0 <= interval && interval <= 1 {
		return &intersectionPoint
	}
	return nil
}

func main() {
	tri1 := triangle{
		point{0, 0.5, 1.5},
		point{1, 0.5, -0.5},
		point{-1, 0.5, -0.5},
	}
	r := ray{
		start: point{0, -0.5, 0},
		end:   point{0, 1.5, 0},
	}

	t1 := time.Now()
	fmt.Println(tri1.collisionPoint(r))
	t2 := time.Now()
	diff := t2.Sub(t1)
	fmt.Println(diff)

}
