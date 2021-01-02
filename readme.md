# GopherNotes

An CLI Note taking application written entirely Golang!

![Test](https://github.com/raygervais/GopherNotes/workflows/Test/badge.svg)

## Usage

| Function    | Command                                           | Comments & Tips                                                                                                                                |
| ----------- | ------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| New Note    | `gn create --note "Your awesome quick note here"` | Date is automatically set based on note entry date                                                                                             |
| Fetch Notes | `gn fetch`                                        | This returns all notes in order of entry                                                                                                       |
| Search Note | `gn search --note "Your search text"`             | This searches based on a per-word criteria. Instead of searching using wildcards, search using part of a complete phrase or word, ex: "Todo:"` |
| Edit Note   | `gn edit --id 20`                                 | The date doesn't change when editing a note                                                                                                    |

## Building & Developing

Requirements:

- GNU Make 4.3
- Golang >= 1.15
- SQLite3

To build the application, run `make` in the root directory. This will compile a standalone binary for your system called `gn` into the `bin` directory. If this directory doesn't exist in the repo, please create.

### Formatting

To format all the codebase to Golang standards, use `make fmt` which calls `go fmt ./...` directly.

### Testing

### Building to Publish

To build for the following systems, run `make publish`:

- Windows
- MacOS
- Linux
- FreeBSD
