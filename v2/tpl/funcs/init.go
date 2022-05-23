package funcs

import (
	"github.com/CarsonSlovoka/go-pkg/v2/tpl/funcs/compare"
)

func GetUtilsFuncMap() map[string]any {
	i18nFunc := func(messageID string, templateData any) string { return messageID } // Just let "i18n" and T is legal. You can override it later.
	return map[string]interface{}{
		"i18n": i18nFunc, "T": i18nFunc,
		"dict":    Dict,
		"list":    List,
		"split":   Split,
		"replace": Replace,

		// 👇 Math
		"add":     Add,
		"sub":     Sub,
		"mul":     Mul,
		"div":     Div,
		"ceil":    Ceil,
		"floor":   Floor,
		"ln":      Ln,
		"log":     Log,
		"sqrt":    Sqrt,
		"mod":     Mod,
		"modBool": ModBool,
		"pow":     Pow,
		"round":   Round,

		"default": compare.Default,
	}
}
