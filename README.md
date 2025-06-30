# RinGo — Simple Redis-like In-Memory Store in Go
RinGo is a small, experimental project implemented in Go to mimic core Redis features such as storing strings, sets, and hashes with optional expiration.   
a learning exercise to understand Redis internals and Go.   
Features
- Store and retrieve strings, string slices (sets), and hash maps   
- Support expiration times on keys (set expiration in seconds)   
- Simple CLI interface to interact with the store: store, get, delete commands   
- Modular command handlers for easy extension   
   
## Usage
**⚠️ IMPORTANT:** there is nothing to do yet, also feel free to contribute to this dumb project.   

## Project Status
⚠️ This project is a work-in-progress and experimental.
Code Structure:
+ main.go — CLI parsing and command dispatch
+ models/ — Core data structures and in-memory storage logic
+ handlers.go — Command handlers implementing store/get/delete logic
+ errs/ - Store Errors

## TODO
- [ ] Add help command
- [x] Implement `hset` command handler
- [ ] Write more tests for edge cases
- [ ] Improve CLI argument parsing
- [ ] Add documentation and examples
- [x] Add into `lists or maps` instead of rewriting
- [ ] implement a Get Handler

### Goals
* Understand Redis data structures and command handling

License
MIT License © veirtex