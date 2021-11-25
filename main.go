package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type ToDo struct {
	Kegiatan string `json:"kegiatan"`
	Waktu string `json:"waktu"`
}

type JSONResponse struct {
	Code int `json:"code"`
	Success bool `json:"success"`
	Message string `json:"message"`
	// Data []ToDo `json:"data"`
	Data interface{} `json:"data"`
}

func main()  {
	daftarKegiatan := []ToDo{}
	daftarKegiatan = append(daftarKegiatan, ToDo{"Traveling to Solo", "2020-11-30"})
	daftarKegiatan = append(daftarKegiatan, ToDo{"Goes to Bali", "2021-12-30"})
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		// GET localhost:8080
		if r.Method == "GET" {
			// do something with GET
			rw.Header().Add("Conteint-Type", "application/json")
			// res := JSONResponse {
			// 	http.StatusOK,
			// 	true,
			// 	"Testing method GET",
			// 	[]ToDo{},
			// }
			// resJSON, err := json.Marshal(res)
			// if err != nil {
			// 	http.Error(rw, "there is an error", http.StatusInternalServerError)
			// }
			// rw.Write(resJSON)
			res := JSONResponse {
				http.StatusOK,
				true,
				"Data activity is loaded",
				daftarKegiatan,
			}
			resJSON, err := json.Marshal(res)
			if err != nil {
				fmt.Println("There is an Error", http.StatusInternalServerError)
				return
			}
			rw.Write(resJSON)
		} else if r.Method == "POST" {
			// do something with POST
			// add new data
			jsonDecode := json.NewDecoder(r.Body)
			aktivitasBaru := ToDo{}
			res := JSONResponse{}

			if err := jsonDecode.Decode(&aktivitasBaru); err != nil {
				fmt.Println("There is an Error")
				http.Error(rw, "There is an Error while read", http.StatusInternalServerError)
				return
			}

			res.Code = http.StatusCreated
			res.Success = true
			res.Message = "Data is loaded and saved"
			res.Data = aktivitasBaru

			daftarKegiatan = append(daftarKegiatan, aktivitasBaru)

			resJSON, err := json.Marshal(res)
			if err != nil {
				fmt.Println("There is an Error while saved", http.StatusInternalServerError)
				return
			}
			rw.Header().Add("Content-Type", "application/json")
			rw.Write(resJSON)
		}
	})
	fmt.Println("Listening on: 8080 ....")
	log.Fatal(http.ListenAndServe(":8080", nil))
}