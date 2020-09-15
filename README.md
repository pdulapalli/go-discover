# go-discover

## Purpose

Enable discovery of local mDNS service broadcasts with zero additional dependencies.
Extremely bare-bones wrapper around [Hashicorp mDNS for Go](https://github.com/hashicorp/mdns)

## Building

NOTE: For the greatest likelihood of compatibility, it is recommended to build
on the same CPU architecture that the intended target environment will use.

Please ensure that you have a working installation of Docker. Locate the
relevant instructions for your Operating System at
[the official Docker website](https://docs.docker.com/install).

### Ubuntu/Debian Targets
```
docker pull golang:1.14.2-stretch
docker run -t -v $(pwd):/workenv -w /workenv golang:1.14.2-stretch go build -o discover discover.go
```

If the above commands are successful, an executable named `discover` should
appear.

## Usage

### Flags

* Network Interface (`-i`), e.g. `"wlan0"`
* Service Type (`-s`), e.g. `"_workstation._tcp_"`
* Duration (`-d`), e.g. `7`

### Invocation Style
```
./discover -i wlan0 -s "_spotify-connect._tcp" -d 5
```

### Report Style

Each item is in the JSON format.

If there are multiple items, a newline will separate the opening brace of the
next item from the closing brace of the previous item.

```
{
	"Name": "Denon\\ AVR-X3400H._spotify-connect._tcp.local.",
	"Host": "Denon-AVR-X3400H.local.",
	"AddrV4": "192.168.1.216",
	"AddrV6": "",
	"Port": 80,
	"Info": "VERSION=1.0|CPath=/spotify",
	"InfoFields": [
		"VERSION=1.0",
		"CPath=/spotify"
	],
	"Addr": "192.168.1.216"
}
```
