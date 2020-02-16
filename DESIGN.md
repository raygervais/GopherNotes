# Design

_Note, this document is meant to describe the design and functions; be-it in incomplete thoughts or concepts. Final design will make it's way into the document later._


## Features
- Consumes a single text file containing all notes, organizes them into a hashtable with the key being date.
- CRUD capabilities for a note
- Application is written in a functional manner
- 100% Unit Test Coverage
- Github Actions integration for releases to various platforms (Windows, MacOS, Linux)

## CLI Commands

- `gn --new --text "We need to follow-up on the onboarding script performance to ensure we are meeting the <2minute requirements as new arrivals are added to the system."`
- `gn --search --date "01/22/2020"`
- `gn --search --text "Onboarding"

## Storagefile Example Entry

`01/12/2020:"We need to follow-up on the onboarding script performance to ensrue we are meeting the <2minute requirements as new arrivals are added to the system."`
