package storage

import (
	pb "github.com/dilshodforever/4-oyimtixon-game-service/genprotos/game"
)

type InitRoot interface {
	Game() GameStorage
}

type GameStorage interface {
	GetLevels(req *pb.GetLevelsRequest) (*pb.GetLevelsResponse, error)
	StartLevel(req *pb.StartLevelRequest) (*pb.StartLevelResponse, error)
	CompleteLevel(req *pb.CompleteLevelRequest) (*pb.CompleteLevelResponse, error)
	GetChallenge(req *pb.GetChallengeRequest) (*pb.Level, error)
	SubmitChallenge(req *pb.SubmitChallengeRequest) (*pb.SubmitChallengeResponse, error)
	GetLeaderboard(req *pb.GetLeaderboardRequest) (*pb.LeaderboardResponse, error)
	GetAchievements(req *pb.GetAchievementsRequest) (*pb.AchievementsResponse, error)
}
