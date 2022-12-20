package postgres

import (
	m "cms/models"
	"context"
	"errors"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_repository_GetCategoryTree(t *testing.T) {
	ctx := context.Background()

	db, sqlMock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    []m.Category
		wantErr bool
		mock    func()
	}{
		{
			name: "scan error",
			args: args{
				ctx: ctx,
			},
			want:    nil,
			wantErr: true,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "slug", "created_at", "updated_at"}).
					AddRow("WRONG TYPE ID", "category 1", "category-1", "2022-12-01 20:29:00", "2022-12-01 20:29:00").
					AddRow("WRONG TYPE ID", "category 2", "category-2", "2022-12-01 20:29:00", "2022-12-01 20:29:00")
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM cms_category ORDER BY id`)).WillReturnRows(rows)
			},
		},
		{
			name: "success",
			args: args{
				ctx: ctx,
			},
			want: []m.Category{
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
			},
			wantErr: false,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "slug", "created_at", "updated_at"}).
					AddRow(1, "category 1", "category-1", "2022-12-01T20:29:00Z", "2022-12-01T20:29:00Z").
					AddRow(2, "category 2", "category-2", "2022-12-01T20:29:00Z", "2022-12-01T20:29:00Z")
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM cms_category ORDER BY id`)).WillReturnRows(rows)
			},
		},
		{
			name: "success with empty categories",
			args: args{
				ctx: ctx,
			},
			want:    []m.Category{},
			wantErr: false,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "slug", "created_at", "updated_at"})
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM cms_category ORDER BY id`)).WillReturnRows(rows)
			},
		},
		{
			name: "query error",
			args: args{
				ctx: ctx,
			},
			want:    nil,
			wantErr: true,
			mock: func() {
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WillReturnError(errors.New("query error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			r := &repository{
				db: db,
			}
			got, err := r.GetCategoryTree(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetCategoryTree() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetCategoryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_GetCategoryByID(t *testing.T) {
	ctx := context.Background()

	db, sqlMock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		args    args
		want    m.Category
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				id:  1,
			},
			want: m.Category{
				Id:        1,
				Title:     "category 1",
				Slug:      "category-1",
				CreatedAt: "2022-12-01 20:29:00",
				UpdatedAt: "2022-12-01 20:29:00",
			},
			wantErr: false,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "slug", "created_at", "updated_at"}).
					AddRow(1, "category 1", "category-1", "2022-12-01T20:29:00Z", "2022-12-01T20:29:00Z")
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM cms_category WHERE id = $1`)).WillReturnRows(rows)
			},
		},
		{
			name: "scan error",
			args: args{
				ctx: ctx,
				id:  1,
			},
			want:    m.Category{},
			wantErr: true,
			mock: func() {
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM cms_category WHERE id = $1`)).WillReturnError(errors.New("error while scanning"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			r := &repository{
				db: db,
			}
			got, err := r.GetCategoryByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetCategoryByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetCategoryByID() = %v, want %v", got, tt.want)
			}
		})
	}
}