# RinGo — Simple Redis-like In-Memory Store in Go
RinGo is a small, experimental project implemented in Go to mimic core Redis features such as storing strings, sets, and hashes with optional expiration.   
a learning exercise to understand Redis internals and Go.   
Features
- Store and retrieve strings, string slices (sets), and hash maps   
- Support expiration times on keys (set expiration in seconds)   
- Simple CLI interface to interact with the store: set, sset, hset, get, delete commands   
- Modular command handlers for easy extension

## Usage
run the app with
```bash
go run .\main.go
```

*COMMANDS*:   
- set: Store a string value 
```bash
set key value 
```

- sset: Store a set (array of strings)
```bash
sset key value1 value2 ...
```

- hset: Store a hash map (dictionary)
```bash
hset key key1 value1 key2 value2 ...
```

- get: Retrieve the value associated with a key
```bash
get key
```

- delete: Remove a key and its value
```bash
delete key
```

- when setting a value use (exp time_in_sec) to add expiration time
```bash
set key value exp 10
```
## Project Status
⚠️ This project is a work-in-progress and experimental.
Code Structure:
+ main.go — command dispatch
+ models/ — Core data structures and in-memory storage logic
+ handlers/ — Command handlers implementing store/get/delete logic
+ errs/ - Store Errors

## TODO
- [ ] Add help command
- [x] Implement `hset` command handler
- [ ] Write more tests for edge cases 
- [x] User input parsing
- [ ] Add documentation and examples
- [x] Add into `lists or maps` instead of rewriting
- [x] implement a Get Handler

### Goals
* Understand Redis data structures and command handling

License
MIT License © veirtex