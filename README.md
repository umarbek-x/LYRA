# 🎵 LYRA

(https://www.vecteezy.com/vector-art/45742674-pixel-art-illustration-jellyfish-pixelated-jellyfish-jellyfish-pixelated-for-the-pixel-art-game-and-icon-for-website-and-game-old-school-retro)

A lightweight terminal music player written in Go.

![Go](https://img.shields.io/badge/Go-1.26-blue)
![License](https://img.shields.io/badge/License-MIT-green)

## Features

- 🎵 Play local MP3 files
- 📂 Browse music from the terminal
- ⚡ Fast and lightweight
- 🖥️ Cross-platform codebase (Linux release available)

## Installation

Download the latest release from the Releases page.

Extract it:

```bash
tar -xzf LYRA_1.0.2_linux_amd64.tar.gz
```

Make it executable:

```bash
chmod +x lyra
```

(Optional) Install globally:

```bash
sudo mv lyra /usr/local/bin/
```

Run:

```bash
lyra
```

## Building from source

Requirements:

- Go 1.26+

Clone the repository:

```bash
git clone git@github.com:umarbek-x/LYRA.git
cd LYRA
```

Build:

```bash
go build -o lyra ./cmd/lyra
```

Run:

```bash
./lyra
```

## Roadmap

Planned features:

- Real-time synchronized lyrics
- Playlist support
- Shuffle and repeat
- Keyboard shortcuts
- Better TUI
- Windows and macOS releases

## Contributing

Contributions, bug reports, and feature requests are welcome.

## License

MIT License.
