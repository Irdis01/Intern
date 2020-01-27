package driver

import "errors"

type command interface {
	Execute() error
	Cancel() error
}

// Driver интерфейс водителя
type Driver interface {
	StoreCommand(newCommand command)
	Execute() (err error)
	Cancel() (err error)
}

type driver struct {
	driveCommands  []command // очередь команд
	currentCommand int
}

func (d *driver) StoreCommand(newCommand command) {
	if len(d.driveCommands) == 0 {
		d.driveCommands = make([]command, 1)
		d.driveCommands[0] = newCommand
	} else {
		d.driveCommands = append(d.driveCommands, newCommand)
	}
}

func (d *driver) Execute() (err error) {
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

func (d *driver) Cancel() (err error) {
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
func NewDriver() Driver {
	return &driver{
		currentCommand: 0,
	}
}
