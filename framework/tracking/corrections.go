package main

import (
	"encoding/json"
	"fmt"
	"time"
)

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
	if sched.BaselinePending == nil {
		sched.BaselinePending = map[string]bool{}
	}

	// Verify track exists
	_, err = loadRoadmap(track)
	if err != nil {
		fmt.Printf(`{"error":"track '%s' nao encontrada"}`, track)
		fmt.Println()
		return
	}

	sched.ActiveTrack = track
	if _, exists := sched.BaselinePending[track]; !exists {
		sched.BaselinePending[track] = true
	}

	if err := saveSchedule(sched); err != nil {
		fmt.Fprintf(osStderr, "erro: %v\n", err)
		return
	}

	fmt.Printf(`{"ok":true,"active_track":"%s"}`, track)
	fmt.Println()
}

func finishCmd(id, result string) {
	sched, err := loadSchedule()
	if err != nil {
		fmt.Fprintf(osStderr, "erro: %v\n", err)
		return
	}
	choice, err := chooseSession(sched, time.Now())
	if err != nil {
		fmt.Fprintf(osStderr, "erro: %v\n", err)
		return
	}
	if choice.ID != id {
		fmt.Printf(`{"error":"sessao esperada '%s', recebido '%s'"}`, choice.ID, id)
		fmt.Println()
		return
	}
	if choice.Action == "challenge" && result != "passed" {
		fmt.Fprintln(osStderr, "desafio exige --pass depois da correcao")
		return
	}

	now := time.Now()
	if choice.Action == "challenge" {
		if id == "DIAG" {
			if sched.BaselinePending == nil {
				sched.BaselinePending = map[string]bool{}
			}
			sched.BaselinePending[sched.ActiveTrack] = false
		} else {
			roadmap, loadErr := loadRoadmap(sched.ActiveTrack)
			if loadErr != nil {
				fmt.Fprintf(osStderr, "erro carregando roadmap: %v\n", loadErr)
				return
			}
			if !markDone(&roadmap, id) {
				fmt.Printf(`{"error":"desafio '%s' nao encontrado ou ja concluido"}`, id)
				fmt.Println()
				return
			}
			if err := saveRoadmap(sched.ActiveTrack, roadmap); err != nil {
				fmt.Fprintf(osStderr, "erro salvando roadmap: %v\n", err)
				return
			}
			scheduleReview(&sched, sched.ActiveTrack, id, now)
		}
		sched.LastSessionKind = "challenge"
		sched.LastSessionID = id
	} else if choice.Action == "review" {
		if result != "passed" && result != "failed" {
			fmt.Fprintln(osStderr, "uma revisão exige --pass ou --fail")
			return
		}
		if !advanceReview(&sched, id, result, now) {
			fmt.Printf(`{"error":"revisao '%s' nao encontrada"}`, id)
			fmt.Println()
			return
		}
		sched.LastSessionKind = "review"
		sched.LastSessionID = id
	}

	sched.SchemaVersion = 2
	if err := saveSchedule(sched); err != nil {
		fmt.Fprintf(osStderr, "erro salvando: %v\n", err)
		return
	}
	data, _ := json.MarshalIndent(map[string]interface{}{
		"ok": true, "id": id, "status": result,
	}, "", "  ")
	fmt.Println(string(data))
}

func rebaselineCmd() {
	sched, err := loadSchedule()
	if err != nil {
		fmt.Fprintf(osStderr, "erro: %v\n", err)
		return
	}
	if sched.ActiveTrack == "" {
		fmt.Fprintln(osStderr, "nenhuma track ativa")
		return
	}
	for i := range sched.Reviews {
		if sched.Reviews[i].Block == "" || sched.Reviews[i].Status == "active" {
			sched.Reviews[i].Status = "archived"
			sched.Reviews[i].NextDue = ""
		}
	}
	if sched.BaselinePending == nil {
		sched.BaselinePending = map[string]bool{}
	}
	sched.BaselinePending[sched.ActiveTrack] = true
	sched.SchemaVersion = 2
	sched.LastSessionKind = ""
	sched.LastSessionID = ""
	if err := saveSchedule(sched); err != nil {
		fmt.Fprintf(osStderr, "erro salvando: %v\n", err)
		return
	}
	fmt.Printf(`{"ok":true,"track":"%s","archived":true}`, sched.ActiveTrack)
	fmt.Println()
}

func advanceReview(s *Schedule, id, result string, now time.Time) bool {
	for i := range s.Reviews {
		r := &s.Reviews[i]
		if r.Track != s.ActiveTrack || r.Block != id || r.Status == "completed" || r.Status == "archived" {
			continue
		}
		if result == "failed" {
			r.Stage = "repair"
			r.NextDue = now.AddDate(0, 0, 1).Format("2006-01-02")
			r.Attempts++
			r.Variant++
			return true
		}
		r.Variant++
		switch r.Stage {
		case "1d", "repair":
			r.Stage = "7d"
			r.NextDue = now.AddDate(0, 0, 7).Format("2006-01-02")
		case "7d":
			r.Stage = "30d"
			r.NextDue = now.AddDate(0, 0, 30).Format("2006-01-02")
		default:
			r.Status = "completed"
			r.NextDue = ""
		}
		return true
	}
	return false
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
