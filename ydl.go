package ydl

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/docker/go-units"
)

type Ydl struct {
	Binary  string
	Options options
}

type Progress struct {
	Percentage float64
	Rate       int64
	Total      int64
	ETA        time.Duration
}

func NewYdl() Ydl {
	return Ydl{
		Binary:  "youtube-dl",
		Options: NewOptions(),
	}
}

func (ydl *Ydl) Download(u string) (chan Progress, error) {
	cmd := exec.Command(ydl.Binary, u) //nolint:gosec
	cmd.Args = append(cmd.Args, strings.Fields(ydl.Options.OptionsToCliParameters())...)

	_, err := cmd.StderrPipe()
	if err != nil {
		return nil, err
	}

	cout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	ch := make(chan Progress)
	go func() {
		scanner := bufio.NewScanner(cout)
		scanner.Split(bufio.ScanWords)

		p := Progress{}
		exp := 0

		for scanner.Scan() {
			m := scanner.Text()
			// fmt.Println(m)

			// [download]
			// 22.0%
			// of
			// 89.64MiB
			// at
			// 499.09KiB/s
			// ETA
			// 02:23

			if exp == 0 {
				switch m {
				case "[download]":
					exp = 1
				case "of":
					exp = 2
				case "at":
					exp = 3
				case "ETA":
					exp = 4
				}
			} else {
				switch exp {
				case 1:
					// percentage
					f, err := strconv.ParseFloat(strings.TrimSuffix(m, "%"), 64)
					if err == nil {
						p.Percentage = f / 100
					}

				case 2:
					// total size
					if total, err := units.FromHumanSize(m); err == nil {
						p.Total = total
					}

				case 3:
					// download rate
					if rate, err := units.FromHumanSize(strings.TrimSuffix(m, "/s")); err == nil {
						p.Rate = rate
					}

				case 4:
					// eta
					d := strings.Split(m, ":")
					switch len(d) {
					case 1:
						m = d[0] + "s"
					case 2:
						m = d[0] + "m" + d[1] + "s"
					case 3:
						m = d[0] + "h" + d[1] + "m" + d[2] + "s"
					}

					if eta, err := time.ParseDuration(m); err == nil {
						p.ETA = eta
					}

					ch <- p
				}

				exp = 0
			}
		}

		_ = cmd.Wait()
		close(ch)
	}()

	return ch, nil
}

// FetchInfo retrieves the metadata for a given URL.
func (ydl *Ydl) FetchInfo(ctx context.Context, u string) (Info, error) {
	// #nosec G204
	cmd := exec.CommandContext(ctx, ydl.Binary, "-J", u) //nolint:gosec
	cout, err := cmd.StdoutPipe()
	if err != nil {
		return Info{}, err
	}

	if err := cmd.Start(); err != nil {
		return Info{}, err
	}

	var info Info
	if err := json.NewDecoder(cout).Decode(&info); err != nil {
		// We need to wait for the command to exit, so we don't end up creating
		// defunct processes.
		_ = cmd.Wait()
		return Info{}, err
	}

	return info, cmd.Wait()
}

// Search returns foobar.
func (ydl *Ydl) Search(ctx context.Context, term string, amount uint) (SearchResult, error) {
	// #nosec G204
	cmd := exec.CommandContext(ctx, ydl.Binary,
		"-J",
		fmt.Sprintf("ytsearch%d:%s", amount, term),
	)
	cout, err := cmd.StdoutPipe()
	if err != nil {
		return SearchResult{}, err
	}

	if err := cmd.Start(); err != nil {
		return SearchResult{}, err
	}

	var result SearchResult
	if err := json.NewDecoder(cout).Decode(&result); err != nil {
		// We need to wait for the command to exit, so we don't end up creating
		// defunct processes.
		_ = cmd.Wait()
		return SearchResult{}, err
	}

	return result, cmd.Wait()
}
