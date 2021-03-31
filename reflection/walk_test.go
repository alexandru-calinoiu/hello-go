package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{"Struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"}},
		{"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"}},
		{"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"}},
		{"Nested fields",
			Person{"Chris", Profile{33, "London"}},
			[]string{"Chris", "London"}},
		{"Pointers to things",
			&Person{"Calin", Profile{33, "Sibiu"}},
			[]string{"Calin", "Sibiu"}},
		{"Slices",
			[]Profile{
				{33, "London"},
				{34, "Sibiu"},
			},
			[]string{"London", "Sibiu"}},
		{"Arrays",
			[2]Profile{
				{33, "London"},
				{34, "Sibiu"},
			},
			[]string{"London", "Sibiu"}},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string

			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	t.Run("with channels", func(t *testing.T) {
		channel := make(chan Profile)
		go func() {
			channel <- Profile{33, "Sibiu"}
			channel <- Profile{34, "Berlin"}
			close(channel)
		}()

		var got []string
		want := []string{"Sibiu", "Berlin"}

		walk(channel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		function := func() (Profile, Profile) {
			return Profile{33, "Sibiu"}, Profile{34, "Berlin"}
		}

		var got []string
		want := []string{"Sibiu", "Berlin"}

		walk(function, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()

	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
			break
		}
	}

	if !contains {
		t.Errorf("expected %+v contain %q but it didn't", haystack, needle)
	}
}
