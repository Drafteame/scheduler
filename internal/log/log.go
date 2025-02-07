package log

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	debugColor = lipgloss.AdaptiveColor{Light: "#6A6A6A", Dark: "#A0A0A0"}
	infoColor  = lipgloss.AdaptiveColor{Light: "#007ACC", Dark: "#80CFFF"}
	warnColor  = lipgloss.AdaptiveColor{Light: "#FFA500", Dark: "#FFCC66"}
	errorColor = lipgloss.AdaptiveColor{Light: "#FF0000", Dark: "#FF6666"}
	tableColor = lipgloss.AdaptiveColor{Light: "#5F5F5F", Dark: "#A8A8A8"}

	tableHeaderBackground = lipgloss.AdaptiveColor{Light: "#F0F0F0", Dark: "#333333"}
	tableCellBackground   = lipgloss.AdaptiveColor{Light: "#FFFFFF", Dark: "#444444"}
)

func Plainf(format string, args ...any) {
	fmt.Println(fmt.Sprintf(format, args...))
}

func Debugf(format string, args ...any) {
	log(debugColor, format, args...)
}

func Infof(format string, args ...any) {
	log(infoColor, format, args...)
}

func Warnf(format string, args ...any) {
	log(warnColor, format, args...)
}

func Errorf(format string, args ...any) {
	log(errorColor, format, args...)
}

func log(color lipgloss.AdaptiveColor, format string, args ...any) {
	style := lipgloss.NewStyle().Foreground(color)
	fmt.Println(style.Render(fmt.Sprintf(format, args...)))
}

func Table(headers []string, rows [][]string) {
	titleStyle := lipgloss.NewStyle().
		Foreground(infoColor).
		Background(tableHeaderBackground).
		Padding(0, 1, 0, 1).
		Margin(0, 0, 0, 0).
		Bold(true)

	cellStyle := lipgloss.NewStyle().
		Foreground(tableColor).
		Background(tableCellBackground).
		Padding(0, 1, 0, 1).
		Margin(0, 0, 0, 0)

	// prepend headers to the data
	rows = append([][]string{headers}, rows...)

	// Calculate the max width of each column to align properly.
	// This helps handle the "dynamic" nature of the data.
	colWidths := make([]int, len(rows[0]))
	for _, row := range rows {
		for i, col := range row {
			if len(col) > colWidths[i] {
				colWidths[i] = len(col)
			}
		}
	}

	// Print each row with spacing based on max column widths.
	for rowIndex, row := range rows {
		for colIndex, col := range row {
			// Use %-*s (left-aligned, variable width)
			style := cellStyle

			if rowIndex == 0 {
				style = titleStyle
			}

			fmt.Print(style.Render(fmt.Sprintf("%-*s", colWidths[colIndex], col)))
		}
		fmt.Println() // Newline after each row
	}
}
