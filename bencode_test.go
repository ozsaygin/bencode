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
		name    string
		args    args
		want    []byte
		wantErr bool
	}

	tests := []testCase{
		{
			name:    "Simple Positive Integer",
			args:    args{14},
			want:    []byte("i14e"),
			wantErr: false,
		},
		{
			name:    "Simple Negative Integer",
			args:    args{-14},
			want:    []byte("i-14e"),
			wantErr: false,
		},
		{
			name:    "Simple String Expression",
			args:    args{"bencode"},
			want:    []byte("7:bencode"),
			wantErr: false,
		},
		{
			name: "Simple Map Expression",
			args: args{map[string]string{
				"Cat":    "Felis catus",
				"Jaguar": "Panthera onca",
				"Lion":   "Panthera leo",
			}},
			want:    []byte("d3:Cat11:Felis catus6:Jaguar13:Panthera onca4:Lion12:Panthera leoe"),
			wantErr: false,
		},
		{
			name:    "Simple String List Expression",
			args:    args{[]string{"A", "B", "C"}},
			want:    []byte("l1:A1:B1:Ce"),
			wantErr: false,
		},
		{
			name:    "Simple Integer List Expression",
			args:    args{[]int{13, 0, -23, 93}},
			want:    []byte("li13ei0ei-23ei93ee"),
			wantErr: false,
		},
	}

	// Run unit tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Marshal(tt.args.v)
			if !reflect.DeepEqual(got, tt.want) && (err != nil) != tt.wantErr {
				t.Errorf("Marshal() = %v, want %v, error = %v, wantError = %v", got, tt.want, (err != nil), tt.wantErr)
			}
		})
	}
}

// func TestUnmarshal(t *testing.T) {

// 	type args struct {
// 		data []byte
// 		v    interface{}
// 	}

// 	type testCase struct {
// 		name    string
// 		args    args
// 		wantErr bool
// 	}

// 	tests := []testCase{
// 		{
// 			name: "Simple String",
// 			args: args{
// 				data: []byte("3:cat"),
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "Simple String Map",
// 			args: args{
// 				data: []byte("d3:cat4:meow3:cow3:mooe"),
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "Simple Map With Integer Value",
// 			args: args{
// 				data: []byte("d3:cat4:meow3:cowi42e"),
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "List",
// 			args: args{
// 				data: []byte("l4:spam4:eggsee"),
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "Map With List",
// 			args: args{
// 				data: []byte("d4:spaml1:a1:bee"),
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "Negatif Integer",
// 			args: args{
// 				data: []byte("i-3e"),
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "Positive Integer",
// 			args: args{
// 				data: []byte("i34e"),
// 			},
// 			wantErr: false,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if err := Unmarshal(tt.args.data, tt.args.v); (err != nil) != tt.wantErr {
// 				t.Errorf("args.v = %v Unmarshal() error = %v, wantErr %v", tt.args.v, err, tt.wantErr)
// 			}
// 		})
// 	}
// }
