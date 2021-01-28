package bencode

import (
	"reflect"
	"testing"
)

// {
// 	name: "single string key value map",
// 	args: args{"d3:cow3:mooe"},
// 	want: map[string]interface{}{
// 		"cow": "moo",
// 	},
// },
// {
// 	name: "multiple string key value map",
// 	args: args{"d3:cat4:meow4:frog5:croak6:pigeon3:cooe"},
// 	want: map[string]interface{}{
// 		"cat":    "meow",
// 		"frog":   "croak",
// 		"pigeon": "coo",
// 	},
// },
// {
// 	name: "map containing single str int key value",
// 	args: args{"d3:cati5ee"},
// 	want: map[string]interface{}{
// 		"cat": 5,
// 	},
// },
// {
// 	name: "map multiple str int key value",
// 	args: args{"d3:cat4:frog2:11"},
// 	want: map[string]interface{}{
// 		"cat":    3,
// 		"frog":   "11",
// 		"pigeon": "coo",
// 		"lion":   18,
// 	},
// },

func TestMarshal(t *testing.T) {

	// All test cases are valid.
	// TODO: Add more some invalid test cases.

	type args struct {
		v interface{}
	}

	type testCase struct {
		name string
		args args
		want []byte
	}

	tests := []testCase{
		{
			name: "Simple Positive Integer",
			args: args{14},
			want: []byte("i14e"),
		},
		{
			name: "Simple Negative Integer",
			args: args{-14},
			want: []byte("i-14e"),
		},
		{
			name: "Simple String Expression",
			args: args{"bencode"},
			want: []byte("7:bencode"),
		},
		{
			name: "Simple Map Expression",
			args: args{map[string]string{
				"Cat":    "Felis catus",
				"Jaguar": "Panthera onca",
				"Lion":   "Panthera leo",
			}},
			want: []byte("d3:Cat11:Felis catus6:Jaguar13:Panthera onca4:Lion12:Panthera leoe"),
		},
		{
			name: "Simple String List Expression",
			args: args{[]string{"A", "B", "C"}},
			want: []byte("l1:A1:B1:Ce"),
		},
		// {
		// 	name: "Simple Integer List Expression",
		// 	args: args{[]int{13, 0, -23, 93}},
		// 	want: []byte("d1:0i13e1:1i0e1:2i-23e1:3i93ee"),
		// },
	}

	// Run unit tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := Marshal(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Marshal() = %v, want %v, err = %v", string(got), string(tt.want), err)
			}
		})
	}
}

// func TestUnmarshal(t *testing.T) {
// 	type args struct {
// 		data []byte
// 		v    interface{}
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantErr bool
// 	}{
// 		{
// 			name: "simple string",
// 			args: args{
// 				data: []byte("3:cat"),
// 			},
// 			wantErr: true,
// 		},
// 		{
// 			name: "simple string map",
// 			args: args{
// 				data: []byte("d3:cat4:meow3:cow3:mooe"),
// 			},
// 			wantErr: true,
// 		},
// 		{
// 			name: "simple dictionary with integer value",
// 			args: args{
// 				data: []byte("d3:cat4:meow3:cowi42e"),
// 			},
// 			wantErr: true,
// 		},
// 		{
// 			name: "list",
// 			args: args{
// 				data: []byte("l4:spam4:eggsee"),
// 			},
// 			wantErr: true,
// 		},
// 		{
// 			name: "map with list",
// 			args: args{
// 				data: []byte("d4:spaml1:a1:bee"),
// 			},
// 			wantErr: true,
// 		},
// 		{
// 			name: "negatif integer",
// 			args: args{
// 				data: []byte("i-3e"),
// 			},
// 			wantErr: true,
// 		},
// 		{
// 			name: "positive integer",
// 			args: args{
// 				data: []byte("i34e"),
// 			},
// 			wantErr: true,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if err := Unmarshal(tt.args.data, tt.args.v); (err != nil) != tt.wantErr {
// 				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
