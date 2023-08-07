package mathlb

import "testing"

func TestAddInt(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test AddInt success 1",
			args: args{3, 4},
			want: 7,
		},
		{
			name: "Test AddInt success 2",
			args: args{10, 4},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddInt(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("AddInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinusInt(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test MinusInt positive",
			args: args{3, 2},
			want: 1,
		},
		{
			name: "Test MinusInt negative",
			args: args{3, 5},
			want: -2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinusInt(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("MinusInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiplyInt(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "TestMultiplyInt positive",
			args: args{3, 4},
			want: 12,
		},
		{
			name: "TestMultiplyInt with positive result",
			args: args{-3, -4},
			want: 12,
		},
		{
			name: "TestMultiplyInt with negative result",
			args: args{-3, 4},
			want: -12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MultiplyInt(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("MultiplyInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivideInt(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test DivideInt rational division",
			args: args{12, 6},
			want: 2,
		},
		{
			name: "Test DivideInt irrational division",
			args: args{6, 7},
			want: 0,
		},
		{
			name: "Test DivideInt with residue",
			args: args{7, 6},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DivideInt(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("DivideInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModule(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Module with zero result",
			args: args{10, 5},
			want: 0,
		},
		{
			name: "Test Module with no zero result",
			args: args{10, 3},
			want: 1,
		},
		{
			name: "Test Module with more than one",
			args: args{7, 4},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Module(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Module() = %v, want %v", got, tt.want)
			}
		})
	}
}
