package handler 

import (
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"encoding/json"
	"net/http"
	"testing"
)

func TestFuncHandlerHello(t *testing.T) {
	// 1º Configuração o Gin para o Modo de test
	gin.SetMode(gin.TestMode)

	// 2° Criar um ResponseRecorder (para capturar a resposta)
	w := httptest.NewRecorder()

	// 3º Criar o Contexto do Gin
	// O gin.New() cria um router. Engine é um Handler
	router := gin.New()

	// 4º Registrar a rota
	// Você registra o handler no seu router (Engine)
	router.GET("/", HandlerHello)

	// 5º Cria a requisição HTTP
	req, _ := http.NewRequest(http.MethodGet, "/", nil)

	// 6º Executar a requisição 
	// O método correto no router (Engine) é ServerHTTP.
	router.ServeHTTP(w, req)

	// --- 7 Asserções (verificações) ----

	// a) Verificar o Statis Code
	if w.Code != http.StatusOK {
		t.Errorf("Esperava status %d, mas recebeu %d", http.StatusOK, w.Code)
	}

	// b) Verifica o conteúdo do corpo (opcional)
	var response map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Erro ao decodificar JSON de resposta: %v", err)
	}

	expectedMessage := "Hello, world!!"
	if response["mensagem"] != expectedMessage {
		t.Errorf("Esperava mensagem '%s', mas recebeu '%s'", expectedMessage, response["mensagem"])
	}
}
