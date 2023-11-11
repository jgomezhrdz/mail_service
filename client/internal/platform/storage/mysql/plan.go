package mysql

const (
	sqlPlanTable = "planes"
)

type sqlPlan struct {
	Id         string `db:"id"`
	Name       string `db:"nombre"`
	QuotaMonth string `db:"quota_month"`
	QuotaDay   string `db:"quota_day"`
}
