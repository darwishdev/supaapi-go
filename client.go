package supaapigo

import (
	"github.com/supabase-community/auth-go"
	storage_go "github.com/supabase-community/storage-go"
)

type SupaapiInterface interface {
}
type Supaapi struct {
	authClient    *auth.Client
	storageClient *storage_go.Client
}

func NewSupaapi(projectRefrence string, storageUrl, apiKey string) SupaapiInterface {
	storageClient := storage_go.NewClient(storageUrl, apiKey, nil)
	authClient := auth.New(projectRefrence, apiKey)
	return &Supaapi{
		storageClient: storageClient,
		authClient:    &authClient,
	}
}
