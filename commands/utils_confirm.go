package commands

import (
    "fmt"
    "log"
)

func askForConfirmation(question string, counts ...int) bool {
    cnt := 0
    if len(counts) > 0 {
        cnt = counts[0]
    }

    fmt.Println(question)
    var response string
    _, err := fmt.Scanln(&response)
    if err != nil {
        log.Fatal(err)
    }
    okResponses := []string{"y", "Y", "yes", "Yes", "YES"}
    ngResponses := []string{"n", "N", "no", "No", "NO"}
    if containsString(okResponses, response) {
        return true
    } else if containsString(ngResponses, response) {
        return false
    } else if cnt >= 2 {
        fmt.Println("Ummmmmm...")
        return false
    } else {
        cnt++
        return askForConfirmation("Please type yes or no and then press enter:", cnt)
    }
}

func posString(slice []string, element string) int {
    for index, elem := range slice {
        if elem == element {
            return index
        }
    }
    return -1
}

func containsString(slice []string, element string) bool {
    return (posString(slice, element) > -1)
}
