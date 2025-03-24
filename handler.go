package handler

import (
    "encoding/json"
    "net/http"
    "strconv"
    "strings"
)

// Converte um número inteiro para número romano
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

// Estrutura da resposta JSON
type Resposta struct {
    Numero int    `json:"numero"`
    Romano string `json:"romano"`
}

// Manipulador da API
func converterHandler(w http.ResponseWriter, r *http.Request) {
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

    resposta := Resposta{Numero: numero, Romano: intParaRomano(numero)}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resposta)
}

// Função exportada para a Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/converter" {
        converterHandler(w, r)
    } else {
        http.NotFound(w, r)
    }
}