package mailing

import "mail_service/internal/platform/shared/types"

// Course is the data structure that represents a course.
type Plan struct {
	id         types.UUID
	nombre     types.NonEmptyString
	quotaMonth types.UnsignedInt
	quotaDay   types.UnsignedInt
}

// NewCourse creates a new course.
func NewPlan(id string, nombre string, quotaMonth int, quotaDay int) (Plan, error) {

	idVO, err := types.NewUUID(id)
	if err != nil {
		return Plan{}, err
	}

	nombreVO, err := types.NewNonEmptyString(nombre)
	if err != nil {
		return Plan{}, err
	}

	quotaMonthVO, err := types.NewUnsignedInt(quotaMonth)
	if err != nil {
		return Plan{}, err
	}

	quotaDayVO, err := types.NewUnsignedInt(quotaDay)
	if err != nil {
		return Plan{}, err
	}

	return Plan{
		id:         idVO,
		nombre:     nombreVO,
		quotaMonth: quotaMonthVO,
		quotaDay:   quotaDayVO,
	}, nil
}

// ID returns the course unique identifier.
func (p Plan) ID() string {
	return p.id.Value()
}

func (p Plan) NOMBRE() string {
	return p.nombre.Value()
}

func (p Plan) QUOTAMONTH() int {
	return p.quotaMonth.Value()
}

func (p Plan) QUOTADAY() int {
	return p.quotaDay.Value()
}

func (p Plan) TOJSON() map[string]interface{} {
	return map[string]interface{}{
		"id":         p.id.Value(),
		"nombre":     p.nombre.Value(),
		"quotaMonth": p.quotaMonth.Value(),
		"quotaDay":   p.quotaDay.Value(),
	}
}
