package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"

	carPkg "github.com/Irdis01/Intern/Number4/pkg/command/car"
	"github.com/Irdis01/Intern/Number4/pkg/command/driver"
)

const (
	rightCycleTest            = "no errors test"
	cancelTest                = "cancel command test"
	wrongStopTest             = "error in srop, cause has speed"
	stopTestError             = "car is moving"
	withoutStartAccTest       = "without start command accelerate test"
	withoutStartTestAccError  = "car can't moving"
	withoutStartSlowTest      = "without start command slow test"
	withoutStartTestSlowError = "speed is zero or car can't moving"
	withoutStartStopTest      = "without start command stop test"
	withoutStartStopTestError = "speed is zero or car can't moving"
)

func TestCommand(t *testing.T) {

	resultRightCycle := []int{10, 20, 10, 0}
	resultCancelTest := 10

	t.Run(rightCycleTest, func(t *testing.T) {
		testDriver := driver.NewDriver()
		testCar := carPkg.NewCar()
		testDriver.StoreCommand(NewStartCommand(testCar))
		testDriver.StoreCommand(NewAccelerateCommand(testCar))
		testDriver.StoreCommand(NewSlowDownCommand(testCar))
		testDriver.StoreCommand(NewStopCommand(testCar))
		for i := 0; i < len(resultRightCycle); i++ {
			err := testDriver.Execute()
			if assert.NoError(t, err, "unexpected error") {
				assert.Equal(t, resultRightCycle[i], testCar.GetSpeed())
			}
		}
	})

	t.Run(cancelTest, func(t *testing.T) {
		testDriver := driver.NewDriver()
		testCar := carPkg.NewCar()
		testDriver.StoreCommand(NewStartCommand(testCar))
		testDriver.StoreCommand(NewAccelerateCommand(testCar))
		testDriver.Execute()
		testDriver.Execute()
		testDriver.Cancel()
		assert.Equal(t, testCar.GetSpeed(), resultCancelTest)
	})

	t.Run(wrongStopTest, func(t *testing.T) {
		testDriver := driver.NewDriver()
		testCar := carPkg.NewCar()
		testDriver.StoreCommand(NewStartCommand(testCar))
		testDriver.StoreCommand(NewAccelerateCommand(testCar))
		testDriver.StoreCommand(NewStopCommand(testCar))
		testDriver.Execute()
		testDriver.Execute()
		err := testDriver.Execute()
		assert.Equal(t, stopTestError, err.Error())
	})

	t.Run(withoutStartAccTest, func(t *testing.T) {
		testDriver := driver.NewDriver()
		testCar := carPkg.NewCar()
		testDriver.StoreCommand(NewAccelerateCommand(testCar))
		testDriver.StoreCommand(NewSlowDownCommand(testCar))
		err := testDriver.Execute()
		assert.Equal(t, withoutStartTestAccError, err.Error())
		err = testDriver.Execute()
		assert.Equal(t, withoutStartTestAccError, err.Error())
	})

	t.Run(withoutStartSlowTest, func(t *testing.T) {
		testDriver := driver.NewDriver()
		testCar := carPkg.NewCar()
		testDriver.StoreCommand(NewSlowDownCommand(testCar))
		err := testDriver.Execute()
		assert.Equal(t, withoutStartTestSlowError, err.Error())
	})

	t.Run(withoutStartStopTest, func(t *testing.T) {
		testDriver := driver.NewDriver()
		testCar := carPkg.NewCar()
		testDriver.StoreCommand(NewSlowDownCommand(testCar))
		err := testDriver.Execute()
		assert.Equal(t, withoutStartStopTestError, err.Error())
	})
}
