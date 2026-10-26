pub struct Solution;
impl Solution {
    pub fn permute_unique(mut nums: Vec<i32>) -> Vec<Vec<i32>> {
        nums.sort_unstable();

        let mut answer = Vec::new();
        let mut row = Vec::with_capacity(nums.len());
        let mut used = vec![false; nums.len()];

        Self::backtrack(&nums, &mut used, &mut row, &mut answer);
        answer
    }

    fn backtrack(nums: &[i32], used: &mut [bool], row: &mut Vec<i32>, answer: &mut Vec<Vec<i32>>) {
        if row.len() == nums.len() {
            answer.push(row.clone());
            return;
        }

        for i in 0..nums.len() {
            if used[i] {
                continue;
            }

            if i > 0 && nums[i] == nums[i - 1] && !used[i - 1] {
                continue;
            }

            used[i] = true;
            row.push(nums[i]);

            Self::backtrack(nums, used, row, answer);

            row.pop();
            used[i] = false;
        }
    }
}
