package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestChooseSessionRotatesReviewAndChallenge(t *testing.T) {
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, "framework", "tracks", "go-backend"), 0755); err != nil {
		t.Fatal(err)
	}
	roadmap := Roadmap{
		Track:   "go-backend",
		Modules: []Module{{ID: "modulo-A", Challenges: []Challenge{{ID: "A1", Name: "one", Status: "pending"}}}},
	}
	data, _ := json.Marshal(roadmap)
	if err := os.WriteFile(filepath.Join(root, "framework", "tracks", "go-backend", "roadmap.json"), data, 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Join(root, "framework", "tracks", "go-backend", "capstones", "C1-calculadora"), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "framework", "tracks", "go-backend", "capstones", "C1-calculadora", "README.md"), []byte("# capstone"), 0644); err != nil {
		t.Fatal(err)
	}
	old, _ := os.Getwd()
	if err := os.Chdir(root); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chdir(old) })

	now := time.Date(2026, 8, 6, 0, 0, 0, 0, time.UTC)
	s := Schedule{
		ActiveTrack: "go-backend",
		Reviews:     []Review{{Track: "go-backend", Block: "RB-go-fundamentos", Stage: "1d", NextDue: "2026-08-05", Status: "active"}},
	}
	choice, err := chooseSession(s, now)
	if err != nil || choice.Action != "review" || choice.Scenario != "undo e redo" {
		t.Fatalf("expected review, got %#v, %v", choice, err)
	}

	s.LastSessionKind = "review"
	choice, err = chooseSession(s, now)
	if err != nil || choice.Action != "challenge" || choice.ID != "A1" {
		t.Fatalf("expected challenge after review, got %#v, %v", choice, err)
	}
	if path := challengeReadmePath("go-backend", "C1", "capstones", "capstone"); path != "framework/tracks/go-backend/capstones/C1-calculadora/README.md" {
		t.Fatalf("expected capstone README, got %q", path)
	}
}

func TestAdvanceReviewIsProgressive(t *testing.T) {
	now := time.Date(2026, 8, 6, 0, 0, 0, 0, time.UTC)
	s := Schedule{ActiveTrack: "go-backend", Reviews: []Review{{
		Track: "go-backend", Block: "RB-go-fundamentos", Stage: "1d", NextDue: "2026-08-06", Status: "active",
	}}}

	if !advanceReview(&s, "RB-go-fundamentos", "passed", now) || s.Reviews[0].Stage != "7d" || s.Reviews[0].NextDue != "2026-08-13" {
		t.Fatalf("1d should advance to 7d: %#v", s.Reviews[0])
	}
	if !advanceReview(&s, "RB-go-fundamentos", "failed", now) || s.Reviews[0].Stage != "repair" || s.Reviews[0].Attempts != 1 {
		t.Fatalf("failure should create one repair: %#v", s.Reviews[0])
	}
}
