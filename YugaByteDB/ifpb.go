//INSTITUTO FEDERAL DA PARAIBA - IFPB
//PROGRAMA DE POS GRADUACAO EM TECNOLOGIA DA INFORMACAO
//DISCIPLINA DE BANCO DE DADOS
//PROFESSORES: DAMIRES E DIEGO
//ALUNOS: SÉRGIO OLIVEIRA, KLIVIO RAFAEL, RAISA PONCE, RUAN VIRGINIO

//DESCRICAO: CRIACAO DE UM BANCO DE DADOS NO YUGABYTES E REALIZACAO DE CRUD.

package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var databaseConnection *sql.DB

type Post struct {
	Id       int
	Conteudo string
	Autor    string
}

func OpenDB() {
	var err error
	configurationString := "host = 127.0.0.1 port = 5433 dbname = yugabyte user = yugabyte password = yugabyte sslmode = disable"
	databaseConnection, err = sql.Open("postgres", configurationString)
	if err != nil {
		log.Println("Error: ", err)
	}

	err = databaseConnection.Ping()
	if err != nil {
		log.Println("Error: ", err)
	} else {
		log.Println("Conexao com o banco de dados bem sucedida!")
	}
}

func (p *Post) CreatePost() (err error) {
	sqlStatement := "INSERT INTO posts (id, content, author) VALUES ($1, $2, $3) RETURNING id"
	sqlRowPtr := databaseConnection.QueryRow(sqlStatement, p.Id, p.Conteudo, p.Autor)
	err = sqlRowPtr.Scan(&p.Id)
	if err == nil {
		log.Println("INSERT realizado com sucesso!")
	}
	return err
}

func ReadPost(id int) (post Post, err error) {
	post = Post{}
	sqlStatement := ("SELECT id, content, author FROM posts WHERE id = $1")
	err = databaseConnection.QueryRow(sqlStatement, id).Scan(&post.Id, &post.Conteudo, &post.Autor)

	return
}

func (p *Post) UpdatePost() (err error) {
	sqlStatement := "UPDATE posts SET content = $2, author = $3 WHERE id = $1"
	_, err = databaseConnection.Exec(sqlStatement, p.Id, p.Conteudo, p.Autor)

	return err
}

func (p *Post) DeletePost() (err error) {
	sqlStatement := "DELETE FROM posts WHERE id = $1"
	databaseConnection.Exec(sqlStatement, p.Id)

	return err
}

func GetAllPosts(limite int) (posts []Post, err error) {
	sqlStatement := "SELECT id, content, author FROM  posts limit $1"
	rows, err := databaseConnection.Query(sqlStatement, limite)
	if err != nil {
		log.Println("Erro: ", err)
	}

	for rows.Next() {
		post := Post{}
		rows.Scan(&post.Id, &post.Conteudo, &post.Autor)
		posts = append(posts, post)
	}
	rows.Close()
	return
}

func main() {
	OpenDB()

	post := Post{
		Id:       3,
		Conteudo: "YugaByte 3",
		Autor:    "IFPB 3",
	}

	//Create
	//post.CreatePost()

	//Read
	//lerPost, _ := ReadPost(1)
	//fmt.Println(lerPost)

	//Update
	//lerPost.Conteudo = "Golang"
	//lerPost.Autor = "Google"
	//lerPost.UpdatePost()

	//Delete
	//post.DeletePost()

	//GetAll
	//posts, _ := GetAllPosts(post.Id)
	//fmt.Println(posts)

}
