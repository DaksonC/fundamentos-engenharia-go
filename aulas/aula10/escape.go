package main

type Config struct {
	Porta   int
	Timeout int
}

func configPadraoValor() Config {
	return Config{Porta: 8080, Timeout: 30}
}

func configPadraoPonteiro() *Config {
	c := Config{Porta: 8080, Timeout: 30}
	return &c
}
