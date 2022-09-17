package tables

import "transporte/src/models"

func ClientesCars_GetSchema() ([]models.Base, string) {
	var ClientesCars []models.Base
	tableName := "ClientesCars"
	ClientesCars = append(ClientesCars, models.Base{
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
	ClientesCars = append(ClientesCars, models.Base{
		Name:        "n_docu",
		Description: "n_docu",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       11,
			Max:       11,
			UpperCase: true,
		},
	})
	ClientesCars = append(ClientesCars, models.Base{
		Name:        "l_marc",
		Description: "l_marc",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       5,
			Max:       50,
			UpperCase: true,
		},
	})
	ClientesCars = append(ClientesCars, models.Base{
		Name:        "l_mode",
		Description: "l_mode",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       5,
			Max:       50,
			UpperCase: true,
		},
	})
	ClientesCars = append(ClientesCars, models.Base{
		Name:        "l_color",
		Description: "l_color",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       7,
			Max:       70,
			UpperCase: true,
		},
	})
	ClientesCars = append(ClientesCars, models.Base{
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
	ClientesCars = append(ClientesCars, models.Base{
		Name:        "c_mode",
		Description: "c_mode",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
			Min:  4,
			Max:  4,
		},
	})
	ClientesCars = append(ClientesCars, models.Base{
		Name:        "n_seri",
		Description: "n_seri",
		Required:    true,
		Type:        "uint64",
		Uint: models.Uints{
			Max: 17,
		},
	})
	ClientesCars = append(ClientesCars, models.Base{
		Name:        "n_pasa",
		Description: "n_pasa",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       2,
			Max:       20,
			UpperCase: true,
		},
	})
	ClientesCars = append(ClientesCars, models.Base{
		Name:        "l_obse",
		Description: "l_obse",
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       5,
			Max:       100,
			UpperCase: true,
		},
	})
	ClientesCars = append(ClientesCars, models.Base{
		Name:        "k_stad",
		Description: "k_stad",
		Required:    true,
		Update:      true,
		Type:        "uint64",
		Uint: models.Uints{
			Max: 10,
		},
	})
	return ClientesCars, tableName
}
