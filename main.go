package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	_ "github.com/mattn/go-sqlite3"
)

func checkerr(err error) {
	if err != nil {
		log.Fatal(err)

	}
}

func initdb(user_filepath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", user_filepath)
	if err != nil {
		return nil, err
	}

	return db, err

}

func getFilePath() (string, error) {
	var user_filepath string

	homedir, err := os.UserHomeDir()
	checkerr(err)

	switch runtime.GOOS {
	case "windows", "darwin", "linux":
		user_filepath = filepath.Join(homedir, "db", "notes.db")
	default:
		return "", fmt.Errorf("unspported platform")
	}
	return user_filepath, nil

}
func createFile() string {
	user_filepath, err := getFilePath()
	checkerr(err)

	err = os.MkdirAll(filepath.Dir(user_filepath), 0755)
	checkerr(err)

	if _, err := os.Stat(user_filepath); os.IsNotExist(err) {
		file, err := os.Create(user_filepath)
		checkerr(err)
		defer file.Close()

	}
	return user_filepath
}
func main() {
	// colors
	reset := "\033[0m"
	red := "\033[31m"
	blue := "\033[34m"
	i := 1
	// color design
	helpStyle_underline := color.New(color.FgHiCyan).SprintFunc()
	helpStyle_working_command := color.New(color.FgWhite).SprintFunc()

	user_filepath := createFile()
	db, err := initdb(user_filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	statement, err := db.Prepare(`create table if not exists notes(id integer not null primary key autoincrement , note text ,  created_date  date  , created_time time default current_time );`)
	checkerr(err)
	statement.Exec()

	if len(os.Args) == 1 {
		rows, err := db.Query("select id, note, created_date, created_time from notes")
		checkerr(err)
		headerStyle := color.New(color.FgCyan, color.Bold).SprintFunc()

		notes_table := table.NewWriter()
		notes_table.SetOutputMirror(os.Stdout)
		notes_table.AppendHeader(table.Row{
			headerStyle("#"),
			headerStyle("Note"),
			headerStyle("Created Date"),
			headerStyle("Created Time"),
		})

		var id string
		var note string

		var created_date string
		var created_time string
		for rows.Next() {
			rows.Scan(&id, &note, &created_date, &created_time)

			if len(created_date) >= 10 {
				created_date = created_date[:10]
			}
			notes_table.AppendRow(table.Row{red + strconv.Itoa(i) + reset, note, blue + created_date + reset, blue + created_time + reset})
			i += 1
		}
		notes_table.SetStyle(table.StyleRounded)
		notes_table.Render()

		// .Add(color.Underline)
		c := color.New(color.FgRed).Add(color.Bold)
		c.Println("Try 'notes help' for usage.")

	}

	if len(os.Args) > 1 {

		for _, v := range os.Args {

			if v == "help" {
				fmt.Println("Usage:")
				fmt.Println("  notes [command]")
				fmt.Println()
				fmt.Println("Available commands:")
				fmt.Println(helpStyle_working_command("help ") + " - " + helpStyle_underline("    		Show this help message"))
				fmt.Println(helpStyle_working_command("add ") + " - " + helpStyle_underline("    		Add notes (Enter to type notes)"))
				fmt.Println(helpStyle_working_command("ls ") + " - " + helpStyle_underline("    		Lists all the notes  "))
				fmt.Println(helpStyle_working_command("rm [note id] ") + " - " + helpStyle_underline("    	Delete notes "))
				fmt.Println(helpStyle_working_command("done ") + " - " + helpStyle_underline("    		Delete when done with one session "))
				return
			}

			if v == "rm" {

				note_delete, err := db.Prepare("DELETE FROM notes WHERE id = ?")
				checkerr(err)
				_, err = note_delete.Exec(os.Args[2])
				checkerr(err)

			}

			if v == "add" {
				reader := bufio.NewReader(os.Stdin)
				fmt.Print(":")
				user_notes, err := reader.ReadString('\n')
				checkerr(err)

				currentTime := time.Now()
				user_current_time := time.Now().Format("3:4 pm")

				date := currentTime.Format("2006-01-02")

				statement, err = db.Prepare("INSERT INTO notes (note , created_date, created_time) VALUES (?,?,?)")
				checkerr(err)
				_, execErr := statement.Exec(user_notes, date, user_current_time)
				checkerr(execErr)

			}

			if v == "ls" {
				rows, err := db.Query("select id, note, created_date, created_time from notes")
				checkerr(err)
				headerStyle := color.New(color.FgCyan, color.Bold).SprintFunc()

				notes_table := table.NewWriter()
				notes_table.SetOutputMirror(os.Stdout)
				notes_table.AppendHeader(table.Row{
					headerStyle("#"),
					headerStyle("Note"),
					headerStyle("Created Date"),
					headerStyle("Created Time"),
				})

				var id string
				var note string

				var created_date string
				var created_time string
				for rows.Next() {
					rows.Scan(&id, &note, &created_date, &created_time)

					if len(created_date) >= 10 {
						created_date = created_date[:10]
					}
					notes_table.AppendRow(table.Row{red + id + reset, note, blue + created_date + reset, blue + created_time + reset})

				}
				notes_table.SetStyle(table.StyleRounded)
				notes_table.Render()

			}
			if v == "done" {
				err := db.Close()
				checkerr(err)
				e := os.Remove(user_filepath)
				checkerr(e)
			}
		}

	}

}
