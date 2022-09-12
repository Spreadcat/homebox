package services

import "github.com/hay-kot/content/backend/internal/repo"

type AllServices struct {
	User     *UserService
	Admin    *AdminService
	Location *LocationService
	Labels   *LabelService
	Items    *ItemService
}

func NewServices(repos *repo.AllRepos) *AllServices {
	return &AllServices{
		User:     &UserService{repos},
		Admin:    &AdminService{repos},
		Location: &LocationService{repos},
		Labels:   &LabelService{repos},
		Items: &ItemService{
			repo:     repos,
			filepath: "/tmp/content",
		},
	}
}
