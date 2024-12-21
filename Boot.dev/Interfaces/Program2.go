package main

import "fmt"

type notification interface {
    importance() int
}

type directMessage struct {
    senderUsername string
    messageContent string
    priorityLevel int
    isUrgent bool
}

type groupMessage struct {
    groupName string
    messageContent string
    priorityLevel int
}

type systemAlert struct {
    alertCode string 
    messageContent string
}

func(dm directMessage) importance() int {
    if dm.isUrgent {
        return 50
    }
    return dm.priorityLevel
}

func(gm groupMessage) importance() int {
    return gm.priorityLevel
}

func(sa systemAlert) importance() int {
    return 100
}

func processNotification(n notification) (string, int) {
    switch ntfn := n.(type) {
    case directMessage:
        return ntfn.senderUsername, ntfn.importance()
    case groupMessage:
        return ntfn.groupName, ntfn.importance()
    case systemAlert:
        return ntfn.alertCode, ntfn.importance()
    default:
        return "", 0
    }
}

func main() {
    dm1 := directMessage{messageContent: "Hello", senderUsername: "Me", priorityLevel: 30, isUrgent: false}
    dm2 := directMessage{messageContent: "Call Me", senderUsername: "Umesh", priorityLevel: 30, isUrgent: true}
    gm := groupMessage{groupName: "Useless",messageContent: "On ?", priorityLevel: 40}
    sa := systemAlert{alertCode: "108" , messageContent: "Fire"}
    fmt.Println(processNotification(dm1))
    fmt.Println(processNotification(dm2))
    fmt.Println(processNotification(gm))
    fmt.Println(processNotification(sa))
}
