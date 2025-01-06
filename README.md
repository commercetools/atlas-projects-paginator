# Atlas Projects Paginator

Small helper to paginate projects from the Atlas API that is used in two of our internal projects.

## Usage

```go
projectIterator := atlaspaginator.NewProjectPaginator(&client.Client)

for projectIterator.Next() {
    projectsPage, err := projectIterator.Value(ctx)
    // do whatever you want with your page of projects
}
```
