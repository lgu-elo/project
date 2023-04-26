package project

import (
	"github.com/lgu-elo/project/internal/project/model"
	"github.com/pkg/errors"
)

type (
	IService interface {
		GetAllProjects() ([]*model.Project, error)
		GetProjectById(id int) (*model.Project, error)
		UpdateProject(project *model.Project) (*model.Project, error)
		DeleteProject(id int) error
		CreateProject(creds *model.Project) error
	}

	service struct {
		repo Repository
	}
)

func NewService(repo Repository) IService {
	return &service{repo}
}

func (s *service) DeleteProject(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetAllProjects() ([]*model.Project, error) {
	projects, err := s.repo.GetAll()
	if err != nil {
		return nil, errors.Wrap(err, "can't get all projects")
	}

	return projects, nil

}

func (s *service) UpdateProject(project *model.Project) (*model.Project, error) {
	if err := s.repo.Update(project); err != nil {
		return nil, errors.Wrap(err, "can't update project")
	}

	project, err := s.GetProjectById(project.ID)
	if err != nil {
		return nil, err
	}

	return project, nil
}

func (s *service) CreateProject(project *model.Project) error {
	err := s.repo.Create(project)
	if err != nil {
		return errors.Wrap(err, "can't create project")
	}

	return nil
}

func (s *service) GetProjectById(id int) (*model.Project, error) {
	project, err := s.repo.GetById(id)
	if err != nil {
		return nil, errors.Wrap(err, "project not found")
	}

	return project, nil
}
