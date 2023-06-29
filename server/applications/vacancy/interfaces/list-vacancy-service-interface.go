package interfaces

import "fullVagas/domain/services/response"

type IListVacancyService interface {
	Run() response.Template
}
