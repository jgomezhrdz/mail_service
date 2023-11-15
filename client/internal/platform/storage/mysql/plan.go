package mysql

import mailing "mail_service/internal"

func (sqlPlan) TableName() string {
	return "planes"
}

type sqlPlan struct {
	Id         string `db:"id"`
	Nombre     string `db:"nombre"`
	QuotaMonth int    `db:"quota_month"`
	QuotaDay   int    `db:"quota_day"`
}

func convertSQLPlanToMailingPlan(sqlPlan sqlPlan) (mailing.Plan, error) {
	cliente, err := mailing.NewPlan(sqlPlan.Id, sqlPlan.Nombre, sqlPlan.QuotaMonth, sqlPlan.QuotaDay)
	if err != nil {
		return cliente, err
	}
	return cliente, nil
}
