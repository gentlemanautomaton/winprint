package portdata

import "fmt"

type marshalError struct {
	Level int
	Msg   string
}

func newMarshalError(level int, format string, a ...any) marshalError {
	if len(a) == 0 {
		return marshalError{
			Level: level,
			Msg:   format,
		}
	}

	return marshalError{
		Level: level,
		Msg:   fmt.Sprintf(format, a...),
	}
}

func (e marshalError) Error() string {
	return fmt.Sprintf("tcpipport: portdata: Level%d.MarshalBinary: %s", e.Level, e.Msg)
}

type unmarshalError struct {
	Level int
	Msg   string
}

func newUnmarshalError(level int, format string, a ...any) unmarshalError {
	if len(a) == 0 {
		return unmarshalError{
			Level: level,
			Msg:   format,
		}
	}

	return unmarshalError{
		Level: level,
		Msg:   fmt.Sprintf(format, a...),
	}
}

func (e unmarshalError) Error() string {
	return fmt.Sprintf("tcpipport: portdata: Level%d.UnmarshalBinary: %s", e.Level, e.Msg)
}
