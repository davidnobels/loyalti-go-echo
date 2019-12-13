package repository

import (
	"database/sql"
	"fmt"

	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
)

func GetCardMerchant(page *int, size *int, id *string) []model.Card {

	db := database.ConnectionDB()
	//db := database.ConnectPostgre()
	var kartu []model.Card
	var rows *sql.Rows
	var err error
	var total int

	if id != nil {
		if page != nil && size != nil {
			fmt.Println("2")
			rows, err = db.Where("id = ?", id).Find(&kartu).Order("title").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				panic(err)
			}
		}
	} else {
		if page != nil && size != nil {
			rows, err = db.Find(&kartu).Order("title").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				panic(err)
			}
		} else {
			rows, err = db.Find(&kartu).Rows()
			if err != nil {
				panic(err)
			}
		}
	}

	result := make([]model.Card, 0)

	for rows.Next() {
		o := new(model.Card)
		err = rows.Scan(
			&o.Id,
			&o.Created,
			&o.CreatedBy,
			&o.Modified,
			&o.ModifiedBy,
			&o.Active,
			&o.IsDeleted,
			&o.Deleted,
			&o.DeletedBy,
			&o.Title,
			&o.Description,
			&o.FontColor,
			&o.TemplateColor,
			&o.IconImage,
			&o.TermsAndCondition,
			&o.Benefit,
			&o.ValidUntil,
			&o.RewardTarget,
			&o.IsValid,
			&o.ProgramId,
			&o.CardType,
			&o.IconImageStamp,
			&o.MerchantId,
		)

		result = append(result, *o)
	}

	db.Close()
	return result
}