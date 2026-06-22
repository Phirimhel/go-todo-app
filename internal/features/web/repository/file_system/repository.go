package web_fs_repository

type WebRepository interface {
	GetFile(path string) ([]byte, error)
}

type webRepository struct{}

func NewWebRepository() *webRepository {
	return &webRepository{}
}
