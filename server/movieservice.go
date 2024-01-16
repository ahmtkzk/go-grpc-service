package main

import (
	"context"
	"go-grpc-service/api"
	r "go-grpc-service/repository"
)

type Movie struct {
	Id               int    `json:"id"`
	OriginalTitle    string `json:"original_title"`
	Overview         string `json:"overview"`
	Title            string `json:"title"`
	OriginalLanguage string `json:"original_language"`
	Adult            bool   `json:"adult"`
}

type movieServiceServer struct {
	api.UnimplementedMovieServiceServer
}

func (s *movieServiceServer) GetMovieList(ctx context.Context, request *api.GetMovieListRequest) (*api.GetMovieListResponse, error) {
	var movieList []*api.Movie
	repo := r.NewGormRepository()
	defer func(repo *r.GormRepository) {
		err := repo.Close()
		if err != nil {

		}
	}(repo)

	repo.Find(&movieList)
	return &api.GetMovieListResponse{MovieList: movieList}, nil
}

func (s *movieServiceServer) GetMovie(ctx context.Context, request *api.GetMovieRequest) (*api.GetMovieResponse, error) {
	var movie *api.Movie
	repo := r.NewGormRepository()
	defer func(repo *r.GormRepository) {
		err := repo.Close()
		if err != nil {

		}
	}(repo)

	repo.Find(&movie, "id = ?", request.Id)
	return &api.GetMovieResponse{Movie: movie}, nil
}
