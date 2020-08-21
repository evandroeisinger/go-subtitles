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
func (p *SRTParser) Parse(r io.Reader) (*Subtitle, error) {
	var line string
	var lineIndex int
	var blockIndex int

	subtitle := NewSubtitle()
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

		subtitle.Blocks = append(subtitle.Blocks, block)
	}

	return subtitle, nil
}

func (p *SRTParser) parseDuration(content string) time.Duration {
	duration := strings.Split(content, ":")

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

// SRTWriter subtitle format
type SRTWriter struct{}

// NewSRTWriter returns SRT format instance
func NewSRTWriter() *SRTWriter {
	return &SRTWriter{}
}

// Write SRT subtitle content
func (w *SRTWriter) Write(s *Subtitle) string {
	var content strings.Builder

	lastLineIndex := len(s.Blocks) - 1
	for index, block := range s.Blocks {
		blockIndex := strconv.Itoa(index + 1)
		content.WriteString(blockIndex)
		content.WriteString("\n")

		startAt := w.formatDuration(block.StartAt)
		finishAt := w.formatDuration(block.FinishAt)

		content.WriteString(startAt)
		content.WriteString(" --> ")
		content.WriteString(finishAt)
		content.WriteString("\n")

		for _, line := range block.Lines {
			content.WriteString(line)
			content.WriteString("\n")
		}

		if index != lastLineIndex {
			content.WriteString("\n")
		}
	}

	return content.String()
}

func (w *SRTWriter) formatDuration(d time.Duration) string {
	h := d / time.Hour

	d -= h * time.Hour
	m := d / time.Minute

	d -= m * time.Minute
	s := d / time.Second

	d -= s * time.Second
	ms := d / time.Millisecond

	return fmt.Sprintf("%02d:%02d:%02d,%03d", h, m, s, ms)
}
