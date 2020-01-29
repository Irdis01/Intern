package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"

	carPkg "github.com/Irdis01/Intern/Number4/pkg/command/car"
	"github.com/Irdis01/Intern/Number4/pkg/command/palyer"
)

const (
	commandSuccessTest        = "no errors test"
	commandErrorTest          = "command error test"
)

func TestCommand(t *testing.T) {

	expectSuccessTest := []int{10, 20, 10, 0}
	expectErrorTest := []string{
		"car is moving",
		"car can't moving",
		"speed is zero or car can't moving",
		"speed is zero or car can't moving",
	}

	preparedDriverForTest := makeErrorCommandTest()

	t.Run(commandSuccessTest, func(t *testing.T) {
		testDriver := palyer.NewDriver()
		testCar := carPkg.NewCar()
		testDriver.StoreCommand(NewStartCommand(testCar))
		testDriver.StoreCommand(NewAccelerateCommand(testCar))
		testDriver.StoreCommand(NewSlowDownCommand(testCar))
		testDriver.StoreCommand(NewStopCommand(testCar))
		for i := 0; i < len(expectSuccessTest); i++ {
			err := testDriver.Execute()
			if assert.NoError(t, err, "unexpected error") {
				assert.Equal(t, expectSuccessTest[i], testCar.GetSpeed())
			}
		}
	})

	t.Run(commandErrorTest, func(t *testing.T) {
		for i := 0; i < len(preparedDriverForTest); i++ {
			err := preparedDriverForTest[i].Execute()
			assert.Equal(t, expectErrorTest[i], err.Error())
		}
	})
}

func makeErrorCommandTest() (testDriver [4]palyer.Player) {

	/// wrongStopTest
	testDriver[0] = palyer.NewDriver()
	testCar := carPkg.NewCar()
	testDriver[0].StoreCommand(NewStartCommand(testCar))
	testDriver[0].StoreCommand(NewAccelerateCommand(testCar))
	testDriver[0].StoreCommand(NewStopCommand(testCar))
	testDriver[0].Execute()
	testDriver[0].Execute()

	/// withoutStartAccTest
	testDriver[1] = palyer.NewDriver()
	testCar = carPkg.NewCar()
	testDriver[1].StoreCommand(NewAccelerateCommand(testCar))

	/// withoutStartSlowTest
	testDriver[2] = palyer.NewDriver()
	testCar = carPkg.NewCar()
	testDriver[2].StoreCommand(NewSlowDownCommand(testCar))

	/// withoutStartStopTest
	testDriver[3] = palyer.NewDriver()
	testCar = carPkg.NewCar()
	testDriver[3].StoreCommand(NewSlowDownCommand(testCar))

	return
}
