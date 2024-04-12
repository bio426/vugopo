package task

import (
	"context"

	"github.com/bio426/vugopo/datasource"
	"github.com/bio426/vugopo/internal/core"
)

type AuthSvc core.Service

func (svc *AuthSvc) List(c context.Context) (CtlListResponse, error) {
	rows, err := datasource.Postgres.QueryContext(
		c,
		`
        select 
          p.id, 
          p.name, 
          p.price, 
          pc.name 
        from 
          products p 
          left join product_categories pc on pc.id = p.category 
        where 
          p.store = $1;
        `,
	)
	if err != nil {
		return CtlListResponse{}, err
	}
	defer rows.Close()

	res := CtlListResponse{Rows: []CtlListRow{}}
	for rows.Next() {
		var row = CtlListRow{}
		if err = rows.Scan(
			&row.Id,
			&row.Name,
			&row.Price,
			&row.Category,
		); err != nil {
			return CtlListResponse{}, err
		}
		res.Rows = append(res.Rows, row)
	}
	return res, nil
}

var Service = &AuthSvc{}
