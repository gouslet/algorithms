package tries

import (
	"reflect"
	"testing"
)

func Test_tries_arr_Put(t *testing.T) {
	type args struct {
		key string
		val any
	}
	tests := []struct {
		name string
		tr   *tries_arr
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tr.Put(tt.args.key, tt.args.val)
		})
	}
}

func Test_tries_arr_Get(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		tr   *tries_arr
		args args
		want any
	}{
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tries_arr.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
