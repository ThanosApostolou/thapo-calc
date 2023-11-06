use std::str::Chars;

pub fn calculate_from_string(expr_string: &String) -> Result<f64, &'static str> {
    if expr_string.is_empty() {
        return Err("expression was empty");
    }
    let mut expr_chars: Vec<char> = expr_string.chars().collect();
    for c in &mut expr_chars {
        println!("{}", c)
    }
    let (res, index) = calc_expression(&expr_chars, 0, false);
    res
}

fn calc_expression(
    expr_chars: &Vec<char>,
    mut index: usize,
    has_parenthesis: bool,
) -> (Result<f64, &'static str>, usize) {
    let mut expr_res: f64 = 0.0;
    let mut myindex = index;

    while myindex < expr_chars.len() {
        if expr_chars[myindex] == '+' {
            let (new_result, newindex) = calc_term(expr_chars, index, has_parenthesis);
            myindex = newindex;
            match new_result {
                Ok(new_res) => expr_res += new_res,
                Err(err) => return (Err(err), index),
            }
        } else if expr_chars[myindex] == '-' {
            let (new_result, newindex) = calc_term(expr_chars, index, has_parenthesis);
            myindex = newindex;
            match new_result {
                Ok(new_res) => expr_res -= new_res,
                Err(err) => return (Err(err), index),
            }
        } else {
            let (new_result, newindex) = calc_term(expr_chars, index, has_parenthesis);
            myindex = newindex;
            match new_result {
                Ok(new_res) => expr_res = new_res,
                Err(err) => return (Err(err), index),
            }
        }
    }

    (Ok(expr_res), index)
}

fn calc_term(
    expr_chars: &Vec<char>,
    mut index: usize,
    has_parenthesis: bool,
) -> (Result<f64, &'static str>, usize) {
    let expr_res: f64 = 0.0;
    let myindex = index;
    (Ok(expr_res), index)
}
