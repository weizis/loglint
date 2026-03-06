package analyzer

import (
	"go/ast"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

var sensitiveWords []string

func SetSensitiveWords(words []string) {
    sensitiveWords = words
}

var specialChars = regexp.MustCompile(`[!@#$%^&*(){}\[\]|\\;:'",.<>/?` + "`" + `]`)

func checkRules(pass *analysis.Pass, node ast.Node, msg string, lit *ast.BasicLit) {
	if msg == "" {
		return
	}

	lowerMsg := strings.ToLower(msg)
	for _, word := range sensitiveWords {
		if strings.Contains(lowerMsg, word) {
			pass.Reportf(node.Pos(), "log message may contain sensitive data (e.g., '%s')", word)
			return 
		}
	}

	for _, r := range msg {
		if r > unicode.MaxASCII {
			pass.Reportf(node.Pos(), "log message should use only English (ASCII) characters")
			return
		}
	}

	if specialChars.MatchString(msg) {
		pass.Reportf(node.Pos(), "log message should not contain special characters or punctuation")
		return
	}

	if len(msg) > 0 {
		firstChar := rune(msg[0])
		if unicode.IsUpper(firstChar) {
			suggestLowercaseFix(pass, node, lit)
			return
		}
	}
}
func suggestLowercaseFix(pass *analysis.Pass, node ast.Node, lit *ast.BasicLit) {
    if lit == nil {
        return
    }
    msg := lit.Value 
    if len(msg) < 3 { 
        return
    }

    firstCharIndex := -1
    for i, r := range msg {
        if i == 0 { 
            continue
        }
        if r != '"' && r != '\\' { 
            firstCharIndex = i
            break
        }
    }

    if firstCharIndex == -1 {
        return
    }

    firstChar := msg[firstCharIndex : firstCharIndex+1]
    lowerFirstChar := strings.ToLower(firstChar)
    if firstChar == lowerFirstChar {
        return 
    }

    newMsg := msg[:firstCharIndex] + lowerFirstChar + msg[firstCharIndex+1:]

    pass.Report(analysis.Diagnostic{
        Pos:     node.Pos(),
        End:     node.End(),
        Message: "log message should start with a lowercase letter",
        SuggestedFixes: []analysis.SuggestedFix{{
            Message: "make first letter lowercase",
            TextEdits: []analysis.TextEdit{{
                Pos:     lit.Pos(),
                End:     lit.End(),
                NewText: []byte(newMsg),
            }},
        }},
    })
}