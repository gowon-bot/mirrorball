package apihelpers

import "sync"

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

func (p Paginator) GetAtPage(page int) {
	currentPage := p.convertCurrentPage(page)

	pagedParams := PagedParams{
		Page:  currentPage,
		Limit: p.PageSize,
	}

	p.Function(pagedParams)
}

// GetNext gets the current page and increments the current page
func (p Paginator) GetNext() {
	p.GetAtPage(p.CurrentPage)
	p.CurrentPage++
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

// GetAllInParallel gets all remaining pages in parallel batches
func (p Paginator) GetAllInParallel(parallelization int) {
	if p.CurrentPage >= p.TotalPages {
		return
	}

	pageChannel := make(chan int)
	var wg sync.WaitGroup
	wg.Add(parallelization)

	for ii := 0; ii < parallelization; ii++ {
		go func(c chan int) {
			for {
				page, more := <-c
				if more == false {
					wg.Done()
					return
				}

				p.GetAtPage(page)
			}
		}(pageChannel)
	}

	for page := p.convertCurrentPage(p.CurrentPage); page <= p.TotalPages; page++ {
		pageChannel <- page
	}

	close(pageChannel)
	wg.Wait()
}
