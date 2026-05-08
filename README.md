# 🔄 go-reloaded

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go&logoColor=white)
![Algorithm](https://img.shields.io/badge/Algorithm-15%25-FF6B6B?style=flat&logo=thealgorithms&logoColor=white)
![Standard Library](https://img.shields.io/badge/Dependencies-Standard%20Library%20Only-00ADD8?style=flat&logo=go&logoColor=white)
![Status](https://img.shields.io/badge/Status-Completed-brightgreen?style=flat)

> A simple yet powerful text completion, editing, and auto-correction tool built in Go — designed to transform and reformat text files using a set of intuitive inline commands.

---

## 📋 Table of Contents

* [About the Project](#-about-the-project)
* [Features](#-features)
* [Tech Stack](#-tech-stack)
* [Getting Started](#-getting-started)
  * [Prerequisites](#prerequisites)
  * [Installation](#installation)
  * [How to Run](#how-to-run)
* [Usage](#-usage)
* [Transformation Rules](#-transformation-rules)
* [Allowed Packages](#-allowed-packages)
* [Learning Outcomes](#-learning-outcomes)
* [Authors](#-authors)

---

## 📖 About the Project

**go-reloaded** is a CLI-based text processing tool that reads an input file, applies a series of text transformations based on inline commands, and writes the corrected output to a new file.

It acts like a lightweight auto-correction engine — handling case conversion, number base conversion, punctuation formatting, quotation cleanup, and grammar correction, all driven by special markers embedded directly in the text.

**Why it was built:**

* To practice string and number manipulation in Go
* To work with Go's file system (`fs`) API
* To build a practical, real-world text transformation pipeline
* To reinforce algorithmic thinking through parsing and processing logic

---

## ✨ Features

* **Hex & Binary Conversion** — inline `(hex)` and `(bin)` markers convert numbers to decimal automatically.
* **Case Transformation** — `(up)`, `(low)`, `(cap)` modify the case of preceding words.
* **Bulk Case Transformation** — `(up, N)`, `(low, N)`, `(cap, N)` apply transformations to N preceding words.
* **Punctuation Formatting** — automatically fixes spacing around `. , ! ? : ;` including grouped punctuation like `...`
* **Quotation Mark Cleanup** — single quotes are placed correctly around enclosed words with no extra spaces.
* **Article Correction** — automatically changes `a` to `an` before words starting with a vowel or `h`.
* **File-Based I/O** — reads from an input file and writes the transformed result to an output file.
* **No External Dependencies** — built strictly with Go's standard library.

---

## 🛠 Tech Stack

| Technology     | Role                | Details                                       |
| -------------- | ------------------- | --------------------------------------------- |
| **Go**         | Core Language       | File I/O, string processing, number parsing   |
| **Algorithms** | Logic               | Parsing, transformation pipelines, edge cases |

---

## 📁 Project Structure

```
```

---

## 🚀 Getting Started

### Prerequisites

Make sure you have installed:

* **Go 1.21+** — [https://golang.org/dl/](https://golang.org/dl/)

Verify Go installation:

```bash
go version
```

---

### Installation

**1. Clone the repository:**

```bash
git clone https://acad.learn2earn.ng/git/lajose/go-reloaded.git
```

**2. Navigate into the project:**

```bash
cd go-reloaded
```

---

### How to Run

**Run the program with an input and output file:**

```bash
go run . <input_file> <output_file>
```

**Example:**

```bash
go run . sample.txt result.txt
```

---

## 💻 Usage

**1. Create or edit your input file:**

```bash
cat sample.txt
```

```
it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom,
it was the age of foolishness (cap, 6) , IT WAS THE (low, 3) winter of despair.
```

**2. Run the tool:**

```bash
go run . sample.txt result.txt
```

**3. View the output:**

```bash
cat result.txt
```

```
It was the best of times, it was the worst of TIMES, it was the age of wisdom,
It Was The Age Of Foolishness, it was the winter of despair.
```

---

## 🔧 Transformation Rules

### Number Conversions

| Marker  | Action                                      | Example                            |
| ------- | ------------------------------------------- | ---------------------------------- |
| `(hex)` | Converts preceding hex number to decimal    | `"1E (hex) files"` → `"30 files"` |
| `(bin)` | Converts preceding binary number to decimal | `"10 (bin) years"` → `"2 years"`  |

### Case Transformations

| Marker     | Action                            | Example                                                   |
| ---------- | --------------------------------- | --------------------------------------------------------- |
| `(up)`     | Uppercases the preceding word     | `"go (up)"` → `"GO"`                                     |
| `(low)`    | Lowercases the preceding word     | `"SHOUTING (low)"` → `"shouting"`                        |
| `(cap)`    | Capitalizes the preceding word    | `"bridge (cap)"` → `"Bridge"`                            |
| `(up, N)`  | Uppercases the preceding N words  | `"so exciting (up, 2)"` → `"SO EXCITING"`                |
| `(low, N)` | Lowercases the preceding N words  | `"IT WAS THE (low, 3)"` → `"it was the"`                 |
| `(cap, N)` | Capitalizes the preceding N words | `"age of foolishness (cap, 6)"` → `"Age Of Foolishness"` |

### Punctuation Formatting

```
"I was sitting over there ,and then BAMM !!"  →  "I was sitting over there, and then BAMM!!"
"I was thinking ... You were right"            →  "I was thinking... You were right"
```

### Quotation Marks

```
"I am exactly how they describe me: ' awesome '"
→ "I am exactly how they describe me: 'awesome'"

"As Elton John said: ' I am the most well-known homosexual in the world '"
→ "As Elton John said: 'I am the most well-known homosexual in the world'"
```

### Article Correction

```
"There it was. A amazing rock!"        →  "There it was. An amazing rock!"
"bearing a untold story inside you."   →  "bearing an untold story inside you."
```

---

## 🧪 Testing

Unit tests are included to validate each transformation individually.

```bash
go test ./...
```

---

## 📦 Allowed Packages

Only Go **standard library** is used:

```
os
fmt
strings
strconv
regexp
log
```

No third-party dependencies.

---

## 🎓 Learning Outcomes

This project helps you understand:

* **Go File System API**
  * Reading and writing files with `os`
* **String Manipulation in Go**
  * Parsing, splitting, replacing, and reformatting text
* **Number Base Conversion**
  * Hexadecimal and binary to decimal using `strconv`
* **Algorithmic Thinking**
  * Building a multi-step text transformation pipeline
* **Unit Testing in Go**
  * Writing and running tests with the `testing` package

---

## 👥 Authors

**Ajose Lawal**
📎 [https://acad.learn2earn.ng/git/lajose](https://acad.learn2earn.ng/git/lajose)


---
