package sseclient

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Event represents a Server-Sent Event
type Event struct {
	Name string
	ID   string
	Data map[string]interface{}
}

// OpenURL opens a connection to a stream of server sent events
func OpenURL(rawurl string) (events chan Event, err error) {
	resp, err := get(rawurl)
	if err != nil {
		return nil, err
	}

	events = make(chan Event)

	var buf bytes.Buffer

	go func() {
		ev := Event{}

		reader := bufio.NewReader(resp.Body)

		for {
			line, err := reader.ReadBytes('\n')
			if err != nil {
				fmt.Fprintf(os.Stderr, "error during resp.Body read:%s\n", err)

				close(events)
			}

			switch {
			// OK line
			case hasPrefix(line, ":ok"):
				// Do nothing

			// id of event
			case hasPrefix(line, "id: "):
				ev.ID = string(line[4:])
			case hasPrefix(line, "id:"):
				ev.ID = string(line[3:])

			// name of event
			case hasPrefix(line, "event: "):
				ev.Name = string(line[7 : len(line)-1])
			case hasPrefix(line, "event:"):
				ev.Name = string(line[6 : len(line)-1])

			// event data
			case hasPrefix(line, "data: "):
				buf.Write(line[6:])
			case hasPrefix(line, "data:"):
				buf.Write(line[5:])

			// end of event
			case bytes.Equal(line, []byte("\n")):
				b := buf.Bytes()

				if hasPrefix(b, "{") {
					var data map[string]interface{}

					err := json.Unmarshal(b, &data)

					if err == nil {
						ev.Data = data
						buf.Reset()
						events <- ev
						ev = Event{}
					}
				}

			default:
				fmt.Fprintf(os.Stderr, "Error: len:%d\n%s", len(line), line)

				close(events)
			}
		}
	}()

	return events, nil
}

func get(rawurl string) (*http.Response, error) {
	resp, err := http.Get(rawurl)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("got response status code %d\n", resp.StatusCode)
	}

	return resp, nil
}

func hasPrefix(s []byte, prefix string) bool {
	return bytes.HasPrefix(s, []byte(prefix))
}
