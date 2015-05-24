package metalarchives

import (
	"reflect"
	"testing"

	"github.com/StalkR/goircbot/lib/metal"
)

func TestSearch(t *testing.T) {
	for _, tt := range []struct {
		name string
		want []metal.Band
	}{
		{
			name: "Convergence",
			want: []metal.Band{
				metal.Band{Name: "Convergence", Genre: "Atmospheric/Industrial Dark/Death Metal", Country: "Austria"},
				metal.Band{Name: "Convergence", Genre: "Melodic Death Metal (early), Nu-metal/Alternative Rock (later)", Country: "Italy"},
				metal.Band{Name: "Convergence from Within", Genre: "Death Metal", Country: "United States"},
			},
		},
		{
			name: "Psycroptic",
			want: []metal.Band{
				metal.Band{Name: "Psycroptic", Genre: "Technical Death Metal", Country: "Australia"},
			},
		},
		{
			name: "sdfgsdfg",
			want: nil,
		},
	} {
		got, err := Search(tt.name)
		if err != nil {
			t.Errorf("Search(%s): err: %v", tt.name, err)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Search(%s): got %s; want %s", tt.name, got, tt.want)
		}
	}
}