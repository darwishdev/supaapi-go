package supaapigo

import (
	"fmt"

	storage_go "github.com/darwishdev/storage-go"
	"github.com/google/uuid"
	"github.com/supabase-community/auth-go"
	"github.com/supabase-community/auth-go/types"
)

// Define constants for valid values
const (
	DEV            Env    = "DEV"
	PROD           Env    = "PROD"
	AUTH_SUFFIX    string = "auth/v1"
	STORAGE_SUFFIX string = "storage/v1"
)

type Supaapi struct {
	StorageUrl     string
	serviceRoleKey string
	anonApiKey    string
	AuthClient     auth.Client
	StorageClient  *storage_go.Client
}

func NewSupaapi(config SupaapiConfig) Supaapi {
	var customURL string
	var baseURL string
	storageURL := fmt.Sprintf("https://%s.supabase.co/%s", config.ProjectRef, STORAGE_SUFFIX)
	authClient := auth.New(config.ProjectRef, config.ApiKey).WithToken(config.ServiceRoleKey)
	if config.Env == DEV {
		baseURL = fmt.Sprintf("http://localhost:%d", config.Port)
		storageURL = fmt.Sprintf("%s/%s", baseURL, STORAGE_SUFFIX)
		customURL = fmt.Sprintf("%s/%s", baseURL, AUTH_SUFFIX)
		authClient = authClient.WithCustomAuthURL(customURL)
	}
	storageClient := storage_go.NewClient(storageURL, config.ServiceRoleKey, nil)
	return Supaapi{
		StorageClient:  storageClient,
		anonApiKey : config.ApiKey
		StorageUrl:     storageURL,
		serviceRoleKey: config.ServiceRoleKey,
		AuthClient:     authClient,
	}
}

func (s *Supaapi) RestartStorageClient() {
	storageClient := storage_go.NewClient(s.StorageUrl, s.serviceRoleKey, nil)
	s.StorageClient = storageClient
}

func (s *Supaapi) UserCreateUpdate(createUpdateReq types.AdminUpdateUserRequest) (*types.User, error) {
	if createUpdateReq.UserID == uuid.Nil {
		user, err := s.AuthClient.Signup(types.SignupRequest{
			Email:    createUpdateReq.Email,
			Phone:    createUpdateReq.Phone,
			Password: createUpdateReq.Password,
		})
		if err != nil {

			return nil, err
		}
		return &user.User, nil
	}
	user, err := s.AuthClient.AdminUpdateUser(createUpdateReq)
	if err != nil {
		return nil, err
	}
	return &user.User, nil
}
func (s *Supaapi) AuthorizeToken(token string) (*types.User, error) {
	resp, err := s.AuthClient.WithToken(token).GetUser()
	if err != nil {
		return nil, err
	}
	return &resp.User, nil
}

func (s *Supaapi) ProviderLogin(provider types.Provider, redirectUrl string) (*types.AuthorizeResponse, error) {
	resp, err := s.AuthClient.Authorize(types.AuthorizeRequest{
		RedirectTo: redirectUrl,
		Provider:   provider,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
