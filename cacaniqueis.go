package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func converte(numero int) string {
	if numero == 0 {
		return "#"
	}
	if numero == 1 {
		return "@"
	}
	if numero == 2 {
		return "*"
	}
	return strconv.Itoa(numero)
}
func main() {
	montante := 100
	semente := rand.NewSource(time.Now().UnixNano())
	randomizador := rand.New(semente)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Bem vindos ao Caça níqueis 2019!")
	fmt.Println("Deseja iniciar na roleta?")
	entrada, _ := reader.ReadString('\n')
	entrada = strings.Replace(entrada, "\r\n", "", -1)
	if strings.Compare("n", entrada) == 0 {
		fmt.Println("Obrigado! Até a próxima!")

		time.Sleep(2 * time.Second)
		return
	}

	for {

		a := randomizador.Intn(3)
		b := randomizador.Intn(3)
		c := randomizador.Intn(3)
		if a == b && b == c {
			fmt.Println("Parabéns você ganhou!")
			montante = montante + 30

		} else {
			fmt.Println("Que pena, você perdeu!")
		}
		montante = montante - 10
		fmt.Printf("Seu montante agora é : %d \n", montante)
		fmt.Printf("%s %s %s\n", converte(a), converte(b), converte(c))
		if montante == 0 {
			fmt.Println("Você não tem mais nada! Que pena!! Volte de novo jogador! ")

			break

		}
		fmt.Println("Deseja continuar?")
		entrada, _ := reader.ReadString('\n')
		entrada = strings.Replace(entrada, "\r\n", "", -1)
		if strings.Compare("n", entrada) == 0 {
			fmt.Println("Obrigado, foi um prazer jogar com você! ")
			break
		}
	}
	entrada, _ = reader.ReadString('\n')
	time.Sleep(2 * time.Second)

}
