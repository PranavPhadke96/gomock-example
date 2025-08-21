#  Greeting Service TDD Example (Go + GoMock)

##  Folder Structure

```
greeting_project/
├── go.mod
└── greeting/
    ├── greeter.go              # Greeter interface and GreetingService
    ├── greeter_test.go         # Test using GoMock
    └── mock_greeter.go         # Pre-generated mock
```

##  Setup Instructions

1. ✅ Open folder in **VS Code**
2. 🧪 Run tests in terminal:
```bash
go test ./...
```

You should see:
```
PASS
ok  	greeting_project/greeting	0.XXXs
```

##  Learning Objectives

- Define and use interfaces in Go
- Inject dependencies using constructor pattern
- Use **GoMock** to isolate and test behavior
- Apply **TDD**: test → fail → code → pass
