package src

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//Roteiro
// 1 - Escreve o primeiro teste, ele quebra porque não tem nada implementado, depois corrige do jeito mais simples possível deixando retornando zero
// 2 - Escreve o segundo teste, ele vai quebrar porque ta retornando zero, faz passar só iterando sobre o array e somando, na refatoração já transforma para table test
// 3 - Segue a analogia do teste 2, na refatoração isola a verificação do spare para um método
// 4 - Mesma coisa do teste 3, na refatoração além de isolar a verificação, pode isolar também o bonus
// 5 - Esse caso sai praticamente sozinho

func TestDeveCalcularPontuacaoAoErrarTodosPinos(t *testing.T) {
	//Arrange
	pontuacaoEsperada := 0
	jogadas := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	game := NewGame()

	//Action
	game.CalcularPontuacao(jogadas)

	//Assert
	assert.Equal(t, pontuacaoEsperada, game.Pontuacao)
}

func TestDeveCalcularPontuacaoQuandoNaoFizerStrikeNemSpare(t *testing.T) {
	//Arrange
	testCases := []struct {
		pontuacaoEsperada int
		jogadas []int
	}{
		{20, []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}},
		{19, []int{5, 2, 0, 0, 5, 0, 3, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{33, []int{5, 2, 0, 0, 5, 0, 3, 4, 1, 0, 5, 0, 5, 3, 0, 0, 0, 0, 0, 0}},
	}
	for _, tc := range testCases{
		game := NewGame()
	
		//Action
		game.CalcularPontuacao(tc.jogadas)
	
		//Assert
		assert.Equal(t, tc.pontuacaoEsperada, game.Pontuacao)
	}
}

func TestDeveCalcularPontuacaoQuandoOcorrerUmSpare(t *testing.T) {
	pontuacaoEsperada := 17
	jogadas := []int{0, 0, 4, 6, 3, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	game := NewGame()

	//Action
	game.CalcularPontuacao(jogadas)

	//Assert
	assert.Equal(t, pontuacaoEsperada, game.Pontuacao)
}

func TestDeveCalcularPontuacaoQuandoOcorrerUmStrike(t *testing.T) {
	pontuacaoEsperada := 24
	jogadas := []int{10, 3, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	game := NewGame()

	//Action
	game.CalcularPontuacao(jogadas)

	//Assert
	assert.Equal(t, pontuacaoEsperada, game.Pontuacao)
}

func TestDeveCalcularPontuacaoQuandoSoOcorrerStrikes(t *testing.T) {
	pontuacaoEsperada := 300
	jogadas := []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	game := NewGame()

	//Action
	game.CalcularPontuacao(jogadas)

	//Assert
	assert.Equal(t, pontuacaoEsperada, game.Pontuacao)
}