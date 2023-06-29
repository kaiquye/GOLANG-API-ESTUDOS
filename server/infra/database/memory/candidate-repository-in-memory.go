package repositoryInMemory

import "fullVagas/domain/models"

type CandidateRepositoryInMemory struct {
	database []models.Candidate
}

func NewCandidateRepositoryInMemory() *CandidateRepositoryInMemory {
	return &CandidateRepositoryInMemory{
		database: make([]models.Candidate, 0),
	}
}

func (rep *CandidateRepositoryInMemory) Apply(data models.Candidate) (*models.Candidate, error) {
	rep.database = append(rep.database, data)
	return &data, nil
}
