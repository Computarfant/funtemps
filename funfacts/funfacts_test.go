package funfacts

import (
	"fmt"
	"testing"
	//"reflect"
)

func TestGetFunFacts(t *testing.T) {
	tests := []struct {
		subject string
	}{
		{"sun"},
		{"luna"},
		{"terra"},
	}

	for _, tt := range tests {
		facts := GetFunFacts(tt.subject)

		if len(facts) != 2 {
			t.Errorf("Expected 2 facts, but got %d", len(facts))
		}

		for _, fact := range facts {
			fmt.Printf("%s%s\n", fact, tt.subject)
		}
	}
}

// /*
// *

// 	Mal for TestGetFunFacts funksjonen.
// 	Definer korrekte typer for input og want,
// 	og sette korrekte testverdier i slice tests.
// */
// func TestGetFunFacts(t *testing.T) {
// 	type test struct {
// 		input // her må du skrive riktig type for input
// 		want  // her må du skrive riktig type for returverdien
// 	}

// 	// Her må du legge inn korrekte testverdier
// 	//tests := []test{
// 	//  {input: , want: },
// 	//}

// 	for _, tc := range tests {
// 		got := GetFunFacts(tc.input)
// 		if !reflect.DeepEqual(tc.want, got) {
// 			t.Errorf("expected: %v, got: %v", tc.want, got)
// 		}
// 	}
// }
