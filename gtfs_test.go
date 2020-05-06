package gtfs

import (
	"fmt"
	"path"
	"reflect"
	"testing"
)

func TestLoad(t *testing.T) {
	gtfs, err := Load(path.Join("gtfs_test", "google"), nil)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	models := map[string]interface{}{
		"Agencies":       gtfs.Agencies,
		"Routes":         gtfs.Routes,
		"Stops":          gtfs.Stops,
		"StopsTimes":     gtfs.StopsTimes,
		"Trips":          gtfs.Trips,
		"Calendars":      gtfs.Calendars,
		"CalendarDates":  gtfs.CalendarDates,
		"Transfers":      gtfs.Transfers,
		"Shapes":         gtfs.Shapes,
		"FareAttributes": gtfs.FareAttributes,
		"FareRules":      gtfs.FareRules,
		"Frequencies":    gtfs.Frequencies,
		"Pathways":       gtfs.Pathways,
		"Levels":         gtfs.Levels,
		"Translations":   gtfs.Translations,
		"Attributions":   gtfs.Attributions,
	}

	for name, model := range models {
		s := reflect.ValueOf(model)
		if s.Len() == 0 {
			fmt.Printf("%s is empty\n", name)
			t.Fail()
		}
	}
}

func TestLoadBad(t *testing.T) {
	_, err := Load(path.Join("gtfs_test", "ratp_bad"), nil)
	if err == nil {
		t.Fail()
	}
}

func TestLoadNoDir(t *testing.T) {
	_, err := Load(path.Join("gtfs_test", "no_dir"), nil)
	if err == nil {
		t.Fail()
	}
}

func TestLoadPartial(t *testing.T) {
	gtfs, err := Load(path.Join("gtfs_test", "google"), map[string]bool{"routes": true})
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	models := map[string]interface{}{
		"Agencies":       gtfs.Agencies,
		"Stops":          gtfs.Stops,
		"StopsTimes":     gtfs.StopsTimes,
		"Trips":          gtfs.Trips,
		"Calendars":      gtfs.Calendars,
		"CalendarDates":  gtfs.CalendarDates,
		"Transfers":      gtfs.Transfers,
		"Shapes":         gtfs.Shapes,
		"FareAttributes": gtfs.FareAttributes,
		"FareRules":      gtfs.FareRules,
		"Frequencies":    gtfs.Frequencies,
		"Pathways":       gtfs.Pathways,
		"Levels":         gtfs.Levels,
		"Translations":   gtfs.Translations,
		"Attributions":   gtfs.Attributions,
	}

	for name, model := range models {
		s := reflect.ValueOf(model)
		if s.Len() != 0 {
			fmt.Printf("%s should be empty\n", name)
			t.Fail()
		}
	}

	if len(gtfs.Routes) == 0 {
		fmt.Println("Routes should not be empty")
		t.Fail()
	}
}

func TestLoadSplitted(t *testing.T) {
	gtfss, err := LoadSplitted(path.Join("gtfs_test", "ratp_splitted"), nil)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if len(gtfss) != 2 {
		t.Fail()
	}
}

func TestLoadSplittedBad(t *testing.T) {
	_, err := LoadSplitted(path.Join("gtfs_test", "ratp_splitted_bad"), nil)
	if err == nil {
		t.Fail()
	}
}

func TestLoadSplittedNoDir(t *testing.T) {
	_, err := LoadSplitted(path.Join("gtfs_test", "no_dir"), nil)
	if err == nil {
		t.Fail()
	}
}
