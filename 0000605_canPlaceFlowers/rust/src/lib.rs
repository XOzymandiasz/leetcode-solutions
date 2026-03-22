pub fn can_place_flowers(mut flowerbed: Vec<i32>,mut n: i32) -> bool {
    for mut i in  0..=flowerbed.len()-1 {
        if flowerbed[i+1] == 1 {
            continue;
        }
        if flowerbed[i] == 1 {
            i+=1;
        }
        if n <= 0 {
            return true;
        }
        n-=1;
        flowerbed[i] = 1;
        i+=1;
    }

    false
}
