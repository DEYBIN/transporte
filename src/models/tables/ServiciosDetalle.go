package tables

import (
	"transporte/src/models"

	"github.com/google/uuid"
)

func ServiciosDetalle_GetSchema() ([]models.Base, string) {
	var ServiciosDetalle []models.Base
	tableName := "_" + "ServiciosDetalle"
	id_serv := uuid.New().String()
	ServiciosDetalle = append(ServiciosDetalle, models.Base{
		Name:        "id_serv",
		Description: "id_serv",
		Required:    true,
		Important:   true,
		Default:     id_serv,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
		},
	})
	ServiciosDetalle = append(ServiciosDetalle, models.Base{
		Name:        "c_year",
		Description: "c_year",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
			Min:  4,
			Max:  4,
		},
	})
	ServiciosDetalle = append(ServiciosDetalle, models.Base{
		Name:        "c_mes",
		Description: "c_mes",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
			Min:  2,
			Max:  2,
		},
	})
	ServiciosDetalle = append(ServiciosDetalle, models.Base{
		Name:        "f_pago",
		Description: "f_pago",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
			Date: true,
		},
	})
	ServiciosDetalle = append(ServiciosDetalle, models.Base{
		Name:        "s_impo",
		Description: "s_impo",
		Required:    true,
		Type:        "float64",
		Float:       models.Floats{},
	})
	ServiciosDetalle = append(ServiciosDetalle, models.Base{
		Name:        "k_stad",
		Description: "k_stad",
		Required:    true,
		Update:      true,
		Type:        "uint64",
		Uint: models.Uints{
			Max: 5,
		},
	})
	return ServiciosDetalle, tableName
}
