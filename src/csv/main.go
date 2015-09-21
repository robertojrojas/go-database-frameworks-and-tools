package main
import (
	"os"
	"log"
	"encoding/csv"
	"fmt"
)

const OUTFILE_NAME string = "src/csv/scraps2.csv"

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <absolute path to file>\n", os.Args[0])
		os.Exit(1)
	}

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Read() // skip headers

	//readAllRecord(reader)
	readOneRecordAtaTime(reader)

	writeCSVToFile()

}

func readAllRecord(reader *csv.Reader) {

	recs, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, row := range recs {
		printRow(row)
	}

}

func readOneRecordAtaTime(reader *csv.Reader) {

	for {
		row, err := reader.Read()
		if err != nil {
			break
		}
		printRow(row)
	}
}

func writeCSVToFile() {

	os.Remove(OUTFILE_NAME)
	file, err := os.Create(OUTFILE_NAME)
	if err != nil {
		log.Fatal(err)
	}

	writer := csv.NewWriter(file)

	data := [][] string{
		[]string{"first", "last", "email"},
		[]string{"roberto", "rojas", "roberto@example.com"},
		[]string{"josette", "rojas", "josette@example.com"},
	}

	writeAllAtOnce(writer, data)
	//writeOneRecordAtaTime(writer, data)
}

func writeAllAtOnce(writer *csv.Writer, data [][]string) {
	writer.WriteAll(data)
	writer.Flush()
}

func writeOneRecordAtaTime(writer *csv.Writer, data [][]string) {

	for _, row := range data {
		writer.Write(row)
	}
	writer.Flush()
}

func printRow(row []string) {
	log.Printf("len(row) %d\n", len(row))
	for i, col := range row {
		log.Printf("[%d]: %s\n", i, col)
	}
}
