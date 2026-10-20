package main

import (
	"log"
	"os"

	"github.com/nishankx/torrentis/torrentfile"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("Usage: %s <torrent-file> <output-file>", os.Args[0])
	}

	inPath := os.Args[1]
	outPath := os.Args[2]

	tf, err := torrentfile.Open(inPath)
	if err != nil {
		log.Fatal("Failed to open .torrent file:", err)
	}

	err = tf.DownloadToFile(outPath)
	if err != nil {
		log.Fatal("Failed to download .torrent file:", err)
	}
}