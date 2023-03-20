package main

import (
	"io/ioutil"
	"net/http"

	"Final_Project/connection"
	"Final_Project/constant"
	"Final_Project/model"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

// type M map[string]interface{}

var connections = make(map[string]*model.WebSocketConnection)

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		content, err := ioutil.ReadFile("index.html")
// 		if err != nil {
// 			http.Error(w, "Could not open requested file", http.StatusInternalServerError)
// 			return
// 		}

// 		fmt.Fprintf(w, "%s", content)
// 	})

// 	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
// 		currentGorillaConn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
// 		if err != nil {
// 			http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
// 		}

// 		username := r.URL.Query().Get("username")
// 		age := r.URL.Query().Get("age")
// 		currentConn := model.WebSocketConnection{Conn: currentGorillaConn, Username: username, Age: age}
// 		connections = append(connections, &currentConn)

// 		go connection.HandleIO(&currentConn, connections)
// 	})

// 	fmt.Println("Server starting at :8080")
// 	http.ListenAndServe(":8080", nil)
// }

func main() {
	e := echo.New()

	e.GET("/", func(ctx echo.Context) error {
		content, err := ioutil.ReadFile(constant.INDEX_PAGE)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, "could not open html")
		}

		return ctx.HTML(http.StatusOK, string(content))
	})
	e.Static("/template", "template")

	e.Any("/ws", func(ctx echo.Context) error {
		upgrader := websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		}
		currentGorillaConn, err := upgrader.Upgrade(ctx.Response().Writer, ctx.Request(), nil)
		if err != nil {
			return ctx.String(http.StatusBadRequest, "Could not open websocket connection")
		}

		username := ctx.Request().URL.Query().Get("username")
		age := ctx.Request().URL.Query().Get("age")

		go connection.HandleIO(&model.WebSocketConnection{
			Conn:     currentGorillaConn,
			Username: username,
			Age:      age,
		}, connections)
		// connections = append(connections, &currentConn)

		// connection.HandleIO(&currentConn, connections)
		return nil
	})

	e.Start(":8080")
}
