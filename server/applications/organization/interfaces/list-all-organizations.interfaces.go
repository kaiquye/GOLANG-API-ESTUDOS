package interfaces

import (
	"fullVagas/domain/services/response"
)

type IListAllOrganizations interface {
	Run() response.Template
}
