package console

import (
	"bufio"
	"context"
	"fmt"
	"go.uber.org/zap"
	"log"
	"os"
	"strings"

	"github.com/kulinsky/guess_word/usecase"
)

type CommandHandler struct {
	ucHandler usecase.Handler
	logger    *zap.Logger
}

func NewCommandHandler(i usecase.Handler, logger *zap.Logger) CommandHandler {
	return CommandHandler{
		ucHandler: i,
		logger:    logger,
	}
}

func (cH *CommandHandler) InitWords(ctx context.Context, words []string) {
	for _, val := range words {
		if _, err := cH.ucHandler.WordCreate(ctx, val); err != nil {
			log.Fatal(err.Error())
		}
	}
}

func (cH *CommandHandler) Start(ctx context.Context, attemptCount int) {
	w, err := cH.ucHandler.GetRandomWord(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	gameID, err := cH.ucHandler.StartNewGame(ctx, attemptCount, w)
	if err != nil {
		log.Fatal(err.Error())
	}

	reader := bufio.NewReader(os.Stdin)

	counter := 1

	for {
		fmt.Print("Guess a letter: ")

		letter, _ := reader.ReadString('\n')

		letter = strings.Trim(letter, "\n")

		guess, err := cH.ucHandler.Guess(ctx, gameID, letter)

		if err != nil {
			panic(err)
		}

		if !guess {
			fmt.Println(fmt.Sprintf("Missed, mistake %d out of %d.", counter, attemptCount))

			counter++
		} else {
			fmt.Println("Hit!")
		}

		stat, err := cH.ucHandler.GetGameStat(ctx, gameID)

		if err != nil {
			panic(err)
		}

		fmt.Println("The word:", stat.CurrentWord)

		if stat.Win {
			fmt.Println("You won!")

			break
		}

		if stat.AttemptCount == 0 {
			fmt.Println("You lost!")

			break
		}
	}
}
