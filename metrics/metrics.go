package metrics

import "time"

type Stroke struct {
	// The character that was typed
	Char string
	// The time it was typed
	Time time.Time
}

func GetRawWpm(strokes []Stroke) int {
	// We're just going to iterate over the strokes, accumulate the times, and add a word everytime we see a space
	// We'll then divide the total time by the number of words and multiply by 60 to get the words per minute
	words := 0
	running_time := 0.0
	for i, stroke := range strokes {
		if stroke.Char == " " {
			words++
		}
		if i == 0 {
			continue
		}
		if stroke.Char == "backspace" {
			continue
		}
		running_time += stroke.Time.Sub(strokes[i-1].Time).Seconds()
	}
	return int(float64(words) / (running_time / 60))
}

func GetWpm(strokes []Stroke, acc float64) int {
	// We'll just get the raw wpm and divide by 5
	return int(float64(GetRawWpm(strokes)) * acc)
}

func GetTimeTaken(strokes []Stroke) float64 {
	// We'll just get the time taken from the first stroke to the last stroke
	return strokes[len(strokes)-1].Time.Sub(strokes[0].Time).Seconds()
}

func TimeLostByFixingMistakes(strokes []Stroke) float64 {
	// We'll just iterate over the strokes and calculate the time lost by fixing mistakes
	// We'll just return the time lost
	time_lost := 0.0
	for i, stroke := range strokes {
		if i == 0 {
			continue
		}
		if stroke.Char == "backspace" {
			time_lost += stroke.Time.Sub(strokes[i-1].Time).Seconds()
		}
	}
	return time_lost
}

func ThinkingTime(strokes []Stroke) float64 {
	// This will calculate the time taken after pressing space and before typing the next word

	time_taken := 0.0
	for i, stroke := range strokes {
		if i == 0 {
			continue
		}
		if strokes[i-1].Char == " " {
			time_taken += stroke.Time.Sub(strokes[i-1].Time).Seconds()
		}
	}
	return time_taken
}
