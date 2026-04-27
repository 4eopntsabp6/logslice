// Package window implements a sliding-window line filter for logslice.
//
// When mode is "sliding", the filter anchors on the timestamp of the first
// line it processes and keeps every subsequent line whose timestamp falls
// within [anchor, anchor+duration). Lines that carry no parseable timestamp
// are kept only before the anchor is established.
//
// Usage:
//
//	f, err := window.New(window.ModeSliding, 30*time.Second)
//	if err != nil { ... }
//	for _, line := range lines {
//	    if f.Keep(line) {
//	        fmt.Println(line)
//	    }
//	}
//
// Call Reset to clear the anchor and restart the window from the next
// timestamped line.
package window
