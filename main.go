package main

import (
	"io"
	"log"
	"net"
	"sync"
	"bufio"
	"bytes"
	"strings"
	"strconv"
	"fmt"
	"os/user"
)

func reader(r io.Reader, group *sync.WaitGroup) {
	defer group.Done()
	buf := make([]byte, 1024)
	n, err := r.Read(buf[:])
	if err != nil {
		log.Fatalf("Failed to read from socket: %v", err)
		return
	}
	reader := bufio.NewReader(bytes.NewReader(buf[0:n]))
	var artist, title string
	var duration, position int
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("Failed to read line: %v", err)
			return
		}
		if strings.HasPrefix(line, "tag artist") {
			artist = line[11 : len(line)-1]
		} else if strings.HasPrefix(line, "tag title") {
			title = line[10 : len(line)-1]
		} else if strings.HasPrefix(line, "duration") {
			duration, err = strconv.Atoi(line[9 : len(line)-1])
			if err != nil {
				log.Fatalf("Failed to convert duration: %v (text: %v)", err, line[9:len(line)-1])
				return
			}
		} else if strings.HasPrefix(line, "position") {
			position, err = strconv.Atoi(line[9 : len(line)-1])
			if err != nil {
				log.Fatalf("Failed to convert position: %v (text: %v)", err, line[9:len(line)-1])
				return
			}
		}
	}
	fmt.Printf("%s - %s [%02d:%02d / %02d:%02d]\n", artist, title, position/60, position%60, duration/60, duration%60)
}

func main() {
	current, err := user.Current()
	if err != nil {
		fmt.Printf("user id environment variable UID is not available %v", err)
		return
	}
	socket := "/run/user/" + current.Uid + "/cmus-socket"
	c, err := net.Dial("unix", socket)
	if err != nil {
		fmt.Printf("socket not available: %v", socket)
		return
	}
	defer c.Close()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go reader(c, wg)
	_, err = c.Write([]byte("status\n"))
	if err != nil {
		fmt.Printf("error writing to socket %v: %v", socket, err)
		return
	}
	wg.Wait()
}
