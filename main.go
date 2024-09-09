package main

import (
	"fmt"
	"net/http"
	"log"
	"database/sql"
	"strings"
	"html/template"
	"encoding/base64"

	_ "github.com/lib/pq"
)

type perfilCandidato struct {
	Foto []byte
	Nome string
	Id string
	Atr1 string
	Atr2 string
	Atr3 string
}

type cardInfo struct {
	Imagem string
	Nome string
	Id string
	Partido string
}

type pageCards struct {
	Cards []cardInfo
}

//Open elections data base
func openDataBase() (*sql.DB, error){
	db, err := sql.Open("postgres", "host=localhost port=5432 user=myuser password=myuser dbname=elections sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}

	err = db.Ping()

	return db, err
}

//Generates an sql command string
func createQueryStr(atr1 string, atr2 string, atr3 string, view_num int) string {
	var queryStr string

	switch view_num {
		case 1:
			if (atr1 != "" && atr2 != "" && atr3 != "") {
				queryStr = fmt.Sprintf("SELECT * FROM view_cor_cargo_partido WHERE cor = '%s' AND cargo = '%s' AND partido = '%s'", atr1, atr2, atr3)
			}
			if (atr1 == "" && atr2 != "" && atr3 != "") {
				queryStr = fmt.Sprintf("SELECT * FROM view_cor_cargo_partido WHERE cargo = '%s' AND partido = '%s'", atr2, atr3)
			}
			if (atr1 == "" && atr2 == "" && atr3 != "") {
				queryStr = fmt.Sprintf("SELECT * FROM view_cor_cargo_partido WHERE partido = '%s'", atr3)
			}
			if (atr1 == "" && atr2 == "" && atr3 == "") {
				queryStr = "SELECT * FROM view_cor_cargo_partido"
			}
			if (atr1 != "" && atr2 == "" && atr3 != "") {
				queryStr = fmt.Sprintf("SELECT * FROM view_cor_cargo_partido WHERE cor = '%s' AND partido = '%s'", atr1, atr3)
			}
			if (atr1 != "" && atr2 == "" && atr3 == "") {
				queryStr = fmt.Sprintf("SELECT * FROM view_cor_cargo_partido WHERE cor = '%s'", atr1)
			}
			if (atr1 != "" && atr2 != "" && atr3 == "") {
				queryStr = fmt.Sprintf("SELECT * FROM view_cor_cargo_partido WHERE cor = '%s' AND cargo = '%s'", atr1, atr2)
			}
			if (atr1 == "" && atr2 != "" && atr3 == "") {
				queryStr = fmt.Sprintf("SELECT * FROM view_cor_cargo_partido WHERE cargo = '%s'", atr2)
			}
		case 2:
			if (atr1 != "" && atr2 != "" && atr3 != "") {
				queryStr = fmt.Sprintf("SELECT * FROM view_federacao_turno_partido WHERE federacao = '%s' AND turno = '%s' AND partido = '%s'", atr1, atr2, atr3)
			}
			if (atr1 == "" && atr2 != "" && atr3 != "") {
				queryStr = fmt.Sprintf("SELECT * FROM view_cor_cargo_partido WHERE turno = '%s' AND partido = '%s'", atr2, atr3)
			}
			if (atr1 == "" && atr2 == "" && atr3 != "") {
				queryStr = fmt.Sprintf("SELECT * FROM view_cor_cargo_partido WHERE partido = '%s'", atr3)
			}
			if (atr1 == "" && atr2 == "" && atr3 == "") {
				queryStr = "SELECT * FROM view_federacao_turno_partido"
			}
			if (atr1 != "" && atr2 == "" && atr3 != "") {
				queryStr = fmt.Sprintf("SELECT * FROM view_federacao_turno_partido WHERE federacao = '%s' AND partido = '%s'", atr1, atr3)
			}
			if (atr1 != "" && atr2 == "" && atr3 == "") {
				queryStr = fmt.Sprintf("SELECT * FROM view_federacao_turno_partido WHERE federacao = '%s'", atr1)
			}
			if (atr1 != "" && atr2 != "" && atr3 == "") {
				queryStr = fmt.Sprintf("SELECT * FROM view_federacao_turno_partido WHERE federacao = '%s' AND turno = '%s'", atr1, atr2)
			}
			if (atr1 == "" && atr2 != "" && atr3 == "") {
				queryStr = fmt.Sprintf("SELECT * FROM view_federacao_turno_partido WHERE turno = '%s'", atr2)
			}
		case 3:
			if (atr1 != "" && atr2 != "" && atr3 != "") {
				queryStr = fmt.Sprintf("SELECT * FROM view_genero_cargo_partido WHERE genero = '%s' AND cargo = '%s' AND partido = '%s'", atr1, atr2, atr3)
			}
			if (atr1 == "" && atr2 != "" && atr3 != "") {
				queryStr = fmt.Sprintf("SELECT * FROM view_genero_cargo_partido WHERE cargo = '%s' AND partido = '%s'", atr2, atr3)
			}
			if (atr1 == "" && atr2 == "" && atr3 != "") {
				queryStr = fmt.Sprintf("SELECT * FROM view_genero_cargo_partido WHERE partido = '%s'", atr3)
			}
			if (atr1 == "" && atr2 == "" && atr3 == "") {
				queryStr = "SELECT * FROM view_genero_cargo_partido"
			}
			if (atr1 != "" && atr2 == "" && atr3 != "") {
				queryStr = fmt.Sprintf("SELECT * FROM view_genero_cargo_partido WHERE genero = '%s' AND partido = '%s'", atr1, atr3)
			}
			if (atr1 != "" && atr2 == "" && atr3 == "") {
				queryStr = fmt.Sprintf("SELECT * FROM view_genero_cargo_partido WHERE genero = '%s'", atr1)
			}
			if (atr1 != "" && atr2 != "" && atr3 == "") {
				queryStr = fmt.Sprintf("SELECT * FROM view_genero_cargo_partido WHERE genero = '%s' AND cargo = '%s'", atr1, atr2)
			}
			if (atr1 == "" && atr2 != "" && atr3 == "") {
				queryStr = fmt.Sprintf("SELECT * FROM view_genero_cargo_partido WHERE cargo = '%s'", atr2)
			}
		case 4:
			if (atr1 != "" && atr2 != "" && atr3 != "") {
				queryStr = fmt.Sprintf("SELECT * FROM view_instrucao_cargo_partido WHERE grau_instrucao = '%s' AND cargo = '%s' AND partido = '%s'", atr1, atr2, atr3)
			}
			if (atr1 == "" && atr2 != "" && atr3 != "") {
				queryStr = fmt.Sprintf("SELECT * FROM view_instrucao_cargo_partido WHERE cargo = '%s' AND partido = '%s'", atr2, atr3)
			}
			if (atr1 == "" && atr2 == "" && atr3 != "") {
				queryStr = fmt.Sprintf("SELECT * FROM view_instrucao_cargo_partido WHERE partido = '%s'", atr3)
			}
			if (atr1 == "" && atr2 == "" && atr3 == "") {
				queryStr = "SELECT * FROM view_instrucao_cargo_partido"
			}
			if (atr1 != "" && atr2 == "" && atr3 != "") {
				queryStr = fmt.Sprintf("SELECT * FROM view_instrucao_cargo_partido WHERE grau_instrucao = '%s' AND partido = '%s'", atr1, atr3)
			}
			if (atr1 != "" && atr2 == "" && atr3 == "") {
				queryStr = fmt.Sprintf("SELECT * FROM view_instrucao_cargo_partido WHERE grau_instrucao = '%s'", atr1)
			}
			if (atr1 != "" && atr2 != "" && atr3 == "") {
				queryStr = fmt.Sprintf("SELECT * FROM view_instrucao_cargo_partido WHERE grau_instrucao = '%s' AND cargo = '%s'", atr1, atr2)
			}
			if (atr1 == "" && atr2 != "" && atr3 == "") {
				queryStr = fmt.Sprintf("SELECT * FROM view_instrucao_cargo_partido WHERE cargo = '%s'", atr2)
			}
		case 5:
			if (atr1 != "" && atr2 != "" && atr3 != "") {
				queryStr = fmt.Sprintf("SELECT * FROM view_instrucao_ocupacao_data WHERE grau_instrucao = '%s' AND ocupacao = '%s' AND data = '%s'", atr1, atr2, atr3)
			}
			if (atr1 == "" && atr2 != "" && atr3 != "") {
				queryStr = fmt.Sprintf("SELECT * FROM view_instrucao_ocupacao_data WHERE ocupacao = '%s' AND data = '%s'", atr2, atr3)
			}
			if (atr1 == "" && atr2 == "" && atr3 != "") {
				queryStr = fmt.Sprintf("SELECT * FROM view_instrucao_ocupacao_data WHERE data = '%s'", atr3)
			}
			if (atr1 == "" && atr2 == "" && atr3 == "") {
				queryStr = "SELECT * FROM view_instrucao_ocupacao_data"
			}
			if (atr1 != "" && atr2 == "" && atr3 != "") {
				queryStr = fmt.Sprintf("SELECT * FROM view_instrucao_ocupacao_data WHERE grau_instrucao = '%s' AND data = '%s'", atr1, atr3)
			}
			if (atr1 != "" && atr2 == "" && atr3 == "") {
				queryStr = fmt.Sprintf("SELECT * FROM view_instrucao_ocupacao_data WHERE grau_instrucao = '%s'", atr1)
			}
			if (atr1 != "" && atr2 != "" && atr3 == "") {
				queryStr = fmt.Sprintf("SELECT * FROM view_instrucao_ocupacao_data WHERE grau_instrucao = '%s' AND ocupacao = '%s'", atr1, atr2)
			}
			if (atr1 == "" && atr2 != "" && atr3 == "") {
				queryStr = fmt.Sprintf("SELECT * FROM view_instrucao_ocupacao_data WHERE ocupacao = '%s'", atr2)
			}
	}

	return queryStr
}

//Gets data from views on the database 
func getFromView(query string) []perfilCandidato {
	db, err := openDataBase()
	if err != nil {
		fmt.Println("Failed to open database. Error:")
		fmt.Println(err)
		return nil
	}

	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Query has failed. Error:")
		fmt.Println(err)
		return nil
	}

	defer rows.Close()
	
	var perfil_slice []perfilCandidato

	for rows.Next() {
		var perfil perfilCandidato 
		
		err = rows.Scan(&perfil.Foto, &perfil.Nome, &perfil.Id, &perfil.Atr1, &perfil.Atr2, &perfil.Atr3)
		if err != nil {
			fmt.Println(err)
		}

		perfil_slice = append(perfil_slice, perfil)
	}	
	
	return perfil_slice
}

//Takes the query result and selects the corresponding results to make a card
func generateCards(query_result [][]perfilCandidato) []cardInfo {
	var cards []cardInfo

	for row_index := range query_result[0] {
		name := query_result[0][row_index].Nome
		table_index := 1
		name_count := 1

		for table_index < len(query_result) {
			for row_index2 := range query_result[table_index] {
				if query_result[table_index][row_index2].Nome == name {
					name_count += 1
				}	
			}

			table_index += 1
		}

		if name_count >= 5 {
			imagem := base64.StdEncoding.EncodeToString(query_result[0][row_index].Foto)	

			current_card := cardInfo{imagem, query_result[0][row_index].Nome, query_result[0][row_index].Id, query_result[0][row_index].Atr3}
			cards = append(cards, current_card)
		}
	}

	return cards
}

//Puts candidates info cards on the page
func executeTemplate(w http.ResponseWriter, page_struct pageCards) {
	templ, err := template.ParseFiles("pages/desktop.html")

	if err != nil {
		fmt.Println(err)
		return
	}
	
	templ.Execute(w, page_struct)
}

//Request handler for "/"
func indexHandler(w http.ResponseWriter, r* http.Request) {
	no_cards := pageCards{nil}
	executeTemplate(w, no_cards)
}

//Request handler for "/search"
func searchHandler(w http.ResponseWriter, r* http.Request) {
	if r.Method == "GET" {
		r.ParseForm()
	}
	
	var query_result [][]perfilCandidato
	var queryStr string
	
	//Query 1
	queryStr = createQueryStr(strings.ToUpper(r.Form["cor-raca"][0]), strings.ToUpper(r.Form["cargo"][0]), strings.ToUpper(r.Form["partido"][0]), 1)
	query_result = append(query_result, getFromView(queryStr))

	//Query 2
	queryStr = createQueryStr(strings.ToUpper(r.Form["federacao"][0]), strings.ToUpper(r.Form["turno"][0]), strings.ToUpper(r.Form["partido"][0]), 2)
	query_result = append(query_result, getFromView(queryStr))
	
	//Query 3
	queryStr = createQueryStr(strings.ToUpper(r.Form["genero"][0]), strings.ToUpper(r.Form["cargo"][0]), strings.ToUpper(r.Form["partido"][0]), 3)
	query_result = append(query_result, getFromView(queryStr))
	
	//Query 4
	queryStr = createQueryStr(strings.ToUpper(r.Form["instrucao"][0]), strings.ToUpper(r.Form["cargo"][0]), strings.ToUpper(r.Form["partido"][0]), 4)
	query_result = append(query_result, getFromView(queryStr))
	
	//Query 5
	queryStr = createQueryStr(strings.ToUpper(r.Form["instrucao"][0]), strings.ToUpper(r.Form["ocupacao"][0]), "", 5)
	query_result = append(query_result, getFromView(queryStr))
	
	cards := generateCards(query_result) 
	page_struct := pageCards{cards}	

	executeTemplate(w, page_struct)
}

func main() {
	//Creating http request handlers
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/search", searchHandler)
	
	//Creating image and audio files handler	
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	//Run http server on port 8080
	fmt.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
