### Golang futures, promise, lazy...etc...

things that help me simplify my code.

### Bugs

- Lazy can leak a goroutine. Maybe provide an interface instead of LazyChan that instructs a user to intuitively close a lazy instance. And do not execute the function when close is called and lazy's expression will not be used.
- 
