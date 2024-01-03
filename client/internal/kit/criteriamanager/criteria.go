package criteriamanager

type Criteria struct {
	filtros [][]Filter
	limit   *int
	offset  *int
}

func NewCriteria(filtros [][]Filter, limit *int, offset *int) Criteria {
	return Criteria{filtros: filtros, limit: limit, offset: offset}
}

func (c Criteria) GETFILTROS() [][]Filter {
	return c.filtros
}

func (c Criteria) GETLIMIT() *int {
	return c.limit
}

func (c Criteria) GETOFFSET() *int {
	return c.offset
}
