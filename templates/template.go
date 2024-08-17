import (
	"bufio"
	"os"
)

filename := ""

func main() {
	f, err := os.Open(filename)

	fs := bufio.NewScanner(f)
}