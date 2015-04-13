package ranges

func main() {
	m := make(map[string]int)
	for k, _ := range m {
		_ = k
	}
}
