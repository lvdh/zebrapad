package server

import (
        "embed"
        "encoding/json"
        "io"
        "io/fs"
        "net/http"

        "github.com/lvdh/zebrapad/internal/printer"
)

//go:embed web/*
var webAssets embed.FS

func NewHTTPServer(p printer.Printer) http.Handler {
        mux := http.NewServeMux()

        // prepare /assets/* -- chroots web/*
        webAssetsFS, err := fs.Sub(webAssets, "web")
        if err != nil {
                panic(err)
        }
        // GET /assets/* -- serves web/*
        mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.FS(webAssetsFS))))

        // GET / -- serves index.html
        mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                if r.URL.Path != "/" {
                        http.NotFound(w, r)
                        return
                }
                f, err := webAssetsFS.Open("index.html")
                if err != nil {
                        http.Error(w, "index.html not found", http.StatusInternalServerError)
                        return
                }
                defer f.Close()
                w.Header().Set("Content-Type", "text/html")
                io.Copy(w, f)
        })

        // POST /print -- sends raw ZPL to the label printer, example request:
        //   curl 'http://localhost:8080/print' \
        //     -X POST \
        //     -H 'Content-Type: application/json' \
        //     --data-raw '{"zpl":"^XA^FO50,50^ADN,36,20^FDTest Label^FS^XZ"}'
        mux.HandleFunc("/print", func(w http.ResponseWriter, r *http.Request) {
                if r.Method != http.MethodPost {
                        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
                        return
                }
                var req struct {
                        Data string `json:"zpl"`
                }
                dec := json.NewDecoder(r.Body)
                err := dec.Decode(&req)
                if err != nil || req.Data == "" {
                        http.Error(w, `{"error":"missing data input"}`, http.StatusBadRequest)
                        return
                }
                err = p.Print([]byte(req.Data))
                if err != nil {
                        http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
                        return
                }
                w.Header().Set("Content-Type", "application/json")
                w.Write([]byte(`{"status":"sent"}`))
        })

        // GET /api/printers -- returns discovered printers in JSON format
        mux.HandleFunc("/api/printers", handlePrintersAPI)

        return mux
        // http.ListenAndServe(":8080", mux)
}
