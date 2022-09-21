package tables

import (
	"transporte/src/models"

	"github.com/google/uuid"
)

func Servicios_GetSchema() ([]models.Base, string) {
	var Servicios []models.Base
	tableName := "Servicios"
	id_serv := uuid.New().String()
	Servicios = append(Servicios, models.Base{
		Name:        "id_serv",
		Description: "id_serv",
		Required:    true,
		Important: true,
		Default: id_serv,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
		},
	})
	Servicios = append(Servicios, models.Base{
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
	Servicios = append(Servicios, models.Base{
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
	Servicios = append(Servicios, models.Base{
		Name:        "n_docu",
		Description: "n_docu",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       8,
			Max:       11,
			UpperCase: true,
		},
	})
	Servicios = append(Servicios, models.Base{
		Name:        "f_fact",
		Description: "f_fact",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
			Date: true,
		},
	})
	Servicios = append(Servicios, models.Base{
		Name:        "s_impo",
		Description: "s_impo",
		Required:    true,
		Update:      true,
		Type:        "float64",
		Float:       models.Floats{},
	})
	Servicios = append(Servicios, models.Base{
		Name:        "c_plac",
		Description: "c_plac",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       6,
			Max:       6,
			UpperCase: true,
		},
	})
	Servicios = append(Servicios, models.Base{
		Name:        "k_stad",
		Description: "k_stad",
		Required:    true,
		Update:      true,
		Type:        "uint64",
		Uint: models.Uints{
			Max: 10,
		},
	})
	Servicios = append(Servicios, models.Base{
		Name:        "f_digi",
		Description: "f_digi",
	})
	return Servicios, tableName
}
