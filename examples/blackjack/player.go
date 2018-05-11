package main

type Player interface {
	Name() string
}

func NewHumanPlayer(name string) Player {
	return nil
}

func NewComputerPlayer(name string) Player {
	return nil
}
