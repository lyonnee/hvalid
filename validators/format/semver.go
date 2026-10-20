package format

import (
	"errors"
	"regexp"

	"github.com/lyonnee/hvalid"
)

var semverPattern = regexp.MustCompile(`^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)

// Semver 要求语义化版本 2.0.0，不接受 v 前缀，数字部分不能有前导零。
func Semver[T ~string]() hvalid.ValidatorFunc[T] {
	return func(value T) error {
		if !semverPattern.MatchString(string(value)) {
			return errors.New("must be a semantic version")
		}
		return nil
	}
}
