# npm-packages-deps-retrieval

## TODOs

- [ ] Add a decent api
- [ ] Add a way to observe components stats
- [ ] Add tests
- [ ] Use a real queue (kafka would make it scale well: a topic with a number of partitions >= # workers)
- [ ] Use a real database (for easier distribution a NoSQL one would fit)

## Ideas

- Use npm registry webhook to point to the `dispatcher`
- Use a graph database that should be much more suitable for this data
- Add an in-memory LRU cache to both workers and APIs
- Add a distributed in-memory LRU cache (redis)
- Before hitting the DB it may worth adding a bloom filter on top
- For each package with unpinned dependencies, consider each compatible ones

## Limitations

- Dependencies prefixed with `~` `^` are threated as pinned
- Dependencies with different qualifiers (`>=`, `<`, etc.) are ignored
