package format

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/yiipu/i9r/runtime/nodes/basenode"
	"github.com/yiipu/i9r/runtime/varpool/basevarpool"
)

type Format struct {
	basenode.BaseNode
	Params FormatParams `json:"params"`
}

type FormatParams struct {
	FormatString string `json:"formatString"`
}

func (n *Format) Execute(vp *basevarpool.IVarPool) error {
	// look for pattern `{{v:变量名}}` or `{{m:变量名}}` in format string
	re, err := regexp.Compile(`{{(v|m):([^}]+)}}`)
	if err != nil {
		return err
	}
	output := n.Params.FormatString
	matches := re.FindAllStringSubmatch(n.Params.FormatString, -1)
	for index, match := range matches {
		if len(match) != 3 {
			continue
		}
		switch match[1] {
		case "v":
			v, err := (*vp).GetVar(match[2])
			if err != nil {
				return err
			}
			output = strings.Replace(output, matches[index][0], fmt.Sprintf("%v", v.Value), -1)
		case "m":
			m, err := (*vp).GetMagicVar(match[2])
			if err != nil {
				return err
			}
			output = strings.Replace(output, matches[index][0], fmt.Sprintf("%v", m.Value), -1)
		}
	}
	// set the result to the output variable
	err = (*vp).SetMagicVar(n.ID, basevarpool.MagicVar{
		BaseVar: basevarpool.BaseVar{
			Value:    output,
			Readonly: false,
		},
	})
	if err != nil {
		return err
	}
	return nil
}
