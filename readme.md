````markdown
# GoApp CLI – Scaffolding Tool for Modular Golang Projects

**GoApp** is a simple CLI tool that helps you scaffold modular structures in your Golang projects — similar to Laravel's Artisan or Breeze.

This tool is designed for developers who want to build REST APIs or enterprise-ready Go applications with clear folder separation and maintainability in mind.

---

## Features

- Scaffold complete modules instantly
- Includes: controller, model, repository, service, route, middleware, config, and utils
- Modular structure: each module lives in its own `mod_<name>` folder
- Lightweight and project-agnostic — works with any Go project
- Installable via `go install` directly from GitHub

---

## Installation

Make sure your Go environment is correctly configured and `$GOPATH/bin` is included in your system’s `PATH`.

```bash
go install github.com/jakewd7/cli@latest
````

You can now use the `gojake` CLI globally from anywhere:

```bash
gojake --help
```

---

## 📦 Available Commands

### `create:module <ModuleName>`

Generates a full-featured module under the `mod_<name>` directory:

```bash
goapp create:module Auth
```

Generated structure:

```
mod_auth/
├── config/
│   └── config.go
├── controllers/
│   └── handler.go
├── middleware/
│   └── auth.go
├── models/
│   └── auth.go
├── repository/
│   └── auth_repository.go
├── routes/
│   └── auth_routes.go
├── service/
│   └── auth_service.go
├── utils/
│   └── hash.go
```

---

## Module Structure Overview

| Directory      | Purpose                                   |
| -------------- | ----------------------------------------- |
| `controllers/` | HTTP handlers using Gin                   |
| `routes/`      | Module-specific route definitions         |
| `models/`      | Structs for data models                   |
| `repository/`  | Database and data access logic            |
| `service/`     | Business logic layer                      |
| `middleware/`  | Middleware for auth, logging, etc         |
| `config/`      | Configuration per module                  |
| `utils/`       | Helper functions (e.g., password hashing) |

---

## Example Usage

```bash
cd your-project/
goapp create:module Product
```

A full `mod_product` folder will be created inside your project.

---

## Planned Features

* `goapp create:model <Name>` — Model + migration + CRUD handler
* `goapp install` — Initialize a complete starter project structure
* Authentication presets (JWT-based login/register flow)
* Template support for `.env`, `.gitignore`, `Dockerfile`, and more

---

## Contributions

Pull requests and feedback are welcome. Please feel free to open an issue or submit a PR if you have improvements or ideas to share.

---

## License

MIT License © 2025 [jakewd7 (Cloudline.dev)](https://github.com/jakewd7)

```
