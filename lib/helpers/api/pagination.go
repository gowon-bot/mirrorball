package helpers

// PagedParams is the struct type for the parameters passed into the pagination function
type PagedParams struct {
	Page, Limit int
}

// PaginatorFunction is any function that accepts PagedParams and returns a pointer
type PaginatorFunction func(PagedParams)

// Paginator iterates over a given series of pages calling a function each time
type Paginator struct {
	CurrentPage, PageSize, TotalPages int
	Function                          PaginatorFunction
	SkipFirstPage                     bool
}

func (p Paginator) convertCurrentPage(page int) int {
	if page == 0 {
		if p.SkipFirstPage == true {
			return 2
		}

		return 1
	}

	return page
}

// GetNext gets the current page and increments the current page
func (p Paginator) GetNext() {
	p.CurrentPage = p.convertCurrentPage(p.CurrentPage)

	pagedParams := PagedParams{
		Page:  p.CurrentPage,
		Limit: p.PageSize,
	}

	p.CurrentPage++

	p.Function(pagedParams)
}

// GetAll gets all remaining pages
func (p Paginator) GetAll() {
	if p.CurrentPage >= p.TotalPages {
		return
	}

	for page := p.convertCurrentPage(p.CurrentPage); page <= p.TotalPages; page++ {
		p.CurrentPage = page

		p.GetNext()
	}
}
