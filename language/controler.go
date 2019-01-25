package language

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/xeipuuv/gojsonschema"
	c "go-api-ws/config"
	"go-api-ws/helpers"
	"net/http"
)

func createLanguage(w http.ResponseWriter, r *http.Request) {
	var schemaLoader = gojsonschema.NewReferenceLoader("file://language/jsonSchemaModels/languageSchema.json")
	var language Language
	_ = json.NewDecoder(r.Body).Decode(&language)
	documentLoader := gojsonschema.NewGoLoader(language)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	helpers.PanicErr(err)

	if result.Valid() {
		db, err := c.Conf.GetDb()
		helpers.PanicErr(err)

		result, err := db.Exec("INSERT INTO languages("+
			"id, "+
			"code, "+
			"code3, "+
			"name, "+
			"nameEn) "+
			" VALUES(?, ?, ?, ?, ?)",
			language.Id,
			language.Code,
			language.Code3,
			language.Name,
			language.NameEn)
		fmt.Println(language.Code3)
		fmt.Println(result)
		helpers.PanicErr(err)
	} else {
		err = json.NewEncoder(w).Encode("There is and error registering currency:")
		helpers.PanicErr(err)
		fmt.Printf("The document is not valid. See errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}

func getLanguage(w http.ResponseWriter, r *http.Request) {
	var language Language
	languageID := chi.URLParam(r, "languageID")
	db, err := c.Conf.GetDb()
	helpers.CheckErr(err)

	err = db.QueryRow("SELECT * FROM languages c WHERE id=?", languageID).
		Scan(&language.Id, &language.Code, &language.Code3, &language.Name, &language.NameEn)
	helpers.CheckErr(err)
	err = json.NewEncoder(w).Encode(language)
	helpers.PanicErr(err)
}

func getLanguageList(w http.ResponseWriter, r *http.Request) {
	var language Language
	var languages []Language
	languages = []Language{}

	db, err := c.Conf.GetDb()
	helpers.CheckErr(err)

	rows, err := db.Query("SELECT id, code, code3, name, nameEn FROM languages")
	helpers.CheckErr(err)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&language.Id,
			&language.Code,
			&language.Code3,
			&language.Name,
			&language.NameEn)
		helpers.CheckErr(err)
		languages = append(languages, language)
	}
	err = rows.Err()
	helpers.CheckErr(err)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(languages)
	helpers.PanicErr(err)
}

func removeLanguage(w http.ResponseWriter, r *http.Request) {
	languageID := chi.URLParam(r, "languageID")
	db, err := c.Conf.GetDb()
	helpers.CheckErr(err)

	res, err := db.Exec("DELETE l FROM languages l WHERE l.id=?", languageID)
	fmt.Println(res)
	helpers.CheckErr(err)
}

func updateLanguage(w http.ResponseWriter, r *http.Request) {
	languageID := chi.URLParam(r, "languageID")
	var language Language
	err := json.NewDecoder(r.Body).Decode(&language)
	helpers.PanicErr(err)

	db, err := c.Conf.GetDb()
	helpers.PanicErr(err)

	query, err := db.Prepare("Update languages set code=?, code3=?, name=?, nameEn=? where id=?")
	helpers.PanicErr(err)

	_, er := query.Exec(language.Code, language.Code3, language.Name, language.NameEn, languageID)
	helpers.PanicErr(er)
	fmt.Println(language.Name + " updated in mysql")
}
