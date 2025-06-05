package common

import (
	"fmt"
	"net"
	"strings"

	"github.com/lyonnee/hvalid/validators/primitive"
)

// IPValidator IP地址验证器
type IPValidator struct {
	*primitive.StringValidator
}

// NewIPValidator 创建一个新的IP地址验证器
func NewIPValidator(fieldName string) *IPValidator {
	return &IPValidator{
		StringValidator: primitive.NewStringValidator(fieldName),
	}
}

// Validate 验证IP地址格式
func (v *IPValidator) Validate() func(string) error {
	return func(s string) error {
		ip := net.ParseIP(s)
		if ip == nil {
			return fmt.Errorf("无效的IP地址格式")
		}
		return nil
	}
}

// ValidateIPv4 验证IPv4地址
func (v *IPValidator) ValidateIPv4() func(string) error {
	return func(s string) error {
		ip := net.ParseIP(s)
		if ip == nil {
			return fmt.Errorf("无效的IP地址格式")
		}

		if ip.To4() == nil {
			return fmt.Errorf("必须是IPv4地址")
		}

		return nil
	}
}

// ValidateIPv6 验证IPv6地址
func (v *IPValidator) ValidateIPv6() func(string) error {
	return func(s string) error {
		ip := net.ParseIP(s)
		if ip == nil {
			return fmt.Errorf("无效的IP地址格式")
		}

		if ip.To4() != nil {
			return fmt.Errorf("必须是IPv6地址")
		}

		return nil
	}
}

// ValidatePrivate 验证是否为私有IP地址
func (v *IPValidator) ValidatePrivate() func(string) error {
	return func(s string) error {
		ip := net.ParseIP(s)
		if ip == nil {
			return fmt.Errorf("无效的IP地址格式")
		}

		// 检查IPv4私有地址范围
		if ip4 := ip.To4(); ip4 != nil {
			// 10.0.0.0/8
			if ip4[0] == 10 {
				return nil
			}
			// 172.16.0.0/12
			if ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31 {
				return nil
			}
			// 192.168.0.0/16
			if ip4[0] == 192 && ip4[1] == 168 {
				return nil
			}
			// 169.254.0.0/16 (链路本地地址)
			if ip4[0] == 169 && ip4[1] == 254 {
				return nil
			}
			// 127.0.0.0/8 (本地回环地址)
			if ip4[0] == 127 {
				return nil
			}
		}

		// 检查IPv6私有地址范围
		if ip.To4() == nil {
			// fc00::/7 (唯一本地地址)
			if ip[0] == 0xfc || ip[0] == 0xfd {
				return nil
			}
			// fe80::/10 (链路本地地址)
			if ip[0] == 0xfe && (ip[1]&0xc0) == 0x80 {
				return nil
			}
			// ::1/128 (本地回环地址)
			if ip.Equal(net.IPv6loopback) {
				return nil
			}
		}

		return fmt.Errorf("必须是私有IP地址")
	}
}

// ValidatePublic 验证是否为公网IP地址
func (v *IPValidator) ValidatePublic() func(string) error {
	return func(s string) error {
		ip := net.ParseIP(s)
		if ip == nil {
			return fmt.Errorf("无效的IP地址格式")
		}

		// 检查是否为私有地址
		if ip4 := ip.To4(); ip4 != nil {
			// 10.0.0.0/8
			if ip4[0] == 10 {
				return fmt.Errorf("不能是私有IP地址")
			}
			// 172.16.0.0/12
			if ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31 {
				return fmt.Errorf("不能是私有IP地址")
			}
			// 192.168.0.0/16
			if ip4[0] == 192 && ip4[1] == 168 {
				return fmt.Errorf("不能是私有IP地址")
			}
			// 169.254.0.0/16 (链路本地地址)
			if ip4[0] == 169 && ip4[1] == 254 {
				return fmt.Errorf("不能是链路本地地址")
			}
			// 127.0.0.0/8 (本地回环地址)
			if ip4[0] == 127 {
				return fmt.Errorf("不能是本地回环地址")
			}
		}

		// 检查IPv6私有地址范围
		if ip.To4() == nil {
			// fc00::/7 (唯一本地地址)
			if ip[0] == 0xfc || ip[0] == 0xfd {
				return fmt.Errorf("不能是唯一本地地址")
			}
			// fe80::/10 (链路本地地址)
			if ip[0] == 0xfe && (ip[1]&0xc0) == 0x80 {
				return fmt.Errorf("不能是链路本地地址")
			}
			// ::1/128 (本地回环地址)
			if ip.Equal(net.IPv6loopback) {
				return fmt.Errorf("不能是本地回环地址")
			}
		}

		return nil
	}
}

// ValidateCIDR 验证CIDR格式
func (v *IPValidator) ValidateCIDR() func(string) error {
	return func(s string) error {
		_, _, err := net.ParseCIDR(s)
		if err != nil {
			return fmt.Errorf("无效的CIDR格式")
		}
		return nil
	}
}

// ValidateInRange 验证IP地址是否在指定范围内
func (v *IPValidator) ValidateInRange(startIP, endIP string) func(string) error {
	return func(s string) error {
		ip := net.ParseIP(s)
		if ip == nil {
			return fmt.Errorf("无效的IP地址格式")
		}

		start := net.ParseIP(startIP)
		if start == nil {
			return fmt.Errorf("无效的起始IP地址格式")
		}

		end := net.ParseIP(endIP)
		if end == nil {
			return fmt.Errorf("无效的结束IP地址格式")
		}

		// 确保所有IP地址都是相同版本
		if (ip.To4() != nil) != (start.To4() != nil) || (ip.To4() != nil) != (end.To4() != nil) {
			return fmt.Errorf("IP地址版本不匹配")
		}

		// 比较IP地址
		if strings.Compare(ip.String(), start.String()) < 0 || strings.Compare(ip.String(), end.String()) > 0 {
			return fmt.Errorf("IP地址不在指定范围内")
		}

		return nil
	}
}
