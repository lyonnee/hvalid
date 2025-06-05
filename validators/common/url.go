package common

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/lyonnee/hvalid/validators/primitive"
)

// URLValidator URL验证器
type URLValidator struct {
	*primitive.StringValidator
}

// NewURLValidator 创建一个新的URL验证器
func NewURLValidator(fieldName string) *URLValidator {
	return &URLValidator{
		StringValidator: primitive.NewStringValidator(fieldName),
	}
}

// Validate 验证URL格式
func (v *URLValidator) Validate() func(string) error {
	return func(s string) error {
		_, err := url.ParseRequestURI(s)
		if err != nil {
			return fmt.Errorf("无效的URL格式")
		}
		return nil
	}
}

// ValidateProtocol 验证URL协议
func (v *URLValidator) ValidateProtocol(protocols []string) func(string) error {
	return func(s string) error {
		parsedURL, err := url.Parse(s)
		if err != nil {
			return fmt.Errorf("无效的URL格式")
		}

		protocol := strings.ToLower(parsedURL.Scheme)
		for _, p := range protocols {
			if protocol == strings.ToLower(p) {
				return nil
			}
		}

		return fmt.Errorf("不支持的URL协议，支持的协议: %v", protocols)
	}
}

// ValidateDomain 验证URL域名
func (v *URLValidator) ValidateDomain(domains []string) func(string) error {
	return func(s string) error {
		parsedURL, err := url.Parse(s)
		if err != nil {
			return fmt.Errorf("无效的URL格式")
		}

		host := strings.ToLower(parsedURL.Host)
		for _, d := range domains {
			if strings.HasSuffix(host, strings.ToLower(d)) {
				return nil
			}
		}

		return fmt.Errorf("不支持的域名，支持的域名: %v", domains)
	}
}

// ValidatePath 验证URL路径
func (v *URLValidator) ValidatePath(prefix string) func(string) error {
	return func(s string) error {
		parsedURL, err := url.Parse(s)
		if err != nil {
			return fmt.Errorf("无效的URL格式")
		}

		if !strings.HasPrefix(parsedURL.Path, prefix) {
			return fmt.Errorf("URL路径必须以 %s 开头", prefix)
		}

		return nil
	}
}

// ValidateQuery 验证URL查询参数
func (v *URLValidator) ValidateQuery(requiredParams []string) func(string) error {
	return func(s string) error {
		parsedURL, err := url.Parse(s)
		if err != nil {
			return fmt.Errorf("无效的URL格式")
		}

		query := parsedURL.Query()
		for _, param := range requiredParams {
			if !query.Has(param) {
				return fmt.Errorf("缺少必需的查询参数: %s", param)
			}
		}

		return nil
	}
}

// ValidateFragment 验证URL片段
func (v *URLValidator) ValidateFragment() func(string) error {
	return func(s string) error {
		parsedURL, err := url.Parse(s)
		if err != nil {
			return fmt.Errorf("无效的URL格式")
		}

		if parsedURL.Fragment == "" {
			return fmt.Errorf("URL必须包含片段")
		}

		return nil
	}
}

// ValidatePort 验证URL端口
func (v *URLValidator) ValidatePort(ports []string) func(string) error {
	return func(s string) error {
		parsedURL, err := url.Parse(s)
		if err != nil {
			return fmt.Errorf("无效的URL格式")
		}

		host := parsedURL.Host
		if !strings.Contains(host, ":") {
			return fmt.Errorf("URL必须指定端口")
		}

		port := strings.Split(host, ":")[1]
		for _, p := range ports {
			if port == p {
				return nil
			}
		}

		return fmt.Errorf("不支持的端口，支持的端口: %v", ports)
	}
}

// ValidateIP 验证URL是否为IP地址
func (v *URLValidator) ValidateIP() func(string) error {
	return func(s string) error {
		parsedURL, err := url.Parse(s)
		if err != nil {
			return fmt.Errorf("无效的URL格式")
		}

		host := parsedURL.Host
		if strings.Contains(host, ":") {
			host = strings.Split(host, ":")[0]
		}

		// 简单的IP地址格式验证
		parts := strings.Split(host, ".")
		if len(parts) != 4 {
			return fmt.Errorf("无效的IP地址格式")
		}

		for _, part := range parts {
			if len(part) == 0 || len(part) > 3 {
				return fmt.Errorf("无效的IP地址格式")
			}
			for _, c := range part {
				if c < '0' || c > '9' {
					return fmt.Errorf("无效的IP地址格式")
				}
			}
			num := 0
			for _, c := range part {
				num = num*10 + int(c-'0')
			}
			if num > 255 {
				return fmt.Errorf("无效的IP地址格式")
			}
		}

		return nil
	}
}
