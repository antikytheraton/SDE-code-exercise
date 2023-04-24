package shipments

import (
	"sort"
)

type Assignation struct {
	SS     float64
	Driver string
	Dest   string
}

func AssignRoutes(destinations, drivers []string) []Assignation {
	assignations := generateScores(destinations, drivers)
	return bestScores(assignations)
}

func generateScores(destinations, drivers []string) []Assignation {
	var out []Assignation
	for _, dest := range destinations {
		for _, driver := range drivers {
			out = append(out, Assignation{
				Dest:   dest,
				Driver: driver,
				SS:     SS(dest, driver),
			})
		}
	}
	return out
}

func bestScores(in []Assignation) []Assignation {
	sort.Slice(in, func(i, j int) bool {
		return in[i].SS > in[j].SS
	})

	unique := make(map[string]bool)
	var out []Assignation
	for _, a := range in {
		_, destOk := unique[a.Dest]
		_, drivOk := unique[a.Driver]
		if !destOk && !drivOk {
			out = append(out, a)
			unique[a.Dest] = true
			unique[a.Driver] = true
		}
	}
	return out
}
