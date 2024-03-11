package audio

func ChangeAmplitude(buffer PcmBuffer, factor float64) PcmBuffer {

	var newBuffer PcmBuffer = make(PcmBuffer, len(buffer))
	for i, s := range buffer {
		newBuffer[i][0] = s[0] * factor

	}

	return buffer
}
