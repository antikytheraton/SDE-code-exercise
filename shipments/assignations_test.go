package shipments_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"code_exercise/shipments"
)

func TestAssignations(t *testing.T) {
	t.Parallel()
	for _, tt := range []struct {
		testName     string
		destinations []string
		drivers      []string
		want         []shipments.Assignation
	}{
		{
			testName: "test 1",
			destinations: []string{
				"First Avenue Street #459, Nashville, Tennessee",
				"Nueva Santa Maria, Azcapotzalco, 02800 Ciudad de Mexico",
				"Cerrada Otono #15, 04102",
			},
			drivers: []string{
				"Daniel Martinez",
				"Alan Moore",
				"Jason Statham",
				"Mad Max",
			},
			want: []shipments.Assignation{
				{
					SS:     13.5,
					Driver: "Daniel Martinez",
					Dest:   "Cerrada Otono #15, 04102",
				},
				{
					SS:     9.0,
					Driver: "Jason Statham",
					Dest:   "First Avenue Street #459, Nashville, Tennessee",
				},
				{
					SS:     7.5,
					Driver: "Alan Moore",
					Dest:   "Nueva Santa Maria, Azcapotzalco, 02800 Ciudad de Mexico",
				},
			},
		},
	} {
		tt := tt
		t.Run(tt.testName, func(t *testing.T) {
			t.Parallel()

			got := shipments.AssignRoutes(tt.destinations, tt.drivers)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}
