package config

// Config is a type which describes the properties which can be in the config
type Config struct {
	// ServiceRegistry *ServiceRegistryConfig `json:"serviceregistry"`
	InstanceID string `json:"instanceId"`
	Name       string `json:"name"`

	AccessToken       string   `json:"access_token" doc:"Bearer access token."`
	RefreshToken      string   `json:"refresh_token" doc:"Offline or refresh token."`
	MasAuthURL        string   `json:"mas_auth_url"`
	MasAccessToken    string   `json:"mas_access_token"`
	MasRefreshToken   string   `json:"mas_refresh_token"`
	APIUrl            string   `json:"api_url" doc:"URL of the API gateway. The value can be the complete URL or an alias. The valid aliases are 'production', 'staging' and 'integration'."`
	AuthURL           string   `json:"auth_url" doc:"URL of the authentication server"`
	ClientID          string   `json:"client_id" doc:"OpenID client identifier."`
	Insecure          bool     `json:"insecure" doc:"Enables insecure communication with the server. This disables verification of TLS certificates and host names."`
	Scopes            []string `json:"scopes" doc:"OpenID scope. If this option is used it will replace completely the default scopes. Can be repeated multiple times to specify multiple scopes."`
	DevPreviewEnabled bool     `json:"dev_preview_enabled" doc:"Enables Developer preview commands"`
}

func (c *Config) HasServiceRegistry() bool {
	return c.Name != ""
}
