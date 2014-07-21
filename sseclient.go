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
func OpenURL(url string) (events chan Event, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("got response status code %d\n", resp.StatusCode)
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
			case bytes.HasPrefix(line, []byte(":ok")):
				// Do nothing

			// id of event
			case bytes.HasPrefix(line, []byte("id:")):
				ev.ID = string(line[4:])

			// name of event
			case bytes.HasPrefix(line, []byte("event:")):
				ev.Name = string(line[7 : len(line)-1])

			// event data
			case bytes.HasPrefix(line, []byte("data:")):
				buf.Write(line[6:])

			// end of event
			case bytes.Equal(line, []byte("\n")):
				b := buf.Bytes()

				if bytes.HasPrefix(b, []byte("{")) {
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
