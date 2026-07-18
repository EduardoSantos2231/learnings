package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"
)

type StatusOutput struct {
	Track         string          `json:"track"`
	ActiveModule  string          `json:"active_module"`
	NextChallenge *NextChallenge  `json:"next_challenge,omitempty"`
	ReviewCount   int             `json:"review_count"`
	PendingReview *ReviewDetail   `json:"pending_review,omitempty"`
}

type NextChallenge struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Module  string `json:"module"`
	Type    string `json:"type"` // "challenge", "mixed_practice", "capstone"
}

type ReviewDetail struct {
	Challenge string `json:"challenge"`
	Interval  string `json:"interval"`
	Date      string `json:"date"`
}

func statusCmd() {
	sched, err := loadSchedule()
	if err != nil {
		fmt.Fprintf(osStderr, "erro: %v\n", err)
		return
	}

	if sched.ActiveTrack == "" {
		fmt.Println(`{"error":"nenhuma track ativa","use":"tracking start <track>"}`)
		return
	}

	roadmap, err := loadRoadmap(sched.ActiveTrack)
	if err != nil {
		fmt.Fprintf(osStderr, "erro: %v\n", err)
		return
	}

	out := StatusOutput{
		Track:        sched.ActiveTrack,
		ActiveModule: findActiveModule(roadmap),
	}

	nc := findNext(roadmap)
	if nc != nil {
		out.NextChallenge = nc
	}

	today := time.Now().Format("2006-01-02")
	for _, r := range sched.Reviews {
		if r.Track != sched.ActiveTrack {
			continue
		}
		for _, key := range reviewOrder {
			iv, ok := r.Intervals[key]
			if !ok || iv.Status != "pending" {
				continue
			}
			if iv.Date <= today {
				out.ReviewCount++
				if out.PendingReview == nil {
					out.PendingReview = &ReviewDetail{
						Challenge: r.Challenge,
						Interval:  key,
						Date:      iv.Date,
					}
				}
			}
		}
	}

	data, _ := json.MarshalIndent(out, "", "  ")
	fmt.Println(string(data))
}

var reviewOrder = []string{"1d", "3d", "7d", "30d"}

func findActiveModule(r Roadmap) string {
	for _, m := range r.Modules {
		for _, c := range m.Challenges {
			if c.Status == "pending" {
				return m.ID
			}
		}
	}
	return r.Modules[len(r.Modules)-1].ID
}

func findNext(r Roadmap) *NextChallenge {
	for _, m := range r.Modules {
		for _, c := range m.Challenges {
			if c.Status == "pending" {
				return &NextChallenge{ID: c.ID, Name: c.Name, Module: m.ID, Type: "challenge"}
			}
		}
	}
	for _, mp := range r.MixedPractice {
		if mp.Status == "pending" {
			return &NextChallenge{ID: mp.ID, Name: mp.Name, Module: "mixed-practice", Type: "mixed_practice"}
		}
	}
	for _, cs := range r.Capstones {
		if cs.Status == "pending" {
			return &NextChallenge{ID: cs.ID, Name: cs.Name, Module: "capstones", Type: "capstone"}
		}
	}
	return nil
}

func doneCmd(challengeID string) {
	sched, err := loadSchedule()
	if err != nil {
		fmt.Fprintf(osStderr, "erro carregando schedule: %v\n", err)
		return
	}
	if sched.ActiveTrack == "" {
		fmt.Println(`{"error":"nenhuma track ativa"}`)
		return
	}

	roadmap, err := loadRoadmap(sched.ActiveTrack)
	if err != nil {
		fmt.Fprintf(osStderr, "erro carregando roadmap: %v\n", err)
		return
	}

	found := markDone(&roadmap, challengeID)
	if !found {
		fmt.Printf(`{"error":"desafio '%s' nao encontrado ou ja concluido"}`, challengeID)
		fmt.Println()
		return
	}

	if err := saveRoadmap(sched.ActiveTrack, roadmap); err != nil {
		fmt.Fprintf(osStderr, "erro salvando roadmap: %v\n", err)
		return
	}

	today := time.Now()
	sched.Reviews = append(sched.Reviews, Review{
		Track:     sched.ActiveTrack,
		Challenge: challengeID,
		Completed: today.Format("2006-01-02"),
		Intervals: map[string]Interval{
			"1d":  {Date: today.AddDate(0, 0, 1).Format("2006-01-02"), Status: "pending"},
			"3d":  {Date: today.AddDate(0, 0, 3).Format("2006-01-02"), Status: "pending"},
			"7d":  {Date: today.AddDate(0, 0, 7).Format("2006-01-02"), Status: "pending"},
			"30d": {Date: today.AddDate(0, 0, 30).Format("2006-01-02"), Status: "pending"},
		},
	})

	sortReviews(&sched)

	if err := saveSchedule(sched); err != nil {
		fmt.Fprintf(osStderr, "erro salvando schedule: %v\n", err)
		return
	}

	result := map[string]interface{}{
		"ok":                true,
		"challenge":         challengeID,
		"reviews_scheduled": []string{"1d", "3d", "7d", "30d"},
	}
	data, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(data))
}

func markDone(r *Roadmap, id string) bool {
	for i, m := range r.Modules {
		for j, c := range m.Challenges {
			if c.ID == id {
				if c.Status == "done" {
					return false
				}
				r.Modules[i].Challenges[j].Status = "done"
				return true
			}
		}
	}
	for i, mp := range r.MixedPractice {
		if mp.ID == id {
			if mp.Status == "done" {
				return false
			}
			r.MixedPractice[i].Status = "done"
			return true
		}
	}
	for i, cs := range r.Capstones {
		if cs.ID == id {
			if cs.Status == "done" {
				return false
			}
			r.Capstones[i].Status = "done"
			return true
		}
	}
	return false
}

func sortReviews(s *Schedule) {
	sort.Slice(s.Reviews, func(i, j int) bool {
		return s.Reviews[i].Completed < s.Reviews[j].Completed
	})
}
