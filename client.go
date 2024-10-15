package supaapigo

import (
	"github.com/supabase-community/auth-go"
	storage_go "github.com/supabase-community/storage-go"
)

type Supaapi struct {
	AuthClient    *auth.Client
	StorageClient *storage_go.Client
}

func NewSupaapi(projectRefrence string, storageUrl, apiKey string) Supaapi {
	storageClient := storage_go.NewClient(storageUrl, apiKey, nil)
	authClient := auth.New(projectRefrence, apiKey)
	return Supaapi{
		StorageClient: storageClient,
		AuthClient:    &authClient,
	}
}
