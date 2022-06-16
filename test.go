package main

import (
	"testing"
  "math"
)

func testMRU(t *testing.T) {
  posInicial := 100.0
  velocidade := 15.5
  tempo := 3.0

  esperado := 146.5
  obtido := MRU(posInicial, velocidade, tempo)
  if (obtido != esperado) {
    t.Errorf("Contrato modificado. MRU() -> obtido %f, esperado %f", obtido, esperado)
  }
}

func testMRUV(t *testing.T) {
  posInicial := 18.8
  velocidadeInicial := 30.7
  tempo := 10.2
  aceleracao := 20.0

  esperado := 1372.34
  obtido := MRUV(posInicial, velocidadeInicial, tempo, aceleracao)
  if (obtido != esperado) {
    t.Errorf("Contrato modificado. MRUV() -> obtido %f, esperado %f", obtido, esperado)
  }
}

func testAceleracaoMedia(t *testing.T) {
  variacaoDaVelocidade := 150.0
  intervaloDeTempo := 3.5

  esperado := 42.857142857142854
  obtido := aceleracaoMedia(variacaoDaVelocidade, intervaloDeTempo)
  if (obtido != esperado) {
    t.Errorf("Contrato modificado. aceleracaoMedia() -> obtido %f, esperado %f", obtido, esperado)
  }
}

func testAceleracaoMedia_divisaoPorZero(t *testing.T) {
  variacaoDaVelocidade := 10.0
  intervaloDeTempo := 0.0

  obtido := aceleracaoMedia(variacaoDaVelocidade, intervaloDeTempo)
  if (obtido == math.Pi) { // divisão inválida
    t.Errorf("intervaloDeTempo == 0 ; aceleracaoMedia() não aceita divisões por zero")
  }
}

func testAceleracaoMedia_divisaoEntreZeroeUm(t *testing.T) {
  variacaoDaVelocidade := 10.0
  intervaloDeTempo := 0.5

  obtido := aceleracaoMedia(variacaoDaVelocidade, intervaloDeTempo)
  if (obtido == math.Pi) { // divisão inválida
    t.Errorf("intervaloDeTempo entre 0 e 1; uso incorreto da função aceleracaoMedia()")
  }
}