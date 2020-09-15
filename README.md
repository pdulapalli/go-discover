# go-discover

Extremely bare-bones wrapper around [Hashicorp mDNS for Go](https://github.com/hashicorp/mdns)

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
