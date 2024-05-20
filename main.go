package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type router struct  {
    Get     *mux.Router 
//    Post    *mux.Router
//    Put     *mux.Router
//    Delete  *mux.Router
}

func NewRouter(mr *mux.Router) router{
  return router{
        Get     : mr.Methods(http.MethodGet).Subrouter(),
//        Post    : mr.Methods(http.MethodPost).Subrouter(),
//        Put     : mr.Methods(http.MethodPut).Subrouter(),
//        Delete  : mr.Methods(http.MethodDelete).Subrouter(),
  }
}


func main(){

    sm := mux.NewRouter()
    routers := NewRouter(sm)
    
    routers.Get.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("content-type", "application/json")
        w.WriteHeader(http.StatusOK)
        jsonEncoder := json.NewEncoder(w)
        if err := jsonEncoder.Encode(struct{Msg  string}{Msg : "Hello World"}); err != nil{
            slog.Error("Unable to Encode and Write body to the client", "Details", err.Error())            
        }

    })

    server := http.Server{
        Addr: ":8080",
        Handler: sm,
        ReadTimeout: 10 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout: 30 * time.Second,
        MaxHeaderBytes: 1 >> 20, // 1 MB
    }
    
    slog.Info("Server Running")
    if err := server.ListenAndServe(); err != nil{
        slog.Error("Unable to Run the server", "Details", err.Error())
    }
    fmt.Println("Hello World")
}
