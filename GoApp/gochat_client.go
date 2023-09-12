package main

func main() {
	myTurn := false
	for {
		if myTurn {
			goSender()
		} else {
			goReceiver()
		}
		myTurn = !myTurn
	}

}
