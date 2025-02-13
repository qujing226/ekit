package option

// Option 是用于 Option 模式的泛型设计，
// 避免在代码中定义很多类似这样的结构体
// 一般情况下 T 应该是一个结构体
type Option[T any] func(t *T)

// Apply 将 opts 应用在 t 之上
func Apply[T any](t *T, opts ...Option[T]) {
	for _, opt := range opts {
		opt(t)
	}
}

// OptionErr 形如 Option, 返回一个error
// 优先使用 Option,除非在设计 option 模式的时候需要进行一些校验
type OptionErr[T any] func(t *T) error

func ApplyErr[T any](t *T, opts ...OptionErr[T]) error {
	for _, opt := range opts {
		if err := opt(t); err != nil {
			return err
		}
	}
	return nil
}
