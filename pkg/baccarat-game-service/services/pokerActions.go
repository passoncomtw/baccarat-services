package services

import (
	"baccarat-services/m/pkg/baccarat-game-service/models"
	"fmt"
	"math/rand"
)

type Suit int

const (
	Spade Suit = iota
	Heart
	Diamond
	Club
)

type CardItemDto struct {
	value int
	point int
	suit  Suit
}

type DrawCardResult struct {
	card            CardItemDto
	nextSourceCards []CardItemDto
}

type PlayerResultDto struct {
	value int
	cards []CardItemDto
}

var validateDrawBankerMapping = [10][10]bool{
	{true, true, true, true, true, true, true, true, true, true},
	{true, true, true, true, true, true, true, true, true, true},
	{true, true, true, true, true, true, true, true, true, true},
	{true, true, true, true, true, true, true, true, false, true},
	{false, false, true, true, true, true, true, true, false, false},
	{false, false, false, false, false, false, true, true, false, false},
	{false, false, false, false, false, false, true, true, false, false},
	{false, false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false, false},
}

func getMin(x int, y int) int {
	if x > y {
		return y
	}
	return x
}

func GetSourceCards() []CardItemDto {
	sourceCards := []CardItemDto{
		{value: 1, point: 1, suit: Spade}, {value: 2, point: 2, suit: Spade}, {value: 3, point: 3, suit: Spade}, {value: 4, point: 4, suit: Spade}, {value: 5, point: 5, suit: Spade}, {value: 6, point: 6, suit: Spade}, {value: 7, point: 7, suit: Spade}, {value: 8, point: 8, suit: Spade}, {value: 9, point: 9, suit: Spade}, {value: 0, point: 10, suit: Spade}, {value: 0, point: 11, suit: Spade}, {value: 0, point: 12, suit: Spade}, {value: 0, point: 13, suit: Spade},
		{value: 1, point: 1, suit: Heart}, {value: 2, point: 2, suit: Heart}, {value: 3, point: 3, suit: Heart}, {value: 4, point: 4, suit: Heart}, {value: 5, point: 5, suit: Heart}, {value: 6, point: 6, suit: Heart}, {value: 7, point: 7, suit: Heart}, {value: 8, point: 8, suit: Heart}, {value: 9, point: 9, suit: Heart}, {value: 0, point: 10, suit: Heart}, {value: 0, point: 11, suit: Heart}, {value: 0, point: 12, suit: Heart}, {value: 0, point: 13, suit: Heart},
		{value: 1, point: 1, suit: Diamond}, {value: 2, point: 2, suit: Diamond}, {value: 3, point: 3, suit: Diamond}, {value: 4, point: 4, suit: Diamond}, {value: 5, point: 5, suit: Diamond}, {value: 6, point: 6, suit: Diamond}, {value: 7, point: 7, suit: Diamond}, {value: 8, point: 8, suit: Diamond}, {value: 9, point: 9, suit: Diamond}, {value: 0, point: 10, suit: Diamond}, {value: 0, point: 11, suit: Diamond}, {value: 0, point: 12, suit: Diamond}, {value: 0, point: 13, suit: Diamond},
		{value: 1, point: 1, suit: Club}, {value: 2, point: 2, suit: Club}, {value: 3, point: 3, suit: Club}, {value: 4, point: 4, suit: Club}, {value: 5, point: 5, suit: Club}, {value: 6, point: 6, suit: Club}, {value: 7, point: 7, suit: Club}, {value: 8, point: 8, suit: Club}, {value: 9, point: 9, suit: Club}, {value: 0, point: 10, suit: Club}, {value: 0, point: 11, suit: Club}, {value: 0, point: 12, suit: Club}, {value: 0, point: 13, suit: Club},
		{value: 1, point: 1, suit: Spade}, {value: 2, point: 2, suit: Spade}, {value: 3, point: 3, suit: Spade}, {value: 4, point: 4, suit: Spade}, {value: 5, point: 5, suit: Spade}, {value: 6, point: 6, suit: Spade}, {value: 7, point: 7, suit: Spade}, {value: 8, point: 8, suit: Spade}, {value: 9, point: 9, suit: Spade}, {value: 0, point: 10, suit: Spade}, {value: 0, point: 11, suit: Spade}, {value: 0, point: 12, suit: Spade}, {value: 0, point: 13, suit: Spade},
		{value: 1, point: 1, suit: Heart}, {value: 2, point: 2, suit: Heart}, {value: 3, point: 3, suit: Heart}, {value: 4, point: 4, suit: Heart}, {value: 5, point: 5, suit: Heart}, {value: 6, point: 6, suit: Heart}, {value: 7, point: 7, suit: Heart}, {value: 8, point: 8, suit: Heart}, {value: 9, point: 9, suit: Heart}, {value: 0, point: 10, suit: Heart}, {value: 0, point: 11, suit: Heart}, {value: 0, point: 12, suit: Heart}, {value: 0, point: 13, suit: Heart},
		{value: 1, point: 1, suit: Diamond}, {value: 2, point: 2, suit: Diamond}, {value: 3, point: 3, suit: Diamond}, {value: 4, point: 4, suit: Diamond}, {value: 5, point: 5, suit: Diamond}, {value: 6, point: 6, suit: Diamond}, {value: 7, point: 7, suit: Diamond}, {value: 8, point: 8, suit: Diamond}, {value: 9, point: 9, suit: Diamond}, {value: 0, point: 10, suit: Diamond}, {value: 0, point: 11, suit: Diamond}, {value: 0, point: 12, suit: Diamond}, {value: 0, point: 13, suit: Diamond},
		{value: 1, point: 1, suit: Club}, {value: 2, point: 2, suit: Club}, {value: 3, point: 3, suit: Club}, {value: 4, point: 4, suit: Club}, {value: 5, point: 5, suit: Club}, {value: 6, point: 6, suit: Club}, {value: 7, point: 7, suit: Club}, {value: 8, point: 8, suit: Club}, {value: 9, point: 9, suit: Club}, {value: 0, point: 10, suit: Club}, {value: 0, point: 11, suit: Club}, {value: 0, point: 12, suit: Club}, {value: 0, point: 13, suit: Club},
		{value: 1, point: 1, suit: Spade}, {value: 2, point: 2, suit: Spade}, {value: 3, point: 3, suit: Spade}, {value: 4, point: 4, suit: Spade}, {value: 5, point: 5, suit: Spade}, {value: 6, point: 6, suit: Spade}, {value: 7, point: 7, suit: Spade}, {value: 8, point: 8, suit: Spade}, {value: 9, point: 9, suit: Spade}, {value: 0, point: 10, suit: Spade}, {value: 0, point: 11, suit: Spade}, {value: 0, point: 12, suit: Spade}, {value: 0, point: 13, suit: Spade},
		{value: 1, point: 1, suit: Heart}, {value: 2, point: 2, suit: Heart}, {value: 3, point: 3, suit: Heart}, {value: 4, point: 4, suit: Heart}, {value: 5, point: 5, suit: Heart}, {value: 6, point: 6, suit: Heart}, {value: 7, point: 7, suit: Heart}, {value: 8, point: 8, suit: Heart}, {value: 9, point: 9, suit: Heart}, {value: 0, point: 10, suit: Heart}, {value: 0, point: 11, suit: Heart}, {value: 0, point: 12, suit: Heart}, {value: 0, point: 13, suit: Heart},
		{value: 1, point: 1, suit: Diamond}, {value: 2, point: 2, suit: Diamond}, {value: 3, point: 3, suit: Diamond}, {value: 4, point: 4, suit: Diamond}, {value: 5, point: 5, suit: Diamond}, {value: 6, point: 6, suit: Diamond}, {value: 7, point: 7, suit: Diamond}, {value: 8, point: 8, suit: Diamond}, {value: 9, point: 9, suit: Diamond}, {value: 0, point: 10, suit: Diamond}, {value: 0, point: 11, suit: Diamond}, {value: 0, point: 12, suit: Diamond}, {value: 0, point: 13, suit: Diamond},
		{value: 1, point: 1, suit: Club}, {value: 2, point: 2, suit: Club}, {value: 3, point: 3, suit: Club}, {value: 4, point: 4, suit: Club}, {value: 5, point: 5, suit: Club}, {value: 6, point: 6, suit: Club}, {value: 7, point: 7, suit: Club}, {value: 8, point: 8, suit: Club}, {value: 9, point: 9, suit: Club}, {value: 0, point: 10, suit: Club}, {value: 0, point: 11, suit: Club}, {value: 0, point: 12, suit: Club}, {value: 0, point: 13, suit: Club},
		{value: 1, point: 1, suit: Spade}, {value: 2, point: 2, suit: Spade}, {value: 3, point: 3, suit: Spade}, {value: 4, point: 4, suit: Spade}, {value: 5, point: 5, suit: Spade}, {value: 6, point: 6, suit: Spade}, {value: 7, point: 7, suit: Spade}, {value: 8, point: 8, suit: Spade}, {value: 9, point: 9, suit: Spade}, {value: 0, point: 10, suit: Spade}, {value: 0, point: 11, suit: Spade}, {value: 0, point: 12, suit: Spade}, {value: 0, point: 13, suit: Spade},
		{value: 1, point: 1, suit: Heart}, {value: 2, point: 2, suit: Heart}, {value: 3, point: 3, suit: Heart}, {value: 4, point: 4, suit: Heart}, {value: 5, point: 5, suit: Heart}, {value: 6, point: 6, suit: Heart}, {value: 7, point: 7, suit: Heart}, {value: 8, point: 8, suit: Heart}, {value: 9, point: 9, suit: Heart}, {value: 0, point: 10, suit: Heart}, {value: 0, point: 11, suit: Heart}, {value: 0, point: 12, suit: Heart}, {value: 0, point: 13, suit: Heart},
		{value: 1, point: 1, suit: Diamond}, {value: 2, point: 2, suit: Diamond}, {value: 3, point: 3, suit: Diamond}, {value: 4, point: 4, suit: Diamond}, {value: 5, point: 5, suit: Diamond}, {value: 6, point: 6, suit: Diamond}, {value: 7, point: 7, suit: Diamond}, {value: 8, point: 8, suit: Diamond}, {value: 9, point: 9, suit: Diamond}, {value: 0, point: 10, suit: Diamond}, {value: 0, point: 11, suit: Diamond}, {value: 0, point: 12, suit: Diamond}, {value: 0, point: 13, suit: Diamond},
		{value: 1, point: 1, suit: Club}, {value: 2, point: 2, suit: Club}, {value: 3, point: 3, suit: Club}, {value: 4, point: 4, suit: Club}, {value: 5, point: 5, suit: Club}, {value: 6, point: 6, suit: Club}, {value: 7, point: 7, suit: Club}, {value: 8, point: 8, suit: Club}, {value: 9, point: 9, suit: Club}, {value: 0, point: 10, suit: Club}, {value: 0, point: 11, suit: Club}, {value: 0, point: 12, suit: Club}, {value: 0, point: 13, suit: Club},
	}
	return sourceCards
}

func DrawCard(sourceCards []CardItemDto) (CardItemDto, []CardItemDto) {
	fmt.Printf("sourceCardsLen: %d\n", len(sourceCards))
	index := rand.Intn(len(sourceCards) - 1)
	card := sourceCards[index]
	nextSourceCards := append(sourceCards[:index], sourceCards[index+1:]...)
	fmt.Printf("nextSourceCardsLen: %d\n", len(nextSourceCards))
	return card, nextSourceCards
}

func InitialRoundSourceCards(sourceCards []CardItemDto) (int, []CardItemDto) {
	/** 牌局開始時會隨機抽掉幾張卡片 */
	var afterThrowCard []CardItemDto
	card, nextSourceCards := DrawCard(sourceCards)
	throwCardCount := getMin(card.value, 9)
	fmt.Printf("throwCardCount: %d\n", throwCardCount)
	for i := 0; i < throwCardCount; i++ {
		_, nextSourceCards = DrawCard(nextSourceCards)
	}

	afterThrowCard = nextSourceCards
	redCardIndex := rand.Intn(len(sourceCards))

	var newGame models.CreateGameRequestDto
	newGame.RedCardIndex = redCardIndex

	result := models.CreateGame(newGame)
	fmt.Printf("create game result: %v", result)

	return redCardIndex, afterThrowCard
}

func GetSingleGameResult(sourceCards []CardItemDto) (PlayerResultDto, PlayerResultDto, []CardItemDto) {
	var bankerCards []CardItemDto
	var playerCards []CardItemDto
	var nextSourceCards []CardItemDto
	var card CardItemDto
	var bankerResult PlayerResultDto
	var playerResult PlayerResultDto
	bankerValue := 0
	playerValue := 0
	card, nextSourceCards = DrawCard(sourceCards)

	bankerCards = append(bankerCards, card)
	bankerValue = (bankerValue + card.value) % 10

	card, nextSourceCards = DrawCard(nextSourceCards)
	bankerCards = append(bankerCards, card)
	bankerValue = (bankerValue + card.value) % 10

	card, nextSourceCards = DrawCard(nextSourceCards)
	playerCards = append(playerCards, card)
	playerValue = (playerValue + card.value) % 10

	card, nextSourceCards = DrawCard(nextSourceCards)
	playerCards = append(playerCards, card)
	playerValue = (playerValue + card.value) % 10
	if bankerValue >= 8 || playerValue >= 8 {
		if playerValue <= 5 {
			card, nextSourceCards = DrawCard(nextSourceCards)
			playerCards = append(playerCards, card)
			playerValue = playerValue + card.value

			shouldDrawToBanker := validateDrawBankerMapping[bankerValue][card.value]
			if shouldDrawToBanker {
				card, nextSourceCards = DrawCard(nextSourceCards)
				bankerCards = append(bankerCards, card)
				bankerValue = bankerValue + card.value
			}
		} else {
			if bankerValue <= 5 {
				card, nextSourceCards = DrawCard(nextSourceCards)
				bankerCards = append(bankerCards, card)
				bankerValue = bankerValue + card.value
			}
		}
	}

	bankerResult.cards = bankerCards
	bankerResult.value = bankerValue

	playerResult.cards = playerCards
	playerResult.value = playerValue

	return bankerResult, playerResult, nextSourceCards
}
