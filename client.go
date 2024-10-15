package supaapigo

import (
	"github.com/darwishdev/auth-go"
	storage_go "github.com/supabase-community/storage-go"
)

type Supaapi struct {
	AuthClient    auth.Client
	StorageClient *storage_go.Client
}

func NewSupaapi(projectRefrence string, storageUrl, apiKey string, isDevelopment bool) Supaapi {
	storageClient := storage_go.NewClient(storageUrl, apiKey, nil)
	authClient := auth.New(projectRefrence, apiKey, isDevelopment)
	return Supaapi{
		StorageClient: storageClient,
		AuthClient:    authClient,
	}
}
