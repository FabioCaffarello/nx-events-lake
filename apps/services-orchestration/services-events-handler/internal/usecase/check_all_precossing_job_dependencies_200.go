package usecase

import (
	apiClient "libs/golang/services/api-clients/services-staging-handler/client"
)

type CheckAllJobDependenciesStatus200UseCase struct {
	stagingHandlerAPIClient apiClient.Client
}

func NewCheckAllJobDependenciesStatus200UseCase() *CheckAllJobDependenciesStatus200UseCase {
     return &CheckAllJobDependenciesStatus200UseCase{
        stagingHandlerAPIClient: *apiClient.NewClient(),
     }
}

func (la *CheckAllJobDependenciesStatus200UseCase) Execute(id string) (bool, error) {
     jobDependencies, err := la.stagingHandlerAPIClient.ListOneProcessingJobDependenciesById(id)
     if err != nil {
          return false, err
     }
     for _, jobDependency := range jobDependencies.JobDependencies {
          if extractStatusCodeRange(jobDependency.StatusCode) != "2XX" {
               return false, nil
          }
     }
     return true, nil
}


func extractStatusCodeRange(statusCode int) string {
     if statusCode >= 200 && statusCode < 300 {
          return "2XX"
     } else if statusCode >= 400 && statusCode < 500 {
          return "4XX"
     } else if statusCode >= 500 && statusCode < 600 {
          return "5XX"
     }
     return "invalid"
}
