package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type StatusOutput struct {
	Track         string         `json:"track"`
	ActiveModule  string         `json:"active_module"`
	NextChallenge *NextChallenge `json:"next_challenge,omitempty"`
	ReviewCount   int            `json:"review_count"`
	PendingReview *ReviewDetail  `json:"pending_review,omitempty"`
}

type NextChallenge struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Module string `json:"module"`
	Type   string `json:"type"` // "challenge", "mixed_practice", "capstone"
}

type ReviewDetail struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Stage string `json:"stage"`
	Date  string `json:"date"`
}

type ReviewBlock struct {
	ID        string
	Title     string
	Path      string
	Sources   []string
	Minutes   int
	Scenarios []string
}

type SessionChoice struct {
	Action   string   `json:"action"`
	ID       string   `json:"id"`
	Title    string   `json:"title"`
	Reason   string   `json:"reason"`
	Path     string   `json:"path,omitempty"`
	Module   string   `json:"module,omitempty"`
	Sources  []string `json:"sources,omitempty"`
	Minutes  int      `json:"estimated_minutes,omitempty"`
	Stage    string   `json:"stage,omitempty"`
	Due      string   `json:"due,omitempty"`
	Scenario string   `json:"scenario,omitempty"`
}

var reviewBlocks = map[string]ReviewBlock{
	"RB-go-fundamentos": {"RB-go-fundamentos", "Fundamentos: erros, ponteiros e interfaces", "framework/tracks/go-backend/mixed-practice/mp1/README.md", []string{"A1", "A2", "A3", "A4"}, 45, []string{"undo e redo", "erros classificaveis", "plugin por interface"}},
	"RB-go-concurrency": {"RB-go-concurrency", "Concorrência: limitar, cancelar e combinar tarefas", "framework/tracks/go-backend/mixed-practice/mp2/README.md", []string{"B1", "B2", "B3", "B4", "B5"}, 45, []string{"processamento limitado", "gateway protegido", "agregador com timeout"}},
	"RB-go-structures":  {"RB-go-structures", "Estruturas e I/O: escolher e medir", "framework/tracks/go-backend/mixed-practice/mp3/README.md", []string{"C1", "C2", "C3", "C4", "C5"}, 45, []string{"historico navegavel", "log maior que a memoria", "ranking em tempo real"}},
	"RB-go-http":        {"RB-go-http", "HTTP: middleware, concorrência e falhas", "framework/tracks/go-backend/mixed-practice/mp4/README.md", []string{"D1", "D2", "D3"}, 45, []string{"middleware fora de ordem", "race no cache", "vazamento de goroutine"}},
	"RB-docker-runtime": {"RB-docker-runtime", "Docker: ciclo de vida e processo principal", "framework/tracks/docker-devops/mixed-practice/mp1/README.md", []string{"D1", "D2", "D3", "D4"}, 30, []string{"container encerra", "PID 1", "porta inacessivel"}},
	"RB-docker-build":   {"RB-docker-build", "Docker: build reproduzivel e imagem pequena", "framework/tracks/docker-devops/mixed-practice/mp2/README.md", []string{"D5", "D6", "D7"}, 30, []string{"cache invalido", "multi-stage", "healthcheck"}},
}

var reviewBlockByChallenge = map[string]string{
	"go-backend:MP1":    "RB-go-fundamentos",
	"go-backend:MP2":    "RB-go-concurrency",
	"go-backend:MP3":    "RB-go-structures",
	"go-backend:MP4":    "RB-go-http",
	"docker-devops:MP1": "RB-docker-runtime",
	"docker-devops:MP2": "RB-docker-build",
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
		if r.Track != sched.ActiveTrack || r.Block == "" || r.Status == "completed" || r.Status == "archived" {
			continue
		}
		if r.NextDue != "" && r.NextDue <= today {
			out.ReviewCount++
			if out.PendingReview == nil {
				block := reviewBlocks[r.Block]
				out.PendingReview = &ReviewDetail{ID: r.Block, Title: block.Title, Stage: r.Stage, Date: r.NextDue}
			}
		}
	}

	data, _ := json.MarshalIndent(out, "", "  ")
	fmt.Println(string(data))
}

func findActiveModule(r Roadmap) string {
	for mi, m := range r.Modules {
		for _, c := range m.Challenges {
			if c.Status == "pending" {
				return m.ID
			}
		}
		if mi < len(r.MixedPractice) && r.MixedPractice[mi].Status == "pending" {
			return m.ID
		}
		if mi < len(r.Capstones) && r.Capstones[mi].Status == "pending" {
			return m.ID
		}
	}
	return r.Modules[len(r.Modules)-1].ID
}

func findNext(r Roadmap) *NextChallenge {
	for mi, m := range r.Modules {
		for _, c := range m.Challenges {
			if c.Status == "pending" {
				return &NextChallenge{ID: c.ID, Name: c.Name, Module: m.ID, Type: "challenge"}
			}
		}
		if mi < len(r.MixedPractice) && r.MixedPractice[mi].Status == "pending" {
			mp := r.MixedPractice[mi]
			return &NextChallenge{ID: mp.ID, Name: mp.Name, Module: "mixed-practice", Type: "mixed_practice"}
		}
		if mi < len(r.Capstones) && r.Capstones[mi].Status == "pending" {
			cs := r.Capstones[mi]
			return &NextChallenge{ID: cs.ID, Name: cs.Name, Module: "capstones", Type: "capstone"}
		}
	}
	for i := len(r.Modules); i < len(r.MixedPractice); i++ {
		if r.MixedPractice[i].Status == "pending" {
			return &NextChallenge{ID: r.MixedPractice[i].ID, Name: r.MixedPractice[i].Name, Module: "mixed-practice", Type: "mixed_practice"}
		}
	}
	for i := len(r.Modules); i < len(r.Capstones); i++ {
		if r.Capstones[i].Status == "pending" {
			return &NextChallenge{ID: r.Capstones[i].ID, Name: r.Capstones[i].Name, Module: "capstones", Type: "capstone"}
		}
	}
	return nil
}

func markDone(r *Roadmap, id string) bool {
	for i, m := range r.Modules {
		for j, c := range m.Challenges {
			if c.ID == id {
				if c.Status == "done" {
					continue
				}
				r.Modules[i].Challenges[j].Status = "done"
				return true
			}
		}
	}
	for i, mp := range r.MixedPractice {
		if mp.ID == id {
			if mp.Status == "done" {
				continue
			}
			r.MixedPractice[i].Status = "done"
			return true
		}
	}
	for i, cs := range r.Capstones {
		if cs.ID == id {
			if cs.Status == "done" {
				continue
			}
			r.Capstones[i].Status = "done"
			return true
		}
	}
	return false
}

func sessionCmd() {
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
	data, _ := json.MarshalIndent(choice, "", "  ")
	fmt.Println(string(data))
}

func chooseSession(s Schedule, now time.Time) (SessionChoice, error) {
	if s.ActiveTrack == "" {
		return SessionChoice{}, fmt.Errorf("nenhuma track ativa")
	}
	if s.BaselinePending != nil && s.BaselinePending[s.ActiveTrack] {
		return SessionChoice{
			Action: "challenge", ID: "DIAG", Title: "Diagnóstico prático", Reason: "baseline_required",
			Path: filepath.Join("framework", "tracks", s.ActiveTrack, "diagnostic.md"), Minutes: 45,
		}, nil
	}
	r, err := loadRoadmap(s.ActiveTrack)
	if err != nil {
		return SessionChoice{}, err
	}
	next := findNext(r)
	review, ok := oldestDueReview(s, now)
	if ok && (s.LastSessionKind != "review" || next == nil) {
		block := reviewBlocks[review.Block]
		return SessionChoice{
			Action: "review", ID: review.Block, Title: block.Title, Reason: reviewReason(review),
			Path: block.Path, Sources: block.Sources, Minutes: block.Minutes, Stage: review.Stage, Due: review.NextDue, Scenario: reviewScenario(block, review),
		}, nil
	}
	if next != nil {
		return SessionChoice{
			Action: "challenge", ID: next.ID, Title: next.Name, Reason: challengeReason(ok),
			Path: challengeReadmePath(s.ActiveTrack, next.ID, next.Module, next.Type), Module: next.Module,
		}, nil
	}
	if ok {
		block := reviewBlocks[review.Block]
		return SessionChoice{
			Action: "review", ID: review.Block, Title: block.Title, Reason: reviewReason(review),
			Path: block.Path, Sources: block.Sources, Minutes: block.Minutes, Stage: review.Stage, Due: review.NextDue, Scenario: reviewScenario(block, review),
		}, nil
	}
	return SessionChoice{Action: "complete", Reason: "roadmap_complete"}, nil
}

func oldestDueReview(s Schedule, now time.Time) (Review, bool) {
	today := now.Format("2006-01-02")
	var due []Review
	for _, r := range s.Reviews {
		if r.Track == s.ActiveTrack && r.Block != "" && r.Status != "completed" && r.Status != "archived" && r.NextDue != "" && r.NextDue <= today {
			due = append(due, r)
		}
	}
	if len(due) == 0 {
		return Review{}, false
	}
	sort.SliceStable(due, func(i, j int) bool {
		if due[i].NextDue != due[j].NextDue {
			return due[i].NextDue < due[j].NextDue
		}
		return due[i].Attempts > due[j].Attempts
	})
	return due[0], true
}

func reviewReason(r Review) string {
	if r.Stage == "repair" {
		return "repair_due"
	}
	return "review_due"
}

func reviewScenario(block ReviewBlock, r Review) string {
	if len(block.Scenarios) == 0 {
		return ""
	}
	return block.Scenarios[r.Variant%len(block.Scenarios)]
}

func challengeReason(hasReview bool) string {
	if hasReview {
		return "rotation_after_review"
	}
	return "no_review_due"
}

func challengeReadmePath(track, id, module, kind string) string {
	root := filepath.Join(basePath(), "framework", "tracks", track)
	var result string
	_ = filepath.WalkDir(root, func(path string, entry os.DirEntry, err error) error {
		if err != nil || result != "" || entry.IsDir() || filepath.Base(path) != "README.md" {
			return nil
		}
		rel, relErr := filepath.Rel(basePath(), path)
		if relErr != nil {
			return nil
		}
		relSlash := filepath.ToSlash(rel)
		if kind == "capstone" && !strings.Contains(relSlash, "/capstones/") {
			return nil
		}
		if kind == "mixed_practice" && !strings.Contains(relSlash, "/mixed-practice/") {
			return nil
		}
		if module != "" && module != "capstones" && module != "mixed-practice" && !strings.Contains(relSlash, "/"+module+"/") {
			return nil
		}
		parent := filepath.Base(filepath.Dir(path))
		parentUpper := strings.ToUpper(parent)
		idUpper := strings.ToUpper(id)
		if strings.EqualFold(parent, id) || strings.HasPrefix(parentUpper, idUpper+"-") {
			result = rel
		}
		return nil
	})
	return filepath.ToSlash(result)
}

func reviewBlockFor(track, challenge string) string {
	return reviewBlockByChallenge[track+":"+challenge]
}

func scheduleReview(s *Schedule, track, challenge string, now time.Time) {
	block := reviewBlockFor(track, challenge)
	if block == "" {
		return
	}
	for _, r := range s.Reviews {
		if r.Track == track && r.Block == block && r.Status != "completed" && r.Status != "archived" {
			return
		}
	}
	s.Reviews = append(s.Reviews, Review{
		Track: track, Challenge: challenge, Block: block, Completed: now.Format("2006-01-02"),
		Stage: "1d", NextDue: now.AddDate(0, 0, 1).Format("2006-01-02"), Status: "active",
	})
}
