package postgresql

import (
	"context"
	"database/sql"
	"github.com/leporo/sqlf"
	"xm-companies/internal/xm-companies/model"
)

type CompaniesRepository struct {
	db *sql.DB
}

const (
	TableCompanies = "companies"
	ColumnCode     = "code"
	ColumnName     = "name"
	ColumnCountry  = "country"
	ColumnPhone    = "phone"
	ColumnWebsite  = "website"
)

func NewCompaniesRepositoryProvider(db *sql.DB) *CompaniesRepository {
	sqlf.SetDialect(sqlf.PostgreSQL)
	return &CompaniesRepository{
		db: db,
	}
}

func (cr *CompaniesRepository) Create(ctx context.Context, company *model.Company) error {
	_, err := sqlf.InsertInto(TableCompanies).
		Set(ColumnCode, company.Code).
		Set(ColumnName, company.Name).
		Set(ColumnCountry, company.Country).
		Set(ColumnPhone, company.Phone).
		Set(ColumnWebsite, company.Website).ExecAndClose(ctx, cr.db)
	return err
}

func (cr *CompaniesRepository) Update(ctx context.Context, company *model.Company) error {
	_, err := sqlf.Update(TableCompanies).
		Set(ColumnName, company.Name).
		Set(ColumnCountry, company.Country).
		Set(ColumnPhone, company.Phone).
		Set(ColumnWebsite, company.Website).
		Where("code = ?", company.Code).ExecAndClose(ctx, cr.db)
	return err
}

func (cr *CompaniesRepository) Delete(ctx context.Context, code string) error {
	_, err := sqlf.DeleteFrom(TableCompanies).Where("code = ?", code).ExecAndClose(ctx, cr.db)
	return err
}

func (cr *CompaniesRepository) FindBy(ctx context.Context, filter model.SearchFilter) ([]*model.Company, error) {
	result := make([]*model.Company, 0)

	err := generateSearchQuery(filter).
		QueryAndClose(ctx, cr.db, func(rows *sql.Rows) {
			comp := &model.Company{}
			err := rows.Scan(
				&comp.Code,
				&comp.Name,
				&comp.Country,
				&comp.Phone,
				&comp.Website,
			)
			if err == nil {
				result = append(result, comp)
			}
		})
	return result, err
}

func generateSearchQuery(f model.SearchFilter) *sqlf.Stmt {
	q := sqlf.From("companies").
		Select("code, name, country, phone, website")
	if f.Code != "" {
		q = q.Where("code = ?", f.Code)
	}
	if f.Name != "" {
		q = q.Where("name LIKE ?", "%"+f.Name+"%")
	}
	if f.Country != "" {
		q = q.Where("country LIKE ?", "%"+f.Country+"%")
	}
	if f.Phone != "" {
		q = q.Where("phone LIKE ?", "%"+f.Phone+"%")
	}
	if f.Website != "" {
		q = q.Where("website LIKE ?", "%"+f.Website+"%")
	}
	return q
}
