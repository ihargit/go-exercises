type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	for {
		n, err := r.r.Read(b)
		for i := 0; i < n; i++ {
			v := b[i]
			if v >= 'A' && v <= 'Z' {
				v += 13
				if v > 'Z' {
					v -= 26
				}
			}
			if v >= 'a' && v <= 'z' {
				v += 13
				if v > 'z' {
					v -= 26
				}
			}
			b[i] = v
		}
		return n, err
	}
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}