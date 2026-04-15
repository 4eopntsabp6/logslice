// Package input provides abstractions for reading log data line-by-line
// from various sources, including regular files and standard input.
//
// # Usage
//
// Open a file for reading:
//
//	reader, err := input.NewFileReader("/var/log/app.log")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer reader.Close()
//
//	for {
//		line, ok := reader.ReadLine()
//		if !ok {
//			break
//		}
//		fmt.Println(reader.LineNum(), line)
//	}
//
// Read from stdin instead:
//
//	reader := input.NewStdinReader()
//	defer reader.Close()
package input
