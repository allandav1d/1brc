# 1️⃣🐝🏎️ The One Billion Row Challenge in GO

Este desafio foi feito na formação GO da Rocketseat, onde foi proposto a implementação do famoso desafio One Billion Row Challenge (1BRC), onde o objetivo é processar um arquivo de 1 bilhão de registros de forma eficiente usando a linguagem Go.


## Desafio

O arquivo de texto contém valores de temperatura para uma série de estações meteorológicas.
Cada linha é uma medição no formato `<string: nome da estação>;<double: medição>`, com o valor da medição tendo exatamente uma casa decimal. O exemplo a seguir mostra dez linhas como exemplo:

```
Hamburg;12.0
Bulawayo;8.9
Palembang;38.8
St. John's;15.2
Cracow;12.6
Bridgetown;26.9
Istanbul;6.2
Roseau;34.4
Conakry;31.2
Istanbul;23.0
```
A tarefa é escrever um programa em Go que lê o arquivo, calcula o valor mínimo, médio e máximo de temperatura por estação meteorológica e emite os resultados no stdout da seguinte forma (ou seja, ordenado alfabeticamente pelo nome da estação e os valores de resultado por estação no formato `<min>/<mean>/<max>`, arredondado para uma casa decimal):

```
{Abha=-23.0/18.0/59.2, Abidjan=-16.2/26.0/67.3, Abéché=-10.0/29.4/69.0, Accra=-10.1/26.4/66.4, Addis Ababa=-23.7/16.0/67.0, Adelaide=-27.8/17.3/58.5, ...}
```

## Prerequisites
[Go](https://golang.org/dl/)
[Git](https://git-scm.com/downloads)
[Python](https://www.python.org/downloads/)

## Getting Started

Este projeto consiste em 2 programas, um em Go e outro em Python. O programa em Go é responsável por ler o arquivo de 1 bilhão de registros e processar os dados, enquanto o programa em Python é responsável por gerar o arquivo de 1 bilhão de registros.

1. Clone o repositório
```bash
git clone git@github.com:allandav1d/1brc.git
```
2. Entre na pasta do projeto
```bash
cd 1brc
```
3. Execute o script em Python para gerar o arquivo de 1 bilhão de registros
```bash
python3 create.py 1_000_000_000
```
(O script em Python irá gerar um arquivo chamado `measurements.txt` na raiz do projeto com 1 bilhão de registros aproximadamente 15GB)

4. Execute o programa em Go para processar o arquivo de 1 bilhão de registros
```bash
go run .
```

(O programa em Go irá processar o arquivo `measurements.txt` no final irá exibir o resultado das medições no formato `<min>/<mean>/<max>` e o tempo de execução do programa)