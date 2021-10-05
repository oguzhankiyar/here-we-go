package main

import (
	"errors"
	"fmt"
	"time"
)

type Speaker struct {
	Id 		string
	SongId	string
}

type Song struct {
	Id string
}

type Command interface {
	Run() error
}

type PlayCommand struct {
	Song	*Song
	Speaker *Speaker
}

func (c PlayCommand) Run() error {
	if c.Song == nil || c.Speaker == nil {
		return errors.New("the command is invalid")
	}

	fmt.Println("Playing the song", c.Song.Id)

	c.Speaker.SongId = c.Song.Id

	return nil
}

type PauseCommand struct {
	 Speaker *Speaker
}

func (c PauseCommand) Run() error {
	if c.Speaker == nil {
		return errors.New("the command is invalid")
	}

	fmt.Println("Pausing the song")

	c.Speaker.SongId = ""

	return nil
}

type Commander struct {

}

func (p Commander) Execute(command Command) error {
	return command.Run()
}

func main() {
	song := &Song{"in_the_end"}
	speaker := &Speaker{"54fg7d", ""}
	commander := Commander{}

	playCommand := PlayCommand{song, speaker}

	err := commander.Execute(playCommand)
	if err != nil {
		fmt.Println("error:", err)
	}

	time.Sleep(time.Millisecond * 500)

	pauseCommand := PauseCommand{speaker}

	err = commander.Execute(pauseCommand)
	if err != nil {
		fmt.Println("error:", err)
	}
}