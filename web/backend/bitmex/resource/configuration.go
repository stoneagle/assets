package resource

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Configuration 配置对象
type Configuration struct {
	UserName      string            `json:"userName,omitempty"`
	Password      string            `json:"password,omitempty"`
	APIKeyPrefix  map[string]string `json:"APIKeyPrefix,omitempty"`
	APIKey        map[string]string `json:"APIKey,omitempty"`
	debug         bool              `json:"debug,omitempty"`
	DebugFile     string            `json:"debugFile,omitempty"`
	OAuthToken    string            `json:"oAuthToken,omitempty"`
	Timeout       int               `json:"timeout,omitempty"`
	BasePath      string            `json:"basePath,omitempty"`
	AuthPath      string            `json:"authPath,omitempty"`
	Host          string            `json:"host,omitempty"`
	Scheme        string            `json:"scheme,omitempty"`
	AccessToken   string            `json:"accessToken,omitempty"`
	DefaultHeader map[string]string `json:"defaultHeader,omitempty"`
	UserAgent     string            `json:"userAgent,omitempty"`
	APIClient     APIClient         `json:"APIClient,omitempty"`
}

// NewConfiguration 创建新配置
func NewConfiguration() *Configuration {
	return &Configuration{
		BasePath:      "https://www.bitmex.com/api/v1",
		AuthPath:      "/api/v1",
		UserName:      "",
		debug:         false,
		DefaultHeader: make(map[string]string),
		APIKey:        make(map[string]string),
		APIKeyPrefix:  make(map[string]string),
		UserAgent:     "Swagger-Codegen/1.0.0/go",
	}
}

// GetBasicAuthEncodedString 获取BA加密字符串
func (c *Configuration) GetBasicAuthEncodedString() string {
	return base64.StdEncoding.EncodeToString([]byte(c.UserName + ":" + c.Password))
}

// AddDefaultHeader 添加默认header
func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}

// AddAPIKey 添加默认header
func (c *Configuration) AddAPIKey(uri string, method string, queryParams url.Values, formParams map[string]string) {
	apiKey := "ayQJFnX-NTLd_Z2xdvnz80gE"
	apiSecret := "mOFPecMfUKrSJejZcf30BLtUGlztvI1QJmdNFHvgE5LghwHm"

	nonce := strconv.FormatInt(time.Now().Unix(), 10)
	mac := hmac.New(sha256.New, []byte(apiSecret))
	path := c.AuthPath + uri
	data := ""

	c.DefaultHeader["api-key"] = apiKey
	c.DefaultHeader["api-nonce"] = nonce

	switch strings.ToUpper(method) {
	case "GET":
		if len(queryParams) > 0 {
			path = path + "?" + queryParams.Encode()
		}
	case "POST", "PUT", "PATCH", "DELETE":
		if len(formParams) > 0 {
		}
	}
	mac.Write([]byte(strings.ToUpper(method) + path + nonce + data))
	signature := mac.Sum(nil)
	c.DefaultHeader["api-signature"] = hex.EncodeToString([]byte(signature))
}

// GetAPIKeyWithPrefix 通过prefix获取APIKey
func (c *Configuration) GetAPIKeyWithPrefix(APIKeyIdentifier string) string {
	if c.APIKeyPrefix[APIKeyIdentifier] != "" {
		return c.APIKeyPrefix[APIKeyIdentifier] + " " + c.APIKey[APIKeyIdentifier]
	}

	return c.APIKey[APIKeyIdentifier]
}

// SetDebug 开启debug
func (c *Configuration) SetDebug(enable bool) {
	c.debug = enable
}

// GetDebug 获取debug状态
func (c *Configuration) GetDebug() bool {
	return c.debug
}
