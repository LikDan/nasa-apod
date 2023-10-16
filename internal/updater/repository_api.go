package updater

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type RepositoryApi interface {
	GetTodayInfo(ctx context.Context) (APODApiResponse, error)
}

type repositoryApi struct {
	client http.Client
	apiKey string
}

func NewRepositoryApi(apiKey string) RepositoryApi {
	return &repositoryApi{
		client: http.Client{},
		apiKey: apiKey,
	}
}

func (r *repositoryApi) GetTodayInfo(ctx context.Context) (APODApiResponse, error) {
	request, err := http.NewRequest(http.MethodGet, "https://api.nasa.gov/planetary/apod?api_key="+r.apiKey, nil)
	if err != nil {
		return APODApiResponse{}, err
	}

	request.WithContext(ctx)

	response, err := r.client.Do(request)
	if err != nil {
		return APODApiResponse{}, err
	}

	str, err := io.ReadAll(response.Body)
	if err != nil {
		return APODApiResponse{}, err
	}

	var responseData APODApiResponse
	if err = json.Unmarshal(str, &responseData); err != nil {
		return APODApiResponse{}, err
	}

	return responseData, nil
}
