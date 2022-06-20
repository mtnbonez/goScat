package goscatgame

import (
	"fmt"
)

// ===================================================================
//
type GameManager struct {

	//List of games managed by the manager
	ActiveGames []*Game

	//List of games to be reported
	ReportGames []*Game

	//Which game are we currently processing?
	ActiveIter int
	ReportIter int
}

// ===================================================================
//
func InitManager(manager *GameManager) {
	manager.ActiveIter = 0
	manager.ReportIter = 0
	manager.ActiveGames = nil
	manager.ReportGames = nil
}

// ===================================================================
//
func (manager *GameManager) GetActiveAndReportingGames() bool {

	return len(manager.ActiveGames) > 0 || len(manager.ReportGames) > 0
}

// ===================================================================
//
func ProcessManager(manager *GameManager) {

	//Do we have any games?
	if len(manager.ActiveGames) < 1 {
		fmt.Printf("No active games!\n")
	} else {
		if !ProcessActiveGame(manager, manager.ActiveGames[manager.ActiveIter]) {
			fmt.Printf("Processing ACTIVE game_id: %s failed!\n", manager.ActiveGames[manager.ActiveIter].ID)
			return
		}

		//Move the ActiveIter for the next game
		if manager.ActiveIter < (len(manager.ActiveGames) - 1) {
			manager.ActiveIter++
		} else {
			manager.ActiveIter = 0
		}
	}

	if len(manager.ReportGames) < 1 {
		fmt.Printf("No reporting games!\n")
	} else {
		if !ProcessReportGame(manager, manager.ReportGames[manager.ReportIter]) {
			fmt.Printf("Process REPORT game_id: %s failed!\n", manager.ReportGames[manager.ReportIter].ID)
		}

		//Move the ReportIter for the next game
		if manager.ReportIter < (len(manager.ReportGames) - 1) {
			manager.ReportIter++
		} else {
			manager.ReportIter = 0
		}
	}

}

// ===================================================================
//
func ProcessActiveGame(manager *GameManager, currGame *Game) bool {

	if manager.ActiveIter > (len(manager.ActiveGames) - 1) {
		fmt.Printf("GameIter (%d) is higher than len-1 of ActiveGames (%d)\n", manager.ActiveIter, (len(manager.ActiveGames) - 1))
		return false
	}

	//Do step functions for each state

	switch currGame.GameState {
	case Playing:
		{
			// Display game board
			Display(currGame)

			// Process input
			Play(currGame)

			// Process game
			Process(currGame, false)

			return true
		}
	case Ending:
		{
			Process(currGame, true)
			// Move game to reporting state
			fmt.Printf("%s:%d to Report\n", currGame.ID.String(), currGame.GameState)
			currGame.GameState = Report
			manager.ReportGames = append(manager.ReportGames, currGame)
			manager.ActiveGames = append(manager.ActiveGames[:manager.ActiveIter], manager.ActiveGames[manager.ActiveIter+1:]...)
			return true
		}
	case Report:
		{
			// Shouldn't be any REPORT states in ActiveGames!
			fmt.Printf("%s is state=%d in ActiveGames! Bad!\n", currGame.ID.String(), currGame.GameState)
			return false
		}
	default:
		{
			fmt.Printf("Unknown state, %d!\n", currGame.GameState)
		}
	}

	fmt.Printf("Problem processing ACTIVE game!\n")
	return false

}

// ===================================================================
//
func ProcessReportGame(manager *GameManager, currGame *Game) bool {

	if manager.ReportIter > (len(manager.ReportGames) - 1) {
		fmt.Printf("ReportIter (%d) is higher than len-1 of ReportGames (%d)\n", manager.ReportIter, (len(manager.ReportGames) - 1))
		return false
	}

	switch currGame.GameState {
	case Report:
		{
			//Do logic for end of game here!
			fmt.Printf("%s has been reported!\n", currGame.ID.String())
			manager.ReportGames = append(manager.ReportGames[:manager.ReportIter], manager.ReportGames[manager.ReportIter+1:]...)
			return true
		}
	}

	fmt.Printf("Problem processing REPORT game!\n")
	return false
}

// ===================================================================
//
func AddGame(manager *GameManager, currGame *Game) {
	manager.ActiveGames = append(manager.ActiveGames, currGame)

	fmt.Printf("Added %+v to the GameManager!\n", currGame.ID)
}
