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
	lineIndex := 0
	blockIndex := 0

	subtitle := NewSubtitle()
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		blockIndex++
		lineIndex++

		// Remove BOM
		lineContent := strings.TrimSpace(scanner.Text())
		if lineIndex == 1 {
			lineContent = StripBOM(lineContent)
		}

		// Block index
		if index, _ := strconv.Atoi(lineContent); index != blockIndex {
			return NewSubtitle(), &ErrInvalidSubtitle{
				format: SRTFormat,
				reason: fmt.Sprintf("Expected index %v at line %v got: %v", blockIndex, lineIndex, lineContent),
			}
		}

		// Next line
		scanner.Scan()
		lineIndex++

		// Block duration: hh:mm:ss,fff --> hh:mm:ss,fff
		lineContent = strings.TrimSpace(scanner.Text())
		blockDurations := SRTBlockDurationPattern.FindStringSubmatch(lineContent)
		if len(blockDurations) == 0 {
			return NewSubtitle(), &ErrInvalidSubtitle{
				format: SRTFormat,
				reason: fmt.Sprintf("Expected duration with pattern (hh:mm:ss,fff --> hh:mm:ss,fff) at line %v got: %v", lineIndex, lineContent),
			}
		}

		startAt := p.parseDuration(blockDurations[1])
		finishAt := p.parseDuration(blockDurations[2])

		// Block lines
		scanner.Scan()
		lineIndex++

		lineContent = strings.TrimSpace(scanner.Text())
		if len(lineContent) == 0 {
			return NewSubtitle(), &ErrInvalidSubtitle{
				format: SRTFormat,
				reason: fmt.Sprintf("Expected text at line %v got: empty line", lineIndex),
			}
		}

		lines := []string{}
		for {
			if len(lineContent) > 0 {
				lines = append(lines, lineContent)

				scanner.Scan()
				lineIndex++

				lineContent = strings.TrimSpace(scanner.Text())
			} else {
				break
			}
		}

		// Empty line
		lineContent = strings.TrimSpace(scanner.Text())
		if len(lineContent) > 0 {
			return NewSubtitle(), &ErrInvalidSubtitle{
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
