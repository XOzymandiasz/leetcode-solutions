const PLANTED: i32 = 1;

pub fn can_place_flowers(mut flowerbed: Vec<i32>,mut n: i32) -> bool {
    let len = flowerbed.len();
    let mut index = 0;
    if n == 0 { return true; }
    while index < len {
        let has_prev_flower = index > 0 && flowerbed[index - 1] == PLANTED;
        let has_next_flower = index < len -1 && flowerbed[index + 1] == PLANTED;

        if flowerbed[index] == PLANTED {
            index += 2;
            continue;
        }
        if has_prev_flower || has_next_flower {
            index += 1;
            continue;
        }
        flowerbed[index] = PLANTED;
        index += 2;
        n-=1;
        if n <= 0 {
            return true;
        }
    }
    false
}
