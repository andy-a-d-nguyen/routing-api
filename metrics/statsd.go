package metrics

import "time"

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6@latest -generate

//counterfeiter:generate . StatSender
type StatSender interface {
	Inc(string, int64, float32) error
	Dec(string, int64, float32) error
	Gauge(string, int64, float32) error
	GaugeDelta(string, int64, float32) error
	Timing(string, int64, float32) error
	TimingDuration(string, time.Duration, float32) error
	Set(string, string, float32) error
	SetInt(string, int64, float32) error
	Raw(string, string, float32) error
}

//counterfeiter:generate . Sender
type Sender interface {
	Send(data []byte) (int, error)
	Close() error
}

//counterfeiter:generate . SamplerFunc
type SamplerFunc func(float32) bool

//counterfeiter:generate . Statter
type Statter interface {
	StatSender
	NewSubStatter(string) SubStatter
	SetPrefix(string)
	Close() error
}

//counterfeiter:generate . SubStatter
type SubStatter interface {
	StatSender
	SetSamplerFunc(SamplerFunc)
	NewSubStatter(string) SubStatter
}
