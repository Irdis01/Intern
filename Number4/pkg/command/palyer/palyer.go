package palyer

import "errors"

type command interface {
	Execute() error
	Cancel() error
}

// Player интерфейс водителя
type Player interface {
	StoreCommand(newCommand command)
	Execute() (err error)
	Cancel() (err error)
}

type palyer struct {
	driveCommands  []command // очередь команд
	currentCommand int
}

func (d *palyer) StoreCommand(newCommand command) {
	if len(d.driveCommands) == 0 {
		d.driveCommands = make([]command, 1)
		d.driveCommands[0] = newCommand
	} else {
		d.driveCommands = append(d.driveCommands, newCommand)
	}
}

func (d *palyer) Execute() (err error) {
	if d.currentCommand < len(d.driveCommands) {
		err = d.driveCommands[d.currentCommand].Execute()
		if err == nil {
			d.currentCommand++
		}
		return
	} else {
		err = errors.New("all command are finished")
		return
	}
}

func (d *palyer) Cancel() (err error) {
	if d.currentCommand != 0 {
		err = d.driveCommands[d.currentCommand-1].Cancel()
		if err == nil {
			d.currentCommand--
		}
		return
	} else {
		err = errors.New("No one command was executed")
		return
	}
}

// NewDriver конструктор нового водителя
func NewDriver() Player {
	return &palyer{
		currentCommand: 0,
	}
}
