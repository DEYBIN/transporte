package date

import "time"

/**
 * Retorna el índice de un elemento en un arreglo de strings
 * date[string] = fecha a validar formato string dd/mm/yyyy
 * return [bool],[error] = [true or false],[descripción del error or  nil]
 */
func CheckDate(date string) error {
	_, err := time.Parse("02/01/2006", date)
	if err != nil {
		return err
	}
	return nil
}

/**
 * Retorna fecha en formato dd/mm/yyyy en zona horaria  (America/Bogota)
 * Return [string] : fecha  formato string dd/mm/yyyy (America/Bogota)
 */
 func GetFechaLocationString() string {
	loc, _ := time.LoadLocation("America/Bogota")
	t := time.Now().In(loc).Format("02/01/2006")

	return t
}