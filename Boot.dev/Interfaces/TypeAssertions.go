package main

import "fmt"

type expense1 interface {
    cost() float64
}

type email1 struct {
    isSubscribed bool
    body string
    toAddress string
}

type sms struct {
    isSubscribed bool
    body string
    toPhoneNumber string
}

type invalid struct {}

func (e email1) cost() float64 {
	if !e.isSubscribed{
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}

func getExpenseReport(e expense1) (string, float64) {
//    eml, ok := e.(email1)
//    if !ok {
//        sm, gg := e.(sms)
//        if !gg {
//            return "", 0.0
//        }
//        return sm.toPhoneNumber, sm.cost()
//    }
//    return eml.toAddress, eml.cost()
    switch exp := e.(type) {
        case email1 :
            return exp.toAddress, exp.cost()
        case sms :
            return exp.toPhoneNumber, exp.cost()
        default :
            return "", .0
    }
}

func main() {
    e := email1{isSubscribed: true, body: "Whoa there!", toAddress: "soldier@monty.com"}
    s := sms{isSubscribed: false, body: "Halt! Who goes there?", toPhoneNumber: "+155555509832"}
    i := invalid{}
    fmt.Println(getExpenseReport(e))
    fmt.Println(getExpenseReport(s))
    fmt.Println(getExpenseReport(i))
}
