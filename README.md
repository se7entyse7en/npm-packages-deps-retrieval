# npm-packages-deps-retrieval

## Running

The requirements are: `docker`, `docker-compose`, `npm`, `make`.
To run the application, you need to clone the repository and run the following commands:
```
$ make prepare
$ make start
```

## TODOs

- [ ] Add diagram of the architecture
- [ ] Add a way to observe components stats
- [ ] Add tests
- [ ] Add error message in web page
- [ ] Dockerize application

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
