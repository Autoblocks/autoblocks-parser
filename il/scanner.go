package il

import(
    "github.com/Autoblocks/go-dsl"
)
func Scan(s *dsl.Scanner) dsl.Token {

}


func sEOF(s *scanner.Scanner) {
	s.Match([]scanner.Match{{"", EOF}})
}

func sStringLiteralDQ(s *scanner.Scanner) {
	s.Expect([]scanner.Expect{{'"', nil}, {'\n', nil}}, scanner.Options{Multiple: true, Invert: true})
	s.Match([]scanner.Match{{"", LITERAL}})
}

func sStringLiteralSQ(s *scanner.Scanner) {
	s.Expect([]scanner.Expect{{'\'', nil}, {'\n', nil}}, scanner.Options{Multiple: true, Invert: true})
	s.Match([]scanner.Match{{"", LITERAL}})
}

// Parse parses a SQL SELECT statement.
func sKWOrIdent(s *scanner.Scanner) {
	s.Expect([]scanner.Expect{{'A', nil}, {'B', nil}, {'C', nil}, {'D', nil}, {'E', nil}, {'F', nil}, {'G', nil}, {'H', nil},
		{'I', nil}, {'J', nil}, {'K', nil}, {'L', nil}, {'M', nil}, {'N', nil}, {'O', nil}, {'P', nil}, {'Q', nil}, {'R', nil},
		{'S', nil}, {'T', nil}, {'U', nil}, {'V', nil}, {'W', nil}, {'X', nil}, {'Y', nil}, {'Z', nil}, {'a', nil}, {'b', nil},
		{'c', nil}, {'d', nil}, {'e', nil}, {'f', nil}, {'g', nil}, {'h', nil}, {'i', nil}, {'j', nil}, {'k', nil}, {'l', nil},
		{'m', nil}, {'n', nil}, {'o', nil}, {'p', nil}, {'q', nil}, {'r', nil}, {'s', nil}, {'t', nil}, {'u', nil}, {'v', nil},
		{'w', nil}, {'x', nil}, {'y', nil}, {'z', nil}, {'_', nil}, {'0', nil}, {'1', nil}, {'2', nil}, {'3', nil}, {'4', nil}, {'5', nil},
		{'6', nil}, {'7', nil}, {'8', nil}, {'9', nil}}, scanner.Options{Multiple: true, Optional: true})
	s.Expect([]scanner.Expect{{'.', sIdent}, {'#', sLit}}, scanner.Options{Multiple: true, Optional: true})
	if !s.Options.NoKeyword {
		s.Match([]scanner.Match{{"LD", IL_LD}, {"LDN", IL_LDN}, {"AND", IL_AND}, {"ANDN", IL_ANDN}, {"OR", IL_OR}, {"ORN", IL_ORN},
			{"XOR", IL_XOR}, {"XORN", IL_XORN}, {"NOT", IL_NOT}, {"ST", IL_ST}, {"STN", IL_STN}, {"S", IL_S}, {"R", IL_R},
			{"ADD", IL_ADD}, {"SUB", IL_SUB}, {"MUL", IL_MUL}, {"DIV", IL_DIV}, {"MOD", IL_MOD}, {"GE", IL_GE}, {"GT", IL_GT},
			{"EQ", IL_EQ}, {"NE", IL_NE}, {"LT", IL_LT}, {"LE", IL_LE}, {"JMP", IL_JMP}, {"JMPC", IL_JMPC},
			{"JMPCN", IL_JMPCN}, {"CAL", IL_CAL}, {"CALC", IL_CALCN}, {"RET", IL_RET}, {"RETC", IL_RETC}, {"RETCN", IL_RETCN}})
	}
	s.Match([]scanner.Match{{"", IDENTIFIER}})
}

func sIdent(s *scanner.Scanner) {
	s.Expect([]scanner.Expect{{'A', nil}, {'B', nil}, {'C', nil}, {'D', nil}, {'E', nil}, {'F', nil}, {'G', nil}, {'H', nil},
		{'I', nil}, {'J', nil}, {'K', nil}, {'L', nil}, {'M', nil}, {'N', nil}, {'O', nil}, {'P', nil}, {'Q', nil}, {'R', nil},
		{'S', nil}, {'T', nil}, {'U', nil}, {'V', nil}, {'W', nil}, {'X', nil}, {'Y', nil}, {'Z', nil}, {'a', nil}, {'b', nil},
		{'c', nil}, {'d', nil}, {'e', nil}, {'f', nil}, {'g', nil}, {'h', nil}, {'i', nil}, {'j', nil}, {'k', nil}, {'l', nil},
		{'m', nil}, {'n', nil}, {'o', nil}, {'p', nil}, {'q', nil}, {'r', nil}, {'s', nil}, {'t', nil}, {'u', nil}, {'v', nil},
		{'w', nil}, {'x', nil}, {'y', nil}, {'z', nil}, {'_', nil}}, scanner.Options{Multiple: true})
	s.Match([]scanner.Match{{"", IDENTIFIER}})
}

func sLitOrBase(s *scanner.Scanner) {
	s.Expect([]scanner.Expect{{'0', nil}, {'1', nil}, {'2', nil}, {'3', nil}, {'4', nil}, {'5', nil},
		{'6', nil}, {'7', nil}, {'8', nil}, {'9', nil}}, scanner.Options{Multiple: true})
	s.Peek([]scanner.ExpectStr{{"..", sInteger}})
	s.Expect([]scanner.Expect{{'.', sFraction}, {'E', sLiteralExponent}, {'#', sLit}}, scanner.Options{Optional: true})
	s.Match([]scanner.Match{{"", LITERAL}})
}

func sLit(s *scanner.Scanner) {
	s.Expect([]scanner.Expect{{'A', nil}, {'B', nil}, {'C', nil}, {'D', nil}, {'E', nil}, {'F', nil}, {'G', nil}, {'H', nil},
		{'I', nil}, {'J', nil}, {'K', nil}, {'L', nil}, {'M', nil}, {'N', nil}, {'O', nil}, {'P', nil}, {'Q', nil}, {'R', nil},
		{'S', nil}, {'T', nil}, {'U', nil}, {'V', nil}, {'W', nil}, {'X', nil}, {'Y', nil}, {'Z', nil}, {'a', nil}, {'b', nil},
		{'c', nil}, {'d', nil}, {'e', nil}, {'f', nil}, {'g', nil}, {'h', nil}, {'i', nil}, {'j', nil}, {'k', nil}, {'l', nil},
		{'m', nil}, {'n', nil}, {'o', nil}, {'p', nil}, {'q', nil}, {'r', nil}, {'s', nil}, {'t', nil}, {'u', nil}, {'v', nil},
		{'w', nil}, {'x', nil}, {'y', nil}, {'z', nil}, {'-', nil}, {'+', nil}, {'0', nil}, {'1', nil}, {'2', nil}, {'3', nil},
		{'4', nil}, {'5', nil}, {'6', nil}, {'7', nil}, {'8', nil}, {'9', nil},
		{':', nil}, {'+', nil}, {',', nil}, {'#', nil}, {'.', nil}}, scanner.Options{Multiple: true})
	s.Match([]scanner.Match{{"", LITERAL}})
}

func sFraction(s *scanner.Scanner) {
	s.Expect([]scanner.Expect{{'0', sLit}, {'1', sLit}, {'2', sLit}, {'3', sLit}, {'4', sLit}, {'5', sLit},
		{'6', sLit}, {'7', sLit}, {'8', sLit}, {'9', sLit}}, scanner.Options{Multiple: true})
	s.Expect([]scanner.Expect{{'E', sLiteralExponent}}, scanner.Options{Optional: true})
}

func sLiteralExponent(s *scanner.Scanner) {
	s.Expect([]scanner.Expect{{'+', sLit}, {'-', sLit}, {'0', sLit}, {'1', sLit}, {'2', sLit}, {'3', sLit},
		{'4', sLit}, {'5', sLit}, {'6', sLit}, {'7', sLit}, {'8', sLit}, {'9', sLit}}, scanner.Options{Multiple: true})
}

func sLeftBracket(s *scanner.Scanner) {
	s.Expect([]scanner.Expect{{'*', sComment}}, scanner.Options{Optional: true})
	s.Match([]scanner.Match{{"", SB_OPEN_PAREN}})
}

func sComment(s *scanner.Scanner) {
	s.Expect([]scanner.Expect{{'*', nil}}, scanner.Options{Invert: true, Multiple: true, Optional: true})
	s.Expect([]scanner.Expect{{'*', nil}})
	s.Expect([]scanner.Expect{{')', nil}})
	s.Match([]scanner.Match{{"", COMMENT}})
}

func sLessThan(s *scanner.Scanner) {
	s.Expect([]scanner.Expect{{'=', nil}, {'>', nil}}, scanner.Options{Optional: true})
	s.Match([]scanner.Match{{"<", SB_LESS_THAN}, {"<=", SB_LESS_THAN_OR_EQUAL}, {"<>", SB_NOT_EQUAL}})
}

func sGreaterThan(s *scanner.Scanner) {
	s.Expect([]scanner.Expect{{'=', nil}}, scanner.Options{Optional: true})
	s.Match([]scanner.Match{{">", SB_GREATER_THAN}, {">=", SB_GREATER_THAN_OR_EQUAL}})
}

func sInteger(s *scanner.Scanner) {
	s.Match([]scanner.Match{{"", LITERAL}})
}

func sStar(s *scanner.Scanner) {
	s.Match([]scanner.Match{{"", LITERAL}})
}

func sAssign(s *scanner.Scanner) {
	s.Expect([]scanner.Expect{{'=', nil}}, scanner.Options{Optional: true})
	s.Expect([]scanner.Expect{{'>', nil}}, scanner.Options{Optional: true})
	s.Match([]scanner.Match{{":", SB_COLON}, {":=", SB_ASSIGN}, {":=>", SB_ASSIGN_OUT}})
}

func sHash(s *scanner.Scanner) {
	s.Match([]scanner.Match{{"", LITERAL}})
}

func sIOVariable(s *scanner.Scanner) {
	s.Expect([]scanner.Expect{{'I', nil}, {'M', nil}, {'Q', nil}})
	s.Expect([]scanner.Expect{{'X', nil}, {'B', nil}, {'W', nil}, {'D', nil}, {'L', nil}, {'*', nil}}, scanner.Options{Optional: true})
	s.Expect([]scanner.Expect{{'0', nil}, {'1', nil}, {'2', nil}, {'3', nil}, {'4', nil}, {'5', nil}, {'6', nil}, {'7', nil}, {'8', nil},
		{'9', nil}}, scanner.Options{Multiple: true})
	s.Expect([]scanner.Expect{{'.', sIOSubVar}}, scanner.Options{Multiple: true, Optional: true})
	s.Match([]scanner.Match{{"", IO}})
}

func sIOSubVar(s *scanner.Scanner) {
	s.Expect([]scanner.Expect{{'0', nil}, {'1', nil}, {'2', nil}, {'3', nil}, {'4', nil}, {'5', nil}, {'6', nil}, {'7', nil}, {'8', nil},
		{'9', nil}}, scanner.Options{Multiple: true})
	s.Expect([]scanner.Expect{{'.', sIOSubVar}}, scanner.Options{Optional: true})
}

func sWindowsNewline(s *scanner.Scanner) {
	s.Expect([]scanner.Expect{{'\n', nil}}, scanner.Options{Optional: true})
	s.Match([]scanner.Match{{"\r", NL}, {"\r\n", NL}})
}
