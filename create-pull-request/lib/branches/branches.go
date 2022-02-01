package branches

import (
	"github.com/go-git/go-git/v5"
	"log"
	"fmt"
)

func NewBranch(branchName string) {

	directory := "/branching"
	url := "github.com/garethjevans/create-pr"

	r, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL: url,
	})
	if err != nil {
		log.Fatalln(err)
	}



	headRef, err := r.Head()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(headRef)
}
