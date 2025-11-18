# passgen

A cryptographically secure random password generator written in Go.

## Features

- **Cryptographically secure** - Uses `crypto/rand` for unpredictable password generation
- **Customizable** - Control password length and character types
- **Fast** - Generates passwords instantly
- **Cross-platform** - Works on Linux, macOS, and Windows
- **Zero dependencies** - Uses only Go standard library

## Installation

### macOS (Homebrew)
```bash
brew tap iambugra/tap
brew install passgen
```

### Using Go
```bash
go install github.com/iambugra/passgen@latest
```

### From Source
```bash
git clone https://github.com/iambugra/passgen.git
cd passgen
go build -o passgen
sudo mv passgen /usr/local/bin/
```

### Manual Install

Download the latest binary from [Releases](https://github.com/iambugra/passgen/releases) and add it to your PATH.

## Usage

### Basic Usage
```bash
# Generate a 16-character password (default: includes uppercase, digits, and symbols)
passgen

# Generate a 20-character password
passgen -length=20

# Generate a 32-character password
passgen -length 32
```

### Custom Character Sets
```bash
# Lowercase and digits only (no uppercase or symbols)
passgen -upper=false -symbols=false

# Only lowercase letters
passgen -upper=false -digits=false -symbols=false

# Uppercase and lowercase only (no digits or symbols)
passgen -digits=false -symbols=false
```

### Command-Line Options

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `-length` | int | 16 | Length of the password |
| `-upper` | bool | true | Include uppercase letters (A-Z) |
| `-digits` | bool | true | Include digits (0-9) |
| `-symbols` | bool | true | Include symbols (!@#$%^&*...) |

### Examples
```bash
# 24-character password without symbols (good for systems with symbol restrictions)
passgen -length=24 -symbols=false

# Simple 12-character alphanumeric password
passgen -length=12 -symbols=false

# Strong 32-character password with everything
passgen -length=32

# PIN-like: 8 digits only
passgen -length=8 -upper=false -symbols=false -digits=true
```

## Character Sets

- **Lowercase**: `abcdefghijklmnopqrstuvwxyz`
- **Uppercase**: `ABCDEFGHIJKLMNOPQRSTUVWXYZ`
- **Digits**: `0123456789`
- **Symbols**: `!@#$%^&*()-_=+[]{}|;:,.<>?`

## Building from Source
```bash
# Clone the repository
git clone https://github.com/iambugra/passgen.git
cd passgen

# Build
go build -o passgen

# Install locally
go install
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

MIT License - see [LICENSE](LICENSE) file for details.

## Security

This tool uses Go's `crypto/rand` package, which provides cryptographically secure random numbers suitable for security-sensitive applications. Never use `math/rand` for password generation!

## Author

[Buğra Alparslan](https://github.com/iambugra)

---

**Note**: Always store generated passwords in a secure password manager. Never share passwords or store them in plain text.