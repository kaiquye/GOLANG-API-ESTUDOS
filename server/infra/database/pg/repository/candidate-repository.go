package repository

import (
	"database/sql"
	"fmt"
	"fullVagas/domain/models"
	"fullVagas/infra/database/pg"
)

type CandidateRepository struct {
	db *sql.DB
}

func NewCandidateRepository() *CandidateRepository {
	return &CandidateRepository{
		db: pg.GetConnection(),
	}
}

func (rep *CandidateRepository) Apply(data models.Candidate) (*models.Candidate, error) {
	stms, err := rep.db.Prepare(`
		INSERT INTO public.tb_candidate
		(candidate_id, "document", email, "level", "location", resume, vacancy_id, score)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8);
   `)
	if err != nil {
		fmt.Println("error when applying to a vacancy")
		return nil, err
	}
	defer stms.Close()

	_, err = stms.Exec(data.Candidate_id, data.Document, data.Email, data.Level, data.Location, data.Resume, data.Vacancy.Vacancy_id, data.Score)
	if err != nil {
		fmt.Println("error when applying to a vacancy")
		return nil, err
	}

	return &data, nil
}
