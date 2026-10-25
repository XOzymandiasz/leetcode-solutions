pub fn reverse(x: i32) -> i32 {
    let mut x_copy = x;
    let mut answer: i32 = 0;

    while x_copy != 0 {
        let digit = x_copy % 10;

        answer = match answer
            .checked_mul(10)
            .and_then(|value| value.checked_add(digit))
        {
            Some(value) => value,
            None => return 0,
        };

        x_copy /= 10;
    }

    answer
}
