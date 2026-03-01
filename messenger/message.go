package messenger

import "github.com/charmbracelet/lipgloss"

type messageStyle struct {
	ColorAll bool
	Icon  string
	Style lipgloss.Style
}

var (
	warn = messageStyle{
		ColorAll: true,
		Icon: "⚠",

		Style: lipgloss.NewStyle().
			Foreground(lipgloss.Color("214")). // orange
			Bold(true),
	}

	success = messageStyle{
		ColorAll: false,
		Icon: "✓",

		Style: lipgloss.NewStyle().
			Foreground(lipgloss.Color("82")). // green
			Bold(true),
	}

	info = messageStyle{
		ColorAll: false,
		Icon: "•",

		Style: lipgloss.NewStyle().Foreground(lipgloss.Color("39")), // blue
	}

	err = messageStyle{
		ColorAll: true,
		Icon: "✗",
		Style: lipgloss.NewStyle().
			Foreground(lipgloss.Color("196")).
			Bold(true),
	}
)

func renderMsg(msg string, style messageStyle) string {
	if style.ColorAll {
		return style.Style.Render(style.Icon + " " + msg)
	}
	return style.Style.Render(style.Icon) + " " + msg
}
