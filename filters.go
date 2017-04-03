package stackexchange

type BaseFilterType string

const (
	DefaultFilter  BaseFilterType = "default"
	WithBodyFilter BaseFilterType = "withbody"
	NoneFilter     BaseFilterType = "none"
	TotalFilter    BaseFilterType = "total"
)

type FilterType int

const (
	Safe   FilterType = iota
	Unsafe
)

type Filter struct {
	Base     BaseFilterType
	Includes []string
	Excludes []string
	Type     FilterType
}

func NewFilter() (*Filter) {
	return &Filter{
		Base: DefaultFilter,
		Type: Safe,
	}
}

func FilterBaseWithMdBody() *Filter {
	return &Filter{
		Base:     NoneFilter,
		Includes: []string{"question.markdown_body"},
		Type:     Safe,
	}
}
