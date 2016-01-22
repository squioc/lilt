package lilt

const LF byte = 0x0A

type LineFeedReader struct{}

func (r *LineFeedReader) Read(p []byte) (n int, err error) {
	p[0] = LF
	return 1, nil
}
