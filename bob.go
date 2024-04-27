// Package bob https://golang.org/doc/effective_go.html#commentary
package bob

//Importing strings
import "strings"

func checkQuestion(remark string) bool {
	return strings.ContainsAny("?", remark)
}

func checkCaps(remark string) bool {
	return remark == strings.ToUpper(remark)
}

func checkEmpty(remark string) bool {
	return remark == ""
}

// Hey - takes a string and responds like a typical teenager.
func Hey(remark string) string {
	var comeback string
	//Check if it is a question first
	if checkEmpty(remark) {
		comeback = "Fine. Be that way!"
	} else { // not empty string
		if checkCaps(remark) { //if in caps
			if checkQuestion(remark) { //if a question AND in CAPS
				comeback = "Calm down, I know what I'm doing!"
			} else { //in caps, but not a question
				comeback = "Whoa, chill out!"
			}
		} else if checkQuestion(remark) { //not in caps, but a question
			comeback = "Sure."
		} else { //not a question, but not empty
			comeback = "Whatever."
		}
	}
	return comeback
}
