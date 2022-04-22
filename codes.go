// read codes.txt
// (c)copyright 2022 by Gerald Wodni <gerald.wodni@gmail.com>

package main

import (
    "bufio"
    "errors"
    "fmt"
    "os"
    "strings"
)

type StatusCode struct {
    Code string
    MessageEN string
    DescriptionEN string
    MessageDE string
    DescriptionDE string
}

func (statusCode StatusCode) String() string {
    return fmt.Sprintf( "%s: '%s' \"%s\"\n     '%s' \"%s\"",
        statusCode.Code, statusCode.MessageEN, statusCode.DescriptionEN,
        statusCode.MessageDE, statusCode.DescriptionDE )
}

type StatusCodeMap map[string]StatusCode

var statusCodes StatusCodeMap = StatusCodeMap{}

func init() {
    readCodes("codes.txt")
}

func readCodes( filename string ) (err error) {
    file, err := os.Open( filename )
    if err != nil {
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner( file )
    scanner.Split( bufio.ScanLines )
    lines := []string{}
    for scanner.Scan() {
        lines = append( lines, scanner.Text() )
    }

    i := 0
    for i < len(lines) {
        // 1st line contains code, then English message
        firstLine := lines[i]; i++
        parts := strings.SplitN( firstLine, " ", 2 )
        code := parts[0]
        messageEN := parts[1]

        // 2nd line contains other English description
        descriptionEN := lines[i]; i++

        // repeat for German, no code here though
        messageDE := lines[i]; i++
        descriptionDE := lines[i]; i++

        statusCode := StatusCode{
            code,
            messageEN,
            descriptionEN,
            messageDE,
            descriptionDE,
        }
        statusCodes[ code ] = statusCode

        if i < len(lines) {
            if lines[i] != "" {
                err = errors.New( "codes.txt parser expects empty line between status codes" )
                return
            }
            i++
        }
    }
    return
}

func GetStatusCode( code string ) StatusCode {
    if value, ok := statusCodes[ code ]; ok {
        return value
    }
    return statusCodes["default"]
}

