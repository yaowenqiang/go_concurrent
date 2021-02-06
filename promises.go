package main

import (
    "fmt"
    "errors"
    "time"
)


type Promise struct {
    successChannel chan interface{}
    failureChannel chan error
}


type PurchaseOrder struct {
    Number int
    Value float64
}

func SavePO(po *PurchaseOrder, shouldFail bool) *Promise {
    result := new(Promise)
    result.successChannel = make(chan interface{}, 1)
    result.failureChannel = make(chan error, 1)

    go func() {
        time.Sleep(2 * time.Second)
        if shouldFail {
            result.failureChannel <- errors.New("Faild to save purchase order!")
        } else {
            po.Number = 1234
            result.successChannel <- po
        }
    }()
    return result
}

func (this *Promise) Then(success func(interface{}) error, failure func(error)) *Promise {
    result := new(Promise)
    result.successChannel = make(chan interface{}, 1)
    result.failureChannel = make(chan error, 1)

    timeout := time.After(1 * time.Second)
    go func() {
        select {
            case obj := <-this.successChannel:
                newErr := success(obj)
                if newErr == nil {
                    result.successChannel <- obj
                } else {
                    result.failureChannel <- newErr
                }
            case err := <- this.failureChannel:
                failure(err)
                result.failureChannel <- err
            case <- timeout:
                failure(errors.New("Promise timed out"))
        }
    }()
    return result
}

func main() {
    po := new(PurchaseOrder)
    po.Value = 42.27

    SavePO(po, false).Then(func(obj interface{}) error{
        po := obj.(*PurchaseOrder)
        fmt.Printf("Purchase Order saved with ID: %d\n", po.Number)
        return nil
        //return errors.New("First promise failed")
    }, func(err error) {
        fmt.Printf("Failed to save Purchase order : " + err.Error() + "\n")
    }).Then(func(obj interface{}) error{
        fmt.Println("Second Promise success")
        return nil
    },func(err error){
        fmt.Println("Second promise failed: " + err.Error())
    })

    fmt.Scanln()
}
