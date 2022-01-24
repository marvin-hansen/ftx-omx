// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package main

const action = 1

func main() {
	switch action {

	case 1:
		generateAuthKey(24)
	case 2:
		printPlaceOrderJsonReq()
	case 3:
		printPlaceBookOrderJsonReq()
	case 4:
		printPlaceTriggerJsonReq()
	case 5:
		printPlaceTrailingJsonReq()

	default:
		println("No action selected")
	}
}
