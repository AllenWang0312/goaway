package spider

import (
	"../../util"
	"bufio"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)


func DownloadFile(client http.Client,durl string, path string, name string) int {
	resp, err := client.Get(durl)
	if err != nil {
		//panic(err)
		println(err.Error())
		return -2
	}
	//log.Println( path + name)
	if resp.StatusCode > 400 {
		return -2
	}
	if resp.ContentLength <= 0 {
		log.Println("content length < 0")
	}

	raw := resp.Body
	defer raw.Close()
	reader := bufio.NewReaderSize(raw, 1024*32)

	file, err := os.Create(path + name)
	if err != nil {
		//panic(err)
		println(err.Error())
	}
	writer := bufio.NewWriter(file)
	buff := make([]byte, 32*1024)
	written := 0
	//go
	func() {
		for {
			nr, er := reader.Read(buff)
			if nr > 0 {
				nw, ew := writer.Write(buff[0:nr])
				if nw > 0 {
					written += nw
				}
				if ew != nil {
					err = ew
					break
				}
				if nr != nw {
					err = io.ErrShortWrite
					break
				}
			}
			if er != nil {
				if er != io.EOF {
					err = er
				}
				break
			}
		}
		if err != nil {
			//panic(err)
			println(err.Error())
		}
	}()

	spaceTime := time.Second * 1
	ticker := time.NewTicker(spaceTime)
	lastWtn := 0
	stop := false
	for {
		select {
		case <-ticker.C:
			speed := written - lastWtn
			log.Printf("speed %s /%s \n", util.BytesToSize(speed), spaceTime.String())
			if written-lastWtn == 0 {
				ticker.Stop()
				stop = true
				break
			}
			lastWtn = written
		}
		if stop {
			break
		}
	}
	return 0
}