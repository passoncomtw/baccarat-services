package services

import (
	"fmt"
	"time"
)

func countDownTimer(num int) {
	timeTicker := time.NewTicker(1 * time.Second)

	for {
		<-timeTicker.C
		fmt.Println("Sec = ", num)
		num--
		if num == 0 {
			timeTicker.Stop()
			break
		}
	}
}

func LoopPokerRound() {
	var redCardIndex int
	var bankerResult PlayerResultDto
	var playerResult PlayerResultDto
	var nextSourceCards []CardItemDto

	fmt.Printf("initial cards\n")
	sourceCards := GetSourceCards()
	fmt.Printf("Change boot start\n")
	countDownTimer(1)
	redCardIndex, nextSourceCards = InitialRoundSourceCards(sourceCards)
	fmt.Printf("\nChange boot end\n")
	fmt.Printf("Game Start redCardIndex: %d\n", redCardIndex)
	for {
		fmt.Printf("Bet Start\n")
		countDownTimer(5)
		fmt.Printf("Bet End\n")
		countDownTimer(5)
		fmt.Printf("DealerDeal\n")
		bankerResult, playerResult, nextSourceCards = GetSingleGameResult(nextSourceCards)
		fmt.Printf("bankerResult: %v\n", bankerResult)
		fmt.Printf("playerResult: %v\n", playerResult)
		countDownTimer(1)
		fmt.Printf("payout\n")
		countDownTimer(5)

		if len(nextSourceCards) < redCardIndex {
			redCardIndex, nextSourceCards = InitialRoundSourceCards(sourceCards)
			fmt.Printf("restart New Round\n")
		} else {
			fmt.Printf("redCardIndex: %d\n", redCardIndex)
			fmt.Printf("nextSourceCardsLen: %d\n", len(nextSourceCards))
			fmt.Printf("continue Round\n")
		}
	}
}
