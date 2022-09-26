package main

import (
	"alukart32.com/tools"
	"alukart32.com/usage/internal/csv"
	"alukart32.com/usage/internal/io/files"
)

func main() {
	tools.TopicConsoleLogger.Log("I/O: fs")
	tools.SubTopicConsoleLogger.Log("I/O: read file and write to os.Stdout")
	files.ReadFile([]byte("../assets/data.txt"))
	files.ReadBufferFile("../assets/data.txt")
	tools.SubTopicConsoleLogger.Log("I/O: seek file and read")
	files.ReadSeekFile("../assets/data1.txt")
	tools.SubTopicConsoleLogger.Log("I/O: io.ReadAtLeast")
	files.ReadAtLeast("../assets/data1.txt")
	tools.SubTopicConsoleLogger.Log("I/O: bufio.Peek")
	files.BufPeek("../assets/data1.txt")
	tools.SubTopicConsoleLogger.Log("I/O: copy src file to dst file")
	files.Copy("../assets/dst.txt", "../assets/src.txt")
	tools.SubTopicConsoleLogger.Log("I/O: create dir if not exists")
	files.MkdirIfNotExists("../pkg")
	tools.SubTopicConsoleLogger.Log("I/O: flush to file")
	files.FlushToFile("This message was flushed to a file")
	tools.SubTopicConsoleLogger.Log("I/O: create and write to a file")
	files.CreateAndFlushToFile("../assets/wr1.txt", []byte("string"))
	tools.SubTopicConsoleLogger.Log("I/O: write to a file")
	files.WriteString([]int{62, 61, 61})
	tools.SubTopicConsoleLogger.Log("I/O: buff write to a file")
	files.BufWrite("aaa")

	// tools.TopicConsoleLogger.Log("I/O: scan")
	// tools.SubTopicConsoleLogger.Log("I/O: line filter (to exit \\q)")
	// scan.ConsoleLineFilter()
	// tools.SubTopicConsoleLogger.Log("I/O: scan console input (to exit \\q)")
	// scan.ScanFromConsole()

	tools.TopicConsoleLogger.Log("CSV files")
	tools.SubTopicConsoleLogger.Log("CSV files : read custom file")
	csv.ReadCustomCSVFile("../assets/csv.txt")
	tools.SubTopicConsoleLogger.Log("CSV files : write custom file")
	csv.WriteCustomCSVFile("../assets/csv1.txt")
}
