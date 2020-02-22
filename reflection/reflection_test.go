package main

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
	t.Run("Table reflection Tests", func(t *testing.T) {
		cases := []struct {
			Name          string
			Input         interface{}
			ExpectedCalls []string
		}{
			{
				"Struct with one string field",
				struct {
					Name string
				}{"Simon"},
				[]string{"Simon"},
			},
			{
				"Struct with two string fields",
				struct {
					Name string
					City string
				}{"Simon", "Chester"},
				[]string{"Simon", "Chester"},
			},
			{
				"Struct with non string field",
				struct {
					Name string
					Age  int
				}{"Simon", 26},
				[]string{"Simon"},
			},
			{
				"Struct with nested fields",
				Person{"Simon", Profile{26, "Chester"}},
				[]string{"Simon", "Chester"},
			},
			{
				"Pointers to things",
				&Person{"Simon", Profile{26, "Chester"}},
				[]string{"Simon", "Chester"},
			},
			{
				"Slices",
				[]Profile{
					{26, "Chester"},
					{27, "Chester"},
				},
				[]string{"Chester", "Chester"},
			},
			{
				"Arrays",
				[2]Profile{
					{26, "Chester"},
					{27, "Chester"},
				},
				[]string{"Chester", "Chester"},
			},
			{
				"Maps",
				map[string]string{
					"foo": "bar",
					"fab": "baz",
				},
				[]string{"bar", "baz"},
			},
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
	})

	t.Run("with Maps", func(t *testing.T) {
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
}

func assertContains(t *testing.T, haystack []string, needle string) {
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it did not", haystack, needle)
	}
}
