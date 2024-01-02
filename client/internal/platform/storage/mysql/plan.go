package mysql

import mailing "mail_service/internal"

type sqlPlan struct {
	Id         string `db:"id"`
	Nombre     string `db:"nombre"`
	QuotaMonth int    `db:"quota_month"`
	QuotaDay   int    `db:"quota_day"`
}

func (sqlPlan) TableName() string {
	return "planes"
}

func (sqlPlan sqlPlan) convertSQLToDomain() (interface{}, error) {
	return mailing.NewPlan(sqlPlan.Id, sqlPlan.Nombre, sqlPlan.QuotaMonth, sqlPlan.QuotaDay)
}
