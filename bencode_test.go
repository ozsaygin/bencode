package bencode

import (
	"fmt"
	"reflect"
	"testing"
)

// Strings are length-prefixed base ten followed by a colon and the string. For example 4:spam corresponds to 'spam'.

// Integers are represented by an 'i' followed by the number in base 10 followed by an 'e'. For example i3e corresponds to 3 and i-3e corresponds to -3. Integers have no size limitation. i-0e is invalid. All encodings with a leading zero, such as i03e, are invalid, other than i0e, which of course corresponds to 0.

// Lists are encoded as an 'l' followed by their elements (also bencoded) followed by an 'e'. For example l4:spam4:eggse corresponds to ['spam', 'eggs'].

// Dictionaries are encoded as a 'd' followed by a list of alternating keys and their corresponding values followed by an 'e'. For example, d3:cow3:moo4:spam4:eggse corresponds to {'cow': 'moo', 'spam': 'eggs'} and d4:spaml1:a1:bee corresponds to {'spam': ['a', 'b']}. Keys must be strings and appear in sorted order (sorted as raw strings, not alphanumerics).

// func TestDecode(t *testing.T) {
// 	type args struct {
// 		data string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want map[string]interface{}
// 	}{
// 		{
// 			name: "single string key value map",
// 			args: args{"d3:cow3:mooe"},
// 			want: map[string]interface{}{
// 				"cow": "moo",
// 			},
// 		},
// 		{
// 			name: "multiple string key value map",
// 			args: args{"d3:cat4:meow4:frog5:croak6:pigeon3:cooe"},
// 			want: map[string]interface{}{
// 				"cat":    "meow",
// 				"frog":   "croak",
// 				"pigeon": "coo",
// 			},
// 		},
// 		{
// 			name: "map containing single str int key value",
// 			args: args{"d3:cati5ee"},
// 			want: map[string]interface{}{
// 				"cat": 5,
// 			},
// 		},
// 		{
// 			name: "map multiple str int key value",
// 			args: args{"d3:cat4:frog2:11"},
// 			want: map[string]interface{}{
// 				"cat":    3,
// 				"frog":   "11",
// 				"pigeon": "coo",
// 				"lion":   18,
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := Decode(tt.args.data); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Decode() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestEncode(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{"spam"},
		},
		{
			name: "",
			args: args{"spam"},
		},
		{
			name: "",
			args: args{map[string]interface{}{
				"name":     "Cat",
				"genre":    "Scific",
				"year":     2015,
				"actrices": []string{"A", "B", "C"},
				"maxLifeSpan": 16
			}},
		},
		// {
		// 	name: "simple string",
		// 	args: args{[]string{"A", "B", "C"}},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Encode(tt.args.v)
			fmt.Println(result)

			if got := Encode(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decode() = %v, want %v", got, tt.want)

			}
		})
	}
}
