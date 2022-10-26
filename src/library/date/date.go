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

/**
 * Retorna fecha actual en zona horaria  (America/Bogota)
 * Return [time.time] : fecha  de ahora
 */
func GetDateLocation() time.Time {
	loc, _ := time.LoadLocation("America/Bogota")
	t := time.Now().In(loc)

	return t
}

/**
 * Retorna la fecha en tipo time
 * @param {string} date: fecha en foprmato DD/MM/YYYY
 * @return {time.Time} fecha en tipo time
 */
func GetDate(date string) time.Time {
	t, _ := time.Parse("02/01/2006", date)
	return t
}

// Obtenga el primer día del mes donde está la hora entrante, es decir, las 0 en punto el primer día de un mes. Si se pasa time.Now (), devuelve la hora a las 0 en punto el primer día del mes actual.
func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return GetZeroTime(d)
}

// Obtenga el último día del mes donde está la hora entrante, es decir, las 0 en punto el último día de un mes. Si se pasa time.Now (), devuelve la hora de las 0 en punto el último día del mes actual.
func GetLastDateOfMonth(d time.Time) time.Time {
	return GetFirstDateOfMonth(d).AddDate(0, 1, -1)
}

// Obtenga la hora a las 0 en punto de un día determinado
func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}
