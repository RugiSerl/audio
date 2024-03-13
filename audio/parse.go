package audio

import (
	"fmt"
	"os"

	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
)

type Parameters struct {
	//Format of the audio
	WaveFormatType int
	//Number of audio channels
	Channel int
	//Amount of samples in a second
	SampleRate int
	//Size of one sample
	BitsPerSample int
}

// parse all the data all at once. Simple but unoptimised
func Parse(fileName string) (*audio.IntBuffer, Parameters, error) {
	//retrieve the stats
	f, err := os.Open(fileName)
	if err != nil {
		return nil, Parameters{}, err
	}
	d := wav.NewDecoder(f)
	d.ReadInfo()
	p := Parameters{
		WaveFormatType: int(d.WavAudioFormat),
		Channel:        int(d.NumChans),
		SampleRate:     int(d.SampleRate),
		BitsPerSample:  int(d.BitDepth),
	}
	samples, err := d.FullPCMBuffer()
	if err != nil {
		return nil, Parameters{}, err
	}

	fmt.Println(p)

	return samples, p, nil
}

// Save the whole buffer
func Save(fileName string, buffer *audio.IntBuffer, param Parameters) error {
	//creating file
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	e := wav.NewEncoder(f, param.SampleRate, param.BitsPerSample, param.Channel, param.WaveFormatType)

	e.Write(buffer)
	e.Close()
	f.Close()

	return nil

}
func initBuffer(param Parameters, size int) *audio.IntBuffer {
	format := &audio.Format{
		NumChannels: int(param.Channel),
		SampleRate:  int(param.SampleRate),
	}

	buf := &audio.IntBuffer{Data: make([]int, size), Format: format, SourceBitDepth: int(param.BitsPerSample)}

	return buf
}

func initBufferWithOthersStat(buf *audio.IntBuffer) *audio.IntBuffer {
	return initBuffer(Parameters{
		WaveFormatType: -1, // this doesn't actually matter
		Channel:        buf.Format.NumChannels,
		SampleRate:     buf.Format.SampleRate,
		BitsPerSample:  buf.SourceBitDepth,
	}, len(buf.Data))

}
