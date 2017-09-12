package main

import (
	"fmt"
	"errors"
)


func main() {
	po := new(PurchaseOrder2)
	po.Value = 42.27

	go SavePOPromise(po, true).Then(func(obj interface{}) error {
		po := obj.(*PurchaseOrder2)
		fmt.Printf("Purchase order saved with ID:  %d\n", po.Number)
		return nil
	}, func(err error) {
		fmt.Printf("Failed to save Purchases Order: " + err.Error() + "\n")
	})

	fmt.Scanln()
}

type PurchaseOrder2 struct {
	Number int
	Value float64
}

type Promise struct {
	successChannel chan interface{}
	failureChannel chan error
}

func SavePOPromise(po *PurchaseOrder2, shouldFail bool) *Promise {
	result := new(Promise)

	result.successChannel = make(chan interface{}, 1)
	result.failureChannel = make(chan error, 1)

	go func() {
		if shouldFail {
			result.failureChannel <- errors.New("Ooops")
		} else {
			po.Number = 123
			result.successChannel <- po
		}
	}()

	return result
}


func (this *Promise) Then(success func(interface{}) error, failure func(error)) *Promise {
	result := new(Promise)

	result.successChannel = make(chan interface{}, 1)
	result.failureChannel = make(chan error, 1)

	go func() {
		select {
			case obj := <- this.successChannel:
				newErr := success(obj)
				if newErr == nil {
					result.successChannel <- obj
				} else {
					result.failureChannel <- newErr
				}
			case err := <- this.failureChannel:
				failure(err)
				result.failureChannel <- err
		}
	}()

	return result
}