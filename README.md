# 🎵 LYRA

(https://www.google.com/search?q=penguin+picture+pexel+style+&client=ubuntu-sn&hs=oUQ&sca_esv=a4aa018d8e4a65dc&channel=fs&udm=2&biw=1920&bih=963&sxsrf=APpeQnuIBHa7olgcKb1vq7Bh_cKcn4DCgA%3A1784602943482&ei=P-FeauiKHaioxc8Pgvv12Aw&ved=0ahUKEwiopouE5OKVAxUoVPEDHYJ9HcsQ4dUDCBE&uact=5&oq=penguin+picture+pexel+style+&gs_lp=Egtnd3Mtd2l6LWltZyIccGVuZ3VpbiBwaWN0dXJlIHBleGVsIHN0eWxlIEj6HlBJWKoecAJ4AJABAZgBkAGgAeYMqgEEMS4xM7gBA8gBAPgBAZgCBKAChQPCAgYQABgHGB7CAgUQABiABMICBxAAGIAEGArCAgYQABgeGArCAgQQABgewgIGEAAYCBgemAMAiAYBkgcDMS4zoAerD7IHAzAuM7gHgAPCBwMyLTTIBxGACAE&sclient=gws-wiz-img#ip=1&sv=CAMSURoyKhBlLWEtUTdGM3hsRDhVaVVNMg5hLVE3RjN4bEQ4VWlVTToOOGhqRWI0OVZlVXJCYU0gBCoXCgFzEhBlLWEtUTdGM3hsRDhVaVVNGAEwARgHIMiD5pYFSggQARgBIAEoAQ)

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
