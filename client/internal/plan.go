package mailing

// Course is the data structure that represents a course.
type Plan struct {
	id         string
	nombre     string
	quotaMonth int
	quotaDay   int
}

// NewCourse creates a new course.
func NewPlan(id string, nombre string, quotaMonth int, quotaDay int) Plan {
	return Plan{
		id:         id,
		nombre:     nombre,
		quotaMonth: quotaMonth,
		quotaDay:   quotaDay,
	}
}

// ID returns the course unique identifier.
func (p Plan) ID() string {
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

func (p Plan) TOJSON() map[string]interface{} {
	return map[string]interface{}{
		"id":         p.id,
		"nombre":     p.nombre,
		"quotaMonth": p.quotaMonth,
		"quotaDay":   p.quotaDay,
	}
}
