package routes

import (
	"fullVagas/applications/candidate"
	ICandidate "fullVagas/applications/candidate/interfaces"
	"fullVagas/applications/organization"
	IOrganization "fullVagas/applications/organization/interfaces"
	"fullVagas/applications/vacancy"
	IVacancy "fullVagas/applications/vacancy/interfaces"
	"fullVagas/infra/http/dto"
	"fullVagas/infra/http/middlewares"
	"github.com/gin-gonic/gin"
)

type Server struct {
	authMiddleware            middlewares.AuthMiddleware
	createOrganizationService IOrganization.ICreateOrganizationService
	loginOrganizationSer      IOrganization.ILoginOrganizationService
	listAllOrganizations      IOrganization.IListAllOrganizations
	createVacancyService      IVacancy.ICreateVacancyService
	listVacancysService       IVacancy.IListVacancyService
	applyCandidate            ICandidate.IApplyToVacancyService
}

func Target() *Server {
	return &Server{
		createOrganizationService: organization.NewCreateOrganizationService(),
		loginOrganizationSer:      organization.NewLoginOrganizationService(),
		createVacancyService:      vacancy.NewCreateVacancyService(),
		applyCandidate:            candidate.NewApplyToVacancyService(),
		authMiddleware:            *middlewares.NewAuthMiddleware(),
		listAllOrganizations:      organization.NewListAllOrganization(),
		listVacancysService:       vacancy.NewListVacancy(),
	}
}

func (server *Server) Init(context *gin.Engine) *gin.Engine {
	func() {
		context.POST("/organization", func(ctx *gin.Context) {
			CreateOrganizationDto := dto.CreateOrganizationDto{}
			err := CreateOrganizationDto.Validate(ctx)

			if err != nil {
				ctx.JSON(400, gin.H{
					"message": err.Error(),
				})

				return
			}

			result := server.createOrganizationService.Run(&CreateOrganizationDto)

			ctx.JSON(result.Status, result)
			return
		})
		context.POST("/organization/login", func(ctx *gin.Context) {
			loginOrganizationDto := dto.LoginOrganizationDto{}
			err := loginOrganizationDto.Validate(ctx)

			if err != nil {
				ctx.JSON(400, gin.H{
					"message": err.Error(),
				})
				return
			}

			result := server.loginOrganizationSer.Run(loginOrganizationDto)

			ctx.JSON(result.Status, result)
			return
		})
		context.GET("/organizations", func(ctx *gin.Context) {
			result := server.listAllOrganizations.Run()
			ctx.JSON(result.Status, result)
			return
		})
	}()

	func() {
		context.POST("/vacancy", server.authMiddleware.ValidateAccessToken(), func(context *gin.Context) {
			CreateVacancyDto := dto.CreateVacancyDto{}
			err := CreateVacancyDto.Validate(context)
			if err != nil {
				context.JSON(400, gin.H{
					"message": err.Error(),
				})
				return
			}

			result := server.createVacancyService.Run(&CreateVacancyDto)

			context.JSON(result.Status, result)
			return
		})
		context.GET("/vacancy", func(context *gin.Context) {
			result := server.listVacancysService.Run()
			context.JSON(result.Status, result)
			return
		})
	}()

	func() {
		context.POST("/candidate", func(context *gin.Context) {
			applyToVacancyDto := dto.ApplyToVacancyDto{}
			err := applyToVacancyDto.Validate(context)
			if err != nil {
				context.JSON(400, gin.H{
					"message": err.Error(),
				})
				return
			}

			result := server.applyCandidate.Run(applyToVacancyDto)

			context.JSON(result.Status, result)
			return
		})

	}()
	return context
}
