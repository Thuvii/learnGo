# LearnGo

This repository contains Go projects built to strengthen my understanding of core Go concepts including concurrency, HTTP requests, JSON handling, file I/O, and structured project design.

## Projects

### 1. Pokedex

A command-line Pokedex application that interacts with a public Pokémon API.

**Features:**

* Fetch Pokémon data from an external API
* Parse and decode JSON responses
* Display structured Pokémon information in the terminal
* Practice working with HTTP requests and error handling

**Concepts Used:**

* `net/http`
* `encoding/json`
* Struct design and data modeling
* CLI interaction
* Basic state management

---

### 2. Web Crawler

A concurrent web crawler that scans web pages and extracts links.

**Features:**

* Crawl web pages starting from a seed URL
* Extract and normalize links
* Avoid duplicate visits
* Write structured output to a file

**Concepts Used:**

* Goroutines and concurrency
* Mutexes / synchronization
* Maps and slices
* Sorting and structured reporting
* File writing and JSON output

---

## Goals of This Repository

* Strengthen understanding of Go fundamentals
* Practice building small but complete CLI applications
* Improve code organization and readability
* Apply structured problem-solving in real projects

---

## How to Run

Navigate into a project folder:

```
cd pokedex
go run .
```

or

```
cd crawler
go run .
```

