package chapter1

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func FetchAll() {
	start := time.Now()
	ch := make(chan string) // unbuffered channel
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	// print to a file
	file, err := os.Create("fetchall.txt")
	if err != nil {
		fmt.Printf("error creating file: %v", err)
	}

	// NOTE - the range ensures that we print only as many lines as there are urls
	// there is no guarantee that the response of the url is printed in the same url iteration
	// whatever is received first is printed
	for range os.Args[1:] {
		message := <-ch
		fmt.Println(message)
		n, err := file.WriteString(message + "\n")
		if err != nil {
			fmt.Printf("error writing to file: %v\n", err)
		}
		fmt.Printf("wrote %v bytes\n", n)
	}

	// print on stdout
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()

	// sending to unbuffered channel will block this go routine until the message is receiver
	// value in only transferred when the receiver is ready to receive
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)

	// NOTE: %7d is used to allign numbers neatly
}
