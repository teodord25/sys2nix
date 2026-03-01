package messenger

import (
	"fmt"

	"github.com/charmbracelet/log"
)

type Messenger struct {
	Success func(string, ...any)
	Warn    func(string, ...any)
	Error   func(string, ...any)
	Info    func(string, ...any)
}

func NewMessenger(logger *log.Logger) Messenger {
	write := func(style messageStyle, lvl log.Level) func(string, ...any) {
		return func(format string, args ...any) {
			msg := fmt.Sprintf(format, args...)
			fmt.Println(renderMsg(msg, style))

			// TODO: just remove this when logging seems pointless
			// TODO: maybe have nix build redirect file logging to XDG
			logger.Log(lvl, msg)
		}
	}

	return Messenger{
		Success: write(success, log.InfoLevel),
		Warn:    write(warn, log.WarnLevel),
		Error:   write(err, log.ErrorLevel),
		Info:    write(info, log.InfoLevel),
	}
}
