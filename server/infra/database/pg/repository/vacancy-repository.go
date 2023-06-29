package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"fullVagas/domain/models"
	"fullVagas/infra/database/pg"
	"github.com/google/uuid"
)

type VacancyRepository struct {
	postgres *sql.DB
}

func NewVacancyRepository() *VacancyRepository {
	return &VacancyRepository{
		postgres: pg.GetConnection(),
	}
}

func (rep *VacancyRepository) Create(data *models.Vacancy) (*models.Vacancy, error) {
	fmt.Println(data.Salary_range)
	stms, err := rep.postgres.Prepare("INSERT INTO tb_vacancy (vacancy_id, title, body, location, level, salary_range, organization_id) VALUES ($1,$2, $3, $4, $5, $6, $7)")
	if err != nil {
		fmt.Println("error during user creation: ", err.Error())
		return nil, errors.New("error creating vacancy, please contact an administrator. internal: 500")
	}
	defer stms.Close()
	_, err = stms.Exec(data.Vacancy_id, data.Tilte, data.Body, data.Location, data.Level, data.Salary_range, data.Organization.Org_id)
	if err != nil {
		fmt.Println("error during user creation: ", err.Error())
		return nil, errors.New("error creating vacancy, please contact an administrator. internal: 500")
	}

	return data, nil
}

func (rep *VacancyRepository) FindById(id uuid.UUID) (*models.Vacancy, error) {
	stms, err := rep.postgres.Prepare(`
		SELECT vacancy_id, title, body, location FROM tb_Vacancy WHERE vacancy_id = $1
	`)
	if err != nil {
		fmt.Println("error during user creation: ", err.Error())
		return nil, errors.New("error find vacancy, please contact an administrator. internal: 500")
	}
	defer stms.Close()

	vacancy := models.Vacancy{}
	row, _ := stms.Query(id)
	for row.Next() {
		row.Scan(&vacancy.Vacancy_id, &vacancy.Tilte, &vacancy.Body, &vacancy.Location)
	}
	if vacancy.Tilte != "" && vacancy.Body != "" {
		return &vacancy, nil
	}
	return nil, nil
}

func (rep *VacancyRepository) FindAll() ([]models.Vacancy, error) {
	stms, err := rep.postgres.Prepare("select vacancy_id, title, body, location, level, salary_range FROM tb_vacancy")
	if err != nil {
		fmt.Println("error during find vacancy: ", err.Error())
		return nil, err
	}
	defer stms.Close()

	row, err := stms.Query()
	vacancys := make([]models.Vacancy, 0)
	for row.Next() {
		vacancy := models.Vacancy{}
		row.Scan(&vacancy.Vacancy_id, &vacancy.Tilte, &vacancy.Body, &vacancy.Location, &vacancy.Level, &vacancy.Salary_range)
		vacancys = append(vacancys, vacancy)
	}

	return vacancys, nil
}
