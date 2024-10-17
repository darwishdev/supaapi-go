package supaapigo

type Env string

type SupaapiConfig struct {
	ProjectRef     string
	ApiKey         string
	ServiceRoleKey string
	Port           uint32
	Env            Env
}
