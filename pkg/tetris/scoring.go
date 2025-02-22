package tetris

import (
	"fmt"
	"os"
	"time"

	"github.com/stuttgart-things/homerun-library"
)

// Scoring is a scoring system for Tetris.
// It keeps track of the current level, total score, and lines cleared.
// It also has options to increase the level, end the game on max level, and end the game on max lines.
type Scoring struct {
	level         int
	maxLevel      int
	increaseLevel bool
	endOnMaxLevel bool

	lines         int
	maxLines      int
	endOnMaxLines bool

	total      int
	backToBack bool

	lastCleared int    // Stores the number of lines cleared in the last action
	playerName  string // ✅ New: Store player's name
}

// LastClearedLines returns the number of lines cleared in the most recent action.
func (s *Scoring) LastClearedLines() int {
	return s.lastCleared
}

// NewScoring creates a new scoring system.
func NewScoring(
	level, maxLevel int,
	increaseLevel, endOnMaxLevel bool,
	maxLines int,
	endOnMaxLines bool) (*Scoring, error) {
	s := &Scoring{
		level:         level,
		maxLevel:      maxLevel,
		increaseLevel: increaseLevel,
		endOnMaxLevel: endOnMaxLevel,

		maxLines:      maxLines,
		endOnMaxLines: endOnMaxLines,
	}
	return s, s.validate()
}

func (s *Scoring) validate() error {
	if s.level <= 0 {
		return fmt.Errorf("invalid level '%d'", s.level)
	}
	if s.maxLevel < 0 {
		return fmt.Errorf("invalid max level '%d'", s.maxLevel)
	}
	if s.maxLines < 0 {
		return fmt.Errorf("invalid max lines '%d'", s.maxLines)
	}
	if s.total < 0 {
		return fmt.Errorf("invalid total '%d'", s.total)
	}
	return nil
}

// Level returns the current level.
func (s *Scoring) Level() int {
	return s.level
}

// Total returns the total score.
func (s *Scoring) Total() int {
	return s.total
}

// Lines returns the total lines cleared.
func (s *Scoring) Lines() int {
	return s.lines
}

// AddSoftDrop adds points for a soft drop.
func (s *Scoring) AddSoftDrop(lines int) {
	s.total += lines
}

// AddHardDrop adds points for a hard drop.
func (s *Scoring) AddHardDrop(lines int) {
	s.total += lines * 2
}

// ProcessAction processes an action and updates the score, lines cleared, level, etc.
// The returned boolean indicates if the game should end.
// ProcessAction processes an action and updates the score, lines cleared, level, etc.
// Define the file path where we store cleared lines
const clearedLinesFile = "cleared_lines.log"

var (
	destination = "https://homerun.homerun-dev.sthings-vsphere.labul.sva.de/generic"
	token       = "IhrGeheimerToken"
	insecure    = true
	dt          = time.Now()
)

func sendNotification(linesCleared int) {
	dt := time.Now()
	messageBody := homerun.Message{
		Title:           "Lines Cleared",
		Message:         fmt.Sprintf("Lines Cleared: %d", linesCleared),
		Severity:        "INFO",
		Author:          "elvis",
		Timestamp:       dt.Format("01-02-2006 15:04:05"),
		System:          "Tetris Game",
		Tags:            "tetris,gameplay",
		AssigneeAddress: "",
		AssigneeName:    "",
		Artifacts:       "",
		Url:             "",
	}

	rendered := homerun.RenderBody(homerun.HomeRunBodyData, messageBody)

	// comment next line and uncomment Print answer lines to debug
	homerun.SendToHomerun(destination, token, []byte(rendered), insecure)
	// Print the answer for debugging purposes
	//answer, resp := homerun.SendToHomerun(destination, token, []byte(rendered), insecure)
	//fmt.Println("ANSWER STATUS: ", resp.Status)
	//fmt.Println("ANSWER BODY: ", string(answer))
}

func (s *Scoring) ProcessAction(a Action) (bool, error) {
	if a == Actions.None {
		return false, nil
	}

	points := float64(a.GetPoints())

	var err error
	var result bool
	backToBack := 0.0
	if result, err = a.EndsBackToBack(); result {
		s.backToBack = false
	} else if result, err = a.StartsBackToBack(); result {
		if s.backToBack {
			backToBack = points * 0.5
		}
		s.backToBack = true
	}
	if err != nil {
		return false, err
	}

	s.total += int(points+backToBack) * s.level

	// Get cleared lines
	linesCleared := a.GetLinesCleared()

	// ✅ Write to file if cleared lines change and it's not 0
	if linesCleared > 0 && linesCleared != s.lastCleared {
		sendNotification(linesCleared)
		//err := appendToFile(clearedLinesFile, fmt.Sprintf("Lines Cleared: %d\n", linesCleared))
		//if err != nil {
		//	fmt.Println("⚠️ Error writing to file:", err)
	}

	// Update last cleared lines
	s.lastCleared = linesCleared
	s.lines += linesCleared

	// Check for max lines
	if s.maxLines > 0 && s.lines >= s.maxLines {
		s.lines = s.maxLines
		if s.endOnMaxLines {
			return true, nil
		}
	}

	// Level up logic
	for s.increaseLevel && s.lines >= s.level*5 {
		s.level++
		if s.maxLevel > 0 && s.level >= s.maxLevel {
			s.level = s.maxLevel
			if s.endOnMaxLevel {
				return true, nil
			}
			break
		}
	}

	return false, nil
}

func appendToFile(filename, content string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}
