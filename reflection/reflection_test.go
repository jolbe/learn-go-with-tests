package reflection

import (
	"slices"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         any
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 42},
			[]string{"Chris"},
		},
		{
			"nested fields",
			Person{
				"Chris",
				Profile{33, "New York"},
			},
			[]string{"Chris", "New York"},
		},
		{
			"pointers to things",
			&Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"slices",
			[]Profile{
				{33, "London"},
				{34, "New York"},
			},
			[]string{"London", "New York"},
		},
		{
			"arrays",
			[2]Profile{
				{33, "London"},
				{34, "New York"},
			},
			[]string{"London", "New York"},
		},
	}

	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {
			var got []string
			Walk(tt.Input, func(input string) {
				got = append(got, input)
			})

			// if len(got) != len(tt.ExpectedCalls) {
			// 	t.Errorf("wrong number of function calls, got %d; want %d", len(got), 1)
			// }

			if !slices.Equal(got, tt.ExpectedCalls) {
				t.Errorf("got %v; want %v", got, tt.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]any{
			"number":  1,
			"boolean": true,
			"cow":     "Moo",
			"sheep":   "Baa",
		}

		var got []string
		Walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Berlin"}
			aChannel <- Profile{34, "Katowice"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "Katowice"}

		Walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !slices.Equal(got, want) {
			t.Errorf("got %v; want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{34, "Katowice"}
		}

		var got []string
		want := []string{"Berlin", "Katowice"}

		Walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !slices.Equal(got, want) {
			t.Errorf("got %v; want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()

	if !slices.Contains(haystack, needle) {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}
