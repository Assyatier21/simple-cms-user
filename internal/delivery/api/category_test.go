package api

import (
	mock_repo "cms/mock/repository/postgres"
	m "cms/models"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
)

func Test_handler_GetCategoryTree(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_repo.NewMockRepository(ctrl)

	type args struct {
		method string
		path   string
	}
	type wants struct {
		statusCode int
	}
	tests := []struct {
		name  string
		args  args
		wants wants
		mock  func()
	}{
		{
			name: "success",
			args: args{
				method: http.MethodGet,
				path:   "/categories",
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				mockRepository.EXPECT().GetCategoryTree(gomock.Any()).Return(
					[]m.Category{
						{
							Id:        1,
							Title:     "category 1",
							Slug:      "category-1",
							CreatedAt: "2022-12-01 20:29:00",
							UpdatedAt: "2022-12-01 20:29:00",
						},
						{
							Id:        2,
							Title:     "category 2",
							Slug:      "category-2",
							CreatedAt: "2022-12-01 20:29:00",
							UpdatedAt: "2022-12-01 20:29:00",
						},
					}, nil)
			},
		},
		{
			name: "repository error",
			args: args{
				method: http.MethodGet,
				path:   "/categories",
			},
			wants: wants{
				statusCode: http.StatusInternalServerError,
			},
			mock: func() {
				mockRepository.EXPECT().GetCategoryTree(gomock.Any()).Return(
					[]m.Category{
						{
							Id:        1,
							Title:     "category 1",
							Slug:      "category-1",
							CreatedAt: "2022-12-01 20:29:00",
							UpdatedAt: "2022-12-01 20:29:00",
						},
						{
							Id:        2,
							Title:     "category 2",
							Slug:      "category-2",
							CreatedAt: "2022-12-01 20:29:00",
							UpdatedAt: "2022-12-01 20:29:00",
						},
					}, errors.New("repository error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(tt.args.method, tt.args.path, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			tt.mock()

			h := &handler{
				repository: mockRepository,
			}

			if err := h.GetCategoryTree(c); err != nil {
				t.Errorf("handler.GetArticles() error = %v", err)
			}

			assert.Equal(t, tt.wants.statusCode, rec.Code)
		})
	}
}

func Test_handler_GetCategoryByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_repo.NewMockRepository(ctrl)

	type args struct {
		method string
		path   string
	}
	type wants struct {
		statusCode int
	}
	tests := []struct {
		name  string
		args  args
		wants wants
		mock  func()
	}{
		{
			name: "success",
			args: args{
				method: http.MethodGet,
				path:   "/category?id=1",
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				mockRepository.EXPECT().GetCategoryByID(gomock.Any(), 1).Return(m.Category{
					Id:        1,
					Title:     "category 1",
					Slug:      "category-1",
					CreatedAt: "2022-12-01 20:29:00",
					UpdatedAt: "2022-12-01 20:29:00",
				}, nil)
			},
		},
		{
			name: "error id not an integer",
			args: args{
				method: http.MethodGet,
				path:   "/category?id=not_integer",
			},
			wants: wants{
				statusCode: http.StatusBadRequest,
			},
			mock: func() {},
		},
		{
			name: "repository error",
			args: args{
				method: http.MethodGet,
				path:   "/category?id=1",
			},
			wants: wants{
				statusCode: http.StatusInternalServerError,
			},
			mock: func() {
				mockRepository.EXPECT().GetCategoryByID(gomock.Any(), 1).Return(m.Category{
					Id:        1,
					Title:     "category 1",
					Slug:      "category-1",
					CreatedAt: "2022-12-01 20:29:00",
					UpdatedAt: "2022-12-01 20:29:00",
				}, errors.New("repository error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(tt.args.method, tt.args.path, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			tt.mock()

			h := &handler{
				repository: mockRepository,
			}

			if err := h.GetCategoryByID(c); err != nil {
				t.Errorf("handler.GetCategoryByID() error = %v", err)
			}

			assert.Equal(t, tt.wants.statusCode, rec.Code)
		})
	}
}
