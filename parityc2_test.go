package parityc

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParityEvenCheck2D(t *testing.T) {
	t.Parallel()

	type args struct {
		i    [][]int
		want [][]int
	}

	cases := []struct {
		name string
		args args
	}{
		{
			name: "collect1",
			args: args{
				i: [][]int{
					{1, 0, 0, 0, 1},
					{1, 0, 0, 0},
				},
				want: [][]int{
					{1, 0, 0, 0, 1},
					{1, 0, 0, 0},
				},
			},
		},
		{
			name: "collect2",
			args: args{
				i: [][]int{
					{1, 1, 0, 1, 1},
					{0, 0, 1, 0, 1},
					{1, 1, 1, 1, 0},
					{1, 1, 0, 0, 0},
					{1, 1, 0, 0},
				},
				want: [][]int{
					{1, 1, 0, 1, 1},
					{0, 0, 1, 0, 1},
					{1, 1, 1, 1, 0},
					{1, 1, 0, 0, 0},
					{1, 1, 0, 0},
				},
			},
		},
		{
			name: "1bit miss1",
			args: args{
				i: [][]int{
					{1, 1, 0, 1, 0},
					{0, 1, 0, 1},
				},
				want: [][]int{
					{0, 1, 0, 1, 0},
					{0, 1, 0, 1},
				},
			},
		},
		{
			name: "1bit miss2",
			args: args{
				i: [][]int{
					{1, 1, 0, 1, 1},
					{0, 0, 1, 0, 1},
					{0, 1, 1, 1, 0},
					{1, 1, 0, 0, 0},
					{1, 1, 0, 0},
				},
				want: [][]int{
					{1, 1, 0, 1, 1},
					{0, 0, 1, 0, 1},
					{1, 1, 1, 1, 0},
					{1, 1, 0, 0, 0},
					{1, 1, 0, 0},
				},
			},
		},
	}

	for n, tt := range cases {
		tt := tt
		n := n
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			t.Parallel()
			got, _ := ParityEvenCheck2D(tt.args.i)
			if diff := cmp.Diff(tt.args.want, got); diff != "" {
				t.Errorf("ParityEvenCheck2D does not correct check: (-got +want)\n%s", diff)
			}
		})
	}
}

func TestParityOddCheck2D(t *testing.T) {
	t.Parallel()

	type args struct {
		i    [][]int
		want [][]int
	}

	cases := []struct {
		name string
		args args
	}{
		{
			name: "collect1",
			args: args{
				i: [][]int{
					{0, 1, 1, 1, 0},
					{1, 0, 0, 0},
				},
				want: [][]int{
					{0, 1, 1, 1, 0},
					{1, 0, 0, 0},
				},
			},
		},
		{
			name: "collect2",
			args: args{
				i: [][]int{
					{1, 1, 0, 1, 0},
					{0, 0, 1, 0, 0},
					{1, 1, 1, 1, 1},
					{1, 1, 0, 0, 1},
					{0, 0, 1, 1},
				},
				want: [][]int{
					{1, 1, 0, 1, 0},
					{0, 0, 1, 0, 0},
					{1, 1, 1, 1, 1},
					{1, 1, 0, 0, 1},
					{0, 0, 1, 1},
				},
			},
		},
		{
			name: "1bit miss1",
			args: args{
				i: [][]int{
					{1, 1, 0, 1, 1},
					{0, 1, 1, 0},
				},
				want: [][]int{
					{1, 0, 0, 1, 1},
					{0, 1, 1, 0},
				},
			},
		},
		{
			name: "1bit miss2",
			args: args{
				i: [][]int{
					{1, 1, 0, 1, 0},
					{0, 0, 1, 0, 0},
					{0, 1, 1, 1, 1},
					{1, 1, 0, 0, 1},
					{1, 0, 0, 1},
				},
				want: [][]int{
					{1, 1, 0, 1, 0},
					{0, 0, 1, 0, 0},
					{0, 1, 0, 1, 1},
					{1, 1, 0, 0, 1},
					{1, 0, 0, 1},
				},
			},
		},
	}

	for n, tt := range cases {
		tt := tt
		n := n
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			t.Parallel()
			got, _ := ParityOddCheck2D(tt.args.i)
			if diff := cmp.Diff(tt.args.want, got); diff != "" {
				t.Errorf("ParityOddCheck2D does not correct check: (-got +want)\n%s", diff)
			}
		})
	}

}
