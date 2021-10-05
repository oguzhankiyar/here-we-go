package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		Password: "",
		DB: 0,
	})
	if err := client.Ping(ctx).Err(); err != nil {
		panic(err)
	}

	err := SaveUser(ctx, client, &User{
		Username: "gopher",
		Points: 100,
		Rank: 1,
	})
	if err != nil {
		panic(err)
	}

	u, err := GetUser(ctx, client, "gopher")
	if err != nil {
		panic(err)
	}

	fmt.Println(u)

	l, err := GetLeaderboard(ctx, client)
	if err != nil {
		panic(err)
	}

	fmt.Println(l)
}

var leaderboardKey = "leaderboard"

type Leaderboard struct {
	Count int `json:"count"`
	Users []*User
}

func GetLeaderboard(ctx context.Context, client *redis.Client) (*Leaderboard, error) {
	scores := client.ZRangeWithScores(ctx, leaderboardKey, 0, -1)
	if scores == nil {
		return nil, nil
	}
	count := len(scores.Val())
	users := make([]*User, count)
	for idx, member := range scores.Val() {
		users[idx] = &User{
			Username: member.Member.(string),
			Points: int(member.Score),
			Rank: idx,
		}
	}
	leaderboard := &Leaderboard{
		Count: count,
		Users: users,
	}
	return leaderboard, nil
}

type User struct {
	Username string `json:"username" binding:"required"`
	Points   int `json:"points" binding:"required"`
	Rank     int    `json:"rank"`
}

func SaveUser(ctx context.Context, client *redis.Client, user *User) error {
	member := &redis.Z{
		Score: float64(user.Points),
		Member: user.Username,
	}
	pipe := client.TxPipeline()
	pipe.ZAdd(ctx, "leaderboard", member)
	rank := pipe.ZRank(ctx, leaderboardKey, user.Username)
	_, err := pipe.Exec(ctx)
	if err != nil {
		return err
	}
	fmt.Println(rank.Val(), err)
	user.Rank = int(rank.Val())
	return nil
}

func GetUser(ctx context.Context, client *redis.Client, username string) (*User, error) {
	pipe := client.TxPipeline()
	score := pipe.ZScore(ctx, leaderboardKey, username)
	rank := pipe.ZRank(ctx, leaderboardKey, username)
	_, err := pipe.Exec(ctx)
	if err != nil {
		return nil, err
	}
	if score == nil {
		return nil, nil
	}
	return &User{
		Username: username,
		Points: int(score.Val()),
		Rank: int(rank.Val()),
	}, nil
}
