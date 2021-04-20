package query

type PositionTranscoder struct {
	width int
}

func InitTranscoder(width int) *PositionTranscoder {
	return &PositionTranscoder{width: width}
}

func (transcoder *PositionTranscoder) CodePosition(x, y int) PositionCode {
	return transcoder.width*y + x
}

func (transcoder *PositionTranscoder) DecodePosition(position PositionCode) (int, int) {
	return position % transcoder.width, position / transcoder.width
}
