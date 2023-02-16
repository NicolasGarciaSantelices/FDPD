package common_sql

import (
	"database/sql"
)

func GenericGet(db *sql.DB) (err error) {

	/* rows, err := db.Query(`SELECT ` +
		`id, code, update_date_time, reddemed_by  ` +
		`FROM public.widecar_codes `)

	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var Code models.WicarCodes
		err = rows.Scan(
			&Code.Id,
			&Code.Code,
			&Code.Update,
			&Code.Redeemed,
		)
		if err == nil {
			if Code.Redeemed == "" {
				Code.Redeemed = "not redeemed"
			}
			(*Codes)[Code.Id] = Code
		}
	}

	if err = rows.Err(); err != nil {
		panic(err)
	} */

	return
}
