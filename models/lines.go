package models

import (
	"bufio"
	"github.com/yangyouwei/newrouter/conf"
	"github.com/yangyouwei/newrouter/db"
	"io"
	"log"
	"os"
	"strings"
)

type Line struct {
	Id      int    `json:"rowid"`
	Ipaddr  string `json:"ipaddr"`
	Comment string `json:"comment"`
}

func (l *Line) AddLines() (id int64, err error) {
	rs, err := db.SqlDB.Exec("INSERT INTO lines(ipaddr, port, comment) VALUES (?, ?, ?)", l.Ipaddr, l.Comment)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}

func (l *Line) GetLines() (lines []Line, err error) {
	lines = make([]Line, 0)
	rows, err := db.SqlDB.Query("SELECT rowid, * FROM lines")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var line Line
		rows.Scan(&line.Id, &line.Ipaddr,&line.Comment)
		lines = append(lines, line)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}

func (l *Line) GetLine() (lines Line, err error) {
	err = db.SqlDB.QueryRow("SELECT rowid, * FROM lines WHERE rowid=?", l.Id).Scan(&lines.Id, &lines.Ipaddr,&lines.Comment )
	return
}

func (l *Line) ModLine() (ra int64, err error) {
	stmt, err := db.SqlDB.Prepare("UPDATE lines SET ipaddr=? , comment=? WHERE rowid=?")
	defer stmt.Close()
	if err != nil {
		return
	}
	rs, err := stmt.Exec(l.Ipaddr, l.Comment,l.Id)
	if err != nil {
		return
	}
	ra, err = rs.RowsAffected()
	return
}

func (l *Line) DelLine() (ra int64, err error) {
	rs, err := db.SqlDB.Exec("DELETE FROM lines WHERE rowid=?", l.Id)
	if err != nil {
		log.Fatalln(err)
	}
	ra, err = rs.RowsAffected()
	return
}

func (l *Line)GetUseLine() {
	f, err := os.Open(conf.RDIRCONF)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		ls, err := rd.ReadString('\n') //以'\n'为结束符读入一行

		if err != nil || io.EOF == err {
			break
		}
		if strings.HasPrefix(string(ls), "Servers=") {
			ipaddport := strings.Split(ls,"=")
			ip := strings.Replace(ipaddport[1], "\n", "", -1)
			l.Ipaddr = ip
		}
	}
}