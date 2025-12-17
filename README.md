# Exploring Golang: A Monorepo Workspace

This repository serves as a learning and exploration space for various concepts and idioms in the Go programming language (Golang).
I am working with Gemini AI to explore concepts of Golang - syntax, concurrency, etc...  From AI discussions, we have summed up
Go as "C with garbage collection and concurrency" - some language features are present to prevent stupid mistakes, but not all
the classic gotchas are filled in, and thus a fair amount of discipline is still required (but not as much as raw C).

## Project Structure and Go Modules

This project uses a **Go Workspace**, which is a feature introduced in Go 1.18 to manage multiple modules within a single repository (a kind of monorepo).

| File/Directory | Purpose | Notes |
| :--- | :--- | :--- |
| **`go.work`** | **Workspace Definition** | Tells the Go toolchain to treat the listed directories as modules within this workspace. |
| **`greetings/`** | **Go Module** | Contains the first application module. |
| **`greetings/go.mod`** | **Module Dependency File** | Defines the module path (`maccergit/go-workspace/greetings`) and tracks external dependencies *for this specific module*. |
| **`greetings/greetings.go`** | **Executable Code** | Contains the `package main` executable and helper functions. |

The example shown here is an initial module developed with the AI - a more advanced "Hello World" that shows some easy Golang features.

## How to Run

1.  Clone the repository.
2.  Navigate to the project root: `cd go-workspace`
3.  Execute the main program using the Go toolchain:
    ```bash
    go run ./greetings <your-name>
    # Example: go run ./greetings John
    ```
### License
The items here are mine, created with the help of the free Gemini AI.  The license materials for this repo were provided by https://github.com/santisoler/cc-licenses:

Shield: [![CC BY-NC-SA 4.0][cc-by-nc-sa-shield]][cc-by-nc-sa]

This work is licensed under a
[Creative Commons Attribution-NonCommercial-ShareAlike 4.0 International License][cc-by-nc-sa].

[![CC BY-NC-SA 4.0][cc-by-nc-sa-image]][cc-by-nc-sa]

[cc-by-nc-sa]: http://creativecommons.org/licenses/by-nc-sa/4.0/
[cc-by-nc-sa-image]: https://licensebuttons.net/l/by-nc-sa/4.0/88x31.png
[cc-by-nc-sa-shield]: https://img.shields.io/badge/License-CC%20BY--NC--SA%204.0-lightgrey.svg
