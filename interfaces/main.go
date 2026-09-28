package main

import "fmt"

type Animal interface {
    Sound() string
}

type Dog struct{}
type Cat struct{}
type Lion struct{}

func (Dog) Sound() string { return "Woof" }
func (Cat) Sound() string { return "Meow" }

func makeSound(a Animal) {
    fmt.Println(a.Sound())
}

// type assertions

type expense interface {
	cost() float64
}

func getExpenseReport(e expense) (string, float64) {
	

	// if em, ok:= e.(email); ok{
	// 	fmt.Println(em.toAddress, em.cost())
	// 	return em.toAddress, em.cost()
	// }

	// if sm, ok:= e.(sms); ok{
	// 	fmt.Println(sm.toPhoneNumber, sm.cost())
	// 	return  sm.toPhoneNumber, sm.cost()
	// }
	// return "", 0.0

	switch v:= e.(type){
		case email:
			fmt.Println(v.toAddress, v.cost())
			return v.toAddress, v.cost()
		case sms:
			fmt.Println(v.toPhoneNumber, v.cost())
			return  v.toPhoneNumber, v.cost()
		default:
			fmt.Println("", 0.0)
			return "", 0.0
	}
	
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func (em email) cost() float64 {
	if !em.isSubscribed {
		return float64(len(em.body)) * .05
	}
	return float64(len(em.body)) * .01
}

func (sm sms) cost() float64 {
	if !sm.isSubscribed {
		return float64(len(sm.body)) * .1
	}
	return float64(len(sm.body)) * .03
}

func (inv invalid) cost() float64 {
	return 0.0
}

//message formatter
type formatter interface{
	format() string
}

type plainText struct{
	message string
} 

func (p plainText) format() string{
	fmt.Println(p.message)
	return p.message
}
type bold struct{
	message string
} 

func (b bold) format() string{
	fmt.Println(`**` + b.message + `**`)
	return `**` + b.message + `**`
}

type code struct{
	message string
} 

func (c code) format() string{
	fmt.Println("`" + c.message + "`")
	return "`" + c.message + "`"
}

func sendMessage(format formatter) string {
	return format.format() // Adjusted to call Format without an argument
}


//process notification

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

func (dm directMessage) importance () int{
	if dm.isUrgent{
		return 50
	}

	return dm.priorityLevel
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

func(gm groupMessage) importance() int{
	return  gm.priorityLevel
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (sa systemAlert) importance() int {
	return  100
}


func processNotification(n notification) (string, int) {
	switch v:= n.(type) {
	case directMessage:
		return v.senderUsername, v.importance()	
	case groupMessage:
		return v.groupName, v.importance()	
	case systemAlert:
		return v.alertCode, v.importance()
	default:
		return "", 0
	}
}


func main(){

makeSound(Dog{}) 
makeSound(Cat{})
// makeSound(Lion{}) -> gives error as it doesnt have the Sound methods

//type assertions
getExpenseReport(email{
	isSubscribed: true,
	body: "Hello",
	toAddress: "oluwa",
})

//message formatter
sendMessage(plainText{
	message: "Oluwafunmbi",
})
}