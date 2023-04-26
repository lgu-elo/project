package project

import (
	"context"

	"github.com/lgu-elo/project/internal/project/model"
	"github.com/lgu-elo/project/pkg/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type (
	projectHandler struct {
		service IService
		log     *logrus.Logger
		server  *grpc.Server
	}

	IHandler interface {
		GetAllProjects(c context.Context, _ *pb.Empty) (*pb.ProjectsList, error)
		GetProjectById(c context.Context, project *pb.ProjectWithID) (*pb.Project, error)
		UpdateProject(c context.Context, project *pb.Project) (*pb.Project, error)
		DeleteProject(c context.Context, project *pb.ProjectWithID) (*pb.Empty, error)
		CreateProject(c context.Context, project *pb.Project) (*pb.Empty, error)
	}
)

func NewHandler(service IService, log *logrus.Logger, server *grpc.Server) IHandler {
	return &projectHandler{service, log, server}
}

func (h *projectHandler) GetAllProjects(c context.Context, _ *pb.Empty) (*pb.ProjectsList, error) {
	projects, err := h.service.GetAllProjects()
	if err != nil {
		return nil, err
	}

	var pbProjects pb.ProjectsList
	for _, project := range projects {
		pbProjects.Projects = append(pbProjects.Projects, &pb.Project{
			Id:          int64(project.ID),
			Name:        project.Name,
			Description: project.Description,
		})
	}

	return &pbProjects, nil
}

func (h *projectHandler) GetProjectById(c context.Context, request *pb.ProjectWithID) (*pb.Project, error) {
	project, err := h.service.GetProjectById(int(request.Id))
	if err != nil {
		return nil, err
	}

	return &pb.Project{
		Id:          int64(project.ID),
		Name:        project.Name,
		Description: project.Description,
	}, nil
}

func (h *projectHandler) UpdateProject(c context.Context, request *pb.Project) (*pb.Project, error) {
	project, err := h.service.UpdateProject(&model.Project{
		ID:          int(request.Id),
		Name:        request.Name,
		Description: request.Description,
	})
	if err != nil {
		return nil, err
	}

	return &pb.Project{
		Id:          int64(project.ID),
		Name:        request.Name,
		Description: request.Description,
	}, nil
}

func (h *projectHandler) DeleteProject(c context.Context, project *pb.ProjectWithID) (*pb.Empty, error) {
	if err := h.service.DeleteProject(int(project.Id)); err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (h *projectHandler) CreateProject(c context.Context, project *pb.Project) (*pb.Empty, error) {
	err := h.service.CreateProject(&model.Project{
		ID:          int(project.Id),
		Name:        project.Name,
		Description: project.Description,
	})

	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}
