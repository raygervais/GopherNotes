# GopherNotes 📓

An CLI Note taking application written entirely Golang!

## Usage

| Function    | Command                                           | Comments & Tips                                                                                                                                |
| ----------- | ------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| New Note    | `gn create --note "Your awesome quick note here"` | Date is automatically set based on note entry date                                                                                             |
| Fetch Notes | `gn fetch`                                        | This returns all notes in order of entry                                                                                                       |
| Search Note | `gn search --note "Your search text"`             | This searches based on a per-word criteria. Instead of searching using wildcards, search using part of a complete phrase or word, ex: "Todo:"` |
| Edit Note   | `gn edit --id 20`                                 | The date doesn't change when editing a note                                                                                                    |

### Filters

The following filters exist for `Fetch` and `Search` as arguments to be provided:

| Function | Argument Type | Argument     | Comments & Tips                                                                   |
| -------- | ------------- | ------------ | --------------------------------------------------------------------------------- |
| Limit    | int           | `--limit 10` | If not provided, the default amount to return is 10                               |
| Sort     | string        | `--sort asc` | If not provided, the default value is `asc`. Other valid is `desc` for descending |

## Building & Developing

Requirements:

- GNU Make 4.3
- Golang >= 1.15
- SQLite3 >= 3.34.0

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
