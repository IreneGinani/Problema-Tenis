package main
import(
	"fmt"
	"math/rand"
	"time"
	"sync"
)

//Funcao que da um numero aleatorio entre um intervalo esperado

func random(min, max int) int {
	return rand.Intn(max - min) + min
	
}

//funcao que determina se o jogador 1 perdeu ou nao a bola

func jogador1(c chan bool){

	perdeu := false
	
	rand.Seed(time.Now().UnixNano())
	bola := random(1, 3)
	//fmt.Println(bola)
	if bola == 1 {
	     perdeu := true
	     c <- perdeu
	}
	
	c <- perdeu
}

//funcao que determina se o jogador 2 perdeu ou nao a bola

func jogador2(c chan bool){

	perdeu := false
	
	rand.Seed(time.Now().UnixNano())
	bola := random(1, 3)
	//fmt.Println(bola)
	if bola == 1 {
	     perdeu := true
	     c <- perdeu
	}
	
	c <- perdeu
}

//funcao principal que faz a dinâmica principal do jogo

func main() {

	
	var mutex = &sync.Mutex{}
	var pontos, game, set int
	
	//variaveis que setam a quantidade de pontos, games e set
	
	pontos = 5
	game = 3
	set = 3

	// canal de comunicacao para saber o resultado das jogadas
	c := make(chan bool)
	
	
	//variaveis que setam a quantidade de pontos, games e set de cada jogador
	pontos_jogador1 := 0
	pontos_jogador2 := 0
	
	games_jogador1 := 0
	games_jogador2 := 0
	
	set_jogador1 := 0
	set_jogador2 := 0

	//laço que executa enquanto os sets não forem completados

	 for set_jogador1 < set && set_jogador2 < set {
	
		//laço que executa enquanto os games não forem completados

		for games_jogador1 < pontos && games_jogador2 < game {
			
			//laço que executa enquanto os pontos não forem completados

			for pontos_jogador1 < pontos && pontos_jogador2 < pontos {
				
				//jogador 1 passando a bola
				
				mutex.Lock()
				go jogador1(c)
				mutex.Unlock()
				
				//teste para saber se o jogador 1 perdeu a bola
			
				if <-c {
				
				   pontos_jogador2 = pontos_jogador2 + 1
				
				   fmt.Println("Jogador 2 fez ponto, total de pontos:")
				   fmt.Println(pontos_jogador2)
				   
				   
				}
				
				//jogador 2 passando a bola
				
				mutex.Lock()
				go jogador2(c)
				mutex.Unlock()
				
				//teste para saber se o jogador 2 pedeu a bola
				
				if <-c {
				   
				   pontos_jogador1 = pontos_jogador1 + 1
				
				   fmt.Println("Jogador 1 fez ponto, total de pontos:")
				   fmt.Println(pontos_jogador1)
				   
				}
				
				//verificacao se a diferença entre pontos é maior que 2 para que o game seja fechado

				if (pontos_jogador1 - pontos_jogador2 == 1 || pontos_jogador2 - pontos_jogador1 == 1) && (pontos-pontos_jogador1 == 1 || pontos - pontos_jogador2  == 1){
					pontos = pontos + 1
				}
			
			}

			//incremento dos games

	 		if pontos_jogador1 == pontos {
	 			games_jogador1 = games_jogador1 + 1
				fmt.Println("Jogador 1 fez game, total de games:")
				fmt.Println(games_jogador1)
	 		} else if pontos_jogador2 == pontos {
	 			games_jogador2 = games_jogador2 + 1
				fmt.Println("Jogador 2 fez game, total de games:")
				fmt.Println(games_jogador2)
	 		}

	 		//zerando pontos para novo game

	 		pontos_jogador2 = 0
	 		pontos_jogador1 = 0

	 	}

	 	//incremento dos sets

	 	if games_jogador1 == game {
	 			set_jogador1 = set_jogador1 + 1
				fmt.Println("Jogador 1 fez set, total de sets:")
				fmt.Println(set_jogador1)
	 		} else if games_jogador2 == game {
				set_jogador2 = set_jogador2 + 1
				fmt.Println("Jogador 2 fez set, total de sets:")
				fmt.Println(set_jogador2)
	 		}
		

		//verificacao se a diferença entre games é maior que 2 para que o set seja fechado
		
		if (games_jogador1 - games_jogador2 == 1 || games_jogador2 - games_jogador1 == 1) && (game-games_jogador1 == 1 || game - games_jogador2  == 1){
					game = game + 1
				}


		//verificacao se a diferença entre sets é maior que 2 para que o jogo seja fechado

		if (set_jogador1 - set_jogador2 == 1 || set_jogador2 - set_jogador1 == 1) && (set-set_jogador1 == 1 || set - set_jogador2  == 1){
					set = set + 1
				}
	 	
			//zerando games para novo set

	 		games_jogador2 = 0
	 		games_jogador1 = 0
	 }

	 //anuncio do vencedor
	 if set_jogador1 == set {
	 			fmt.Println("jogador 1 venceu")
	 		} else if set_jogador2 == set {
	 			fmt.Println("jogador 2 venceu")
	 		}
}