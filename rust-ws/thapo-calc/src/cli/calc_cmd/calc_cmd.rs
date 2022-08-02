use clap::{Args, Parser, Subcommand};
use serde::{Deserialize, Serialize};
use std::{default, fmt::Debug};

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
}
