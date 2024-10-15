package supaapigo

import (
	"github.com/supabase-community/auth-go"
	storage_go "github.com/supabase-community/storage-go"
)

type SupaapiUsecaseInterface interface {
}
type SupaapiUsecase struct {
	authClient    *auth.Client
	storageClient *storage_go.Client
}

func NewSupaapiUsecase(projectRefrence string, storageUrl, apiKey string) SupaapiUsecaseInterface {
	storageClient := storage_go.NewClient(storageUrl, apiKey, nil)
	authClient := auth.New(projectRefrence, apiKey)
	return &SupaapiUsecase{
		storageClient: storageClient,
		authClient:    &authClient,
	}
}
