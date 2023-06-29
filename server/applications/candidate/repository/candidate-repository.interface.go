package IRepository

import "fullVagas/domain/models"

type ICandidateRepository interface {
	Apply(data models.Candidate) (*models.Candidate, error)
}
