package main

import (
	"fmt"
	"image"
	"image/color"
	"io"
	"math"
	"os"
	"strings"

	"golang.org/x/tour/pic"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("Nil value for T")
		return
	}
	fmt.Printf("Value %+v\n", *t)
}

type Vehicle interface {
	accelerate()
}

type Car struct {
	speed int
}

func (car *Car) accelerate() {
	car.speed += 10
}

func do(d interface{}) {
	switch t := d.(type) {
	case int:
		fmt.Printf("The type is int and value %d\n", t)
	case string:
		fmt.Printf("Length of string : %d\n", len(t))
	default:
		fmt.Println("I dont knw the type")
	}
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	value := fmt.Sprintf("cannot Sqrt negative number: %d", int(e))
	return value
}

func reader_demo() {
	r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 8)

	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

func main() {
	image_exercise()
}

type MyStruct struct{}

func (my MyStruct) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func (r *rot13Reader) Read(b []byte) (int, error) {
	reader := r.r

	n, err := reader.Read(b)

	if err != nil {
		return n, err
	}

	for i := 0; i < n; i++ {
		b[i] = rot13Transform(b[i])
	}

	return n, nil
}

func rot13Transform(b byte) byte {
	if b >= 'a' && b <= 'z' {
		return 'a' + (b-'a'+13)%26
	}
	if b >= 'A' && b <= 'Z' {
		return 'A' + (b-'A'+13)%26
	}
	return b
}

func reader_exercise() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

type IPAddr [4]byte

type rot13Reader struct {
	r io.Reader
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

type CustomImage struct{}

func (customImage CustomImage) Bounds() image.Rectangle {
	return image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{100, 100}}
}

func (customImage CustomImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (customImage CustomImage) At(x, y int) color.Color {
	z := uint8(x ^ y)
	return color.RGBA{z, z, 255, 255}
}

func image_exercise() {
	m := CustomImage{}
	pic.ShowImage(m)
}
