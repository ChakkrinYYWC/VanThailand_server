package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"van_thailand_server/services"
)

func HandleRequest(ctx context.Context) {
	http.HandleFunc("/vanSchedule", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			paramId := r.URL.Query().Get("id")
			paramVanId := r.URL.Query().Get("van_id")
			if paramId != "" {
				result := services.GetVanSchedule(ctx, paramId)
				if result != nil {
					json.NewEncoder(w).Encode(result)
					return
				} else {
					ReturnFailed(w)
					return
				}
			} else if paramVanId != "" {
				log.Println("Get van schedule by van id: ", paramVanId)
				results := services.GetVanSchedules(ctx, paramVanId)
				log.Println("Result: ", results)
				if results != nil {
					json.NewEncoder(w).Encode(results)
					return
				} else {
					ReturnFailed(w)
					return
				}
			} else {
				log.Println("No params given.")
				ReturnFailed(w)
				return
			}
		} else if r.Method == http.MethodPost {
			vanId := r.FormValue("vanId")
			date := r.FormValue("date")
			destination := r.FormValue("destination")
			if vanId != "" && date != "" && destination != "" {
				result := services.CreateVanSchedule(ctx, vanId, date, destination)
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
			vanId := r.FormValue("vanId")
			date := r.FormValue("date")
			destination := r.FormValue("destination")
			if (vanId != "" || date != "" || destination != "") && paramId != "" {
				result := services.UpdateSchedule(ctx, paramId, vanId, date, destination)
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
		if r.Method == http.MethodGet {
			param := r.URL.Query().Get("id")
			if param != "" {
				results := services.GetVan(ctx, param)
				if results != nil {
					json.NewEncoder(w).Encode(results)
					return
				} else {
					ReturnFailed(w)
					return
				}
			} else {
				results := services.GetVans(ctx)
				if results != nil {
					json.NewEncoder(w).Encode(results)
					return
				} else {
					ReturnFailed(w)
					return
				}
			}
		} else if r.Method == http.MethodPost {
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
			name := r.FormValue("name")
			code := r.FormValue("code")
			desc := r.FormValue("desc")
			files := r.MultipartForm.File["images"]
			imagePosition := r.FormValue("imagePosition")
			if paramId != "" && (name != "" || code != "" || desc != "" || (files != nil && imagePosition != "")) {
				result := services.UpdateVan(ctx, paramId, name, code, desc, files, imagePosition)
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
