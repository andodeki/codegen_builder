package util



/*
=======================================
|| CONTEXT.go                        ||
=======================================
*/

/*
=======================================
|| ULID.go                        ||
=======================================
*/
import (
	"fmt"
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

func New() Generator {
	return &generator{}
}
type Generator interface {
	Generate() string
	Parse(ulid string) error
}

type generator struct{}

func (g *generator) Generate() string {

	// timestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	// tstmp := time.Now().UTC().UnixNano()
	t := time.Now().UTC()

	// t := time.Unix(1000000, 0)
	// fmt.Printf("timestamp: %v\n", timestamp)
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)

	return fmt.Sprint(ulid.MustNew(ulid.Timestamp(t), entropy)) //uuid.New().String()
}

func (g *generator) Parse(ulidStr string) error {
	_, err := ulid.Parse(ulidStr) //uuid.Parse(uuidStr)
	return err
}

/*
=======================================
|| ERRORS.go                        ||
=======================================
*/
/*
=======================================
|| LOGGER.go                        ||
=======================================
*/
