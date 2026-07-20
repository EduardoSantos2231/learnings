package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type BookEntry struct {
	Slug          string `json:"slug"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	TotalChapters int    `json:"total_chapters"`
	ChaptersRead  int    `json:"chapters_read"`
	Category      string `json:"category"`
}

type BookIndex struct {
	CurrentBook string      `json:"current_book"`
	Books       []BookEntry `json:"books"`
}

type ChapterEntry struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
	Status string `json:"status"`
	ReadAt string `json:"read_at,omitempty"`
}

type BookMeta struct {
	Title         string         `json:"title"`
	Author        string         `json:"author"`
	Category      string         `json:"category"`
	TotalChapters int            `json:"total_chapters"`
	Chapters      []ChapterEntry `json:"chapters"`
}

func bookCmd(args []string) {
	if len(args) == 0 {
		fmt.Fprintln(osStderr, "book: subcomando requerido: list, start, status, note, done, reflect, update")
		return
	}
	switch args[0] {
	case "list":
		bookList()
	case "start":
		if len(args) < 6 {
			fmt.Fprintln(osStderr, "uso: tracking book start <slug> <titulo> <autor> <total-cap> <categoria>")
			return
		}
		total, _ := strconv.Atoi(args[4])
		bookStart(args[1], args[2], args[3], total, args[5])
	case "status":
		bookStatus()
	case "note":
		if len(args) < 2 {
			fmt.Fprintln(osStderr, "uso: tracking book note <capitulo>")
			return
		}
		cap, _ := strconv.Atoi(args[1])
		bookNote(cap)
	case "done":
		if len(args) < 2 {
			fmt.Fprintln(osStderr, "uso: tracking book done <capitulo>")
			return
		}
		cap, _ := strconv.Atoi(args[1])
		bookDone(cap)
	case "reflect":
		bookReflect()
	case "switch":
		if len(args) < 2 {
			fmt.Fprintln(osStderr, "uso: tracking book switch <slug>")
			return
		}
		bookSwitch(args[1])
	case "update":
		if len(args) < 4 {
			fmt.Fprintln(osStderr, "uso: tracking book update <slug> <campo> <valor>")
			return
		}
		bookUpdate(args[1], args[2], args[3])
	default:
		fmt.Fprintf(osStderr, "book: subcomando desconhecido: %s\n", args[0])
	}
}

func booksPath() string {
	return filepath.Join(basePath(), "books")
}

func loadBookIndex() (BookIndex, error) {
	var idx BookIndex
	path := filepath.Join(booksPath(), "index.json")
	data, err := os.ReadFile(path)
	if err != nil {
		return idx, err
	}
	err = json.Unmarshal(data, &idx)
	return idx, err
}

func saveBookIndex(idx BookIndex) error {
	path := filepath.Join(booksPath(), "index.json")
	data, err := json.MarshalIndent(idx, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

func loadBookMeta(slug string) (BookMeta, error) {
	var m BookMeta
	path := filepath.Join(booksPath(), slug, "meta.json")
	data, err := os.ReadFile(path)
	if err != nil {
		return m, err
	}
	err = json.Unmarshal(data, &m)
	return m, err
}

func saveBookMeta(slug string, m BookMeta) error {
	path := filepath.Join(booksPath(), slug, "meta.json")
	data, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

func findActiveBook() (string, error) {
	idx, err := loadBookIndex()
	if err != nil {
		return "", err
	}
	if idx.CurrentBook != "" {
		for _, b := range idx.Books {
			if b.Slug == idx.CurrentBook {
				return idx.CurrentBook, nil
			}
		}
	}
	if len(idx.Books) > 0 {
		return idx.Books[0].Slug, nil
	}
	return "", fmt.Errorf("nenhum livro registrado")
}

func bookList() {
	idx, err := loadBookIndex()
	if err != nil {
		fmt.Println(`{"error":"index.json nao encontrado"}`)
		return
	}
	for _, b := range idx.Books {
		pct := 0
		if b.TotalChapters > 0 {
			pct = b.ChaptersRead * 100 / b.TotalChapters
		}
		marker := " "
		if b.Slug == idx.CurrentBook {
			marker = "*"
		}
		fmt.Printf("%s %-35s %3d/%-3d %3d%% %s\n", marker, b.Slug, b.ChaptersRead, b.TotalChapters, pct, b.Category)
	}
}

func bookStart(slug, title, author string, total int, category string) {
	booksPath := booksPath()
	bookDir := filepath.Join(booksPath, slug)
	os.MkdirAll(filepath.Join(bookDir, "anotacoes"), 0755)

	chapters := make([]ChapterEntry, total)
	for i := 0; i < total; i++ {
		chapters[i] = ChapterEntry{
			Number: i + 1,
			Title:  fmt.Sprintf("Capítulo %d", i+1),
			Status: "pending",
		}
	}

	meta := BookMeta{
		Title:         title,
		Author:        author,
		Category:      category,
		TotalChapters: total,
		Chapters:      chapters,
	}
	saveBookMeta(slug, meta)

	idx, _ := loadBookIndex()
	idx.Books = append(idx.Books, BookEntry{
		Slug:          slug,
		Title:         title,
		Author:        author,
		TotalChapters: total,
		ChaptersRead:  0,
		Category:      category,
	})
	saveBookIndex(idx)

	fmt.Printf(`{"ok":true,"slug":"%s","title":"%s","total":%d,"category":"%s"}`, slug, title, total, category)
	fmt.Println()
}

func bookStatus() {
	slug, err := findActiveBook()
	if err != nil {
		fmt.Println(`{"error":"nenhum livro encontrado"}`)
		return
	}

	meta, err := loadBookMeta(slug)
	if err != nil {
		fmt.Println(`{"error":"meta.json nao encontrado"}`)
		return
	}

	var current, next ChapterEntry
	done := 0
	for _, c := range meta.Chapters {
		switch c.Status {
		case "done":
			done++
		case "reading":
			current = c
		}
		if current.Number == 0 && next.Number == 0 && c.Status == "pending" {
			next = c
		}
	}
	if current.Number == 0 {
		current = next
	}

	pct := 0
	if meta.TotalChapters > 0 {
		pct = done * 100 / meta.TotalChapters
	}

	out := map[string]interface{}{
		"slug":            slug,
		"title":           meta.Title,
		"author":          meta.Author,
		"category":        meta.Category,
		"total_chapters":  meta.TotalChapters,
		"chapters_read":   done,
		"progress_pct":    pct,
		"current_chapter": map[string]interface{}{"number": current.Number, "title": current.Title, "status": current.Status},
	}
	data, _ := json.MarshalIndent(out, "", "  ")
	fmt.Println(string(data))
}

func bookNote(chNumber int) {
	slug, err := findActiveBook()
	if err != nil {
		fmt.Println(`{"error":"nenhum livro encontrado"}`)
		return
	}

	meta, err := loadBookMeta(slug)
	if err != nil {
		fmt.Println(`{"error":"meta.json nao encontrado"}`)
		return
	}

	if chNumber < 1 || chNumber > meta.TotalChapters {
		fmt.Printf(`{"error":"capitulo %d fora do intervalo (1-%d)"}`, chNumber, meta.TotalChapters)
		fmt.Println()
		return
	}

	ch := meta.Chapters[chNumber-1]
	noteDir := filepath.Join(booksPath(), slug, "anotacoes")
	filename := fmt.Sprintf("%02d-%s.md", chNumber, slugify(ch.Title))
	notePath := filepath.Join(noteDir, filename)

	if _, err := os.Stat(notePath); err == nil {
		fmt.Printf(`{"error":"arquivo ja existe: %s"}`, filename)
		fmt.Println()
		return
	}

	template := getTemplate(meta.Category, chNumber, ch.Title)
	if err := os.WriteFile(notePath, []byte(template), 0644); err != nil {
		fmt.Printf(`{"error":"erro ao escrever arquivo: %v"}`, err)
		fmt.Println()
		return
	}

	if meta.Chapters[chNumber-1].Status == "pending" {
		meta.Chapters[chNumber-1].Status = "reading"
		saveBookMeta(slug, meta)
	}

	fmt.Printf(`{"ok":true,"file":"%s","chapter":%d,"title":"%s"}`, filename, chNumber, ch.Title)
	fmt.Println()
}

func bookDone(chNumber int) {
	slug, err := findActiveBook()
	if err != nil {
		fmt.Println(`{"error":"nenhum livro encontrado"}`)
		return
	}

	meta, err := loadBookMeta(slug)
	if err != nil {
		fmt.Println(`{"error":"meta.json nao encontrado"}`)
		return
	}

	if chNumber < 1 || chNumber > meta.TotalChapters {
		fmt.Printf(`{"error":"capitulo %d fora do intervalo (1-%d)"}`, chNumber, meta.TotalChapters)
		fmt.Println()
		return
	}

	today := time.Now().Format("2006-01-02")
	meta.Chapters[chNumber-1].Status = "done"
	meta.Chapters[chNumber-1].ReadAt = today
	saveBookMeta(slug, meta)

	done := 0
	for _, c := range meta.Chapters {
		if c.Status == "done" {
			done++
		}
	}

	idx, _ := loadBookIndex()
	for i, b := range idx.Books {
		if b.Slug == slug {
			idx.Books[i].ChaptersRead = done
			break
		}
	}
	saveBookIndex(idx)

	pct := 0
	if meta.TotalChapters > 0 {
		pct = done * 100 / meta.TotalChapters
	}

	fmt.Printf(`{"ok":true,"chapter":%d,"read":%d,"total":%d,"progress":%d}`, chNumber, done, meta.TotalChapters, pct)
	fmt.Println()
}

func bookReflect() {
	slug, err := findActiveBook()
	if err != nil {
		fmt.Println(`{"error":"nenhum livro encontrado"}`)
		return
	}

	meta, err := loadBookMeta(slug)
	if err != nil {
		fmt.Println(`{"error":"meta.json nao encontrado"}`)
		return
	}

	var recent []ChapterEntry
	for i := len(meta.Chapters) - 1; i >= 0 && len(recent) < 3; i-- {
		if meta.Chapters[i].Status == "done" {
			recent = append(recent, meta.Chapters[i])
		}
	}

	if len(recent) == 0 {
		fmt.Println("Nenhum capítulo concluído ainda.")
		return
	}

	fmt.Printf("=== %s (%s) ===\n\n", meta.Title, meta.Category)
	fmt.Println("Últimos capítulos lidos:")
	for i := len(recent) - 1; i >= 0; i-- {
		fmt.Printf("  %d. %s — lido em %s\n", recent[i].Number, recent[i].Title, recent[i].ReadAt)
	}

	fmt.Println()
	fmt.Println("Pergunta para reflexão:")
	fmt.Println(reflectQuestion(meta.Category))
}

func bookSwitch(slug string) {
	idx, err := loadBookIndex()
	if err != nil {
		fmt.Println(`{"error":"index.json nao encontrado"}`)
		return
	}
	found := false
	for _, b := range idx.Books {
		if b.Slug == slug {
			found = true
			break
		}
	}
	if !found {
		fmt.Printf(`{"error":"livro '%s' nao encontrado"}`, slug)
		fmt.Println()
		return
	}
	idx.CurrentBook = slug
	saveBookIndex(idx)
	fmt.Printf(`{"ok":true,"current_book":"%s"}`, slug)
	fmt.Println()
}

func bookUpdate(slug, field, value string) {
	idx, err := loadBookIndex()
	if err != nil {
		fmt.Println(`{"error":"index.json nao encontrado"}`)
		return
	}

	found := false
	for i, b := range idx.Books {
		if b.Slug == slug {
			switch field {
			case "title":
				idx.Books[i].Title = value
			case "author":
				idx.Books[i].Author = value
			case "total":
				n, _ := strconv.Atoi(value)
				idx.Books[i].TotalChapters = n
			case "category":
				idx.Books[i].Category = value
			default:
				fmt.Printf(`{"error":"campo desconhecido: %s"}`, field)
				fmt.Println()
				return
			}
			found = true
			break
		}
	}

	if !found {
		fmt.Printf(`{"error":"livro '%s' nao encontrado"}`, slug)
		fmt.Println()
		return
	}

	saveBookIndex(idx)

	meta, err := loadBookMeta(slug)
	if err == nil {
		switch field {
		case "title":
			meta.Title = value
		case "author":
			meta.Author = value
		case "category":
			meta.Category = value
		case "total":
			n, _ := strconv.Atoi(value)
			if n != meta.TotalChapters {
				chapters := make([]ChapterEntry, n)
				for j := 0; j < n; j++ {
					chapters[j] = ChapterEntry{Number: j + 1, Title: fmt.Sprintf("Capítulo %d", j+1), Status: "pending"}
				}
				meta.TotalChapters = n
				meta.Chapters = chapters
			}
		}
		saveBookMeta(slug, meta)
	}

	fmt.Println(`{"ok":true}`)
}

func getTemplate(category string, number int, title string) string {
	header := fmt.Sprintf("# %d — %s\n\n> Lido em: [data]\n\n", number, title)

	var body string
	switch category {
	case "technical":
		body = `## Ideias principais
1. 
2. 
3. 

## O que me surpreendeu


## Como aplicar no meu trabalho


## Conexões com tracks de exercícios


## Perguntas que ficaram

- 
- 
`
	case "philosophy":
		body = `## Tese central do capítulo


## Argumentos que o autor usa


## Com o que eu concordo


## Com o que eu discordo (e por quê)


## Conexões com outras leituras


## Perguntas que ficaram

- 
- 
`
	case "fiction":
		body = `## O que aconteceu


## Personagens e motivações


## Temas que o autor explora


## Passagem que mais me impactou


## Previsão: o que acontece a seguir?


`
	case "history":
		body = `## Fatos principais


## Contexto da época


## Viés do autor (o que ele omite?)


## Paralelo com o presente


## Perguntas que ficaram

- 
- 
`
	case "self-help":
		body = `## Ideia principal


## O que eu já faço (ou não)


## Uma ação concreta que vou testar


## Ceticismo: o que me parece exagerado?


## Perguntas que ficaram

- 
- 
`
	case "biography":
		body = `## Período da vida do biografado


## Decisão crucial que ele(a) tomou


## O que eu faria diferente no lugar dele(a)?


## Lição que levo para minha vida


## Perguntas que ficaram

- 
- 
`
	default:
		body = `## Ideias principais


## O que me surpreendeu


## Conexões com coisas que já sei


## Perguntas que ficaram

- 
- 
`
	}
	return header + body
}

func reflectQuestion(category string) string {
	questions := map[string]string{
		"technical":  "Das ideias dos últimos 3 capítulos, qual você já consegue ensinar para alguém? Descreva como ensinaria.",
		"philosophy": "Se o autor estivesse aqui agora, qual objeção você levantaria contra o argumento dele? O que ele responderia?",
		"fiction":    "Qual personagem dos últimos capítulos você teria tomado uma decisão diferente? Qual seria a consequência?",
		"history":    "Se os eventos desses capítulos acontecessem hoje, o desfecho seria diferente? Por quê?",
		"self-help":  "Das ações que você anotou nos últimos capítulos, qual você de fato testou? O resultado foi o esperado?",
		"biography":  "Qual traço de personalidade do biografado foi determinante nesses capítulos? Você compartilha desse traço?",
	}
	if q, ok := questions[category]; ok {
		return q
	}
	return "Qual a ideia mais útil ou instigante dos últimos capítulos? Por quê?"
}

func slugify(s string) string {
	s = strings.ToLower(s)
	s = strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' {
			return r
		}
		if r == ' ' || r == '_' {
			return '-'
		}
		return -1
	}, s)
	return strings.Trim(s, "-")
}
