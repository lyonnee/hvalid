package complex

import (
	"sync"

	"github.com/lyonnee/hvalid"
)

// CacheValidator 缓存验证器结构体
type CacheValidator[T any] struct {
	FieldName string // 字段名称
	cache     sync.Map
}

// NewCacheValidator 创建缓存验证器
func NewCacheValidator[T any](fieldName string) *CacheValidator[T] {
	return &CacheValidator[T]{
		FieldName: fieldName,
	}
}

// WithCache 使用缓存执行验证
func (v *CacheValidator[T]) WithCache(validator hvalid.ValidatorFunc[T]) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(value T) error {
		// 尝试从缓存中获取结果
		if cachedErr, ok := v.cache.Load(value); ok {
			if err, ok := cachedErr.(error); ok {
				return err
			}
		}

		// 执行验证
		err := validator(value)

		// 缓存结果
		v.cache.Store(value, err)

		return err
	})
}

// WithTTL 使用带过期时间的缓存执行验证
func (v *CacheValidator[T]) WithTTL(validator hvalid.ValidatorFunc[T], ttl int64) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(value T) error {
		// 尝试从缓存中获取结果
		if cachedResult, ok := v.cache.Load(value); ok {
			if result, ok := cachedResult.(struct {
				err error
				ttl int64
			}); ok {
				if result.ttl > 0 {
					return result.err
				}
			}
		}

		// 执行验证
		err := validator(value)

		// 缓存结果
		v.cache.Store(value, struct {
			err error
			ttl int64
		}{
			err: err,
			ttl: ttl,
		})

		return err
	})
}

// ClearCache 清除缓存
func (v *CacheValidator[T]) ClearCache() {
	v.cache = sync.Map{}
}

// RemoveFromCache 从缓存中移除特定值
func (v *CacheValidator[T]) RemoveFromCache(value T) {
	v.cache.Delete(value)
}
