# golang
A Tour of Goで書いたコードを羅列していく

https://go-tour-jp.appspot.com/list

gopacketでの実装した際のメモ

shimakaBookpuro:golang shimakatsuya$ ls
README.md	array.go	go_packe.go	go_packet	goslack.go	point.go	webhello.go
shimakaBookpuro:golang shimakatsuya$ go run go_packe
package go_packe: cannot find package "go_packe" in any of:
	/usr/local/Cellar/go/1.12.5/libexec/src/go_packe (from $GOROOT)
	/Users/shimakatsuya/go/src/go_packe (from $GOPATH)
shimakaBookpuro:golang shimakatsuya$ go run go_packe.go
go_packe.go:7:2: cannot find package "github.com/awalterschulze/gographviz" in any of:
	/usr/local/Cellar/go/1.12.5/libexec/src/github.com/awalterschulze/gographviz (from $GOROOT)
	/Users/shimakatsuya/go/src/github.com/awalterschulze/gographviz (from $GOPATH)
go_packe.go:8:2: cannot find package "github.com/google/gopacket" in any of:
	/usr/local/Cellar/go/1.12.5/libexec/src/github.com/google/gopacket (from $GOROOT)
	/Users/shimakatsuya/go/src/github.com/google/gopacket (from $GOPATH)
go_packe.go:9:2: cannot find package "github.com/google/gopacket/pcap" in any of:
	/usr/local/Cellar/go/1.12.5/libexec/src/github.com/google/gopacket/pcap (from $GOROOT)
	/Users/shimakatsuya/go/src/github.com/google/gopacket/pcap (from $GOPATH)
shimakaBookpuro:golang shimakatsuya$ go env GOROOT
/usr/local/Cellar/go/1.12.5/libexec
shimakaBookpuro:golang shimakatsuya$ export GOPATH=/Users/1000ch/workspace/go
shimakaBookpuro:golang shimakatsuya$ go install github.com/awalterschulze/gographviz
can't load package: package github.com/awalterschulze/gographviz: cannot find package "github.com/awalterschulze/gographviz" in any of:
	/usr/local/Cellar/go/1.12.5/libexec/src/github.com/awalterschulze/gographviz (from $GOROOT)
	/Users/1000ch/workspace/go/src/github.com/awalterschulze/gographviz (from $GOPATH)
shimakaBookpuro:golang shimakatsuya$ go get github.com/awalterschulze/gographviz
package github.com/awalterschulze/gographviz: mkdir /Users/1000ch: permission denied
shimakaBookpuro:golang shimakatsuya$ go get github.com/google/gopacket
package github.com/google/gopacket: mkdir /Users/1000ch: permission denied
shimakaBookpuro:golang shimakatsuya$ mkdir /Users/1000ch
mkdir: /Users/1000ch: Permission denied
shimakaBookpuro:golang shimakatsuya$ sudo mkdir /Users/1000ch
Password:
shimakaBookpuro:golang shimakatsuya$ sudo mkdir /Users/1000ch
mkdir: /Users/1000ch: File exists
shimakaBookpuro:golang shimakatsuya$ go get github.com/google/gopacket
package github.com/google/gopacket: mkdir /Users/1000ch/workspace: permission denied
shimakaBookpuro:golang shimakatsuya$ chmod 777 /Users/1000ch
chmod: Unable to change file mode on /Users/1000ch: Operation not permitted
shimakaBookpuro:golang shimakatsuya$ sudo chmod 777 /Users/1000ch
shimakaBookpuro:golang shimakatsuya$ go get github.com/google/gopacket
shimakaBookpuro:golang shimakatsuya$ go get github.com/awalterschulze/gographviz
shimakaBookpuro:golang shimakatsuya$ go get github.com/google/gopacket/pcap
shimakaBookpuro:golang shimakatsuya$ go run go_packe.go
2019/05/20 01:03:56 ./file.pcap: No such file or directory
exit status 1
shimakaBookpuro:golang shimakatsuya$ go run go_packe.go
2019/05/20 01:04:54 ./file.pcap: No such file or directory
exit status 1
shimakaBookpuro:golang shimakatsuya$ go run go_packe.go
shimakaBookpuro:golang shimakatsuya$ dot -V
dot - graphviz version 2.36.0 (20140111.2315)
shimakaBookpuro:golang shimakatsuya$ dot -T png ./digraph.dot -o ./digraph.png
shimakaBookpuro:golang shimakatsuya$ ls
README.md			digraph.dot			file.pcap			go_packet			point.go			名称未設定フォルダ
array.go			digraph.png			go_packe.go			goslack.go			webhello.go
