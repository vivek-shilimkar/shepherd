package cloudcredentials

// The json/yaml config key for the alibaba cloud credential config
const AlibabaCredentialConfigurationFileKey = "alibabaCredentials"

// AlibabaCredentialConfig is configuration need to create an alibaba cloud credential
type AlibabaCredentialConfig struct {
	AccessKey string `json:"accessKeyId" yaml:"accessKeyId"`
	SecretKey string `json:"accessKeySecret" yaml:"accessKeySecret"`
}
