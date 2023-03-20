package connection

import (
	"Final_Project/constant"
	"Final_Project/model"
	"fmt"
	"log"
	"strings"
)

func HandleIO(currentConn *model.WebSocketConnection, connections map[string]*model.WebSocketConnection) {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		log.Println("ERROR", fmt.Sprintf("%v", r))
	// 	}
	// }()

	connections[currentConn.Username] = currentConn
	broadcastMessage(connections, currentConn, constant.MESSAGE_NEW_USER, "")

	for {
		payload := model.SocketPayload{}
		err := currentConn.ReadJSON(&payload)
		if err != nil {
			if strings.Contains(err.Error(), "websocket: close") {
				broadcastMessage(connections, currentConn, constant.MESSAGE_LEAVE, "")
				ejectConnection(connections, currentConn)
				return
			}

			log.Println("ERROR", err.Error())
			continue
		}

		broadcastMessage(connections, currentConn, constant.MESSAGE_CHAT, payload.Message)
	}
}

func broadcastMessage(connections map[string]*model.WebSocketConnection, currentConn *model.WebSocketConnection, kind, message string) {
	for _, eachConn := range connections {

		if eachConn == currentConn {
			continue
		}

		eachConn.WriteJSON(model.SocketResponse{
			From:    fmt.Sprintf(currentConn.Username + " Age: " + currentConn.Age),
			Type:    kind,
			Message: message,
		})
	}
}

func ejectConnection(connections map[string]*model.WebSocketConnection, currentConn *model.WebSocketConnection) {
	// var newConn []*model.WebSocketConnection
	// for _, conn := range connections {
	// 	if conn != currentConn {
	// 		newConn = append(newConn, conn)
	// 	}
	// }
	// connections = newConn

	delete(connections, currentConn.Username)
}
