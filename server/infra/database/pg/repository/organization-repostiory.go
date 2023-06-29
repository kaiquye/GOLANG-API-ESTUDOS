package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"fullVagas/domain/models"
	"fullVagas/infra/database/pg"
	"github.com/google/uuid"
)

type OrganizationRepository struct {
	postgres *sql.DB
}

func NewOrganizationRepository() *OrganizationRepository {
	return &OrganizationRepository{
		postgres: pg.GetConnection(),
	}
}

func (rep *OrganizationRepository) Save(data *models.Organization) (*models.Organization, error) {
	fmt.Println(data.Password)
	stms, err := rep.postgres.Prepare("INSERT INTO tb_organization (org_id, name, document, logo, password) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		fmt.Println("error creating organization: ", err.Error())
		return nil, errors.New("error creating organization, please contact an administrator. internal: 500")
	}
	defer stms.Close()
	if _, err := stms.Exec(data.Org_id, data.Name, data.Document, data.Logo, data.Password); err != nil {
		fmt.Println("error creating organization: ", err.Error())
		return nil, errors.New("error creating organization, please contact an administrator. internal: 500")
	}

	return data, nil
}

func (rep *OrganizationRepository) Exist(document string) (bool, error) {
	stms, err := rep.postgres.Prepare("SELECT Org.Org_id from tb_organization as org WHERE org.document = $1")
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	defer stms.Close()

	row, err := stms.Query(document)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	if row.Next() {
		return true, nil
	}

	return false, nil
}

func (rep *OrganizationRepository) FindById(id uuid.UUID) (bool, error) {
	stms, err := rep.postgres.Prepare("select name from tb_organization where org_id = $1")
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	defer stms.Close()

	row, _ := stms.Query(id)
	if row.Next() {
		return true, nil
	}
	return false, nil
}

func (rep *OrganizationRepository) FindByDocument(document string) (*models.Organization, error) {
	stms, err := rep.postgres.Prepare("SELECT org.Org_id, org.Password, org.Name from tb_organization as org WHERE org.document = $1")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer stms.Close()

	row, err := stms.Query(document)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	org := models.Organization{}
	for row.Next() {
		row.Scan(&org.Org_id, &org.Password, &org.Name)
	}

	return &org, nil
}

func (rep *OrganizationRepository) FindAll() ([]models.Organization, error) {
	stms, err := rep.postgres.Prepare(`select org_id, name, logo, document from tb_Organization`)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer stms.Close()

	row, err := stms.Query()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	organizations := make([]models.Organization, 0)
	for row.Next() {
		organization := models.Organization{}
		row.Scan(&organization.Org_id, &organization.Name, &organization.Logo, &organization.Document)
		organizations = append(organizations, organization)
	}

	return organizations, nil
}
