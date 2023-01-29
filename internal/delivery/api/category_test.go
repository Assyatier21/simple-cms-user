package api

import (
	mock_usecase "cms-user/mock/usecase"
	m "cms-user/models"
	"database/sql"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
)

func Test_handler_GetCategoryTree(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock_usecase.NewMockUsecaseHandler(ctrl)

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
				data := []m.Category{
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
				}
				var categories []interface{}
				for _, v := range data {
					categories = append(categories, v)
				}
				mockUsecase.EXPECT().GetCategoryTree(gomock.Any()).Return(categories, nil)
			},
		},
		{
			name: "usecase error",
			args: args{
				method: http.MethodGet,
				path:   "/categories",
			},
			wants: wants{
				statusCode: http.StatusInternalServerError,
			},
			mock: func() {
				data := []m.Category{
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
				}
				var categories []interface{}
				for _, v := range data {
					categories = append(categories, v)
				}
				mockUsecase.EXPECT().GetCategoryTree(gomock.Any()).Return(categories, errors.New("usecase error"))
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
				usecase: mockUsecase,
			}

			if err := h.GetCategoryTree(c); err != nil {
				t.Errorf("handler.GetCategoryTree() error = %v", err)
			}

			assert.Equal(t, tt.wants.statusCode, rec.Code)
		})
	}
}
func Test_handler_GetCategoryDetails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock_usecase.NewMockUsecaseHandler(ctrl)

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
				data := []m.Category{
					{
						Id:        1,
						Title:     "category 1",
						Slug:      "category-1",
						CreatedAt: "2022-12-01 20:29:00",
						UpdatedAt: "2022-12-01 20:29:00",
					},
				}
				var category []interface{}
				for _, v := range data {
					category = append(category, v)
				}
				mockUsecase.EXPECT().GetCategoryDetails(gomock.Any(), 1).Return(category, nil)
			},
		},
		{
			name: "error id not an integer",
			args: args{
				method: http.MethodGet,
				path:   "/category?id=not_number",
			},
			wants: wants{
				statusCode: http.StatusBadRequest,
			},
			mock: func() {},
		},
		{
			name: "error sql no rows affected",
			args: args{
				method: http.MethodGet,
				path:   "/category?id=1",
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				data := []m.Category{
					{
						Id:        1,
						Title:     "category 1",
						Slug:      "category-1",
						CreatedAt: "2022-12-01 20:29:00",
						UpdatedAt: "2022-12-01 20:29:00",
					},
				}
				var category []interface{}
				for _, v := range data {
					category = append(category, v)
				}
				mockUsecase.EXPECT().GetCategoryDetails(gomock.Any(), 1).Return(category, sql.ErrNoRows)
			},
		},
		{
			name: "error usecase",
			args: args{
				method: http.MethodGet,
				path:   "/category?id=1",
			},
			wants: wants{
				statusCode: http.StatusInternalServerError,
			},
			mock: func() {
				data := []m.Category{
					{
						Id:        1,
						Title:     "category 1",
						Slug:      "category-1",
						CreatedAt: "2022-12-01 20:29:00",
						UpdatedAt: "2022-12-01 20:29:00",
					},
				}
				var category []interface{}
				for _, v := range data {
					category = append(category, v)
				}
				mockUsecase.EXPECT().GetCategoryDetails(gomock.Any(), 1).Return(category, errors.New("usecase error"))
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
				usecase: mockUsecase,
			}

			if err := h.GetCategoryDetails(c); err != nil {
				t.Errorf("handler.GetCategoryDetails() error = %v", err)
			}

			assert.Equal(t, tt.wants.statusCode, rec.Code)
		})
	}
}
