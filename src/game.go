package src

type Game struct {
	Pontuacao int
}

func (g *Game) CalcularPontuacao(jogadas []int) {
	pontuacao := 0
	numeroDaJogada := 0
	for frame := 0; frame < 10; frame++ {
		if ehStrike(jogadas, numeroDaJogada){
			pontuacao += 10 + calcularBonusDoStrike(jogadas, numeroDaJogada)
			numeroDaJogada++
		} else if ehSpare(jogadas, numeroDaJogada) {
			pontuacao += 10 + calcularBonusDoSpare(jogadas, numeroDaJogada)
			numeroDaJogada += 2
		} else {
			pontuacao += jogadas[numeroDaJogada] + jogadas[numeroDaJogada+1]
			numeroDaJogada += 2
		}
	}
	g.Pontuacao = pontuacao
}

func calcularBonusDoStrike(jogadas []int, numeroDaJogada int) int {
	return jogadas[numeroDaJogada+1] + jogadas[numeroDaJogada+2]
}

func calcularBonusDoSpare(jogadas []int, numeroDaJogada int) int {
	return jogadas[numeroDaJogada+2]
}

func ehSpare(jogadas []int, numeroDaJogada int) bool {
	return jogadas[numeroDaJogada]+jogadas[numeroDaJogada+1] == 10
}

func ehStrike(jogadas []int, numeroDaJogada int) bool {
	return jogadas[numeroDaJogada] == 10
}

func NewGame() *Game {
	return &Game{
		Pontuacao: 0,
	}
}
