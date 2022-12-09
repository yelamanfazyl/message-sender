package server

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"encoding/json"
	"upgradeFinal/startbot/cmd/bot"
)

type Message struct {
	Message string `json:"message"`
}

const keyServerAddr = "serverAddr"

type Server struct {
	upgradeBot bot.UpgradeBot
}

func (s *Server) getRoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Printf("%s: got / request\n", ctx.Value(keyServerAddr))
	io.WriteString(w, "This is my website!\n")
}

func (s *Server) getHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	resp := make(map[string]string)
	resp["status"] = "ok"
	
	jsonResp, err := json.Marshal(resp)
	
	if err != nil {
		fmt.Printf("Error happened in JSON marshal. Err: %s", err)
	}
	
	w.Write(jsonResp)
	return
}

func (s *Server) sendMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var message Message

		err := json.NewDecoder(r.Body).Decode(&message)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			w.WriteHeader(http.StatusCreated)
			w.Header().Set("Content-Type", "application/json")

			resp := make(map[string]string)
			resp["error"] = "something went wrong" + err.Error()
			
			jsonResp, err := json.Marshal(resp)
			
			if err != nil {
				fmt.Printf("Error happened in JSON marshal. Err: %s", err)
			}
			
			w.Write(jsonResp)
			return
		}

		if(message.Message == "" || message.Message == " ") {
			http.Error(w, "Message is empty", http.StatusBadRequest)

			w.WriteHeader(http.StatusCreated)
			w.Header().Set("Content-Type", "application/json")

			resp := make(map[string]string)
			resp["error"] = "something went wrong, message is empty"
			
			jsonResp, err := json.Marshal(resp)
			
			if err != nil {
				fmt.Printf("Error happened in JSON marshal. Err: %s", err)
			}
			
			w.Write(jsonResp)
			
			return
		}

		fmt.Println("Message: ", message.Message)

		err = s.upgradeBot.SendAll(message.Message)

		if err != nil {
			fmt.Println("Error happened in sending message: ", err)
			w.WriteHeader(http.StatusCreated)
			w.Header().Set("Content-Type", "application/json")

			resp := make(map[string]string)
			resp["error"] = "something went wrong" + err.Error()
			
			jsonResp, err := json.Marshal(resp)
			
			if err != nil {
				fmt.Printf("Error happened in JSON marshal. Err: %s", err)
			}
			
			w.Write(jsonResp)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")

		resp := make(map[string]string)
		resp["status"] = "ok"
			
		jsonResp, err := json.Marshal(resp)
			
		if err != nil {
			fmt.Printf("Error happened in JSON marshal. Err: %s", err)
		}
			
		w.Write(jsonResp)
		return
	}
}

func (s *Server) StartServer(upgradeBot bot.UpgradeBot) {
	s.upgradeBot = upgradeBot

	mux := http.NewServeMux()
	mux.HandleFunc("/", s.getRoot)
	mux.HandleFunc("/health", s.getHealth)
	mux.HandleFunc("/message", s.sendMessage)

	ctx := context.Background()
	server := &http.Server{
		Addr:    ":8888",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
	}
	
	err := server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error listening for server: %s\n", err)
	}
}