package kafka

import (
	"context"
	"encoding/json"
	"log"

	pb "github.com/dilshodforever/4-oyimtixon-game-service/genprotos/game"
	"github.com/dilshodforever/4-oyimtixon-game-service/service"
)

func StartLevel(rootService *service.GameService) func(message []byte) {
	return func(message []byte) {
		var app pb.StartLevelRequest
		if err := json.Unmarshal(message, &app); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		appEval, err := rootService.StartLevel(context.Background(), &app)
		if err != nil {
			log.Printf("Cannot create evaluation via Kafka: %v", err)
			return
		}
		log.Printf("Created evaluation: %+v", appEval)
	}
}

// func EvaluationUpdateHandler(evalService *service.EvaluationService) func(message []byte) {
// 	return func(message []byte) {
// 		var eval pb.EvaluationCreate
// 		if err := json.Unmarshal(message, &eval); err != nil {
// 			log.Printf("Cannot unmarshal JSON: %v", err)
// 			return
// 		}

// 		respEval, err := evalService.(context.Background(), &eval)
// 		if err != nil {
// 			log.Printf("Cannot create evaluation via Kafka: %v", err)
// 			return
// 		}
// 		log.Printf("Created evaluation: %+v", respEval)
// 	}
// }
