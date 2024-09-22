package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	_ "github.com/mattn/go-sqlite3"
)

func checkerr(err error) {
	if err != nil {
		log.Printf("error: %v", err)
		log.Fatal(err)

	}
}

func initdb() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "notes.db")
	if err != nil {
		return nil, err
	}
	return db, err
}

func main() {
	// colors
	reset := "\033[0m"
	red := "\033[31m"
	// green := "\033[32m"
	// yellow := "\033[33m"
	blue := "\033[34m"
	// purple := "\033[35m"
	// cyan := "\033[36m"
	// gray := "\033[37m"
	// white := "\033[97m"

	// color design
	helpStyle_underline := color.New(color.FgHiCyan).Add(color.Underline).SprintFunc()
	helpStyle_notes_func_command := color.New(color.FgGreen).SprintFunc()

	db, err := initdb()
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
			notes_table.AppendRow(table.Row{red + id + reset, note, blue + created_date + reset, blue + created_time + reset})

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
				fmt.Println(helpStyle_underline("help") + "   - Show this help message")
				fmt.Println(helpStyle_notes_func_command("add") + " -  Add notes (Enter to type notes)")
				fmt.Println(helpStyle_notes_func_command("show") + " -  Show notes  ")
				fmt.Println(helpStyle_notes_func_command("rm") + " -  Delete notes ")
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

			if v == "show" {
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
		}

	}

}
