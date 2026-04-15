// Package tail implements real-time log following for logslice.
//
// It provides a Tailer that opens a file, seeks to the end, and
// continuously polls for new lines, sending them over a channel.
// Callers control the lifetime of following via a context.Context;
// cancelling the context stops the goroutine and closes the channel.
//
// Example usage:
//
//	tr := tail.New("/var/log/app.log", tail.DefaultPollInterval)
//	lines, err := tr.Follow(ctx)
//	if err != nil {
//		log.Fatal(err)
//	}
//	for line := range lines {
//		fmt.Println(line)
//	}
package tail
