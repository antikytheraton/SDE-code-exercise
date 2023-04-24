package shipments_test

import (
	"code_exercise/shipments"
	"fmt"
	"testing"
)

func TestGCF(t *testing.T) {
	t.Parallel()
	for _, tt := range []struct {
		a    int
		b    int
		want int
	}{
		{a: 98, b: 56, want: 14},
		{a: 5, b: 7, want: 1},
		{a: 2252, b: 4082, want: 2},
		{a: 5, b: 10, want: 5},
	} {
		tt := tt
		t.Run(fmt.Sprintf(`GDC of %d and %d is %d`, tt.a, tt.b, tt.want),
			func(t *testing.T) {
				t.Parallel()
				got := shipments.GCF(tt.a, tt.b)
				if tt.want != got {
					t.Errorf(`want: %d, got: %d`, tt.want, got)
				}
			})
	}
}
