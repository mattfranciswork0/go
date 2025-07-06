package graph

import (
	"context"
	"reflect"
	"testing"

	"github.com/mattfranciswork0/go/graph/model"
)

func Test_queryResolver_Albums(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		r       *queryResolver
		args    args
		want    []*model.Album
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.Albums(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("queryResolver.Albums() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("queryResolver.Albums() = %v, want %v", got, tt.want)
			}
		})
	}
}
