pub struct Solution;
impl Solution {
    pub fn permute(nums: Vec<i32>) -> Vec<Vec<i32>> {
        let mut answer: Vec<Vec<i32>> = Vec::new();
        Self::get_permutations(&mut answer, nums, Vec::new());
        answer
    }

    fn get_permutations(answer: &mut Vec<Vec<i32>>, nums: Vec<i32>, row: Vec<i32>) {
        if nums.is_empty() {
            answer.push(row.clone());
            return;
        }
        for i in 0..nums.len() {
            let mut remaining = nums.clone();
            let value = remaining.remove(i);
            let mut next_row: Vec<i32> = row.clone();
            next_row.push(value);
            Self::get_permutations(answer, remaining, next_row);
        }
    }
}
