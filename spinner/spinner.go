package spinner

import (
	"fmt"
	"sync/atomic"
	"time"
)

// ClearLine go to the beggining of the line and clear it
const ClearLine = "\r\033[K"

// Spin type
var Spin = `⠋⠙⠹⠸⠼⠴⠦⠧⠇⠏`

// Spinner main type
type Spinner struct {
	frames []rune
	pos    int
	active uint64
	text   string
}

// New Spinner with args
func New(text string) *Spinner {
	s := &Spinner{
		text: ClearLine + text,
	}
	s.Set(Spin)
	return s
}

// Set frames to the given string which must not use spaces.
func (s *Spinner) Set(frames string) {
	s.frames = []rune(frames)
}

// Start shows the spinner.
func (s *Spinner) Start() *Spinner {
	if atomic.LoadUint64(&s.active) > 0 {
		return s
	}
	atomic.StoreUint64(&s.active, 1)
	go func() {
		for atomic.LoadUint64(&s.active) > 0 {
			fmt.Printf(s.text, s.next())
			time.Sleep(100 * time.Millisecond)
		}
	}()
	return s
}

// Stop hides the spinner.
func (s *Spinner) Stop() bool {
	if x := atomic.SwapUint64(&s.active, 0); x > 0 {
		fmt.Printf(ClearLine)
		return true
	}
	return false
}

func (s *Spinner) next() string {
	r := s.frames[s.pos%len(s.frames)]
	s.pos++
	return string(r)
}
