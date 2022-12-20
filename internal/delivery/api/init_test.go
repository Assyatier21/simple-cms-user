package api

import (
	"cms/internal/repository/postgres"
	mock_repo "cms/mock/repository/postgres"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestNew(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_repo.NewMockRepository(ctrl)

	type args struct {
		repository postgres.Repository
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Success",
			args: args{
				repository: mockRepository,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.repository)
			_, ok := got.(Handler)
			if !ok {
				t.Errorf("Not Handler interface")
			}
		})
	}
}
