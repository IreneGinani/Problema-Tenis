package main
import(
	"fmt"
	"math/rand"
	//"time"
	"sync"
)

func random(min, max int) int {
	return rand.Intn(max - min) + min
	
}

func jogador1(c chan bool){

	perdeu := false
	
	rand.Seed(0)
	bola := random(1, 3)
	//fmt.Println(bola)
	if bola == 1 {
	     perdeu := true
	     c <- perdeu
	}
	
	c <- perdeu
}

func jogador2(c chan bool){

	perdeu := false
	
	rand.Seed(1)
	bola := random(1, 3)
	//fmt.Println(bola)
	if bola == 1 {
	     perdeu := true
	     c <- perdeu
	}
	
	c <- perdeu
}

func main() {

	
	var mutex = &sync.Mutex{}
	var pontos, game, set int
	//fmt.Println("Digite de quantos pontos é feito o game")
	
	//fmt.Scan(&pontos)
	
	pontos = 5
	game = 3
	set = 3
	c := make(chan bool)
	
	
	pontos_jogador1 := 0
	pontos_jogador2 := 0
	
	 games_jogador1 := 0
	 games_jogador2 := 0
	
	 set_jogador1 := 0
	 set_jogador2 := 0

	 for set_jogador1 < set && set_jogador2 < set {
		
		for games_jogador1 < pontos && games_jogador2 < game {
	
			for pontos_jogador1 < pontos && pontos_jogador2 < pontos {
				
				//jogador 1 passando a bola
				
				mutex.Lock()
				go jogador1(c)
				mutex.Unlock()
				
				//teste para saber se o jogador 1 perdeu a bola
			
				if <-c {
				
				   pontos_jogador2 = pontos_jogador2 + 1
				   
				   
				}
				
				//jogador 2 passando a bola
				
				mutex.Lock()
				go jogador2(c)
				mutex.Unlock()
				
				//teste para saber se o jogador 2 pedeu a bola
				
				if <-c {
				   
				   pontos_jogador1 = pontos_jogador1 + 1
				   
				}
			
				fmt.Println("pontos jogador 1:")
				fmt.Println(pontos_jogador1)
				fmt.Println("pontos jogador 2:")
				fmt.Println(pontos_jogador2)
				//pontos_jogador1 = 5
			
			}

	 		if pontos_jogador1 == pontos {
	 			games_jogador1 = games_jogador1 + 1
				fmt.Println("games jogador 1:")
				fmt.Println(games_jogador1)
				fmt.Println("games jogador 2:")
				fmt.Println(games_jogador2)
	 		} else if pontos_jogador2 == pontos {
	 			games_jogador2 = games_jogador2 + 1
				fmt.Println("games jogador 1:")
				fmt.Println(games_jogador1)
				fmt.Println("games jogador 2:")
				fmt.Println(games_jogador2)
	 		}

	 		pontos_jogador2 = 0
	 		pontos_jogador1 = 0

	 	}

	 	if games_jogador1 == game {
	 			set_jogador1 = set_jogador1 + 1
				fmt.Println("sets jogador 1:")
				fmt.Println(set_jogador1)
				fmt.Println("sets jogador 2:")
				fmt.Println(set_jogador2)
	 		} else if games_jogador2 == game {
	 			set_jogador2 = set_jogador2 + 1
				fmt.Println("sets jogador 1:")
				fmt.Println(set_jogador1)
				fmt.Println("sets jogador 2:")
				fmt.Println(set_jogador2)
	 		}

	 		games_jogador2 = 0
	 		games_jogador1 = 0
	 }

	 if set_jogador1 == set {
	 			fmt.Println("jogador 1 venceu")
	 		} else if set_jogador2 == set {
	 			fmt.Println("jogador 2 venceu")
	 		}
}