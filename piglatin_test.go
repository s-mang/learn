package learn

import "testing"

type testCase struct {
	description string
	in, want    string
}

// pigLatinTestCases
//
// About Table Driven Tests:
// this type of structuring your tests is called 'Table Driven Testing'.
// It is called that because you can see below it's just like
// a db table with columns (description, in, out) and data or
// values.
// This is the preferred way of testing in Go because, it "reduces the amount of repetitive code compared to repeating the same code for each test and makes it straightforward to add more test cases."
// See https://blog.golang.org/subtests
//
// Essentially, we put all the data up front, and then all the
// functional stuff below, separated.
// The alternative would be something like this:
//
// func TestPigLatinOneWordLowerCase(t *testing.T) {
// 	if Translate("hello") != "ellohay" {
//		t.Fatal("...")
//	}
// }
// ...
// func TestPigLatinPhraseWithCapitalLetter(t *testing.T) {
// 	if Translate("Hey you") != "Eyhay ouyay" {
//		t.Fatal("...")
//	}
// }
//

// I came up with 8 test cases, let's see how many you can get!

var pigLatinTestCases = []testCase{
	{
		description: "one word, lower case",
		in:          "hello",
		want:        "ellohay",
	},
}

func TestPigLatinTranslate(t *testing.T) {
	// This is the basic way of doing table driven testing
	// in Go:
	//
	// for _, tc := range pigLatinTestCases {
	// 	out := Translate(tc.in)
	// 	if out != tc.want {
	// 		t.Errorf("%s: got '%s',	want: '%s'", tc.description, out, tc.want)
	// 	}
	// }

	// Feel free to uncomment the above to see how these two methods
	// print out tests differently to the console.

	// Subtests:
	// This is the new way, with 'subtests' (came out with Go 1.7).
	// See https://blog.golang.org/subtests
	//
	// This way is slightly preferred:
	for _, tc := range pigLatinTestCases {
		t.Run(tc.description, func(t *testing.T) {
			out := Translate(tc.in)
			if out != tc.want {
				t.Fatalf("got '%s',	want: '%s'", out, tc.want)
			}
		})
	}
}
