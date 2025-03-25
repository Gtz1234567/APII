package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"
    "strings"
)

func intParaRomano(num int) string {
    valores := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
    simbolos := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

    var resultado strings.Builder
    for i := 0; i < len(valores); i++ {
        for num >= valores[i] {
            resultado.WriteString(simbolos[i])
            num -= valores[i]
        }
    }
    return resultado.String()
}

func ConverterHandler(w http.ResponseWriter, r *http.Request) {
    numeroStr := r.URL.Query().Get("numero")
    if numeroStr == "" {
        http.Error(w, "O parâmetro 'numero' é obrigatório", http.StatusBadRequest)
        return
    }

    numero, err := strconv.Atoi(numeroStr)
    if err != nil || numero <= 0 || numero > 3999 {
        http.Error(w, "Número inválido. Digite um número entre 1 e 3999.", http.StatusBadRequest)
        return
    }

    resposta := struct {
        Numero int    `json:"numero"`
        Romano string `json:"romano"`
    }{
        Numero: numero,
        Romano: intParaRomano(numero),
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resposta)
}

func handler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/converter" {
        ConverterHandler(w, r)
    } else {
        http.Error(w, "Rota não encontrada", http.StatusNotFound)
    }
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("API rodando em http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}