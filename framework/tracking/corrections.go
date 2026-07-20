package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func reviewCmd(challengeID, interval, result string) {
	sched, err := loadSchedule()
	if err != nil {
		fmt.Fprintf(osStderr, "erro: %v\n", err)
		return
	}

	for i, r := range sched.Reviews {
		if r.Track != sched.ActiveTrack || r.Challenge != challengeID {
			continue
		}
		iv, ok := r.Intervals[interval]
		if !ok {
			fmt.Printf(`{"error":"intervalo '%s' nao encontrado"}`, interval)
			fmt.Println()
			return
		}
		iv.Status = result
		sched.Reviews[i].Intervals[interval] = iv

		if result == "failed" {
			today := time.Now()
			sched.Reviews[i].Intervals["1d"] = Interval{Date: today.AddDate(0, 0, 1).Format("2006-01-02"), Status: "pending"}
			sched.Reviews[i].Intervals["3d"] = Interval{Date: today.AddDate(0, 0, 3).Format("2006-01-02"), Status: "pending"}
			sched.Reviews[i].Intervals["7d"] = Interval{Date: today.AddDate(0, 0, 7).Format("2006-01-02"), Status: "pending"}
			sched.Reviews[i].Intervals["30d"] = Interval{Date: today.AddDate(0, 0, 30).Format("2006-01-02"), Status: "pending"}
		}
	}

	if err := saveSchedule(sched); err != nil {
		fmt.Fprintf(osStderr, "erro salvando: %v\n", err)
		return
	}

	out := map[string]interface{}{
		"ok":        true,
		"challenge": challengeID,
		"interval":  interval,
		"status":    result,
	}
	data, _ := json.MarshalIndent(out, "", "  ")
	fmt.Println(string(data))
}

func addErrorCmd(challengeID, category, description string) {
	sched, err := loadSchedule()
	if err != nil {
		fmt.Fprintf(osStderr, "erro: %v\n", err)
		return
	}

	c, err := loadCorrections(sched.ActiveTrack)
	if err != nil {
		c = Corrections{Entries: []CorrectionEntry{}}
	}

	c.Entries = append(c.Entries, CorrectionEntry{
		Challenge: challengeID,
		Category:  category,
		Error:     description,
	})

	if err := saveCorrections(sched.ActiveTrack, c); err != nil {
		fmt.Fprintf(osStderr, "erro salvando: %v\n", err)
		return
	}

	fmt.Println(`{"ok":true}`)
}

func checkRecurrenceCmd(category string) {
	sched, err := loadSchedule()
	if err != nil {
		fmt.Fprintf(osStderr, "erro: %v\n", err)
		return
	}

	c, err := loadCorrections(sched.ActiveTrack)
	if err != nil {
		fmt.Println(`{"found":false}`)
		return
	}

	var matches []CorrectionEntry
	for _, e := range c.Entries {
		if e.Category == category {
			matches = append(matches, e)
		}
	}

	out := map[string]interface{}{
		"found":  len(matches) > 0,
		"count":  len(matches),
		"errors": matches,
	}
	data, _ := json.MarshalIndent(out, "", "  ")
	fmt.Println(string(data))
}

func startCmd(track string) {
	sched, err := loadSchedule()
	if err != nil {
		sched = Schedule{Reviews: []Review{}}
	}

	// Verify track exists
	_, err = loadRoadmap(track)
	if err != nil {
		fmt.Printf(`{"error":"track '%s' nao encontrada"}`, track)
		fmt.Println()
		return
	}

	sched.ActiveTrack = track

	if err := saveSchedule(sched); err != nil {
		fmt.Fprintf(osStderr, "erro: %v\n", err)
		return
	}

	fmt.Printf(`{"ok":true,"active_track":"%s"}`, track)
	fmt.Println()
}

func renderRoadmapCmd() {
	sched, err := loadSchedule()
	if err != nil {
		fmt.Fprintf(osStderr, "erro: %v\n", err)
		return
	}

	r, err := loadRoadmap(sched.ActiveTrack)
	if err != nil {
		fmt.Fprintf(osStderr, "erro: %v\n", err)
		return
	}

	fmt.Printf("# Roadmap — %s\n\n", r.Track)
	fmt.Println("## Módulos")
	for _, m := range r.Modules {
		fmt.Printf("### %s (%s)\n\n", m.ID, m.Scaffolding)
		fmt.Println("| # | Desafio | Status |")
		fmt.Println("|---|---------|--------|")
		for _, c := range m.Challenges {
			status := "⬜"
			if c.Status == "done" {
				status = "✅"
			}
			fmt.Printf("| %s | %s | %s |\n", c.ID, c.Name, status)
		}
		fmt.Println()
	}
	if len(r.MixedPractice) > 0 {
		fmt.Println("## Mixed Practice")
	fmt.Println()
		fmt.Println("| # | Desafio | Status |")
		fmt.Println("|---|---------|--------|")
		for _, mp := range r.MixedPractice {
			status := "⬜"
			if mp.Status == "done" {
				status = "✅"
			}
			fmt.Printf("| %s | %s | %s |\n", mp.ID, mp.Name, status)
		}
		fmt.Println()
	}
	if len(r.Capstones) > 0 {
		fmt.Println("## Capstones")
	fmt.Println()
		fmt.Println("| # | Projeto | Status |")
		fmt.Println("|---|---------|--------|")
		for _, cs := range r.Capstones {
			status := "⬜"
			if cs.Status == "done" {
				status = "✅"
			}
			fmt.Printf("| %s | %s | %s |\n", cs.ID, cs.Name, status)
		}
		fmt.Println()
	}
}
