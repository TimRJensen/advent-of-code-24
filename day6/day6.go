package main

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
)

type vec [3]float64

func (v vec) sub(w vec) vec {
	return vec{v[0] - w[0], v[1] - w[1]}
}

func (v vec) length() float64 {
	return math.Sqrt(v[0]*v[0] + v[1]*v[1])
}

func (v vec) project(dist, theta float64) vec {
	return vec{(v[0] + dist*math.Cos(theta)), (v[1] + dist*math.Sin(theta))}
}

func (v vec) cross(w vec) float64 {
	// A cross product on 2d vectors is technically not defined.
	// To achieve it, one should pad the 2d vectors to 3d vectors,
	// cross it and return the z component. The below is practically
	// the same.
	return v[0]*w[1] - v[1]*w[0]
}

func incercepts(v, r, w, s vec) bool {
	// https://stackoverflow.com/a/565282/16857618
	rxs := r.sub(v).cross(s.sub(w))
	if rxs == 0 {
		return false
	}

	t := w.sub(v).cross(s.sub(w)) / rxs
	u := w.sub(v).cross(r.sub(v)) / rxs
	if (t >= 0 && t <= 1) && (u >= 0 && u <= 1) {
		return true
	}
	return false
}

func task1(org vec, vs []vec, bounds [4]float64) int {
	thetas := []float64{0.5 * math.Pi, 0, 1.5 * math.Pi, math.Pi}
	path := make(map[complex128]bool, 124)
	ok := true
	dir := 0
	for ok {
		for _, v := range vs {
			if ok = incercepts(
				org,
				org.project(v.sub(org).length()+0.5, thetas[dir]),
				v.project(-0.5, thetas[(dir+1)%2]),
				v.project(0.5, thetas[(dir+1)%2]),
			); !ok {
				continue
			}
			for i := 0; i < 2; i++ {
				if org[i] != v[i] {
					continue
				}
				a, b := org[(i+1)%2], v[(i+1)%2]
				if a > b {
					a, b = b+1, a
				}
				for j := a; j < b; j++ {
					if i == 0 {
						path[complex(org[i], j)] = true
					} else {
						path[complex(j, org[i])] = true
					}
				}
			}
			org = org.project(v.sub(org).length()-1, thetas[dir])
			org[0], org[1] = math.Round(org[0]), math.Round(org[1])
			dir = (dir + 1) % 4
			break
		}
		if !ok {
			for i := 0; i < 2; i++ {
				if dir%2 != i {
					continue
				}
				a, b := org[(i+1)%2], bounds[dir]
				if a > b {
					a, b = b+1, a
				}
				for j := a; j < b; j++ {
					if i == 0 {
						path[complex(j, org[i])] = true
					} else {
						path[complex(org[i], j)] = true
					}
				}
			}
		}
	}

	return len(path)
}

func parse(path string) (vec, []vec, [4]float64) {
	buff, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	ns := make([]float64, 2)
	vs := make([]vec, 0, 124)
	org := vec{0, 0}
	for i, buff := range bytes.Split(buff, []byte{'\n'}) {
		if buff = bytes.TrimSpace(buff); len(buff) == 0 {
			continue
		}
		if ns[0] == 0 {
			ns[0] = float64(len(buff))
		}
		for j, b := range buff {
			switch b {
			case '#':
				vs = append(vs, vec{float64(j), float64(i)})
			case '^':
				org = vec{float64(j), float64(i)}
			}
		}
		ns[1]++
	}

	for i := range vs {
		vs[i][0] = vs[i][0] - org[0]
		vs[i][1] = org[1] - vs[i][1]
	}

	return vec{0, 0}, vs, [4]float64{org[1], ns[0] - org[0], org[1] - ns[1], -org[0]}
}

func main() {
	origin, points, bounds := parse("input.txt")
	fmt.Printf("Task %v: %v\n", 1, task1(origin, points, bounds))
}
