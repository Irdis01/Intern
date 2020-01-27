package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"

	carPkg "github.com/Irdis01/Intern/Number4/pkg/command/car"
	"github.com/Irdis01/Intern/Number4/pkg/command/driver"
)

const (
	commandSuccessTest        = "no errors test"
	commandErrorTest          = "command error test"
	cancelTest                = "cancel command test"
	wrongStopTest             = "error in stop, cause has speed"
	stopTestError             = "car is moving"
	withoutStartAccTest       = "without start command accelerate test"
	withoutStartTestAccError  = "car can't moving"
	withoutStartSlowTest      = "without start command slow test"
	withoutStartTestSlowError = "speed is zero or car can't moving"
	withoutStartStopTest      = "without start command stop test"
	withoutStartStopTestError = "speed is zero or car can't moving"
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
		testDriver := driver.NewDriver()
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

func makeErrorCommandTest() (testDriver [4]driver.Driver) {

	/// wrongStopTest
	testDriver[0] = driver.NewDriver()
	testCar := carPkg.NewCar()
	testDriver[0].StoreCommand(NewStartCommand(testCar))
	testDriver[0].StoreCommand(NewAccelerateCommand(testCar))
	testDriver[0].StoreCommand(NewStopCommand(testCar))
	testDriver[0].Execute()
	testDriver[0].Execute()

	/// withoutStartAccTest
	testDriver[1] = driver.NewDriver()
	testCar = carPkg.NewCar()
	testDriver[1].StoreCommand(NewAccelerateCommand(testCar))

	/// withoutStartSlowTest
	testDriver[2] = driver.NewDriver()
	testCar = carPkg.NewCar()
	testDriver[2].StoreCommand(NewSlowDownCommand(testCar))

	/// withoutStartStopTest
	testDriver[3] = driver.NewDriver()
	testCar = carPkg.NewCar()
	testDriver[3].StoreCommand(NewSlowDownCommand(testCar))

	return
}
