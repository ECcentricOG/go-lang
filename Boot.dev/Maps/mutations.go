package main

import (
	"errors"
	"fmt"
)

func deleteIfNecessary(users map[string]user1, name string) (deleted bool, err error) {
    usr, ok := users[name]    //checks if exist
    if !ok {
        return false, errors.New("not found")
        // umesh = users["Umesh"]   to get user
    }
    if usr.scheduledForDeletion {
        delete(users, name)     // delete from map
        return false, nil
    }
    return true, nil
}

type user1 struct {
	name                 string
	number               int
	scheduledForDeletion bool
}

func main() {
    usr := make(map[string]user1)
    usr["Umesh"] = user1{name: "Umesh", number: 7767039576, scheduledForDeletion: false} // add to map
    usr["Levi"] = user1{name: "Levi", number: 36257635, scheduledForDeletion: true} 
    usr["Erwin"] = user1{name: "Erwin", number:12345678, scheduledForDeletion: false}
    fmt.Println(deleteIfNecessary(usr, "Umesh"))
    fmt.Println(deleteIfNecessary(usr, "Levi"))
    fmt.Println(deleteIfNecessary(usr, "Erwin"))
    fmt.Println(deleteIfNecessary(usr, "ECcentricOG"))
}
