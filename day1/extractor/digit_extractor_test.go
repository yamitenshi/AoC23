package extractor

import "testing"

func TestFirstAndLastDigitExtractor_extractFromString(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1abc2",
			args: args{input: "1abc2"},
			want: 12,
		},
		{
			name: "pqr3stu8vwx",
			args: args{input: "pqr3stu8vwx"},
			want: 38,
		},
		{
			name: "a1b2c3d4e5f",
			args: args{input: "a1b2c3d4e5f"},
			want: 15,
		},
		{
			name: "treb7uchet",
			args: args{input: "treb7uchet"},
			want: 77,
		},
		{
			name: "nodigitshere",
			args: args{input: "nodigitshere"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := FirstAndLastDigitExtractor{}
			if got := f.ExtractFromString(tt.args.input); got != tt.want {
				t.Errorf("ExtractFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getFirstDigitFromString(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1abc2",
			args: args{input: "1abc2"},
			want: 1,
		},
		{
			name: "pqr3stu8vwx",
			args: args{input: "pqr3stu8vwx"},
			want: 3,
		},
		{
			name: "a1b2c3d4e5f",
			args: args{input: "a1b2c3d4e5f"},
			want: 1,
		},
		{
			name: "treb7uchet",
			args: args{input: "treb7uchet"},
			want: 7,
		},
		{
			name: "nodigitshere",
			args: args{input: "nodigitshere"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFirstDigitFromString(tt.args.input); got != tt.want {
				t.Errorf("getFirstDigitFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getLastDigitFromString(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1abc2",
			args: args{input: "1abc2"},
			want: 2,
		},
		{
			name: "pqr3stu8vwx",
			args: args{input: "pqr3stu8vwx"},
			want: 8,
		},
		{
			name: "a1b2c3d4e5f",
			args: args{input: "a1b2c3d4e5f"},
			want: 5,
		},
		{
			name: "treb7uchet",
			args: args{input: "treb7uchet"},
			want: 7,
		},
		{
			name: "nodigitshere",
			args: args{input: "nodigitshere"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLastDigitFromString(tt.args.input); got != tt.want {
				t.Errorf("getLastDigitFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}
