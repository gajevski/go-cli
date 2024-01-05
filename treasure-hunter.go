// NOTE: 62 CHARS ASCII

package main

import (
	"fmt"

	"github.com/liamg/gobless"
)

const (
	maxBoardX = 10
	minBoardX = -10
	maxBoardY = 10
	minBoardY = -10
)

var (
	gui           *gobless.GUI
	controlBox    *gobless.TextBox
	controlText   string
	gameText      string
	userPositionX int
	userPositionY int
)

func main() {
	startGame()
}

func startGame() {
	gui = gobless.NewGUI()
	if err := gui.Init(); err != nil {
		panic(err)
	}
	defer gui.Close()

	controlText = ""
	gameText = ""
	userPositionX = 0
	userPositionY = 0

	refreshGUI()

	closeGame()
	moveNorth()
	moveSouth()
	moveWest()
	moveEast()

	gui.HandleResize(func(event gobless.ResizeEvent) {
		refreshGUI()
	})

	gui.Loop()
}

func refreshGUI() {
	rows := []gobless.Component{
		gobless.NewRow(
			gobless.GridSizeThreeQuarters,
			gobless.NewColumn(
				gobless.GridSizeTwoThirds,
				renderGameBox(gameText),
			),
			gobless.NewColumn(
				gobless.GridSizeOneThird,
				gobless.NewRow(
					gobless.GridSizeFull,
					gobless.NewColumn(
						gobless.GridSizeFull,
						renderInformationBox(),
						renderItemsBox(userPositionX, userPositionY),
					),
				),
			),
		),
		gobless.NewRow(
			gobless.GridSizeOneQuarter,
			gobless.NewColumn(
				gobless.GridSizeOneSixth,
				renderControlBox(controlText),
			),
		),
	}

	gui.Render(rows...)
}

func moveNorth() {
	gui.HandleKeyPress(gobless.KeyUp, func(event gobless.KeyPressEvent) {
		controlText = `
		 _  _    ___    ___   _____   _  _ 
		| \| |  / _ \  | _ \ |_   _| | || |
		| .  | | (_) | |   /   | |   | __ |
		|_|\_|  \___/  |_|_\   |_|   |_||_|
	   `
		gameText = `
		 ..............................................................
		 ..............................................................
		 .............................=*++**+..........................
		 .............................@#+=+#%:.........................
		 ........................:-=#%%#***%%%#*-:.....................
		 .........................:+*#%%%%%%%#*+:......................
		 .............................+@@@@@+:.........................
		 ............................:-%####*:.........................
		 ..........................-===-:-**##*-.......................
		 ........................:*+-::-+**+*::+-......................
		 ......................:==**==**#+=*@+--+=.....................
		 ......................=++%#**%%+**#@#+-=+:....................
		 ....................:+==#@%###%***%@@%*-=+=...................
		 ...................:==*#%#%#%%#*##*%@=##*-+*-.................
		 ...................-+**%*-%@##**###@@:=%##+==.................
		 ...................:=+#=-#%%*##*%#%@%:.:=#=---................
		 ..................-++*-.+@#*+++=+#%##*==:.=+==-...............
		 ..................+%%#-.+#*++%*--=+#@@%*+..+@#*-..............
		 .................-+#@@*=###%%@@%##**@@#%#..+@%*=:.............
		 .................=+**+#%%#*+*@@%#@%#@@#%+..:+**+-.............
		 ...................::=#%#+####%..+%@#%@@#=..:+*-..............
		 ......................**=-=*#@:....%@@@%#+....................
		 ....................:***=+*%#:.....-*%%%#*=:..................
		 ..................:=++++*#%%=.......:%@%%#*+:.................
		 .................-+=-:-**#%=.........+@@###**:................
		 ................=%%%@%++%@*...........+@%%#+==:...............
		 ................=*=+#@%#*:...........:*%%#*+=+=...............
		 ..............:=++*%@@*-..............:%@#**#%=.-++:..........
		 .............-#%##%%@=..................:=%@@%##*++#-.........
		 .............:*@@@%%#-....................=@@@%%%%%#:.........
		 .......:.......+@@@%%%*:...................+@%%%%%=..-........
		 ...--+*#%#==+*+*@@@@@%@+=---------=--------+@@@@*--+#=*%#*++-.
		 ...:=-====+++====================++=====++=================--:
		 .......:...:......::.:..:::.:..:::..:...::........::.:...:....
	`
		if userPositionY < maxBoardY {
			userPositionY++
		}
		refreshGUI()
	})
}

func moveSouth() {
	gui.HandleKeyPress(gobless.KeyDown, func(event gobless.KeyPressEvent) {
		controlText = `
		 ___    ___    _   _   _____   _  _ 
		/ __|  / _ \  | | | | |_   _| | || |
		\__ \ | (_) | | |_| |   | |   | __ |
		|___/  \___/   \___/    |_|   |_||_|
	   `
		gameText = `
	 	..............................................................
	 	..............................................................
	 	.............................:................................
	 	...........................-=--:-++...........................
	 	...........................+-===++*...........................
	 	........................-===--====+++=:.......................
	 	.......................-+##**====+*##*-.......................
	 	...........................-*::::++...........................
	 	............................=****+............................
	 	.........................:=+++=-*+++-:........................
	 	........................=:=+==+++-++:=:.......................
	 	.......................:-:+*-.=:..**-:=.......................
	 	.......................=::**:.*...+*=:+.......................
	 	......................:=-=*+..:..:=#+---......................
	 	.....................:===+*+-::::-=#*==+-.....................
	 	......................==+*#+-:----=#**==:.....................
	 	.....................:=-**#***+*******=-=.....................
	 	.....................==+*+**++++++*#**+=+.....................
	 	.....................++**#*=-::::--+*-**+-....................
	 	.....................==++*=--+**+--+*-+=+:....................
	 	.....................+#=+*=-:=*++::+*=*#+:....................
	 	......................-=+*=-:+=:*-:=+:........................
	 	......................::-+-:-*:.*-:=+-........................
	 	........................:=-:-*..+=-=+=........................
	 	........................-=--++..-*==++........................
	 	........................=++++=...*=+++........................
	 	........................+=--*:...*=-++:.......................
	 	.......................:+=-=*....*===+-.......................
	 	.......................=+=++*....+*+=+*.......................
	 	.......................:++=#+....-#++*=.......................
	 	.......................-****=....:#***+.......................
	 	......................+++**#-.....#**++*:.....................
	 	.............::.::-:.:------------------::.::..::.............
	 	..............................................................
	  `
		if userPositionY > minBoardY {
			userPositionY--
		}
		refreshGUI()
	})
}

func moveEast() {
	gui.HandleKeyPress(gobless.KeyRight, func(event gobless.KeyPressEvent) {
		controlText = `
		 ___     _     ___   _____ 
		| __|   /_\   / __| |_   _|
		| _|   / _ \  \__ \   | |  
		|___| /_/ \_\ |___/   |_|  
		`
		gameText = `
		 ..............................................................
		 ................................:--::.........................
		 ...............................+@#*=--:.......................
		 ............................:::%@@%##+=:......................
		 ..........................:#@@@@@@@@%#+=++=:..................
		 ...........................::-*@@%%+-=++-::...................
		 ...........................:-=+*#=#%+--.......................
		 .........................:*+--=*+=*+--:.......................
		 ......................:-=+=:::-=##%*-.......:.................
		 ....................:+***=++++*#=##*.......:=-:--:............
		 ...................:=##**=##%@%+-##*......::::=*=:............
		 ...................:+*#+:+@@****###+....:==-=*+:..............
		 .....................:*#*=-##**##%##=::-+##*+=:...............
		 .......................:-+#+-=#**@*+=++**++#*-................
		 .....................--==#%*==%##@##*+--=+==:.................
		 ....................=*##%@%%@#%##-::..........................
		 ..................:=**#@@@@%%%====::..........................
		 ..................:+#%@@@%#*+==-:::==--::.....................
		 ....................::-#@@@##**++=+=+--+-.....................
		 .....................=%%#%@@@#%#%%#%+*#::--:..................
		 ....................-#%##%%@*:.::*%###++*+=-..................
		 ....................:#@%#%%%:.....:--:=#+--++:................
		 .............:=**+++++=*%%*-..........:+#+--:--:..............
		 ............:+%+=+******+:............:=%#**+=+=..............
		 ..........:-+*##**#%*==-:..............-#@##%%*-:.....::......
		 .......:-=*##%@%#%#=:...................-**#@%*++=:::--+-:....
		 .....:=#*+#%%%=:.::........................:=#%#***#*+*%-.....
		 ......:+%%#%#-...............................:*%%**%%%#-......
		 ........-+%%%*-:..:......::......:...:........-#%%@*-::.......
		 .....:=::-#@%#%*:-=-::---==-.....-++=:--:-..::::*#-...-=+:....
		 ..:=-=*#*##**+**++*#**#*###**%##*++**##*#%%#**####%%#++====--:
		 ...:::-::-::-=------:::--:-=-:-:-----=:--::-:::-===-:---:::::.
		 ..............................................................
		 ..............................................................
		`
		if userPositionX < maxBoardX {
			userPositionX++
		}
		refreshGUI()
	})
}

func moveWest() {
	gui.HandleKeyPress(gobless.KeyLeft, func(event gobless.KeyPressEvent) {
		controlText = `
		__      __  ___   ___   _____ 
		\ \    / / | __| / __| |_   _|
		 \ \/\/ /  | _|  \__ \   | |  
		  \_/\_/   |___| |___/   |_|  
		`
		gameText = `
		 ..............................................................
		 .........................::--:................................
		 .......................:--=+#%+...............................
		 ......................:=+*#%@@%:::............................
		 ..................:=++=+#%@@@@@@@@%:..........................
		 ...................::-++=-+%%@@*-::...........................
		 .......................--=##+#*++-:...........................
		 .......................::-+*=+*==-+*-.........................
		 .................:.......-+%##+-:::=+==:......................
		 ............:--:-=-.......*##=#*++++=***+:....................
		 ............:=*=::::......*##-=#@@##=**##=:...................
		 ..............:+*=-==:....=###****@@+:+#*+:...................
		 ...............:=+*##+-::=##%##**##-=*#*:.....................
		 ................-*#*+**++=+*%**#=-+#*-:.......................
		 .................:=++=--+*##@##%==*%#==--.....................
		 ..........................::-##@#@%%@%##*=....................
		 ..........................::====%%%@@@@#**=:..................
		 .....................::--==:::-==+*#%@@@%#+:..................
		 .....................-=--+=+=++**##@@@%-::....................
		 ..................:--::#*+##%%#%#@@@%#%%=.....................
		 ..................-=+*++###%#-:..+@%%##%#-....................
		 ................:+*--+#=:--:.....:%#%#%@%:....................
		 ..............:--:--+#+:..........:*%%*=+++++**=:.............
		 ..............=+=+**#%=:............:+#*****+=+%+:............
		 .......:.....:-*%%##%#-..............:-==*%#**##*+-:..........
		 ....:-+--:::-++*%@#**-...................:=#%#%@%##*=-:.......
		 .....-%*+*#***#%#=:........................::.:=#%%#+*#=:.....
		 ......-*%%%**#%*:...............................-*%#%%*:......
		 ........:-*@%%#-........:...:......::......:..:-+%%%+-:.......
		 ....:+=-...-#*::::..----:-++-.....:==---::--=:+%#%@#-::-:.....
		 :--====++*%%####**#%%#**#**++*##%**###*#*+#*++**+***#*#*=-=:..
		 .:::::---:-===-:::-::--:=-----:-:-=-:--:::------=-::-::-:::...
		 ..............................................................
		 ..............................................................
		 `
		if userPositionX > minBoardX {
			userPositionX--
		}
		refreshGUI()
	})
}

func closeGame() {
	gui.HandleKeyPress(gobless.KeyCtrlC, func(event gobless.KeyPressEvent) {
		gui.Close()
	})
}

func renderGameBox(text string) *gobless.TextBox {
	gameBox := gobless.NewTextBox()
	gameBox.SetBorderColor(gobless.ColorGreen)
	gameBox.SetTitle("Game")
	gameBox.SetText(text)
	gameBox.SetTextWrap(true)
	return gameBox
}

func renderControlBox(text string) *gobless.TextBox {
	controlBox = gobless.NewTextBox()
	controlBox.SetText(text)
	controlBox.SetBorderColor(gobless.ColorDarkRed)
	return controlBox
}

func renderItemsBox(boardX int, boardY int) *gobless.TextBox {
	itemsBox := gobless.NewTextBox()
	itemsBox.SetTitle("Items")
	formattedText := fmt.Sprintf("X: %d, Y: %d", boardX, boardY)
	itemsBox.SetText(formattedText)
	itemsBox.SetBorderColor(gobless.Color100)
	return itemsBox
}

func renderInformationBox() *gobless.TextBox {
	informationBox := gobless.NewTextBox()
	informationBox.SetTitle("Information")
	informationBox.SetText(`Objective: Find the hidden treasure.
    
    Controls:
        
    "Key ↑" to move North
    "Key ↓" to move South
    "Key ←" to move West
    "Key →" to move East
    
    Navigate wisely and watch for clues. Your adventure begins now! Good luck!`)
	return informationBox
}
