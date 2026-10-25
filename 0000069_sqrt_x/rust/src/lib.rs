pub struct Solution;
impl Solution {
    pub fn my_sqrt(x: i32) -> i32 {
        let x = i64::from(x);
        let mut x_1 = 0;
        let mut x_2 = x;
        let mut y = 0;

        while x_1 <= x_2 {
            let mid = x_1 + (x_2 - x_1) / 2;

            if mid * mid <= x {
                y = mid;
                x_1 = mid + 1;
            } else {
                x_2 = mid - 1;
            }
        }

        y as i32
    }
}
