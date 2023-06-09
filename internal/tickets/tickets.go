package tickets

import (
	"errors"
	"strconv"
	"strings"
)

type Tickets struct {
	Tickets []Ticket
}

type Ticket struct {
	Id            string
	Nombre        string
	Email         string
	PaisDeDestino string
	HoraDelVuelo  string
	Precio        string
}

/*************************Consigna 1*************************/
func (tickets Tickets) GetTotalTickets(destination string) (int, error) {
	var totalPersonasQueViajan int
	totalPersonasQueViajan = 0

	if destination == "" {
		return totalPersonasQueViajan, errors.New("No se ingresó ningún lugar de destino")
	}

	for _, ticket := range tickets.Tickets {
		if destination == ticket.PaisDeDestino {
			totalPersonasQueViajan++
		}
	}

	return totalPersonasQueViajan, nil
}

/*************************Consigna 2*************************/
func cantidadDePersonasPorRango(tickets Tickets, horaInicio, horaFin int64) (int, error) {
	var totalPersonasPorRango int
	totalPersonasPorRango = 0
	if horaInicio < 0 || horaInicio > 23 || horaFin < 0 || horaFin > 23 || horaInicio >= horaFin {
		return totalPersonasPorRango, errors.New("El rango de horas ingresado es inválido")
	}
	for _, ticket := range tickets.Tickets {
		horaExactaSplit := strings.Split(ticket.HoraDelVuelo, ":")
		hora, _ := strconv.ParseInt(horaExactaSplit[0], 0, 64)

		if hora >= horaInicio && hora <= horaFin {
			totalPersonasPorRango++
		}
	}

	return totalPersonasPorRango, nil
}

func (tickets Tickets) GetCountByPeriod(time string) (int, error) {
	switch strings.ToLower(time) {
	case "madrugada":
		return cantidadDePersonasPorRango(tickets, 0, 6)
	case "mañana":
		return cantidadDePersonasPorRango(tickets, 7, 12)
	case "tarde":
		return cantidadDePersonasPorRango(tickets, 13, 19)
	case "noche":
		return cantidadDePersonasPorRango(tickets, 20, 23)
	default:
		return 0, errors.New("Período inexistente de viaje")
	}
}

/*************************Consigna 3*************************/
func (tickets Tickets) AverageDestination(destination string) (float64, error) {
	var totalPersonas int
	var totalPorcentaje float64
	totalPersonasPorDestino, errorTotalTicketsPorDestino := tickets.GetTotalTickets(destination)
	totalPersonas = len(tickets.Tickets)
	if errorTotalTicketsPorDestino != nil {
		return 0, errorTotalTicketsPorDestino
	}
	if totalPersonas == 0 {
		return 0, errors.New("No hay ningún ticket vendido aún")
	}

	totalPorcentaje = (float64(totalPersonasPorDestino) / float64(totalPersonas)) * 100

	return totalPorcentaje, nil
}
