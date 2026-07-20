package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func initAllCmd() {
	fmt.Println("gerando roadmaps.json...")
	writeGoBackend()
	writeDockerDevops()
	writeRedesPratica()
	writeLinuxSystems()
	writeTypeScript()
	writeDSA()
	fmt.Println("gerando corrections.json...")
	writeCorrections("go-backend", []CorrectionEntry{
		{Challenge: "A2", Category: "errors-wrapping", Error: "Uso de == em vez de errors.Is para comparar erro wrappeado"},
		{Challenge: "A3", Category: "nil-pointer", Error: "Acesso a campo de node nil sem verificacao previa"},
		{Challenge: "B4", Category: "concurrency-mutex", Error: "Mutex copiado por valor ao passar struct para funcao"},
	})
	writeCorrections("docker-devops", []CorrectionEntry{
		{Challenge: "D1", Category: "docker-basics", Error: "Confusao entre docker run e docker start — run cria container novo"},
		{Challenge: "D3", Category: "dockerfile-order", Error: "Layer de COPY antes de RUN apt-get — cache invalida desnecessariamente"},
	})
	writeCorrections("redes-pratica", nil)
	writeCorrections("linux-systems", nil)
	writeCorrections("typescript", nil)
	writeCorrections("dsa", nil)
	fmt.Println("gerando schedule.json...")
	s := Schedule{ActiveTrack: "redes-pratica", Reviews: []Review{}}
	writeJSON(basePath()+"/spaced-repetition/schedule.json", s)
	fmt.Println("pronto.")
}

func writeCorrections(track string, entries []CorrectionEntry) {
	if entries == nil {
		entries = []CorrectionEntry{}
	}
	writeJSON(basePath()+"/framework/tracks/"+track+"/corrections.json", Corrections{Entries: entries})
}

func writeJSON(path string, v interface{}) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Fprintf(osStderr, "erro marshal %s: %v\n", path, err)
		return
	}
	if err := os.WriteFile(path, data, 0644); err != nil {
		fmt.Fprintf(osStderr, "erro salvando %s: %v\n", path, err)
	}
}

func writeGoBackend() {
	r := Roadmap{
		Track: "go-backend",
		Modules: []Module{
			{ID: "modulo-A-fundamentos", Scaffolding: "alto", Challenges: []Challenge{
				{ID: "A1", Name: "currency-conversor", Template: "implementacao", Concepts: []string{"CLI", "sentinel errors", "errors.Is"}, Status: "done"},
				{ID: "A2", Name: "bank-account", Template: "implementacao", Concepts: []string{"pointer receiver", "erro tipado", "errors.As"}, Status: "done"},
				{ID: "A3", Name: "linked-list", Template: "implementacao", Concepts: []string{"ponteiros", "nil seguro", "Remove/Reverse"}, Status: "done"},
				{ID: "A4", Name: "shape-interface", Template: "implementacao", Concepts: []string{"interface implicita", "type switch"}, Status: "done"},
			}},
			{ID: "modulo-B-concorrencia", Scaffolding: "alto", Challenges: []Challenge{
				{ID: "B1", Name: "worker-pool", Template: "implementacao", Concepts: []string{"goroutines", "channels", "WaitGroup", "context"}, Status: "done"},
				{ID: "B2", Name: "parallel-query", Template: "debug", Concepts: []string{"fan-in simples", "time.Duration"}, Status: "done"},
				{ID: "B3", Name: "fan-in", Template: "implementacao", Concepts: []string{"select multi-canal", "merger", "graceful shutdown"}, Status: "done"},
				{ID: "B4", Name: "rate-limiter", Template: "implementacao", Concepts: []string{"sync.Mutex", "token bucket"}, Status: "done"},
				{ID: "B5", Name: "select-sem-default", Template: "otimizacao", Concepts: []string{"select sem default", "ctx.Done()"}, Status: "done"},
			}},
			{ID: "modulo-C-estruturas", Scaffolding: "medio", Challenges: []Challenge{
				{ID: "C1", Name: "stack-queue", Template: "explicacao", Concepts: []string{"thread-safe", "pub/priv split"}, Status: "done"},
				{ID: "C2", Name: "io-reader-writer", Template: "implementacao", Concepts: []string{"io.Reader/Writer", "delegacao", "io.Copy"}, Status: "done"},
				{ID: "C3", Name: "bst", Template: "implementacao", Concepts: []string{"arvore binaria", "recursao", "Delete 3 casos"}, Status: "done"},
				{ID: "C4", Name: "error-is-as", Template: "debug", Concepts: []string{"errors.Is vs errors.As", "%w wrapping"}, Status: "done"},
				{ID: "C5", Name: "slice-leak", Template: "debug", Concepts: []string{"backing array", "memory leak"}, Status: "done"},
			}},
			{ID: "modulo-D-http-apis", Scaffolding: "medio", Challenges: []Challenge{
				{ID: "D1", Name: "cache-ttl", Template: "design", Concepts: []string{"RWMutex", "lazy eviction", "cleanup goroutine"}, Status: "done"},
				{ID: "D2", Name: "products-api", Template: "implementacao", Concepts: []string{"HTTP CRUD", "ServeMux", "middleware"}, Status: "done"},
				{ID: "D3", Name: "middleware-chain", Template: "implementacao", Concepts: []string{"Logger", "Recoverer", "Auth", "CORS"}, Status: "done"},
			}},
			{ID: "modulo-E-armadilhas", Scaffolding: "baixo", Challenges: []Challenge{
				{ID: "E1", Name: "nil-interface", Template: "debug", Concepts: []string{"interface (type, value) pair", "nil pointer vs nil interface"}, Status: "done"},
				{ID: "E2", Name: "nil-interface-revisao", Template: "explicacao", Concepts: []string{"revisao aprofundada"}, Status: "done"},
			}},
		},
		MixedPractice: []Challenge{
			{ID: "MP1", Name: "escolha-ferramentas", Template: "mixed-practice", Status: "pending"},
			{ID: "MP2", Name: "padrao-concorrencia", Template: "mixed-practice", Status: "pending"},
			{ID: "MP3", Name: "estrutura-certa", Template: "mixed-practice", Status: "pending"},
			{ID: "MP4", Name: "debugging-apis", Template: "mixed-practice", Status: "pending"},
		},
		Capstones: []Challenge{
			{ID: "C1", Name: "calculadora", Template: "capstone", Concepts: []string{"A1", "A2", "A3", "A4"}, Status: "pending"},
			{ID: "C2", Name: "crawler", Template: "capstone", Concepts: []string{"B1", "B2", "B3", "B4", "B5"}, Status: "pending"},
			{ID: "C3", Name: "cache-persistente", Template: "capstone", Concepts: []string{"C1", "C2", "C3", "C4", "C5"}, Status: "pending"},
			{ID: "C4", Name: "api-blog", Template: "capstone", Concepts: []string{"D1", "D2", "D3"}, Status: "pending"},
			{ID: "C5", Name: "auditoria", Template: "capstone", Concepts: []string{"E1", "E2", "C5", "B5", "C4"}, Status: "pending"},
		},
	}
	writeJSON(basePath()+"/framework/tracks/go-backend/roadmap.json", r)
}

func writeDockerDevops() {
	r := Roadmap{
		Track: "docker-devops",
		Modules: []Module{
			{ID: "modulo-1-containers", Scaffolding: "alto", Challenges: []Challenge{
				{ID: "D1", Name: "hello-world", Template: "implementacao", Concepts: []string{"docker run", "imagem vs container", "docker ps"}, Status: "done"},
				{ID: "D2", Name: "interactive-shell", Template: "implementacao", Concepts: []string{"-it flags", "docker exec", "PID 1"}, Status: "done"},
				{ID: "D3", Name: "first-dockerfile", Template: "implementacao", Concepts: []string{"Dockerfile", "FROM", "COPY", "CMD"}, Status: "done"},
				{ID: "D4", Name: "entrypoint-cmd", Template: "implementacao", Concepts: []string{"ENTRYPOINT", "argumentos em runtime"}, Status: "done"},
			}},
			{ID: "modulo-2-build", Scaffolding: "alto", Challenges: []Challenge{
				{ID: "D5", Name: "dockerignore-layers", Template: "otimizacao", Concepts: []string{".dockerignore", "ordem de layers", "cache"}, Status: "done"},
				{ID: "D6", Name: "multi-stage", Template: "implementacao", Concepts: []string{"multi-stage builds", "reducao de tamanho"}, Status: "pending"},
				{ID: "D7", Name: "healthcheck", Template: "implementacao", Concepts: []string{"HEALTHCHECK", "restart policies"}, Status: "pending"},
			}},
			{ID: "modulo-3-persistencia", Scaffolding: "medio", Challenges: []Challenge{
				{ID: "D8", Name: "volumes-bind", Template: "implementacao", Concepts: []string{"bind mount", "named volume", "persistencia"}, Status: "done"},
				{ID: "D9", Name: "docker-networks", Template: "implementacao", Concepts: []string{"bridge", "DNS interno"}, Status: "pending"},
				{ID: "D10", Name: "docker-compose", Template: "implementacao", Concepts: []string{"docker-compose.yml", "services"}, Status: "pending"},
				{ID: "D11", Name: "compose-multi", Template: "design", Concepts: []string{"multiplos servicos", "volumes", "networks"}, Status: "pending"},
			}},
		},
		MixedPractice: []Challenge{
			{ID: "MP1", Name: "diagnostico-containers", Template: "mixed-practice", Status: "pending"},
			{ID: "MP2", Name: "otimizacao-dockerfile", Template: "mixed-practice", Status: "pending"},
		},
		Capstones: []Challenge{
			{ID: "C1", Name: "container-multi", Template: "capstone", Concepts: []string{"D1", "D2", "D3", "D4"}, Status: "pending"},
			{ID: "C2", Name: "imagem-producao", Template: "capstone", Concepts: []string{"D5", "D6", "D7"}, Status: "pending"},
			{ID: "C3", Name: "stack-completa", Template: "capstone", Concepts: []string{"D8", "D9", "D10", "D11"}, Status: "pending"},
		},
	}
	writeJSON(basePath()+"/framework/tracks/docker-devops/roadmap.json", r)
}

func writeRedesPratica() {
	r := Roadmap{
		Track: "redes-pratica",
		Modules: []Module{
			{ID: "modulo-1-sockets", Scaffolding: "alto", Challenges: []Challenge{
				{ID: "R1", Name: "echo-server", Template: "implementacao", Concepts: []string{"net.Listen", "net.Dial", "TCP handshake"}, Status: "pending"},
				{ID: "R2", Name: "chat-tcp", Template: "implementacao", Concepts: []string{"goroutines por conexao", "broadcast"}, Status: "pending"},
				{ID: "R3", Name: "udp-echo", Template: "implementacao", Concepts: []string{"net.ListenUDP", "datagramas"}, Status: "pending"},
				{ID: "R4", Name: "timeout-retry", Template: "otimizacao", Concepts: []string{"deadline", "retry", "buffer sizing"}, Status: "pending"},
			}},
			{ID: "modulo-2-http", Scaffolding: "alto", Challenges: []Challenge{
				{ID: "R5", Name: "http-parser", Template: "implementacao", Concepts: []string{"parse HTTP/1.1", "method", "headers", "body"}, Status: "pending"},
				{ID: "R6", Name: "http-server", Template: "implementacao", Concepts: []string{"servidor HTTP", "resposta GET", "rotas"}, Status: "pending"},
				{ID: "R7", Name: "keep-alive", Template: "design", Concepts: []string{"conexoes persistentes"}, Status: "pending"},
				{ID: "R8", Name: "chunked", Template: "implementacao", Concepts: []string{"Transfer-Encoding: chunked"}, Status: "pending"},
			}},
			{ID: "modulo-3-dns", Scaffolding: "medio", Challenges: []Challenge{
				{ID: "R9", Name: "dns-wire", Template: "implementacao", Concepts: []string{"query DNS", "wire format"}, Status: "pending"},
				{ID: "R10", Name: "dns-resolver", Template: "implementacao", Concepts: []string{"resolucao iterativa", "root → TLD"}, Status: "pending"},
				{ID: "R11", Name: "dns-cache", Template: "design", Concepts: []string{"cache respostas", "TTL"}, Status: "pending"},
			}},
			{ID: "modulo-4-tls", Scaffolding: "medio", Challenges: []Challenge{
				{ID: "R12", Name: "certs", Template: "explicacao", Concepts: []string{"X.509", "CA root", "cadeia confianca"}, Status: "pending"},
				{ID: "R13", Name: "tls-server", Template: "implementacao", Concepts: []string{"crypto/tls", "certificado"}, Status: "pending"},
				{ID: "R14", Name: "tls-mitm", Template: "debug", Concepts: []string{"proxy MITM", "certificado on-the-fly"}, Status: "pending"},
			}},
		},
		MixedPractice: []Challenge{
			{ID: "MP1", Name: "tcp-vs-udp", Template: "mixed-practice", Status: "pending"},
			{ID: "MP2", Name: "cliente-http", Template: "mixed-practice", Status: "pending"},
			{ID: "MP3", Name: "debugging-dns", Template: "mixed-practice", Status: "pending"},
		},
		Capstones: []Challenge{
			{ID: "C1", Name: "proxy-http", Template: "capstone", Status: "pending"},
			{ID: "C2", Name: "netcat-tls", Template: "capstone", Status: "pending"},
		},
	}
	writeJSON(basePath()+"/framework/tracks/redes-pratica/roadmap.json", r)
}

func writeLinuxSystems() {
	r := Roadmap{
		Track: "linux-systems",
		Modules: []Module{
			{ID: "modulo-1-processos", Scaffolding: "alto", Challenges: []Challenge{
				{ID: "L1", Name: "spawn-wait", Template: "implementacao", Concepts: []string{"spawn processo", "wait", "exit code"}, Status: "pending"},
				{ID: "L2", Name: "signals", Template: "implementacao", Concepts: []string{"SIGTERM", "SIGINT", "graceful shutdown"}, Status: "pending"},
				{ID: "L3", Name: "daemon", Template: "design", Concepts: []string{"fork", "setsid", "chdir"}, Status: "pending"},
				{ID: "L4", Name: "zombies", Template: "debug", Concepts: []string{"zumbis", "wait", "PPID"}, Status: "pending"},
			}},
			{ID: "modulo-2-fd-io", Scaffolding: "alto", Challenges: []Challenge{
				{ID: "L5", Name: "pipes", Template: "implementacao", Concepts: []string{"pipe", "pai↔filho"}, Status: "pending"},
				{ID: "L6", Name: "redirection", Template: "implementacao", Concepts: []string{"dup2", "stdin/stdout/stderr"}, Status: "pending"},
				{ID: "L7", Name: "poll-select", Template: "design", Concepts: []string{"select", "epoll", "FDs multiplos"}, Status: "pending"},
			}},
			{ID: "modulo-3-filesystem", Scaffolding: "medio", Challenges: []Challenge{
				{ID: "L8", Name: "walk", Template: "implementacao", Concepts: []string{"walk recursivo", "permissoes", "tipos"}, Status: "pending"},
				{ID: "L9", Name: "proc", Template: "explicacao", Concepts: []string{"/proc", "PID", "mem", "fds"}, Status: "pending"},
				{ID: "L10", Name: "inode", Template: "debug", Concepts: []string{"inodes", "hardlink vs symlink"}, Status: "pending"},
			}},
			{ID: "modulo-4-mini-shell", Scaffolding: "baixo", Challenges: []Challenge{
				{ID: "L11", Name: "parser", Template: "implementacao", Concepts: []string{"tokenizer", "AST simples"}, Status: "pending"},
				{ID: "L12", Name: "execution", Template: "implementacao", Concepts: []string{"PATH lookup", "builtins"}, Status: "pending"},
				{ID: "L13", Name: "pipes-shell", Template: "implementacao", Concepts: []string{"pipes no shell"}, Status: "pending"},
				{ID: "L14", Name: "redirect-shell", Template: "implementacao", Concepts: []string{">", "<", "2>", "2>&1"}, Status: "pending"},
			}},
		},
		MixedPractice: []Challenge{
			{ID: "MP1", Name: "gerenciamento-processos", Template: "mixed-practice", Status: "pending"},
			{ID: "MP2", Name: "tee-tail", Template: "mixed-practice", Status: "pending"},
			{ID: "MP3", Name: "ferramenta-find", Template: "mixed-practice", Status: "pending"},
		},
		Capstones: []Challenge{
			{ID: "C1", Name: "tee-tail", Template: "capstone", Status: "pending"},
			{ID: "C2", Name: "shell", Template: "capstone", Status: "pending"},
		},
	}
	writeJSON(basePath()+"/framework/tracks/linux-systems/roadmap.json", r)
}

func writeTypeScript() {
	r := Roadmap{
		Track: "typescript",
		Modules: []Module{
			{ID: "modulo-1-tipos", Scaffolding: "alto", Challenges: []Challenge{
				{ID: "TS1", Name: "generics-constraints", Template: "implementacao", Concepts: []string{"extends", "keyof", "infer"}, Status: "pending"},
				{ID: "TS2", Name: "conditional-types", Template: "implementacao", Concepts: []string{"conditional types", "never", "distributive"}, Status: "pending"},
				{ID: "TS3", Name: "mapped-types", Template: "implementacao", Concepts: []string{"mapped types", "template literals", "as clause"}, Status: "pending"},
				{ID: "TS4", Name: "utility-types", Template: "explicacao", Concepts: []string{"Partial", "Required", "Pick", "Omit", "Readonly"}, Status: "pending"},
			}},
			{ID: "modulo-2-padroes", Scaffolding: "medio", Challenges: []Challenge{
				{ID: "TS5", Name: "discriminated-unions", Template: "implementacao", Concepts: []string{"union discriminada", "never check", "exhaustive"}, Status: "pending"},
				{ID: "TS6", Name: "branded-types", Template: "design", Concepts: []string{"opaque types", "type branding", "nominal"}, Status: "pending"},
				{ID: "TS7", Name: "builder-pattern", Template: "implementacao", Concepts: []string{"type-safe builder", "fluent API", "chaining"}, Status: "pending"},
				{ID: "TS8", Name: "result-either", Template: "implementacao", Concepts: []string{"Result type", "Either", "match", "unwrap"}, Status: "pending"},
			}},
			{ID: "modulo-3-type-level", Scaffolding: "baixo", Challenges: []Challenge{
				{ID: "TS9", Name: "type-challenges-1", Template: "otimizacao", Concepts: []string{"type-challenges easy/medium", "recursive types"}, Status: "pending"},
				{ID: "TS10", Name: "type-challenges-2", Template: "debug", Concepts: []string{"type-challenges hard", "conditional recursion"}, Status: "pending"},
				{ID: "TS11", Name: "sdk-typed", Template: "design", Concepts: []string{"SDK design", "user-facing types", "DX"}, Status: "pending"},
			}},
		},
		MixedPractice: []Challenge{
			{ID: "MP1", Name: "tipos-avancados", Template: "mixed-practice", Status: "pending"},
			{ID: "MP2", Name: "padroes-reais", Template: "mixed-practice", Status: "pending"},
		},
		Capstones: []Challenge{
			{ID: "C1", Name: "sdk-fortemente-tipada", Template: "capstone", Concepts: []string{"TS1-TS11"}, Status: "pending"},
		},
	}
	writeJSON(basePath()+"/framework/tracks/typescript/roadmap.json", r)
}

func writeDSA() {
	r := Roadmap{
		Track: "dsa",
		Modules: []Module{
			{ID: "modulo-1-arrays", Scaffolding: "alto", Challenges: []Challenge{
				{ID: "D1", Name: "two-pointers", Template: "implementacao", Concepts: []string{"two pointers", "in-place", "O(n)"}, Status: "pending"},
				{ID: "D2", Name: "sliding-window", Template: "implementacao", Concepts: []string{"janela deslizante", "subarray", "O(n)"}, Status: "pending"},
				{ID: "D3", Name: "prefix-sum", Template: "implementacao", Concepts: []string{"prefix sum array", "range query O(1)"}, Status: "pending"},
				{ID: "D4", Name: "string-search", Template: "otimizacao", Concepts: []string{"busca em string", "KMP intro"}, Status: "pending"},
			}},
			{ID: "modulo-2-estruturas", Scaffolding: "alto", Challenges: []Challenge{
				{ID: "D5", Name: "linked-list-full", Template: "implementacao", Concepts: []string{"singly/doubly", "reverse", "merge", "detect cycle"}, Status: "pending"},
				{ID: "D6", Name: "stack-queue-impl", Template: "implementacao", Concepts: []string{"array-based", "linked-based", "thread-safe"}, Status: "pending"},
				{ID: "D7", Name: "deque", Template: "implementacao", Concepts: []string{"double-ended queue", "circular buffer"}, Status: "pending"},
				{ID: "D8", Name: "lru-cache", Template: "design", Concepts: []string{"LRU", "doubly linked list", "hash map", "O(1)"}, Status: "pending"},
			}},
			{ID: "modulo-3-arvores", Scaffolding: "medio", Challenges: []Challenge{
				{ID: "D9", Name: "bst-full", Template: "implementacao", Concepts: []string{"insert", "search", "delete 3 casos", "successor"}, Status: "pending"},
				{ID: "D10", Name: "tree-traversals", Template: "implementacao", Concepts: []string{"pre/in/post", "BFS/DFS", "iterativo", "recursivo"}, Status: "pending"},
				{ID: "D11", Name: "heap", Template: "implementacao", Concepts: []string{"min-heap", "heapify", "heap sort"}, Status: "pending"},
				{ID: "D12", Name: "trie", Template: "implementacao", Concepts: []string{"trie", "insert", "search", "prefix"}, Status: "pending"},
			}},
			{ID: "modulo-4-grafos", Scaffolding: "medio", Challenges: []Challenge{
				{ID: "D13", Name: "graph-repr", Template: "implementacao", Concepts: []string{"adjacency list", "adjacency matrix", "edge list"}, Status: "pending"},
				{ID: "D14", Name: "bfs-dfs", Template: "implementacao", Concepts: []string{"BFS iterativo", "DFS recursive/iterative", "connected components"}, Status: "pending"},
				{ID: "D15", Name: "dijkstra", Template: "implementacao", Concepts: []string{"shortest path", "priority queue", "greedy"}, Status: "pending"},
				{ID: "D16", Name: "topological-sort", Template: "implementacao", Concepts: []string{"Kahn algorithm", "DFS-based", "DAG"}, Status: "pending"},
			}},
			{ID: "modulo-5-algoritmos", Scaffolding: "baixo", Challenges: []Challenge{
				{ID: "D17", Name: "sorting", Template: "implementacao", Concepts: []string{"merge sort", "quick sort", "heap sort", "complexity"}, Status: "pending"},
				{ID: "D18", Name: "binary-search", Template: "otimizacao", Concepts: []string{"lower/upper bound", "rotated array", "binary answer"}, Status: "pending"},
				{ID: "D19", Name: "backtracking", Template: "implementacao", Concepts: []string{"n-queens", "subsets", "permutations", "pruning"}, Status: "pending"},
				{ID: "D20", Name: "dp-intro", Template: "explicacao", Concepts: []string{"memoization", "tabulation", "coin change", "knapsack"}, Status: "pending"},
			}},
		},
		MixedPractice: []Challenge{
			{ID: "MP1", Name: "estrutura-certa", Template: "mixed-practice", Status: "pending"},
			{ID: "MP2", Name: "algoritmo-certo", Template: "mixed-practice", Status: "pending"},
		},
		Capstones: []Challenge{
			{ID: "C1", Name: "cache-persistente", Template: "capstone", Concepts: []string{"LRU", "heap", "BST"}, Status: "pending"},
			{ID: "C2", Name: "roteador-caminhos", Template: "capstone", Concepts: []string{"graph", "Dijkstra", "trie"}, Status: "pending"},
		},
	}
	writeJSON(basePath()+"/framework/tracks/dsa/roadmap.json", r)
}
