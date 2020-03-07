# network-speed-test
This package measure your network card can handle how much packet per a second.

### Installation

    go get github.com/batuberksahin/network-speed-test
    cd $GOPATH/src/github.com/batuberksahin/network-speed-test
    go run main.go -f [pcap destination]

### Example

    $ go run main.go -f test.pcap
    PCAP took 4m43.466932s
    136.31 packet per second

