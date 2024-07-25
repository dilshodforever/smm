package service

import (
	"context"
	"fmt"
	"log"

	pb "github.com/dilshodforever/4-oyimtixon-game-service/genprotos/game"
	s "github.com/dilshodforever/4-oyimtixon-game-service/storage"
)

type GameService struct {
	stg s.InitRoot
	pb.UnimplementedGameServiceServer
}

func NewGameService(stg s.InitRoot) *GameService {
	return &GameService{stg: stg}
}

func (s *GameService) GetLevels(ctx context.Context, req *pb.GetLevelsRequest) (*pb.GetLevelsResponse, error) {
	resp, err := s.stg.Game().GetLevels(req)
	if err != nil {
		log.Print(err)
	}
	return resp, err
}

func (s *GameService) StartLevel(ctx context.Context, req *pb.StartLevelRequest) (*pb.StartLevelResponse, error) {
	resp, err := s.stg.Game().StartLevel(req)
	fmt.Println(11111111111)
	if err != nil {
		log.Print(err)
	}
	return resp, err
}

func (s *GameService) CompleteLevel(ctx context.Context, req *pb.CompleteLevelRequest) (*pb.CompleteLevelResponse, error) {
	resp, err := s.stg.Game().CompleteLevel(req)
	if err != nil {
		log.Print(err)
	}
	return resp, err
}

func (s *GameService) GetChallenge(ctx context.Context, req *pb.GetChallengeRequest) (*pb.Level, error) {
	resp, err := s.stg.Game().GetChallenge(req)
	if err != nil {
		log.Print(err)
	}
	return resp, err
}

func (s *GameService) SubmitChallenge(ctx context.Context, req *pb.SubmitChallengeRequest) (*pb.SubmitChallengeResponse, error) {
	resp, err := s.stg.Game().SubmitChallenge(req)
	if err != nil {
		log.Print(err)
	}
	return resp, err
}

func (s *GameService) GetLeaderboard(ctx context.Context, req *pb.GetLeaderboardRequest) (*pb.LeaderboardResponse, error) {
	resp, err := s.stg.Game().GetLeaderboard(req)
	if err != nil {
		log.Print(err)
	}
	return resp, err
}

func (s *GameService) GetAchievements(ctx context.Context, req *pb.GetAchievementsRequest) (*pb.AchievementsResponse, error) {
	resp, err := s.stg.Game().GetAchievements(req)
	if err != nil {
		log.Print(err)
	}
	return resp, err
}
