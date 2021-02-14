package paging

type Paging interface {
	Next() string
	Prev() string
}
