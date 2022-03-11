package funcs

import (
	"carson.io/pkg/tpl/funcs/compare"
)

func GetUtilsFuncMap() map[string]interface{} {
	i18nFunc := func(messageID string, templateData interface{}) string { return messageID } // Just let "i18n" and T is legal. You can override it later.
	return map[string]interface{}{
		"i18n": i18nFunc, "T": i18nFunc,
		"dict":    Dict,
		"slice":   Slice, // Let 1st char uppercase since "slice" was defined already.
		"split":   Split,
		"replace": Replace,

		// 👇 Math
		"add":     Add,
		"sub":     Sub,
		"mul":     Mul,
		"div":     Div,
		"ceil":    Ceil,
		"floor":   Floor,
		"log":     Log,
		"sqrt":    Sqrt,
		"mod":     Mod,
		"modBool": ModBool,
		"pow":     Pow,
		"round":   Round,

		"default": compare.Default,
	}
}
