package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"van_thailand_server/models"
	"van_thailand_server/services"
)

func HandleRequest(ctx context.Context) {
	http.HandleFunc("/vanSchedule", func(w http.ResponseWriter, r *http.Request) {
		var targetSchedule *models.ScheduleStruct
		if r.Method == "GET" {
			paramId := r.URL.Query().Get("id")
			paramVanId := r.URL.Query().Get("van_id")
			if paramId != "" {
				result := services.GetVanSchedule(ctx, paramId)
				json.NewEncoder(w).Encode(result)
			} else if paramVanId != "" {
				results := services.GetVanSchedules(ctx, paramVanId)
				json.NewEncoder(w).Encode(results)
			} else {
				resp := models.Response{
					Message: "fail",
				}
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(resp)
			}
		} else if r.Method == "POST" {
			body, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Println("POST vanSchedule recieve invalid body: ", err)
				return
			}
			err = json.Unmarshal(body, &targetSchedule)
			if err != nil {
				fmt.Println("Unmarshal failed: ", err)
				return
			}
			result := services.CreateVanSchedule(ctx, targetSchedule)
			if result.InsertedID != "" {
				resp := models.Response{
					Message: "success",
				}
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(resp)
			} else {
				resp := models.Response{
					Message: "fail",
				}
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(resp)
			}
		} else if r.Method == "PATCH" {
			paramId := r.URL.Query().Get("id")
			body, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Println("PATCH vanSchedule recieve invalid body: ", err)
				return
			}
			err = json.Unmarshal(body, &targetSchedule)
			if err != nil {
				fmt.Println("Unmarshal failed: ", err)
				return
			}
			if targetSchedule != nil {
				result := services.UpdateSchedule(ctx, paramId, targetSchedule)
				if result != 0 {
					resp := models.Response{
						Message: "success",
					}
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(resp)
				} else {
					resp := models.Response{
						Message: "fail",
					}
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(resp)
				}
			}
		} else if r.Method == "DELETE" {
			paramId := r.URL.Query().Get("id")
			if paramId != "" {
				result := services.DeleteSchedule(ctx, paramId)
				if result != 0 {
					resp := models.Response{
						Message: "success",
					}
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(resp)
				} else {
					resp := models.Response{
						Message: "fail",
					}
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(resp)
				}
			} else {
				results := services.GetVans(ctx)
				json.NewEncoder(w).Encode(results)
			}
		} else {
			fmt.Fprintf(w, "Method not allowed\n")
		}
	})
	// ################################################################################################
	http.HandleFunc("/vanManagement", func(w http.ResponseWriter, r *http.Request) {
		var targetVan *models.VansStruct
		if r.Method == "GET" {
			param := r.URL.Query().Get("id")
			if param != "" {
				results := services.GetVan(ctx, param)
				json.NewEncoder(w).Encode(results)
			} else {
				results := services.GetVans(ctx)
				json.NewEncoder(w).Encode(results)
			}
		} else if r.Method == "POST" {
			body, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Println("POST vanManagement recieve invalid body: ", err)
				return
			}
			err = json.Unmarshal(body, &targetVan)
			if err != nil {
				fmt.Println("Unmarshal failed: ", err)
				return
			}
			result := services.CreateVan(ctx, targetVan)
			if result.InsertedID != "" {
				resp := models.Response{
					Message: "success",
				}
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(resp)
			} else {
				resp := models.Response{
					Message: "fail",
				}
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(resp)
			}
		} else if r.Method == "PATCH" {
			paramId := r.URL.Query().Get("id")
			body, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Println("PATCH vanManagement recieve invalid body: ", err)
				return
			}
			err = json.Unmarshal(body, &targetVan)
			if err != nil {
				fmt.Println("Unmarshal failed: ", err)
				return
			}
			if targetVan != nil {
				result := services.UpdateVan(ctx, paramId, targetVan)
				if result != 0 {
					resp := models.Response{
						Message: "success",
					}
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(resp)
				} else {
					resp := models.Response{
						Message: "fail",
					}
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(resp)
				}
			}
		} else if r.Method == "DELETE" {
			paramId := r.URL.Query().Get("id")
			if paramId != "" {
				result := services.DeleteVan(ctx, paramId)
				if result != 0 {
					resp := models.Response{
						Message: "success",
					}
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(resp)
				} else {
					resp := models.Response{
						Message: "fail",
					}
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(resp)
				}
			} else {
				results := services.GetVans(ctx)
				json.NewEncoder(w).Encode(results)
			}
		} else {
			fmt.Fprintf(w, "Method not allowed\n")
		}
	})
	http.ListenAndServe(":8080", nil)
}
