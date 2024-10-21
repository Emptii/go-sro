package navmesh

import "github.com/Emptii/go-sro/framework/math"

type ObjectCell struct {
	*math.Triangle
	Index int
	Flag  uint16
}
