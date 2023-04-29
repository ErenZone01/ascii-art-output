package main

import (
	"os"
)

func main() {
	if len(os.Args) >= 2 && len(os.Args) <= 4 {
		var arg3 string
		arg1 := os.Args[1]
		arg2 := DefineArg()

		if len(os.Args) == 4 || len(os.Args) == 3 {
			arg3 = os.Args[1]
			if IsFile(arg3) {
				arg1 = os.Args[2]
				arg2 = DefineArg()
			} else {
				arg3 = ""
			}
		}
		if CaseResolved(arg1, arg3) {
			return
		}
		texte := foundFile(arg2)
		asciiArt(arg1, arg3, texte)
	}
}
