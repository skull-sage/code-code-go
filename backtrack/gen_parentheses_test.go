package backtrack

type StrList []string

func genSeq(str string, open_count int, close_count int, n int, result *StrList) {

	if open_count == n && close_count == n {
		*result = append(*result, str)
		return
	}

	if open_count < n {
		genSeq(str+"(", open_count+1, close_count, n, result)
	}

	if close_count < open_count {
		genSeq(str+")", open_count, close_count+1, n, result)
	}

}

func generateParenthesis(n int) []string {
	result := make(StrList, 0)
	if n == 1 {
		return []string{"()"}
	}

	genSeq("", 0, 0, n, &result)
	return result

}
