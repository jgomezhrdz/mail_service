package mailing

// Course is the data structure that represents a course.
type Plan struct {
	id         *int
	nombre     string
	quotaMonth int
	quotaDay   int
}

// NewCourse creates a new course.
func NewPlan(id *int, nombre string, quotaMonth int, quotaDay int) Plan {
	return Plan{
		id:         id,
		nombre:     nombre,
		quotaMonth: quotaMonth,
		quotaDay:   quotaDay,
	}
}

// ID returns the course unique identifier.
func (p Plan) ID() *int {
	return p.id
}

func (p Plan) NOMBRE() string {
	return p.nombre
}

func (p Plan) QUOTAMONTH() int {
	return p.quotaMonth
}

func (p Plan) QUOTADAY() int {
	return p.quotaDay
}
