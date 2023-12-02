package utils

import "regexp"

func IsDay(e string) bool {
    emailRegex := regexp.MustCompile("^(sun|Sun|mon|Mon|t(ues|hurs)|(T(ues|hurs))|Fri|fri)(day|\\.)?$|wed(\\.|nesday)?$|Wed(\\.|nesday)?$|Sat(\\.|urday)?$|sat(\\.|urday)?$|t((ue?)|(hu?r?))\\.?$|T((ue?)|(hu?r?))\\.?$")
    return emailRegex.MatchString(e)
}