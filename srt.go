package subtitles

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// SRTFormat extension for format
const SRTFormat string = "SRT"

// SRTExtension extension for format
const SRTExtension string = ".srt"

// SRTBlockDurationPattern regex pattern
var SRTBlockDurationPattern = regexp.MustCompile(`(\d{2}:\d{2}:\d{2},\d{3})\s-->\s(\d{2}:\d{2}:\d{2},\d{3})`)

// SRTParser subtitle format
type SRTParser struct{}

// NewSRTParser returns SRT format instance
func NewSRTParser() *SRTParser {
	return &SRTParser{}
}

// Parse SRT subtitle content
func (p *SRTParser) Parse(r io.Reader) (sub *Subtitle, err error) {
	var line string
	var lineIndex int
	var blockIndex int

	sub = NewSubtitle()
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		blockIndex++
		lineIndex++

		// Remove BOM
		line = strings.TrimSpace(scanner.Text())
		if lineIndex == 1 {
			line = StripBOM(line)
		}

		// Block index
		if index, _ := strconv.Atoi(line); index != blockIndex {
			return nil, &ErrInvalidSubtitle{
				format: SRTFormat,
				reason: fmt.Sprintf("Expected index %v at line %v got: %v", blockIndex, lineIndex, line),
			}
		}

		// Next line
		scanner.Scan()
		lineIndex++

		// Block duration: hh:mm:ss,fff --> hh:mm:ss,fff
		line = strings.TrimSpace(scanner.Text())
		blockDurations := SRTBlockDurationPattern.FindStringSubmatch(line)
		if len(blockDurations) == 0 {
			return nil, &ErrInvalidSubtitle{
				format: SRTFormat,
				reason: fmt.Sprintf("Expected duration with pattern (hh:mm:ss,fff --> hh:mm:ss,fff) at line %v got: %v", lineIndex, line),
			}
		}

		startAt := p.parseDuration(blockDurations[1])
		finishAt := p.parseDuration(blockDurations[2])

		// Block lines
		scanner.Scan()
		lineIndex++

		line = strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			return nil, &ErrInvalidSubtitle{
				format: SRTFormat,
				reason: fmt.Sprintf("Expected text at line %v got: empty line", lineIndex),
			}
		}

		lines := []string{}
		for {
			if len(line) > 0 {
				lines = append(lines, line)

				scanner.Scan()
				lineIndex++

				line = strings.TrimSpace(scanner.Text())
			} else {
				break
			}
		}

		// Empty line
		line = strings.TrimSpace(scanner.Text())
		if len(line) > 0 {
			return nil, &ErrInvalidSubtitle{
				format: SRTFormat,
				reason: fmt.Sprintf("Expected empty line at %v", lineIndex),
			}
		}

		block := NewBlock()
		block.StartAt = startAt
		block.FinishAt = finishAt
		block.Lines = lines

		sub.Blocks = append(sub.Blocks, block)
	}

	return sub, nil
}

func (p *SRTParser) parseDuration(c string) time.Duration {
	duration := strings.Split(c, ":")

	parsedHours, _ := strconv.Atoi(duration[0])
	hours := time.Duration(parsedHours) * time.Hour

	parsedMinutes, _ := strconv.Atoi(duration[1])
	minutes := time.Duration(parsedMinutes) * time.Minute

	parsedSeconds, _ := strconv.Atoi(strings.Split(duration[2], ",")[0])
	seconds := time.Duration(parsedSeconds) * time.Second

	parsedMilliseconds, _ := strconv.Atoi(strings.Split(duration[2], ",")[1])
	milliseconds := time.Duration(parsedMilliseconds) * time.Millisecond

	return hours + minutes + seconds + milliseconds
}

// SRTFormatter subtitle format
type SRTFormatter struct {
	content strings.Builder
}

// NewSRTFormatter returns SRT format instance
func NewSRTFormatter() *SRTFormatter {
	return &SRTFormatter{}
}

// Format SRT subtitle content
func (f *SRTFormatter) Format(sub *Subtitle) (string, error) {
	blockCount := len(sub.Blocks)
	if blockCount == 0 {
		return "", &ErrInvalidSubtitle{
			format: SRTFormat,
			reason: fmt.Sprintf("Empty blocks"),
		}
	}

	lastBlockIndex := blockCount - 1
	for index, block := range sub.Blocks {
		// block index
		blockIndex := strconv.Itoa(index + 1)
		f.content.WriteString(blockIndex)
		f.content.WriteString("\n")

		// duration: hh:mm:ss,fff --> hh:mm:ss,fff
		startAt := f.formatDuration(block.StartAt)
		finishAt := f.formatDuration(block.FinishAt)

		f.content.WriteString(startAt)
		f.content.WriteString(" --> ")
		f.content.WriteString(finishAt)
		f.content.WriteString("\n")

		// lines
		for _, line := range block.Lines {
			f.content.WriteString(line)
			f.content.WriteString("\n")
		}

		// blank line between blocks
		if index != lastBlockIndex {
			f.content.WriteString("\n")
		}
	}

	return f.content.String(), nil
}

func (f *SRTFormatter) formatDuration(d time.Duration) string {
	// hours
	h := d / time.Hour
	// minutes
	d -= h * time.Hour
	m := d / time.Minute
	// seconds
	d -= m * time.Minute
	s := d / time.Second
	// milliseconds
	d -= s * time.Second
	ms := d / time.Millisecond

	return fmt.Sprintf("%02d:%02d:%02d,%03d", h, m, s, ms)
}
