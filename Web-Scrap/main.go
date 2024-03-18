/*INSTITUTO FEDERAL DA PARAIBA - IFPB
UNIDADE ACADEMICA DE INFORMACAO E COMUNICACAO
MESTRADO PROFISSIONAL EM TECNOLOGIA DA INFORMACAO

DISCIPLINA: BANCO DE DADOS
PROFESSORES: DAMIRES E DIEGO

ALUNO: SERGIO RICARDO DE OLIVEIRA BRANDAO

ATIVIDADE 02, QUESITO 03: Crie um Scraper para ler os dados dos docentes da página do PPGTI (https://www.ifpb.edu.br/ppgti/programa/corpo-docente) - construa um dataframe que liste o nome, linha de pesquisa, url do lattes e e-mail de cada professor.

SOLUCAO IMPLEMENTADA EM GOLANG*/


package main

import (
    "fmt"
    "github.com/gocolly/colly"
    "strings"
)

type Professor struct {
    Nome           string
    LinhaPesquisa  string
    URLLattes      string
    Email          string
}

func main() {
    c := colly.NewCollector()

    var professores []Professor

    c.OnHTML("div#parent-fieldname-text", func(e *colly.HTMLElement) {
        e.ForEach("h4", func(_ int, el *colly.HTMLElement) {
            nome := el.Text
            details := strings.Split(el.DOM.Next().Text(), "\n")
            linhaPesquisa := strings.TrimSpace(strings.Replace(details[0], "Linha de Pesquisa:", "", -1))
            urlLattes := ""
            email := ""
            if len(details) > 1 {
                urlAndEmail := strings.Split(details[1], "E-mail:")
                urlLattes = strings.TrimSpace(strings.Split(urlAndEmail[0], " ")[0])
                if len(urlAndEmail) > 1 {
                    email = strings.TrimSpace(urlAndEmail[1])
                }
            }

            professor := Professor{
                Nome:          strings.TrimSpace(nome),
                LinhaPesquisa: strings.TrimSpace(linhaPesquisa),
                URLLattes:     strings.TrimSpace(urlLattes),
                Email:         strings.TrimSpace(email),
            }
            professores = append(professores, professor)
        })
    })

    c.OnRequest(func(r *colly.Request) {
        fmt.Println("Visiting", r.URL)
    })

    c.Visit("https://www.ifpb.edu.br/ppgti/programa/corpo-docente")

    for _, professor := range professores {
        fmt.Println("Nome: ", professor.Nome)
        fmt.Println("Linha de Pesquisa: ", professor.LinhaPesquisa)
        fmt.Println("URL do Lattes: ", professor.URLLattes)
        fmt.Println("Email: ", professor.Email)
        fmt.Println("-----------------------------")
    }
}