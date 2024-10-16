package supaapigo

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/supabase-community/auth-go"
	"github.com/supabase-community/auth-go/types"
	storage_go "github.com/supabase-community/storage-go"
)

// Define constants for valid values
const (
	DEV            Env    = "DEV"
	PROD           Env    = "PROD"
	AUTH_SUFFIX    string = "auth/v1/"
	STORAGE_SUFFIX string = "storage/v1/"
)

type Supaapi struct {
	AuthClient            auth.Client
	OAuthLoginCallback    string
	OAuthRegisterCallback string

	StorageClient storage_go.Client
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
		StorageClient:         *storageClient,
		OAuthLoginCallback:    config.OAuthLoginCallback,
		OAuthRegisterCallback: config.OAuthRegisterCallback,
		AuthClient:            authClient,
	}
}
func (s *Supaapi) UserCreateUpdate(createUpdateReq types.AdminUpdateUserRequest) (*types.User, error) {
	if createUpdateReq.UserID == uuid.Nil {
		user, err := s.AuthClient.AdminCreateUser(types.AdminCreateUserRequest{
			Aud:          createUpdateReq.Aud,
			Role:         createUpdateReq.Role,
			Email:        createUpdateReq.Email,
			Phone:        createUpdateReq.Phone,
			Password:     &createUpdateReq.Password,
			EmailConfirm: createUpdateReq.EmailConfirm,
			PhoneConfirm: createUpdateReq.PhoneConfirm,
			UserMetadata: createUpdateReq.UserMetadata,
			AppMetadata:  createUpdateReq.AppMetadata,
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

func (s *Supaapi) GoogleLogin() (*types.AuthorizeResponse, error) {
	resp, err := s.AuthClient.Authorize(types.AuthorizeRequest{
		RedirectTo: s.OAuthLoginCallback,
		Provider:   "google",
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
