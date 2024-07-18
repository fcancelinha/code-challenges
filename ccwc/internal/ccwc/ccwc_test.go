package ccwc_test

import (
	"io/fs"
	"testing"
	"testing/fstest"

	"github.com/fcancelinha/code-challenge-ccwc/internal/ccwc"
)

func TestWordCountBy(t *testing.T) {
	t.Run("Number of Bytes output", func(t *testing.T) {
		files := []fstest.MapFS{
			{"file1": {Data: []byte("lorem")}},
			{"file2": {Data: []byte("ipsum")}},
			{"file3": {Data: []byte("abc")}},
			{"file4": {Data: []byte("")}},
			{"file5": {Data: []byte(" ")}},
			{"file6": {Data: []byte("\n")}},
			{"file7": {Data: []byte("\r\n")}},
			{"file8": {Data: []byte("Sun Tzŭ")}},
			{"file9": {Data: []byte("***")}},
			{"file10": {Data: []byte("\n\n\n\n\n\n")}},
			{"file11": {Data: []byte("Project Gutenberg™")}},
			{"file12": {Data: []byte("Chapter XIII. The Use of Spies")}},
		}

		tt := []struct {
			filename string
			want     int
		}{
			{
				filename: "file1",
				want:     5,
			},
			{
				filename: "file2",
				want:     5,
			},
			{
				filename: "file3",
				want:     3,
			},
			{
				filename: "file4",
				want:     0,
			},
			{
				filename: "file5",
				want:     1,
			},
			{
				filename: "file6",
				want:     1,
			},
			{
				filename: "file7",
				want:     2,
			},
			{
				filename: "file8",
				want:     8,
			},
			{
				filename: "file9",
				want:     3,
			},
			{
				filename: "file10",
				want:     6,
			},
			{
				filename: "file11",
				want:     20,
			},
			{
				filename: "file12",
				want:     30,
			},
		}

		for k, file := range files {
			f, err := fs.ReadFile(file, tt[k].filename)
			if err != nil {
				t.Fatal(err)
			}

			res := ccwc.ByteCount()(f)

			if tt[k].want != res {
				t.Errorf("Got %d , want %d in '%s'", res, tt[k].want, tt[k].filename)
			}
		}
	})

	t.Run("Number of Lines output", func(t *testing.T) {
		files := []fstest.MapFS{
			{"file1": {Data: []byte("\n")}},
			{"file2": {Data: []byte("")}},
			{"file3": {Data: []byte("\n\n")}},
			{"file4": {Data: []byte("\n\n\n")}},
			{"file5": {Data: []byte("\r\n")}},
			{"file6": {Data: []byte("\n          \n")}},
			{"file7": {Data: []byte("\n\n\n\n\n\n")}},
		}

		tt := []struct {
			filename string
			want     int
		}{
			{
				filename: "file1",
				want:     1,
			},
			{
				filename: "file2",
				want:     0,
			},
			{
				filename: "file3",
				want:     2,
			},
			{
				filename: "file4",
				want:     3,
			},
			{
				filename: "file5",
				want:     1,
			},
			{
				filename: "file6",
				want:     2,
			},
			{
				filename: "file7",
				want:     6,
			},
		}

		for k, file := range files {
			f, err := fs.ReadFile(file, tt[k].filename)
			if err != nil {
				t.Fatal(err)
			}

			res := ccwc.LineCount()(f)

			if tt[k].want != res {
				t.Errorf("Got %d , want %d in '%s'", res, tt[k].want, tt[k].filename)
			}
		}
	})

	t.Run("Number of Words output", func(t *testing.T) {
		files := []fstest.MapFS{
			{"file1": {Data: []byte("word")}},
			{"file2": {Data: []byte("word word word")}},
			{"file3": {Data: []byte("word \n word \n word")}},
			{"file4": {Data: []byte("reallylongwordthatperhapsisgerman")}},
			{"file5": {Data: []byte("\rword\rword\rword\rword")}},
			{"file6": {Data: []byte("™™™ ¶¶ üüüçççã ã ã ã ã")}},
		}

		tt := []struct {
			filename string
			want     int
		}{
			{
				filename: "file1",
				want:     1,
			},
			{
				filename: "file2",
				want:     3,
			},
			{
				filename: "file3",
				want:     3,
			},
			{
				filename: "file4",
				want:     1,
			},
			{
				filename: "file5",
				want:     4,
			},
			{
				filename: "file6",
				want:     7,
			},
		}

		for k, file := range files {
			f, err := fs.ReadFile(file, tt[k].filename)
			if err != nil {
				t.Fatal(err)
			}

			res := ccwc.WordCount()(f)

			if tt[k].want != res {
				t.Errorf("Got %d , want %d in '%s'", res, tt[k].want, tt[k].filename)
			}
		}
	})

	t.Run("Number of Characters output", func(t *testing.T) {
		files := []fstest.MapFS{
			{"file1": {Data: []byte("word")}},
			{"file2": {Data: []byte("word word word")}},
			{"file3": {Data: []byte("word \n word \n word")}},
			{"file4": {Data: []byte("reallylongwordthatperhapsisgerman")}},
			{"file5": {Data: []byte("\rword\rword\rword\rword")}},
			{"file6": {Data: []byte("™™™ üüüçççã ã ã ã ã")}},
		}
		tt := []struct {
			filename string
			want     int
		}{
			{
				filename: "file1",
				want:     4,
			},
			{
				filename: "file2",
				want:     14,
			},
			{
				filename: "file3",
				want:     18,
			},
			{
				filename: "file4",
				want:     33,
			},
			{
				filename: "file5",
				want:     20,
			},
			{
				filename: "file6",
				want:     19,
			},
		}

		for k, file := range files {
			f, err := fs.ReadFile(file, tt[k].filename)
			if err != nil {
				t.Fatal(err)
			}

			c := ccwc.CharCount()(f)

			if tt[k].want != c {
				t.Errorf("Got %d , want %d in '%s'", c, tt[k].want, tt[k].filename)
			}
		}
	})
}
