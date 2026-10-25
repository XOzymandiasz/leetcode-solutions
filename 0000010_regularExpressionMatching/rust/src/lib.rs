pub struct Solution;
impl Solution {
    pub fn is_match(s: String, p: String) -> bool {
        let text: Vec<char> = s.chars().collect();
        let pattern: Vec<char> = p.chars().collect();
        let mut memo = vec![vec![None; pattern.len() + 1]; text.len() + 1];

        Self::matches_from(&text, &pattern, 0, 0, &mut memo)
    }

    fn matches_from(
        text: &[char],
        pattern: &[char],
        i: usize,
        j: usize,
        memo: &mut [Vec<Option<bool>>],
    ) -> bool {
        if let Some(result) = memo[i][j] {
            return result;
        }

        let result = if j == pattern.len() {
            i == text.len()
        } else {
            let first_matches = i < text.len() && (pattern[j] == '.' || pattern[j] == text[i]);

            if pattern.get(j + 1) == Some(&'*') {
                Self::matches_from(text, pattern, i, j + 2, memo)
                    || (first_matches && Self::matches_from(text, pattern, i + 1, j, memo))
            } else {
                first_matches && Self::matches_from(text, pattern, i + 1, j + 1, memo)
            }
        };

        memo[i][j] = Some(result);
        result
    }
}
