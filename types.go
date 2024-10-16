package supaapigo

type Env string

type SupaapiConfig struct {
	ProjectRef            string
	ApiKey                string
	ServiceRoleKey        string
	OAuthRegisterCallback string
	OAuthLoginCallback    string
	Port                  uint32
	Env                   Env
}
