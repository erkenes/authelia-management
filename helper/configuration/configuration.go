package configuration

import (
	"authelia-users/helper/basics"
	"fmt"
	"gopkg.in/yaml.v3"
)

type AutheliaServerStruct struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
type AutheliaLogStruct struct {
	Level      string `yaml:"level"`
	FilePath   string `yaml:"file_path"`
	KeepStdout bool   `yaml:"keep_stdout"`
}
type AutheliaNtpStruct struct {
	Address             string `yaml:"address"`
	Version             int    `yaml:"version"`
	MaxDesync           string `yaml:"max_desync"`
	DisableStartupCheck bool   `yaml:"disable_startup_check"`
	DisableFailure      bool   `yaml:"disable_failure"`
}
type AutheliaTotpStruct struct {
	Disable    bool   `yaml:"disable"`
	Issuer     string `yaml:"issuer"`
	Algorithm  string `yaml:"algorithm"`
	Digits     int    `yaml:"digits"`
	Period     int    `yaml:"period"`
	Skew       int    `yaml:"skew"`
	SecretSize int    `yaml:"secret_size"`
}
type AutheliaWebauthnStruct struct {
	Disable                         bool   `yaml:"disable"`
	DisplayName                     string `yaml:"display_name"`
	AttestationConveyancePreference string `yaml:"attestation_conveyance_preference"`
	UserVerification                string `yaml:"user_verification"`
	Timeout                         string `yaml:"timeout"`
}

type AutheliaAuthBackendPasswordStruct struct {
	Disable bool `yaml:"disable"`
}
type AutheliaAuthBackendFilePasswordStruct struct {
	Algorithm   string `yaml:"algorithm"`
	Iterations  int    `yaml:"iterations"`
	KeyLength   int    `yaml:"key_length"`
	SaltLength  int    `yaml:"salt_length"`
	Parallelism int    `yaml:"parallelism"`
	Memory      int    `yaml:"memory"`
}
type AutheliaAuthBackendFileStruct struct {
	Path     string                                `yaml:"path"`
	Password AutheliaAuthBackendFilePasswordStruct `yaml:"password"`
}
type AutheliaAuthenticationBackendStruct struct {
	PasswordReset AutheliaAuthBackendPasswordStruct `yaml:"password_reset"`
	File          AutheliaAuthBackendFileStruct     `yaml:"file"`
}

type AutheliaAccessControlRulesStruct struct {
	Domain    []string `yaml:"domain"`
	Policy    string   `yaml:"policy"`
	Subject   []string `yaml:"subject,omitempty"` // ToDo: Support multiple requirements: https://www.authelia.com/configuration/security/access-control/#subject
	Resources []string `yaml:"resources,omitempty"`
}
type AutheliaAccessControlStruct struct {
	DefaultPolicy string                             `yaml:"default_policy"`
	Rules         []AutheliaAccessControlRulesStruct `yaml:"rules"`
}

type AutheliaSessionRedisStruct struct {
	Host                     string `yaml:"host"`
	Port                     int    `yaml:"port"`
	DatabaseIndex            int    `yaml:"database_index"`
	MaximumActiveConnections int    `yaml:"maximum_active_connections"`
	MinimumIdleConnections   int    `yaml:"minimum_idle_connections"`
}
type AutheliaSessionStruct struct {
	Redis      AutheliaSessionRedisStruct `yaml:"redis"`
	Name       string                     `yaml:"name"`
	Expiration int                        `yaml:"expiration"`
	Inactivity int                        `yaml:"inactivity"`
	Domain     string                     `yaml:"domain"`
}

type AutheliaRegulationStruct struct {
	MaxRetries int    `yaml:"max_retries"`
	FindTime   string `yaml:"find_time"`
	BanTime    string `yaml:"ban_time"`
}

type AutheliaPasswordPolicyStandardStruct struct {
	Enabled          bool `yaml:"enabled"`
	MinLength        int  `yaml:"min_length"`
	MaxLength        int  `yaml:"max_length"`
	RequireUppercase bool `yaml:"require_uppercase"`
	RequireLowercase bool `yaml:"require_lowercase"`
	RequireNumber    bool `yaml:"require_number"`
	RequireSpecial   bool `yaml:"require_special"`
}
type AutheliaPasswordPolicyStruct struct {
	Standard AutheliaPasswordPolicyStandardStruct `yaml:"standard"`
}

type AutheliaStorageMysqlStruct struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
}
type AutheliaStorageStruct struct {
	Mysql AutheliaStorageMysqlStruct `yaml:"mysql"`
}

type AutheliaNotifierSmtpStruct struct {
	Username            string `yaml:"username"`
	Host                string `yaml:"host"`
	Port                int    `yaml:"port"`
	Sender              string `yaml:"sender"`
	Subject             string `yaml:"subject"`
	StartupCheckAddress string `yaml:"startup_check_address"`
}
type AutheliaNotifierStruct struct {
	DisableStartupCheck bool                       `yaml:"disable_startup_check"`
	Smtp                AutheliaNotifierSmtpStruct `yaml:"smtp"`
}

type AutheliaConfigStruct struct {
	Server                AutheliaServerStruct                `yaml:"server"`
	Log                   AutheliaLogStruct                   `yaml:"log"`
	DefaultRedirectionUrl string                              `yaml:"default_redirection_url"`
	Ntp                   AutheliaNtpStruct                   `yaml:"ntp"`
	Totp                  AutheliaTotpStruct                  `yaml:"totp"`
	WebAuthn              AutheliaWebauthnStruct              `yaml:"webauthn"`
	AuthenticationBackend AutheliaAuthenticationBackendStruct `yaml:"authentication_backend"`
	AccessControl         AutheliaAccessControlStruct         `yaml:"access_control"`
	Session               AutheliaSessionStruct               `yaml:"session"`
	Regulation            AutheliaRegulationStruct            `yaml:"regulation"`
	PasswordPolicy        AutheliaPasswordPolicyStruct        `yaml:"password_policy"`
	Storage               AutheliaStorageStruct               `yaml:"storage"`
	Notifier              AutheliaNotifierStruct              `yaml:"notifier"`
}

// getCurrentUsers Get all users
func getConfiguration() AutheliaConfigStruct {
	content := basics.ReadFile(".", "configuration.yml")

	var err error
	var m AutheliaConfigStruct

	err = yaml.Unmarshal(content, &m)
	if err != nil {
		panic(err)
	}

	return m
}

func saveToFile(configuration AutheliaConfigStruct) {
	yamlData, err := yaml.Marshal(configuration)

	if err != nil {
		fmt.Printf("Error while Marshaling. %v", err)
	}

	basics.WriteFile(".", "configuration.yml", yamlData, 0644)

}

// hasAccessRuleForDomain checks if a rule exists for a specific domain
func hasAccessRuleForDomain(domain string) bool {
	domains := getConfiguration()

	for i := 0; i < len(domains.AccessControl.Rules); i++ {
		accessRule := domains.AccessControl.Rules[i]

		if basics.StringSliceContains(accessRule.Domain, domain) {
			return true
		}
	}

	return false
}

func addAccessRuleForDomain(domain []string, policy string, subjects []string, resources []string) AutheliaConfigStruct {
	newRule := AutheliaConfigStruct{
		AccessControl: AutheliaAccessControlStruct{
			Rules: []AutheliaAccessControlRulesStruct{
				{
					Domain:    domain,
					Policy:    policy,
					Subject:   subjects,
					Resources: resources,
				},
			},
		},
	}

	return newRule
}
