package main
import "os"
import "bufio"

func main() {
	inreader := bufio.NewReader(os.Stdin)
	outwriter := bufio.NewWriter(os.Stdout)

	Decrypt(inreader, outwriter)
}

func Decrypt(inreader *bufio.Reader, outwriter *bufio.Writer) {
	buff := make([]byte, 1024)
	key := NewKeyState("xxx")

	for {
		read, _ := inreader.Read(buff)
		if read == 0 {
			break
		}

		for i := 0; i < read; i++ {
			buff[i] = buff[i] ^ key.NextByte()
		}

		_, err := outwriter.Write(buff[:read])
		if err != nil {
			panic(err)
			break
		}
	}

	outwriter.Flush()
}
