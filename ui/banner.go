package ui

import "fmt"

const HelloBanner = `
╔════════════════════════════════════╗
║         MPV PLAYER CONTROL         ║
║           by MrPhoenix             ║
╚════════════════════════════════════╝

Hello! What would you like to do?

[1] View saved list  
[2] Add a new link  
[3] Play audio  
[4] Play video  
[5] Launch Cava  
[6] Launch EasyEffects  
[7] Settings  
[0] Exit

`
const InputStr = "> "

func ShowMainMenu() {
	fmt.Print(HelloBanner)
	//fmt.Print(InputStr)
}
