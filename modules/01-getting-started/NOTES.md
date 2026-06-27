# Go Notes

Go was created at Google in **2007** by **Robert Griesemer**, **Rob Pike**, **Ken Thompson** and released publicly in **2009**.

Its guiding philosophy is **simplicity over cleverness**.

---

Every executable Go program has a **Package** and a `func main()` - this is the entry point.

## What is a Package?

- A collection of Go files in the same directory that are compiled together as one unit.
- Packages create boundaries. Without them, every function would have one global namespace.
  - `os.Open()`
  - `zip.Open()`
- There's one special package named **main**.
  - It tells Go that **this is an executable program**.
  - There must be a `main()` function — that's the entry point.
  - Without it, Go has nowhere to begin execution.
- Files inside the same package can freely access each other — **no import required**.
- Different packages require imports.

## Go Modules

Go Modules are how Go manages dependencies and versioning.

A module is defined by a `go.mod` file at the root of our project.

```bash
go mod init github.com/yourname/projectname
go get github.com/some/dependency
go mod tidy    # Cleans up unused / missing dependencies
```
