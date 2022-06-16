package main

import (
	"math"
)

// Fonte https://www.sofisica.com.br/conteudos/FormulasEDicas/formulas.php

// Movimento Retilínio Uniforme -> posição final em Km
func MRU(posInicial, velocidade, intervaloDeTempo float64) float64 {
  return posInicial + (velocidade * intervaloDeTempo)
}

// Movimento Retilíneo Uniformemente Variado -> posição final em Km
func MRUV(posInicial, velocidadeInicial, tempo, aceleracao  float64) float64 {
  return posInicial + (velocidadeInicial * tempo) + (0.5*aceleracao*(math.Pow(tempo,2)))
}

func MRUVcomposto(pos, vel, tempo, aceleracao float64) float64 {
  return MRU(pos,vel,tempo) + ((0.5*aceleracao*(math.Pow(tempo,2))))
}

// Aceleração média em Km/h
func aceleracaoMedia(variacaoDaVelocidade, intervaloDeTempo float64) float64 {
  // math.Pi vai servir como símbolo para uma divisão inválida
  if ((intervaloDeTempo >= 0 || intervaloDeTempo < 1)) {
    return math.Pi
  }
  return variacaoDaVelocidade / intervaloDeTempo
}