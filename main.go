package main

import (
	"fmt"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"log"
	"net/http"
	"time"
)

func main() {
	var funGaming = true
	var waiting = true
	for waiting {
		registerAccountRequest := RegisterPlayerJSONRequestBody{
			EmojiAlias: ":bullettrain_front:",
			Name:       "dithomas",
			Password:   "bananenpannekoek",
		}

		request, err := NewRegisterPlayerRequest(server, registerAccountRequest)
		if err != nil {
			log.Fatalln(err)
		}

		response, err := client.Do(request)
		if err != nil {
			log.Fatalln(err)
		}
		if response.StatusCode == 200 {
			break
		}
		if response.StatusCode == 201 {
			break
		}
		if response.StatusCode == 202 {
			break
		}
		if response.StatusCode == 203 {
			break
		}
	}
	for funGaming == true {

		var registration = register()
		var playerId = registration.JSON200.Id

		var currentGame = getCurrentGame(*playerId).JSON200
		var phase = *currentGame.GamePhase
		var gameId = currentGame.GameId

		for phase == GameStateDTOGamePhaseFinished {
			fmt.Println("IDLE sleep for 5 seconds")
			time.Sleep(5 * time.Second)
			var currentGame = getCurrentGame(*playerId).JSON200
			if *currentGame.GamePhase != GameStateDTOGamePhaseFinished {
				break
			}
		}
		for phase != GameStateDTOGamePhaseFinished {
			var currentGame = getCurrentGame(*playerId).JSON200

			var nextState = nextMove(*playerId, *gameId, *currentGame)
			fmt.Println(nextState)
			//resp, err := nextState.JSON200

			phase := *getCurrentGame(*playerId).JSON200.GamePhase
			if phase == GameStateDTOGamePhaseFinished {
				break
			}
			time.Sleep(50 * time.Millisecond)
		}
		time.Sleep(5 * time.Second)
	}
}

const server = "https://raspberry-runaround-yjiibwucma-ez.a.run.app/raspberry-runaround/"

var previousMove = "none"

var client = http.Client{}

func getCurrentGame(playerId openapi_types.UUID) GetCurrentGameStateResponse {
	request, err := NewGetCurrentGameStateRequest(server, playerId)
	request.Header.Set("Authorization", "bananenpannekoek")
	if err != nil {
		log.Fatalln(err)
	}
	response, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	result, err := ParseGetCurrentGameStateResponse(response)
	if err != nil {
		log.Fatalln(err)
	}
	return *result
}

func nextMove(playerId openapi_types.UUID, gameId openapi_types.UUID, gameState GameStateDTO) MoveResponse {
	if previousMove == "Right" {
		return rightward(playerId, gameId, gameState)
	}
	if previousMove == "Up" {
		return upward(playerId, gameId, gameState)

	}
	if previousMove == "Left" {
		return leftward(playerId, gameId, gameState)

	}
	return downward(playerId, gameId, gameState)

}
func downward(playerId openapi_types.UUID, gameId openapi_types.UUID, gameState GameStateDTO) MoveResponse {
	var currentWalls = *gameState.Walls
	fmt.Println(currentWalls)
	if !containsMove(currentWalls, "Right") && !lastMove(currentWalls, "Right") {
		resp := moveBitch(playerId, gameId, MoveRequestDirectionRight)
		previousMove = "Right"
		return resp
	} else if !containsMove(currentWalls, "Down") && !lastMove(currentWalls, "Down") {
		resp := moveBitch(playerId, gameId, MoveRequestDirectionDown)
		previousMove = "Down"
		return resp
	} else if !containsMove(currentWalls, "Left") && !lastMove(currentWalls, "Left") {
		resp := moveBitch(playerId, gameId, MoveRequestDirectionLeft)
		previousMove = "Left"
		return resp
	} else {
		resp := moveBitch(playerId, gameId, MoveRequestDirectionUp)
		previousMove = "Up"
		return resp
	}
}

func upward(playerId openapi_types.UUID, gameId openapi_types.UUID, gameState GameStateDTO) MoveResponse {
	var currentWalls = *gameState.Walls
	fmt.Println(currentWalls)
	if !containsMove(currentWalls, "Left") && !lastMove(currentWalls, "Left") {
		resp := moveBitch(playerId, gameId, MoveRequestDirectionLeft)
		previousMove = "Left"
		return resp
	} else if !containsMove(currentWalls, "Up") && !lastMove(currentWalls, "Up") {
		resp := moveBitch(playerId, gameId, MoveRequestDirectionUp)
		previousMove = "Up"
		return resp
	} else if !containsMove(currentWalls, "Right") && !lastMove(currentWalls, "Right") {
		resp := moveBitch(playerId, gameId, MoveRequestDirectionRight)
		previousMove = "Right"
		return resp
	} else {
		resp := moveBitch(playerId, gameId, MoveRequestDirectionDown)
		previousMove = "Down"
		return resp
	}
}

func leftward(playerId openapi_types.UUID, gameId openapi_types.UUID, gameState GameStateDTO) MoveResponse {
	var currentWalls = *gameState.Walls
	fmt.Println(currentWalls)
	if !containsMove(currentWalls, "Down") && !lastMove(currentWalls, "Down") {
		resp := moveBitch(playerId, gameId, MoveRequestDirectionDown)
		previousMove = "Down"
		return resp
	} else if !containsMove(currentWalls, "Left") && !lastMove(currentWalls, "Left") {
		resp := moveBitch(playerId, gameId, MoveRequestDirectionLeft)
		previousMove = "Left"
		return resp
	} else if !containsMove(currentWalls, "Up") && !lastMove(currentWalls, "Up") {
		resp := moveBitch(playerId, gameId, MoveRequestDirectionUp)
		previousMove = "Up"
		return resp
	} else {
		resp := moveBitch(playerId, gameId, MoveRequestDirectionRight)
		previousMove = "Right"
		return resp
	}
}
func rightward(playerId openapi_types.UUID, gameId openapi_types.UUID, gameState GameStateDTO) MoveResponse {
	var currentWalls = *gameState.Walls
	fmt.Println(currentWalls)
	if !containsMove(currentWalls, "Up") && !lastMove(currentWalls, "Up") {
		resp := moveBitch(playerId, gameId, MoveRequestDirectionUp)
		previousMove = "Up"
		return resp
	} else if !containsMove(currentWalls, "Right") && !lastMove(currentWalls, "Right") {
		resp := moveBitch(playerId, gameId, MoveRequestDirectionRight)
		previousMove = "Right"
		return resp
	} else if !containsMove(currentWalls, "Down") && !lastMove(currentWalls, "Down") {
		resp := moveBitch(playerId, gameId, MoveRequestDirectionDown)
		previousMove = "Down"
		return resp
	} else {
		resp := moveBitch(playerId, gameId, MoveRequestDirectionUp)
		previousMove = "Up"
		return resp
	}
}
func lastMove(walls []GameStateDTOWalls, expected GameStateDTOWalls) bool {
	if len(walls) == 3 {
		return false
	}
	if lastMoveIsOpposite(expected) {
		return true
	}
	return false
}

func lastMoveIsOpposite(expected GameStateDTOWalls) bool {
	fmt.Println(previousMove)
	if expected == "Left" {

		return previousMove == "Right"
	}
	if expected == "Up" {
		return previousMove == "Down"
	}
	if expected == "Down" {
		return previousMove == "UP"
	} else {
		return previousMove == "Left"
	}
}
func containsMove(walls []GameStateDTOWalls, expected GameStateDTOWalls) bool {
	for i := 0; i < len(walls); i++ {
		if walls[i] == expected {
			return true
		}
		// perform an operation
	}
	return false
}
func moveBitch(playerId openapi_types.UUID, gameId openapi_types.UUID, moveDirection MoveRequestDirection) MoveResponse {
	fmt.Println("THE BITCH MOVED: " + moveDirection)
	var requestBody = MoveJSONRequestBody{
		Direction: moveDirection,
		GameId:    gameId,
		PlayerId:  playerId,
	}
	request, err := NewMoveRequest(server, requestBody)
	request.Header.Set("Authorization", "bananenpannekoek")
	if err != nil {
		log.Fatalln(err)
	}
	response, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	result, err := ParseMoveResponse(response)
	if err != nil {
		log.Fatalln(err)
	}
	return *result
}
func register() RegisterPlayerResponse {
	registerAccountRequest := RegisterPlayerJSONRequestBody{
		EmojiAlias: ":bullettrain_front:",
		Name:       "dithomas",
		Password:   "bananenpannekoek",
	}

	request, err := NewRegisterPlayerRequest(server, registerAccountRequest)
	if err != nil {
		log.Fatalln(err)
	}

	response, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	result, err := ParseRegisterPlayerResponse(response)
	if err != nil {
		log.Fatalln(err)
	}
	return *result

}
