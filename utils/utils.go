package utils

import "fmt"

const (
	BookAuthorIds = "book_authorIDs"
	AuthorBookIds = "author_bookIDs"
)

type ResolverKey struct {
	Key interface{}
}

func NewResolverKey(key interface{}) *ResolverKey {
	return &ResolverKey{
		Key: key,
	}
}

func (rk *ResolverKey) String() string {
	return fmt.Sprintf("%v", rk.Key)
}

func (rk *ResolverKey) Raw() interface{} {
	return rk.Key
}

func BoolP(boolValue bool) *bool {
	return &boolValue
}