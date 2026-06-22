package web_service

import (
	"fmt"
	"os"
	"path"
)

func (s *webService) GetMainPage() ([]byte, error) {

	htmlFilePath := path.Join(
		os.Getenv("PROJECT_ROOT"),
		"/public/index.html",
	)

	html, err := s.webRepository.GetFile(htmlFilePath)
	if err != nil {
		return nil, fmt.Errorf("[service]: failed to get main page %w", err)
	}

	return html, nil
}

func (s *webService) GetFile(path string) ([]byte, error) {
	return nil, nil
}
