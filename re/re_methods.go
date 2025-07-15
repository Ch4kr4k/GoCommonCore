package re

import (
	re "regexp"
	"strings"
)

func CreateRe() Rgx {
	return Rgx{}
}

func (r *Rgx) In(rgxPttrn string, line string, case_sensitive bool) (bool, error) {

	if !case_sensitive {
		rgx := re.MustCompile(rgxPttrn)
		matches := rgx.FindStringSubmatch(line)
		if len(matches) == 0 {
			return false, nil
		} else {
			return true, nil
		}
	}
	return re.Match(line, []byte(rgxPttrn))
}

func (r *Rgx) Replace(rgxPttrn string, rep string, line string) (string, error) {
	rgx, err := re.Compile(rgxPttrn)
	return rgx.ReplaceAllLiteralString(line, rep), err
}

func (r *Rgx) GetMaches(rgxPttrn string, line string) []string {
	return re.MustCompile(rgxPttrn).FindStringSubmatch(line)
}

func (r *Rgx) CleanInput(line string) []string {
	return strings.Fields(line)
}
