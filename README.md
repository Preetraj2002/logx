# logx 🎨

A minimalist, colorized logging utility for Go — built with ❤️ for devs who want clarity and control without importing big logging frameworks.

## ✨ Features

- Color-coded logs for better readability
- `IsDebug` toggle to control verbosity (great for production!)
- Custom `Prefix` support for distinguishing modules/services
- Simpler than `log/slog` but just as powerful for small-to-mid projects

---

## 📦 Installation

```bash
go get github.com/Preetraj2002/logx

```

## 🚀 Usage

```go
package main

import "github.com/Preetraj2002/logx"

func main() {
    logx.IsDebug = true
    logx.Prefix = "[KEY-SERVER]"

    logx.Info("Starting service on port %d...", 12345)
    logx.Success("Server keys generated")
    logx.Warning("No custom config provided")
    logx.Error("Failed to load OTP module: %v", "missing file")
    // logx.Fatal stops execution
    // logx.Fatal("Critical failure: %s", "port already in use")
}
```

---

## 🎛️ Config Options

| Option     | Type    | Default | Description                           |
|------------|---------|---------|---------------------------------------|
| `IsDebug`  | `bool`  | `true`  | If `false`, suppresses Info/Success/Warning |
| `Prefix`   | `string`| `""`    | Prepended before log level label      |

---

## 🎨 Log Levels

| Method      | Color     | Description               |
|-------------|-----------|---------------------------|
| `Info`      | Magenta   | General status updates    |
| `Success`   | Green     | Completed tasks           |
| `Warning`   | Yellow    | Something unexpected      |
| `Error`     | Red       | Recoverable failures      |
| `Fatal`     | Bold Red  | Exits the program         |

---

## 🧪 Example Output

```text
2025/04/18 12:45:17 [INFO]     OTP sent to user (email & phone)
2025/04/18 12:45:22 [SUCCESS]  User Verified
2025/04/18 12:45:22 [WARNING]  Config not found, using defaults
2025/04/18 12:45:22 [ERROR]    Failed to save public key
```

---

## 📜 License

MIT — do what you want, just don't claim you wrote it.

---

## 🤝 Contribute

Feel free to suggest ideas.

Happy logging! 🔐🚦

