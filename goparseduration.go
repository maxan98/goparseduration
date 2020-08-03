package goparseduration

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const hoursinweek = 168
const hoursinday = 24
type ExtendedDuration struct {
	weeks   int
	days    int
	hours   int
	minutes int
	seconds int
}


func ParseDuration(str string) (*ExtendedDuration, error) {
	var r = regexp.MustCompile(`(?P<weeks>\d+w)?(?P<days>\d+d)?(?P<hours>\d+h)?(?P<minutes>\d+m)?(?P<seconds>\d+s)?`)
	var result = &ExtendedDuration{}
	var err error
	fmt.Printf("%#v\n", r.FindAllStringSubmatch(str, 10))
	//
	fmt.Printf("%#v\n", r.SubexpNames())
	res := r.FindAllStringSubmatch(`2w3d4m14h`, -1)
	for idx, _ := range res {
		if  result.weeks == 0 && strings.ContainsRune(res[idx][1],'w') {
			result.weeks, err = strconv.Atoi(res[idx][1][:len(res[idx][1])-1])
			if err != nil {
				log.Errorf("")
			}
		}
		if  result.days == 0 && strings.ContainsRune(res[idx][2],'d') {
			result.days, err = strconv.Atoi(res[idx][2][:len(res[idx][2])-1])
			if err != nil {
				log.Errorf("")
			}
		}
		if  result.hours == 0 && strings.ContainsRune(res[idx][3],'h') {
			result.hours, err = strconv.Atoi(res[idx][3][:len(res[idx][3])-1])
			if err != nil {
				log.Errorf("")
			}
		}
		if  result.minutes == 0 && strings.ContainsRune(res[idx][4],'m') {
			result.minutes, err = strconv.Atoi(res[idx][4][:len(res[idx][4])-1])
			if err != nil {
				log.Errorf("")
			}
		}
		if  result.seconds == 0 && strings.ContainsRune(res[idx][5],'s') {
			result.seconds, err = strconv.Atoi(res[idx][5][:len(res[idx][5])-1])
			if err != nil {
				log.Errorf("")
			}
		}
	}
	return result, err
}

func (s ExtendedDuration) Duration() time.Duration {
	var dur time.Duration
	dur += time.Duration(s.hours) * time.Hour
	dur += time.Duration(s.days*hoursinday) * time.Hour
	dur += time.Duration(s.weeks*hoursinweek) * time.Hour
	dur += time.Duration(s.minutes) * time.Minute
	dur += time.Duration(s.seconds) * time.Second
	return dur
}
