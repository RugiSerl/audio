package audio

import (
	"fmt"

	"github.com/zenwerk/go-wave"
)

type PcmBuffer [][]float64

// parse all the data all at once. Simple but unoptimised
func Parse(fileName string) PcmBuffer {
	reader, err := wave.NewReader(fileName)

	if err != nil {
		panic(err)
	}

	var e error = nil
	var s PcmBuffer = PcmBuffer{}

	for i := 0; e == nil; i++ {
		s[i], e = reader.ReadSample()
		fmt.Println(s)
	}

	return s

}

// func Save(fileName string, buffer PcmBuffer, param wave.WriterParam) {
// 	w, err := wave.NewWriter(param)
// 	if err != nil {
// 		panic(err)
// 	}

// }
