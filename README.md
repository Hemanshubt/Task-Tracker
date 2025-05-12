# 📝 Task Tracker CLI (Golang)

A simple command-line task manager built with Go. This CLI helps you manage your to-dos, track their status, and stores tasks in a local JSON file.

---

## 📌 Features

- Add new tasks
- Update existing tasks
- Delete tasks
- Mark tasks as:
  - `todo` (default)
  - `in-progress`
  - `done`
- List tasks:
  - All tasks
  - Filtered by status (todo, in-progress, done)
- Tasks stored persistently in a `task.json` file

---

## 🚀 Getting Started

### ✅ Prerequisites

- Go installed on your system  
  [Download Go](https://go.dev/dl/)

### 📁 Build the Project

```bash
go build -o task-cli
