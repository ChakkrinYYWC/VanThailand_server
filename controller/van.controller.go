package controller

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"van_thailand_server/models"
	"van_thailand_server/services"
)

func HandleRequest(ctx context.Context) {
	http.HandleFunc("/vanSchedule", func(w http.ResponseWriter, r *http.Request) {
		var targetSchedule *models.ScheduleStruct
		if r.Method == http.MethodGet {
			paramId := r.URL.Query().Get("id")
			paramVanId := r.URL.Query().Get("van_id")
			if paramId != "" {
				result := services.GetVanSchedule(ctx, paramId)
				json.NewEncoder(w).Encode(result)
				return
			} else if paramVanId != "" {
				results := services.GetVanSchedules(ctx, paramVanId)
				json.NewEncoder(w).Encode(results)
				return
			} else {
				log.Println("No params given.")
				ReturnFailed(w)
				return
			}
		} else if r.Method == http.MethodPost {
			body, err := io.ReadAll(r.Body)
			if err != nil {
				log.Println("Read body failed: ", err)
				ReturnFailed(w)
				return
			}
			err = json.Unmarshal(body, &targetSchedule)
			if err != nil {
				log.Println("Unmarshal failed: ", err)
				ReturnFailed(w)
				return
			}
			if targetSchedule != nil {
				result := services.CreateVanSchedule(ctx, targetSchedule)
				if result.InsertedID != "" {
					ReturnSuccess(w)
					return
				} else {
					log.Println("Can not create schedule.")
					ReturnFailed(w)
					return
				}
			} else {
				log.Println("No body given.")
				ReturnFailed(w)
				return
			}
		} else if r.Method == http.MethodPatch {
			paramId := r.URL.Query().Get("id")
			body, err := io.ReadAll(r.Body)
			if err != nil {
				log.Println("Invalid body: ", err)
				ReturnFailed(w)
				return
			}
			err = json.Unmarshal(body, &targetSchedule)
			if err != nil {
				log.Println("Unmarshal failed: ", err)
				ReturnFailed(w)
				return
			}
			if targetSchedule != nil {
				result := services.UpdateSchedule(ctx, paramId, targetSchedule)
				if result != 0 {
					ReturnSuccess(w)
					return
				} else {
					log.Println("Can not update schedule.")
					ReturnFailed(w)
					return
				}
			} else {
				log.Println("No body given.")
				ReturnFailed(w)
				return
			}
		} else if r.Method == http.MethodDelete {
			paramId := r.URL.Query().Get("id")
			if paramId != "" {
				result := services.DeleteSchedule(ctx, paramId)
				if result != 0 {
					ReturnSuccess(w)
					return
				} else {
					log.Println("Can not delete schedule.")
					ReturnFailed(w)
					return
				}
			} else {
				log.Println("No params given.")
				ReturnFailed(w)
				return
			}
		} else {
			log.Println(w, "Method not allowed.")
			return
		}
	})
	// ################################################################################################
	http.HandleFunc("/vanManagement", func(w http.ResponseWriter, r *http.Request) {
		var targetVan *models.RecieveVansStruct
		if r.Method == http.MethodGet {
			param := r.URL.Query().Get("id")
			if param != "" {
				results := services.GetVan(ctx, param)
				json.NewEncoder(w).Encode(results)
				return
			} else {
				results := services.GetVans(ctx)
				json.NewEncoder(w).Encode(results)
				return
			}
		} else if r.Method == http.MethodPost {
			// body, err := io.ReadAll(r.Body)
			// if err != nil {
			// 	log.Println("Invalid body: ", err)
			// 	return
			// }
			// err = json.Unmarshal(body, &targetVan)
			// if err != nil {
			// 	log.Println("Unmarshal failed: ", err)
			// 	return
			// }
			name := r.FormValue("name")
			code := r.FormValue("code")
			desc := r.FormValue("desc")
			files := r.MultipartForm.File["images"]
			if name != "" && code != "" {
				result := services.CreateVan(ctx, name, code, desc, files)
				if result.InsertedID != "" {
					ReturnSuccess(w)
					return
				} else {
					log.Println("Can not create van.")
					ReturnFailed(w)
					return
				}
			} else {
				log.Println("body invalid.")
				ReturnFailed(w)
				return
			}
		} else if r.Method == http.MethodPatch {
			paramId := r.URL.Query().Get("id")
			body, err := io.ReadAll(r.Body)
			if err != nil {
				log.Println("PATCH vanManagement recieve invalid body: ", err)
				ReturnFailed(w)
				return
			}
			err = json.Unmarshal(body, &targetVan)
			if err != nil {
				log.Println("Unmarshal failed: ", err)
				ReturnFailed(w)
				return
			}
			if targetVan != nil {
				result := services.UpdateVan(ctx, paramId, targetVan)
				if result != 0 {
					ReturnSuccess(w)
					return
				} else {
					log.Println("Can not update van.")
					ReturnFailed(w)
					return
				}
			} else {
				log.Println("No body given.")
				ReturnFailed(w)
				return
			}
		} else if r.Method == http.MethodDelete {
			paramId := r.URL.Query().Get("id")
			if paramId != "" {
				result := services.DeleteVan(ctx, paramId)
				if result != 0 {
					ReturnSuccess(w)
					return
				} else {
					log.Println("Can not delete van.")
					ReturnFailed(w)
					return
				}
			} else {
				log.Println("No params given.")
				ReturnFailed(w)
				return
			}
		} else {
			log.Println(w, "Method not allowed.")
			return
		}
	})
}
