package utils

import (
	"errors"
	"fmt"
	"time"
)

var (
	PathArticles   = "/articles"
	PathArticle    = "/article"
	PathCategories = "/categories"
	PathCategory   = "/category"
	ErrNotFound    = errors.New("data not found")
	NoRowsAffected = errors.New("no rows affected")
	jakartaLoc, _  = time.LoadLocation("Asia/Jakarta")
	TimeNow        = fmt.Sprintf("%d-%d-%d %d:%d:%d", time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second())
)
