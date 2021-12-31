// Dinkur the task time tracking utility.
// <https://github.com/dinkur/dinkur>
//
// SPDX-FileCopyrightText: 2021 Kalle Fagerberg
// SPDX-License-Identifier: GPL-3.0-or-later
//
// This program is free software: you can redistribute it and/or modify it under
// the terms of the GNU General Public License as published by the
// Free Software Foundation, either version 3 of the License, or (at your option)
// any later version.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
// FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more
// details.
//
// You should have received a copy of the GNU General Public License along with
// this program.  If not, see <http://www.gnu.org/licenses/>.

// Package console contains code to pretty-print different types to the console.
package console

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/dinkur/dinkur/pkg/dinkur"
	"github.com/fatih/color"
	"github.com/mattn/go-colorable"
)

var (
	stdout          = colorable.NewColorableStdout()
	stderr          = colorable.NewColorableStderr()
	timeFormatLong  = "Jan 02 15:04"
	timeFormatShort = "15:04"

	taskIDColor        = color.New(color.FgHiBlack)
	taskLabelColor     = color.New(color.FgWhite, color.Italic)
	taskNameColor      = color.New(color.FgHiYellow, color.Bold)
	taskTimeDelimColor = color.New(color.FgHiBlack)
	taskDateColor      = color.New(color.FgGreen)
	taskStartColor     = color.New(color.FgGreen)
	taskEndColor       = color.New(color.FgGreen)
	taskEndNilColor    = color.New(color.FgHiGreen, color.Italic)
	taskEndNilText     = "now…"
	taskEndNilTextLen  = utf8.RuneCountInString(taskEndNilText)
	taskDurationColor  = color.New(color.FgCyan)
	taskEditDelimColor = color.New(color.FgHiMagenta)
	taskEditNoneColor  = color.New(color.FgHiBlack, color.Italic)

	fatalLabelColor = color.New(color.FgHiRed, color.Bold)
	fatalValueColor = color.New(color.FgRed)

	tableEmptyColor   = color.New(color.FgHiBlack, color.Italic)
	tableEmptyText    = "No results to display."
	tableHeaderColor  = color.New(color.FgWhite, color.Underline, color.Bold)
	tableSummaryColor = color.New(color.FgHiBlack, color.Italic)

	usageHeaderColor = color.New(color.FgYellow, color.Underline, color.Italic)
	usageHelpColor   = color.New(color.FgHiBlack, color.Italic)
)

// PrintTaskWithDuration writes a label string followed by a formatted task,
// including its elapsed duration, to STDOUT.
func PrintTaskWithDuration(label string, task dinkur.Task) {
	var sb strings.Builder
	taskLabelColor.Fprint(&sb, label)
	sb.WriteByte(' ')
	taskIDColor.Fprint(&sb, "#", task.ID)
	sb.WriteByte(' ')
	writeTaskName(&sb, task.Name)
	sb.WriteByte(' ')
	writeTaskTimeSpan(&sb, task.Start, task.End)
	sb.WriteByte(' ')
	writeTaskDuration(&sb, task.Elapsed())
	fmt.Fprintln(stdout, sb.String())
}

// PrintTask writes a label string followed by a formatted task to STDOUT.
func PrintTask(label string, task dinkur.Task) {
	var sb strings.Builder
	taskLabelColor.Fprint(&sb, label)
	sb.WriteByte(' ')
	taskIDColor.Fprint(&sb, "#", task.ID)
	sb.WriteByte(' ')
	writeTaskName(&sb, task.Name)
	sb.WriteByte(' ')
	writeTaskTimeSpan(&sb, task.Start, task.End)
	fmt.Fprintln(stdout, sb.String())
}

func writeTaskName(w io.Writer, name string) {
	taskNameColor.Fprint(w, `"`, name, `"`)
}

func writeTaskTimeSpan(w io.Writer, start time.Time, end *time.Time) {
	today := newDate(time.Now().Date())
	layout := timeFormatShort
	if today != newDate(start.Date()) ||
		(end != nil && newDate(end.Date()) != today) {
		// also, if start date != end date, also use long format.
		// This still applies, through transitivity
		layout = timeFormatLong
	}
	taskStartColor.Fprintf(w, start.Format(layout))
	taskTimeDelimColor.Fprint(w, " - ")
	if end != nil {
		taskEndColor.Fprintf(w, end.Format(layout))
	} else {
		taskEndNilColor.Fprintf(w, taskEndNilText)
	}
}

func writeTaskDuration(w io.Writer, dur time.Duration) {
	taskTimeDelimColor.Fprint(w, "(")
	taskDurationColor.Fprint(w, FormatDuration(dur))
	taskTimeDelimColor.Fprint(w, ")")
}

// PrintFatal writes a label and some error value to STDERR and then exits the
// application with status code 1.
func PrintFatal(label string, v interface{}) {
	var sb strings.Builder
	fatalLabelColor.Fprint(&sb, label)
	sb.WriteByte(' ')
	fatalValueColor.Fprint(&sb, v)
	fmt.Fprintln(stderr, sb.String())
	os.Exit(1)
}

// PrintTaskEdit writes a formatted task and highlights any edits made to it,
// by diffing the before and after tasks, to STDOUT.
func PrintTaskEdit(update dinkur.UpdatedTask) {
	const editPrefix = "  "
	const editDelim = "   =>   "
	var anyEdit bool
	var sb strings.Builder
	taskLabelColor.Fprint(&sb, "Updated task ")
	taskIDColor.Fprint(&sb, "#", update.Updated.ID)
	sb.WriteByte(' ')
	writeTaskName(&sb, update.Updated.Name)
	taskLabelColor.Fprint(&sb, ":")
	fmt.Fprintln(&sb)
	if update.Old.Name != update.Updated.Name {
		sb.WriteString(editPrefix)
		writeTaskName(&sb, update.Old.Name)
		taskEditDelimColor.Fprint(&sb, editDelim)
		writeTaskName(&sb, update.Updated.Name)
		fmt.Fprintln(&sb)
		anyEdit = true
	}
	var (
		oldStartUnix = update.Old.Start.UnixMilli()
		oldEndUnix   int64
		newStartUnix = update.Updated.Start.UnixMilli()
		newEndUnix   int64
	)
	if update.Old.End != nil {
		oldEndUnix = update.Old.End.Unix()
	}
	if update.Updated.End != nil {
		newEndUnix = update.Updated.End.Unix()
	}
	if oldStartUnix != newStartUnix || oldEndUnix != newEndUnix {
		sb.WriteString(editPrefix)
		writeTaskTimeSpan(&sb, update.Old.Start, update.Old.End)
		sb.WriteByte(' ')
		writeTaskDuration(&sb, update.Old.Elapsed())
		taskEditDelimColor.Fprint(&sb, editDelim)
		writeTaskTimeSpan(&sb, update.Updated.Start, update.Updated.End)
		sb.WriteByte(' ')
		writeTaskDuration(&sb, update.Updated.Elapsed())
		fmt.Fprintln(&sb)
		anyEdit = true
	}
	if !anyEdit {
		sb.WriteString(editPrefix)
		taskEditNoneColor.Fprint(&sb, "No changes were applied.")
		fmt.Fprintln(&sb)
	}
	fmt.Fprint(stdout, sb.String())
}

// PrintTaskList writes a table for a list of tasks, grouped by the date
// (year, month, day), to STDOUT.
func PrintTaskList(tasks []dinkur.Task) {
	if len(tasks) == 0 {
		tableEmptyColor.Fprintln(stdout, tableEmptyText)
		return
	}
	var t table
	t.SetSpacing("  ")
	t.SetPrefix("  ")
	t.WriteColoredRow(tableHeaderColor, "ID", "NAME", "DAY", "START", "END", "DURATION")
	for i, group := range groupTasksByDate(tasks) {
		if i > 0 {
			t.CommitRow() // commit empty delimiting row
		}
		for i, task := range group.tasks {
			t.WriteCellWidth(taskIDColor.Sprintf("#%d", task.ID), uintWidth(task.ID)+1)
			t.WriteCellWidth(taskNameColor.Sprint(task.Name), utf8.RuneCountInString(task.Name))
			if i == 0 {
				dateStr := group.date.String()
				t.WriteCellWidth(taskDateColor.Sprint(dateStr), len(dateStr))
			} else {
				t.WriteCell("")
			}
			t.WriteCellWidth(taskStartColor.Sprint(task.Start.Format(timeFormatShort)), len(timeFormatShort))
			if task.End != nil {
				var endStr string
				if newDate(task.End.Date()) != group.date {
					endStr = task.End.Format(timeFormatLong)
				} else {
					endStr = task.End.Format(timeFormatShort)
				}
				t.WriteCellWidth(taskEndColor.Sprint(endStr), len(endStr))
			} else {
				t.WriteCellWidth(taskEndNilColor.Sprint(taskEndNilText), taskEndNilTextLen)
			}
			dur := FormatDuration(task.Elapsed())
			t.WriteCellWidth(taskDurationColor.Sprint(dur), len(dur))
			t.CommitRow()
		}
	}
	sum := sumTasks(tasks)
	t.CommitRow() // commit empty delimiting row
	endStr := taskEndNilText
	if sum.end != nil {
		endStr = sum.end.Format(timeFormatShort)
	}
	t.WriteColoredRow(tableSummaryColor,
		"", // ID
		fmt.Sprintf("TOTAL: %d tasks", len(tasks)), // NAME
		"",                                // DAY
		sum.start.Format(timeFormatShort), // START
		endStr,                            // END
		FormatDuration(sum.duration),      // DURATION
	)
	t.Fprintln(stdout)
}

// UsageTemplate returns a lightly colored usage template for Cobra.
func UsageTemplate() string {
	var sb strings.Builder
	usageHeaderColor.Fprint(&sb, "Usage:")
	sb.WriteString(`{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}

`)
	usageHeaderColor.Fprint(&sb, "Aliases:")
	sb.WriteString(`
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

`)
	usageHeaderColor.Fprint(&sb, "Examples:")
	sb.WriteString(`
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}

`)
	usageHeaderColor.Fprint(&sb, "Available Commands:")
	sb.WriteString(`{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

`)
	usageHeaderColor.Fprint(&sb, "Flags:")
	sb.WriteString(`
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

`)
	usageHeaderColor.Fprint(&sb, "Global Flags:")
	sb.WriteString(`
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

`)
	usageHeaderColor.Fprint(&sb, "Additional help topics:")
	sb.WriteString(`{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

`)
	usageHelpColor.Fprint(&sb, `Use "{{.CommandPath}} [command] --help" for more information about a command.`)
	sb.WriteString(`{{end}}`)
	sb.WriteByte('\n')
	return sb.String()
}
