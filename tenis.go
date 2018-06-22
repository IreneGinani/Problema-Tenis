package main
import(
	"fmt"
	"math/rand"
	"time"
	"sync"
)

func random(min, max int) int {
	return rand.Intn(max - min) + min
	
}

func jogador1(c chan bool){

	perdeu := false
	
	rand.Seed(time.Now().Unix())
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
	
	rand.Seed(time.Now().Unix())
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
	var pontos int
	//fmt.Println("Digite de quantos pontos é feito o game")
	
	//fmt.Scan(&pontos)
	
	pontos = 5;
	c := make(chan bool)
	
	
	pontos_jogador1 := 0
	pontos_jogador2 := 0
	
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
	
		fmt.Println("pontos 1:")
		fmt.Println(pontos_jogador1)
		fmt.Println("pontos 2:")
		fmt.Println(pontos_jogador2)
		//pontos_jogador1 = 5
	
	}
	
}