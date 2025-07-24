# Go File Renamer

A Go utility for batch renaming media files from the format `"Name YYYY (X of Y).ext"` to `"YYYY - Name - X of Y.ext"`.

## Installation

```bash
git clone https://github.com/bkandh30/goRename.git
cd goRename
```

## Usage

### Basic Usage (Dry Run)

```bash
go run main.go
```

By default, the program runs in dry-run mode, showing what would be renamed without making actual changes.

### Real Execution

```bash
go run main.go -dry=false
```

### Command Line Options

- `-dry` (default: `true`) - Set to `false` to perform actual file renames

## File Pattern

The program matches files with this pattern:

```
Name YYYY (X of Y).extension
```

Examples of matching files:

- `The Dark Knight 2008 (1 of 2).mp4`
- `Breaking Bad S01E01 2008 (1 of 62).mkv`
- `Documentary Series 2023 (5 of 10).avi`
