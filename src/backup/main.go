package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func main() {

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	str1 := `
	CREATE TYPE user_role AS ENUM('admin', 'member', 'memberIsTarget', 'anonym');

-- schema for roles table
CREATE TABLE roles(r_user_id UUID NOT NULL REFERENCES users, r_role user_role NOT NULL,r_created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),PRIMARY KEY(r_user_id, r_role));

-- create index for roles
CREATE INDEX index_roles_table ON roles(r_user_id, r_user_id)
	`
	str := `
	CREATE TYPE user_role AS ENUM(
'admin'
-- 'member',
-- 'memberIsTarget',
-- 'anonym',
); 

-- schema for roles table
CREATE TABLE roles(
r_user_id UUID NOT NULL REFERENCES users,
r_role user_role NOT NULL,
r_created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
PRIMARY KEY(r_user_id, r_role)
); 

-- create index for roles
CREATE INDEX index_roles_table
ON roles(r_user_id, r_user_id)
`
	fmt.Printf("str: %v\n", str)
	fmt.Printf("str1: %v\n", str1)

	var tableName string
	// var tableComments string
	var tableFields []string
	var typeName string
	// var typeComments string
	// var typels []string
	var typeFields []string
	var indexName string
	var indexFields []string
	// var indexls []string

	fmt.Print("=======================================================================\n")
	// var lines []string
	jsonString := strings.Split(string(str1), ";")
	fmt.Printf("jsonString: %v\n", len(jsonString))
	vals := []string{"TABLE ", "TYPE ", "INDEX "}

	for i, j := range jsonString {
		jsonString[i] = j + ";"
	}
	lines := strings.Split(strings.Join(jsonString, "#"), "\n")
	// fmt.Printf("lines: %q\n", lines)
	for i, line := range lines {
		if strings.Contains(line, "--") {
			lines[i] = ""
		}
	}
	var delNewLines []string
	lines1 := delete_empty(lines)
	// fmt.Printf("lines1: %q\n", lines1)
	for _, v := range lines1 {
		delNewLines = append(delNewLines, strings.TrimSpace(strings.Replace(v, "\t", " ", -1)))
		// fmt.Printf("v: %q\n", strings.Replace(v, "\t", " ", -1))
	}
	jsonString = strings.Split(strings.Join(delNewLines, "@"), "#")

	// fmt.Printf("jsonString: %v\n", jsonString)

	for i, v := range jsonString {
		if strings.Contains(v, "@)") {
			jsonString[i] = strings.Replace(v, "@)", ")", -1)
		}
	}
	for i, v := range jsonString {
		if strings.Contains(v, "(@") {
			jsonString[i] = strings.Replace(v, "(@", "(", -1)
		}
	}
	for i, v := range jsonString {
		if strings.Contains(v, ",@") {
			jsonString[i] = strings.Replace(v, ",@", ", ", -1)
		}
	}
	for i, v := range jsonString {
		if strings.Contains(v, "@") {
			jsonString[i] = strings.Replace(v, "@", " ", -1)
		}
	}
	for _, pos := range vals {
		for _, jpos := range jsonString {
			switch pos {
			case "TABLE ":
				if strings.Contains(jpos, "TABLE ") {
					ls := delete_empty(strings.Split(jpos, " "))
					// fmt.Printf("ls: %v\n", ls)
					for i := range ls {
						// fmt.Printf("v: %v\n", v)
						if ls[i] == strings.TrimSpace(pos) {
							tableName = before(strings.TrimSpace(ls[i+1]), "(") //strings.Replace(strings.TrimSpace(ls[i+1]), "(", "", -1)
							// fmt.Printf("tableName: %v\n", tableName)
						} else if strings.Contains(ls[i], tableName+"(") { //ls[i] == strings.TrimSpace("ENUM(") { //ls[i+len(ls)]
							// fmt.Printf("recurcive(i+1, len(ls), ls): %q\n", recurcive(i+1, len(ls), ls))
							for j, _ := range ls[i:] {
								if !strings.Contains(ls[i], ");") {
									// fmt.Printf("i: %v\n", ls[i+j])
									tableFields = append(tableFields, strings.TrimSpace(ls[i+j]))
								}
							}
						}
					}
					// fmt.Printf("strings.Join(tableFields, \"\"): %q\n", strings.Join(tableFields, " "))
					// lines := strings.Split((strings.Join(tableFields, " ")), ",")
					lines := strings.Split(before(after(strings.Join(tableFields, " "), tableName), "PRIMARY KEY("), ",")
					// lines := strings.Split(after(before(strings.Join(tableFields, " "), ",PRIMARY KEY("), tableName), ",")
					for i, j := range lines {
						lines[i] = j + ","
					}
					for _, v := range lines {
						if len(v) > 4 {
							lines := delete_empty(strings.Split(strings.Replace(v, "(", "", -1), " "))
							fmt.Printf("lines: %q\n", (lines[:2]))
							// fmt.Printf("v: %v\n", v[:2])
						}
					}
					// fmt.Printf("lines: %q\n\n\n", (lines))

				}
			case "TYPE ":
				if strings.Contains(jpos, "TYPE ") {
					ls := delete_empty(strings.Split(jpos, " "))
					for i := range ls {
						// fmt.Printf("v: %v\n", v)
						if ls[i] == strings.TrimSpace(pos) {
							typeName = strings.Replace(strings.TrimSpace(ls[i+1]), "(", "", -1)
							// fmt.Printf("typeName: %v\n", typeName)
						} else if strings.Contains(ls[i], "ENUM") { //ls[i] == strings.TrimSpace("ENUM(") { //ls[i+len(ls)]
							for j, _ := range ls[i:] {
								if !strings.Contains(ls[i], ");") {
									// fmt.Printf("i: %v\n", ls[i+j])
									typeFields = append(typeFields, strings.TrimSpace(ls[i+j]))
								} else if strings.Contains(ls[i], "ENUM") {
									typeFields = append(typeFields, strings.TrimSpace(ls[i]))
								}
							}

						}
					}
					linesType := strings.Split((after(strings.Join(typeFields, " "), "ENUM")), ",")
					// lines := strings.Split(after(before(strings.Join(tableFields, " "), ",PRIMARY KEY("), tableName), ",")
					fmt.Printf("linesType: %v\n", linesType)
					for i, j := range linesType {
						linesType[i] = j + ","
					}
					for _, v := range linesType {
						fmt.Printf("v: %q\n", reg.ReplaceAllString(strings.Join(delete_empty(strings.Split(strings.Replace(v, "(", "", -1), " ")), ""), ""))
						// if len(v) > 0 {
						// 	linesType := delete_empty(strings.Split(strings.Replace(v, "(", "", -1), " "))
						// 	fmt.Printf("lines: %q\n", (linesType[:2]))
						// 	// 	// fmt.Printf("v: %v\n", v[:2])
						// }
					}
				}
			case "INDEX ":
				if strings.Contains(jpos, "INDEX ") {
					ls := delete_empty(strings.Split(jpos, " "))
					for i := range ls {
						if ls[i] == strings.TrimSpace(pos) {
							indexName = strings.TrimSpace(ls[i+1])
						} else if ls[i] == strings.TrimSpace("ON") { //ls[i+len(ls)]
							for j, _ := range ls[i:] {
								if !strings.Contains(ls[i], ");") {
									indexFields = append(indexFields, strings.TrimSpace(ls[i+j]))
								} else if strings.Contains(ls[i], "ON") {
									indexFields = append(indexFields, strings.TrimSpace(ls[i]))
								}
							}
						}
					}
				}
			}
		}
	}
	fmt.Printf("tableName: %v\n", tableName)
	fmt.Printf("tableFields: %v\n", tableFields)
	fmt.Printf("typeName: %v\n", typeName)
	fmt.Printf("typeFields: %v\n", typeFields)
	fmt.Printf("indexName: %v\n", indexName)
	fmt.Printf("indexFields: %v\n", indexFields)
	// fmt.Printf("jsonString: %v\n", (jsonString))
	// fmt.Printf("delNewLines: %q\n", (strings.Join(delNewLines, "@")))
	// fmt.Printf("lines: %v\n", len(lines))
	// fmt.Printf("jsonString: %v\n", jsonString)
	// lines := strings.Split(jpos, "\n")

	// min, max := MinMax(str, "ON")
	// st := betweenSpecial(str, "INDEX ", "ON", min)
	// fmt.Printf("st: %v\n", st)
	// // jsonString := strings.Split(string(str), "CREATE ")
	// // jsonString = delete_empty(jsonString)
	// // fmt.Printf("jsonString: %v\n", len(jsonString[1:]))
	// // posFirst := []string{"TABLE ", "TYPE ", "INDEX "}
	// // for _, pos := range posFirst {
	// // 	for _, jpos := range jsonString[1:] {
	// // 		// fmt.Printf("jsonString: %v\n", len(jpos))
	// // 		// fmt.Printf("jpos:<<%v>> %v\n", (jpos), pos)
	// // 		switch pos {
	// // 		case "TABLE ":
	// // 			// fmt.Printf("pos: %v\n", pos)
	// // 			fmt.Printf("jpos: %v\n", len(jpos))
	// // 			// fmt.Printf("jpos: %v\n", (jpos))
	// // 		}
	// // 	}
	// // }
	// fmt.Printf("min: %v<<max: %v>>\n", min, max)
	// // test := between(str, "TABLE ", "u")
	// fmt.Printf("test: %v\n", test)
	// posFirst := strings.Index(str, "(")
	// if posFirst == -1 {
	// 	// return ""
	// }
	// fmt.Printf("posFirst: %v\n", posFirst)
	// posLast := strings.Index(str, "(")
	// if posLast == -1 {
	// 	// return ""
	// }
	// posFirstAdjusted := posFirst + len("TABLE")
	// if posFirstAdjusted >= posLast {
	// 	// return ""
	// }
	// fmt.Printf("posFirst: %v\n", str[posFirstAdjusted:posLast])
}

func betweenSpecial(value string, a string, b string, size int) string {
	// Get substring between two strings.
	posFirst := strings.Index(value, a)
	if posFirst == -1 {
		return ""
	}
	// fmt.Printf("posFirst: %v\n", posFirst)
	var posLast int
	var nextChr string
	if posFirst+size <= len(value) {
		nextChr = value[posFirst : posFirst+size]
		// fmt.Printf("size: %v\n", size)
		// fmt.Printf("nextChr: %v\n", nextChr)
	} else if posFirst+size >= len(value) {
		// fmt.Printf("nextChr: %v\n", nextChr)
		nextChr = value[posFirst:size]
	}
	// fmt.Printf("nextChr: %v\n", nextChr)
	if strings.Contains(nextChr, b) { // b = "("
		posLast = strings.Index(nextChr, b)
		if posLast == -1 {
			return ""
		}
	}
	posFirstAdjusted := posFirst + len(a)
	if posFirstAdjusted >= posFirst+posLast {
		str := fmt.Sprintf("<<<\nposFirst: %d\nposFirstAdjusted: %d\nposFirst+posLast: %d>>>", posFirst, posFirstAdjusted, posFirst+posLast)
		return str //fmt.Sprintf("\nposFirst: %d\nposFirstAdjusted: %d\nposFirst+posLast: %d", posFirst, posFirstAdjusted, posFirst+posLast)
	}

	return value[posFirstAdjusted : posFirst+posLast]
}
func delete_empty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
func MinMax(str string, char string) (min int, max int) {
	var input []int
	for i, j := range strings.Split(str, "") {
		// fmt.Printf("j: %v\n", j)
		if strings.Contains(j, char) {
			fmt.Printf("j: %v\n", j)
			input = append(input, i)
		}
	}
	if len(char) > 1 {
		min = len(char)
		max = len(char)
	} else if len(char) <= 1 {
		min = input[0]
		max = input[0]
	}

	for _, value := range input {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}
func between(value string, a string, b string) string {
	// Get substring between two strings.
	posFirst := strings.Index(value, a)
	if posFirst == -1 {
		return ""
	}
	posLast := strings.Index(value, b)
	if posLast == -1 {
		return ""
	}
	posFirstAdjusted := posFirst + len(a)
	if posFirstAdjusted >= posLast {
		return ""
	}
	return value[posFirstAdjusted:posLast]
}

func after(value string, a string) string {
	// Get substring after a string.
	pos := strings.LastIndex(value, a)
	if pos == -1 {
		return ""
	}
	adjustedPos := pos + len(a)
	if adjustedPos >= len(value) {
		return ""
	}
	return value[adjustedPos:len(value)]
}

func before(value string, a string) string {
	// Get substring before a string.
	pos := strings.Index(value, a)
	if pos == -1 {
		return ""
	}
	return value[0:pos]
}
