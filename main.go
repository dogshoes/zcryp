package main
import "os"
import "bufio"
import "fmt"
import "flag"

func main() {
	input, output, mode, key, keylength := ReadFlags()
	PrintFlags(input, output, mode, key, keylength)

	if key == nil {
		fmt.Fprintln(os.Stderr, "A key must be supplied.")
		os.Exit(-1)
	}

	buffsize := 1024
	keystate := NewKeyState(*key)

	var inreader *bufio.Reader
	var outwriter *bufio.Writer

	if input != nil && output != nil {
		infile, inerr := os.Open(*input)
		if inerr != nil {
			fmt.Fprintln(os.Stderr, "can't get the stat of file %s", input)
			os.Exit(-1)
		}

		inreader = bufio.NewReader(infile)
		defer infile.Close()

		outfile, outerr := os.Create(*output)
		if outerr != nil {
			fmt.Fprintln(os.Stderr, "Can not creat temp file '%s'", output)
			os.Exit(-1)
		}

		outwriter = bufio.NewWriter(outfile)
		defer outfile.Close()

		fmt.Fprintln(os.Stderr, fmt.Sprintf("file mode, ZCRYP_BUFSIZ:%d len:%d", buffsize, keystate.keylen))
	} else {
		inreader = bufio.NewReader(os.Stdin)
		outwriter = bufio.NewWriter(os.Stdout)

		fmt.Fprintln(os.Stderr, fmt.Sprintf("pipe mode, ZCRYP_BUFSIZ:%d len:%d", buffsize, keystate.keylen))
	}

	err := Decrypt(inreader, outwriter, keystate, buffsize)
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(os.Stderr, "Result: OK")
	os.Exit(0)
}

func ReadFlags() (*string, *string, *string, *string, *string) {
	var input, output, mode, key, length *string

	flag.String("i", "", "If using file mode, the input file to read.")
	flag.String("o", "", "If using file mode, the output file to create.")
	flag.String("m", "", "The encryption mode (unused, for compatibility).")
	flag.String("k", "", "The encryption key.")
	flag.String("l", "", "The encryption key length (unused, for compatibility).")

	// We're doing this wacky visitor bit here to ensure we have nil values for unset flags, which
	// allows us to detect unset flags and display (null) in the output.
	visitor := func(a *flag.Flag) {
		switch a.Name {
			case "i":
				temp := a.Value.String()
				input = &temp
				break
			case "o":
				temp := a.Value.String()
				output = &temp
				break
			case "m":
				temp := a.Value.String()
				mode = &temp
				break
			case "k":
				temp := a.Value.String()
				key = &temp
				break
			case "l":
				temp := a.Value.String()
				length = &temp
				break
		}
	}

	flag.Parse()
	flag.Visit(visitor)

	return input, output, mode, key, length
}

func PrintFlags(input *string, output *string, mode *string, key *string, length *string) {
	// Preserve the (null) display behavior of stock zcryp.
	printableinput := "(null)"
	printableoutput := "(null)"
	printablemode := "(null)"
	printablekey := "(null)"
	printablelength := "(null)"

	if input != nil {
		printableinput = *input
	}

	if output != nil {
		printableoutput = *output
	}

	if mode != nil {
		printablemode = *mode
	}

	if key != nil {
		printablekey = *key
	}
	
	if length != nil {
		printablelength = *length
	}

	fmt.Fprintln(os.Stderr, fmt.Sprintf("input->%s , output->%s , mode->%s , key->%s, len->%s", printableinput, printableoutput, printablemode, printablekey, printablelength))
}

func Decrypt(inreader *bufio.Reader, outwriter *bufio.Writer, key *KeyState, buffersize int) (error) {
	buff := make([]byte, buffersize)

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
			return err
		}
	}

	outwriter.Flush()

	return nil
}
