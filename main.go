package main
import "os"
import "bufio"

func main() {
	inreader := bufio.NewReader(os.Stdin)
	outwriter := bufio.NewWriter(os.Stdout)

	buff := make([]byte, 1024)

	for {
		read, _ := inreader.Read(buff)
		if read == 0 {
			break
		}

		_, err := outwriter.Write(buff[:read])
		if err != nil {
			panic(err)
			break
		}
	}

	outwriter.Flush()
}
