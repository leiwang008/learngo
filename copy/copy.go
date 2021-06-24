package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/leiwang008/utils"
)

func main() {
	debugmsg := utils.Debugmsg(true)

	fmt.Printf(debugmsg+"parameters %v\n", os.Args)

	source := os.Args[1]
	dest := os.Args[2]
	utils.Verbose, _ = strconv.ParseBool(os.Args[3])
	checkDate, _ := strconv.ParseBool(os.Args[4])
	copy, _ := strconv.ParseBool(os.Args[5])

	utils.HandleFiles(source, dest, checkDate, copy)

}
