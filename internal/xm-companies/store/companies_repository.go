package store

import (
	"context"
	"xm-companies/internal/xm-companies/model"
)

type CompaniesRepository interface {
	Create(ctx context.Context, company *model.Company) error
	Update(ctx context.Context, company *model.Company) error
	Delete(ctx context.Context, name string) error
	FindBy(ctx context.Context, filter model.SearchFilter) ([]*model.Company, error)
}
