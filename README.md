# Go + gRPC
Operations created using [Protocol Buffer](https://developers.google.com/protocol-buffers):

- CreateCategory: Create a category entity using bidirectional data streaming
- GetCategoryBy(Id): Return a category by the identifier
- GetAllCategories: Return all categories

# Generating the code from Protocol Buffer

```
protoc --go_out=. --go-grpc_out=. proto/course_category.proto
```

# Client to interact with Server
For interacting it's possible use [Evans](https://github.com/ktr0731/evans). If you run this project using [Dev Containers](https://code.visualstudio.com/docs/devcontainers/containers) Evans it's already installed in the docker image.

```
evans -r repl
```