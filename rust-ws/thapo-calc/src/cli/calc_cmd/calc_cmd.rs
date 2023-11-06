use clap::Args;
use serde::{Deserialize, Serialize};
use std::fmt::Debug;

use crate::domain::calculation::calculate_from_string;

#[derive(Clone, Debug, Serialize, Deserialize, Hash, PartialEq, PartialOrd, Args)]
pub struct CalcCmd {
    #[clap(value_parser)]
    expression: String,

    #[clap(short, long, value_parser)]
    var: Vec<String>,
}

pub fn calc_cmd_handler(calc_cmd: &CalcCmd) {
    println!("calculating expression: {:?}", calc_cmd.expression);
    for variable in calc_cmd.var.iter() {
        println!("variable: {:?}", variable);
    }
    let result: Result<f64, &str> = calculate_from_string(&calc_cmd.expression);
    match result {
        Ok(res) => println!("{}", res),
        Err(err) => eprintln!("{}", err),
    }
}
