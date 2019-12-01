package search

import "short/app/entity"

type Search struct {
}

type Pagination struct {
	pageNum int
	numRows int
}

type Filter struct{
	public bool
	private bool
}

type Query struct {
	keyword *string
	filter *Filter
	pagination *Pagination
}

func (s Search) SearchURL(query Query) []entity.URL {
	return []entity.URL{}
}

func (s Search) fetchPrivateURLsByUser() {

}