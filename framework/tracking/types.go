package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Challenge struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Template string   `json:"template"`
	Concepts []string `json:"concepts"`
	Status   string   `json:"status"` // "done" or "pending"
}

type Module struct {
	ID          string      `json:"id"`
	Scaffolding string      `json:"scaffolding"`
	Challenges  []Challenge `json:"challenges"`
}

type Roadmap struct {
	Track         string      `json:"track"`
	CurrentModule string      `json:"current_module"`
	Modules       []Module    `json:"modules"`
	MixedPractice []Challenge `json:"mixed_practice"`
	Capstones     []Challenge `json:"capstones"`
}

type Interval struct {
	Date   string `json:"date"`
	Status string `json:"status"` // "pending", "passed", "failed"
}

type Review struct {
	Track     string              `json:"track"`
	Challenge string              `json:"challenge"`
	Completed string              `json:"completed"`
	Intervals map[string]Interval `json:"intervals"`
}

type Schedule struct {
	ActiveTrack string   `json:"active_track"`
	Reviews     []Review `json:"reviews"`
}

type CorrectionEntry struct {
	Challenge    string   `json:"challenge"`
	Category     string   `json:"category"`
	Error        string   `json:"error"`
	Recurrences  []string `json:"recurrences,omitempty"`
}

type Corrections struct {
	Entries []CorrectionEntry `json:"entries"`
}

func basePath() string {
	dir, err := os.Getwd()
	if err != nil {
		dir = "."
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "spaced-repetition")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	wd, _ := os.Getwd()
	return wd
}

func loadRoadmap(track string) (Roadmap, error) {
	var r Roadmap
	path := filepath.Join(basePath(), "framework", "tracks", track, "roadmap.json")
	data, err := os.ReadFile(path)
	if err != nil {
		return r, err
	}
	err = json.Unmarshal(data, &r)
	return r, err
}

func saveRoadmap(track string, r Roadmap) error {
	path := filepath.Join(basePath(), "framework", "tracks", track, "roadmap.json")
	data, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

func loadSchedule() (Schedule, error) {
	var s Schedule
	path := filepath.Join(basePath(), "spaced-repetition", "schedule.json")
	data, err := os.ReadFile(path)
	if err != nil {
		return s, err
	}
	err = json.Unmarshal(data, &s)
	return s, err
}

func saveSchedule(s Schedule) error {
	path := filepath.Join(basePath(), "spaced-repetition", "schedule.json")
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

func loadCorrections(track string) (Corrections, error) {
	var c Corrections
	path := filepath.Join(basePath(), "framework", "tracks", track, "corrections.json")
	data, err := os.ReadFile(path)
	if err != nil {
		return c, err
	}
	err = json.Unmarshal(data, &c)
	return c, err
}

func saveCorrections(track string, c Corrections) error {
	path := filepath.Join(basePath(), "framework", "tracks", track, "corrections.json")
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
