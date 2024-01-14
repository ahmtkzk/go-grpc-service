package main

import (
	"context"
	"go-grpc-service/api"
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
	return &api.GetMovieListResponse{
		MovieList: []*api.Movie{
			{
				Id:               1,
				OriginalTitle:    "Demo Movie 1",
				Overview:         "Demo Overview 1",
				Title:            "Demo Title 1",
				OriginalLanguage: "Demo Language 1",
				Adult:            true,
			},
			{
				Id:               2,
				OriginalTitle:    "Demo Movie 2",
				Overview:         "Demo Overview 2",
				Title:            "Demo Title 2",
				OriginalLanguage: "Demo Language 2",
				Adult:            false,
			},
		},
	}, nil
}

func (s *movieServiceServer) GetMovie(ctx context.Context, request *api.GetMovieRequest) (*api.GetMovieResponse, error) {
	return &api.GetMovieResponse{
		Movie: &api.Movie{
			Id:               request.Id,
			OriginalTitle:    "Demo Movie",
			Overview:         "Demo Overview",
			Title:            "Demo Title",
			OriginalLanguage: "Demo Language",
			Adult:            true,
		},
	}, nil
}
