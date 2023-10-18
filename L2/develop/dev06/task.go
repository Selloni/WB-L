package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type flags struct {
	f int
	d string
	s bool
}

type Args struct {
	fl  flags
	str []string
}

func main() {
	argc := getArgc()
	output, err := readString(argc)
	if err != nil {
		log.Fatalln(err)
	}
	for _, str := range output {
		fmt.Println(str)
	}
}

func readString(argc *Args) ([]string, error) {
	var output []string
	if argc.fl.f == 0 {
		return nil, fmt.Errorf("вы должны задать список полей")
	}
	for _, str := range argc.str {
		isSplit := strings.Contains(str, argc.fl.d)
		arr := strings.Split(str, argc.fl.d)
		if isSplit {
			if len(arr) > argc.fl.f-1 {
				output = append(output, arr[argc.fl.f-1])
			}
		} else if !(argc.fl.s) {
			output = append(output, str)
		}
	}
	return output, nil
}

func getArgc() *Args {
	ff := flag.Int("f", 0, "fields - выбрать поля (колонки)")
	fd := flag.String("d", "\t", "delimiter - использовать другой разделитель")
	fs := flag.Bool("s", false, "separated - только строки с разделителем")
	flag.Parse()
	return &Args{
		fl: flags{
			*ff,
			*fd,
			*fs,
		},
		str: flag.Args(), // передаю все строки после команды и флагов
	}
}
