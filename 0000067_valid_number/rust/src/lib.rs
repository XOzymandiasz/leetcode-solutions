const FIRST_VALID_SIGNS: &[char] = &['-', '+', '.'];
const AFTER_SIGN_VALID_SIGNS: &[char] = &['.'];
const AFTER_NUMBER_VALID_SIGNS: &[char] = &['.', 'e', 'E'];
const AFTER_DOT_VALID_SIGNS: &[char] = &['e', 'E'];
const AFTER_EXPONENT_VALID_SIGNS: &[char] = &['-', '+'];
const NUMBERS: &[char] = &['0', '1', '2', '3', '4', '5', '6', '7', '8', '9'];

pub fn is_number(param: String) -> bool {
    let mut valid = FIRST_VALID_SIGNS;
    let mut has_digit = false;
    let mut has_dot = false;
    let mut has_exponent = false;

    for sign in param.chars() {
        if NUMBERS.contains(&sign) {
            has_digit = true;

            valid = if has_exponent {
                &[]
            } else if has_dot {
                AFTER_DOT_VALID_SIGNS
            } else {
                AFTER_NUMBER_VALID_SIGNS
            };

            continue;
        }

        if !valid.contains(&sign) {
            return false;
        }

        match sign {
            '+' | '-' => {
                valid = if has_exponent {
                    &[]
                } else {
                    AFTER_SIGN_VALID_SIGNS
                };
            }
            '.' => {
                has_dot = true;
                valid = AFTER_DOT_VALID_SIGNS;
            }
            'e' | 'E' => {
                if !has_digit {
                    return false;
                }

                has_exponent = true;
                has_digit = false;
                valid = AFTER_EXPONENT_VALID_SIGNS;
            }
            _ => return false,
        }
    }

    has_digit
}
