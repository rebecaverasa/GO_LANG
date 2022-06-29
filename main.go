package main

import (
	"fmt"
	"strings"
)

func main() {
	var nomedaConferencia = "Conferencia Go"
	const ticketsTotal = 50
	var remainingTickets uint = 50
	var reservas []string

	fmt.Printf("nomedaConferencia é %T, ticketsTotal é %T, remainingTickets é %T. \n \n \n", nomedaConferencia, ticketsTotal, remainingTickets)

	greetUsers(nomedaConferencia, ticketsTotal, remainingTickets)

	for remainingTickets > 0 && len(reservas) < 50 {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(remainingTickets, userTickets, reservas, firstName, lastName, email, nomedaConferencia)			
			
			firstNames := getFirstNames(reservas)
			fmt.Printf("Os primeiros nomes que fizeram reservas são: %v.\n", firstNames)
			
			if remainingTickets == 0 {
				//terminar o programa
				fmt.Println("Nossa conferencia esgotou, volte proximo ano.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("O seu nome ou sobrenome são muito pequenos.")
			}
			if !isValidEmail {
				fmt.Println("O email que você digitou é inválido, por favor, inserir @.")
			}
			if !isValidTicketNumber {
				fmt.Println("O número de tickets que você digitou é inválido.")
			}
			fmt.Println("Você inseriu dados errados, tente novamente.")
		}
	}
}

func greetUsers(nomedaConferencia string, ticketsTotal uint, remainingTickets uint) {
	fmt.Printf("Bem-vindos à %v!\n", nomedaConferencia)
	fmt.Printf("No início da conferência, foram disponibilizados %v tickets. Atualmente há %v tickets disponíveis para compra.\n", ticketsTotal, remainingTickets)
	fmt.Println("Adquira seus ingressos para participar!")
}

func getFirstNames(reservas []string) []string {
	firstNames := []string{}
	for _, reserva := range reservas {
		var names = strings.Fields(reserva)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Escreva seu primeiro nome:")
	fmt.Scan(&firstName)

	fmt.Println("Escreva seu sobrenome:")
	fmt.Scan(&lastName)

	fmt.Println("Escreva seu email:")
	fmt.Scan(&email)

	fmt.Println("Digite o numero de tickets que quer comprar: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(remainingTickets uint, userTickets uint, reservas [] string, firstName string, lastName string, email string, nomedaConferencia string) {
	remainingTickets = remainingTickets - userTickets
	reservas = append(reservas, firstName+" "+lastName)

	fmt.Printf("Obrigada %v %v por comprar %v tickets. Seu comprovante chegará no seu email: %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets sobrando para %v.\n", remainingTickets, nomedaConferencia)

}
