# MongoDB Atlas Iterators 📃

Small helper to iterate over resources from the Atlas API that is used in a couple of our internal projects.

## Usage

```go
projectsService := mongodbatlas.Client{...}.Projects
ctx := context.Background()

for project, err := range atlasiter.AllProjects(ctx, projectsService) {
    if err != nil {
        return nil, err
    }

    // do whatever you want with your project
}
```

## License

[MIT](./LICENSE-MIT)
