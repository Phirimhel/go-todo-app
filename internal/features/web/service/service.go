package web_service

import web_fs_repository "github.com/Phirimhel/go-todo-app/internal/features/web/repository/file_system"

type WebService interface {
	GetMainPage() ([]byte, error)
	GetFile(path string) ([]byte, error)
}

type webService struct {
	webRepository web_fs_repository.WebRepository
}

func NewWebService(webRepo web_fs_repository.WebRepository) *webService {
	return &webService{
		webRepository: webRepo,
	}
}
