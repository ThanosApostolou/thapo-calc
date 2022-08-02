use clap::{Args, Parser, Subcommand};
use serde::{Deserialize, Serialize};
use std::{default, fmt::Debug};

#[derive(Clone, Debug, Serialize, Deserialize, Hash, PartialEq, PartialOrd, Args)]
pub struct TuiCmd {}

pub fn tui_cmd_handler(tui_cmd: &TuiCmd) {
    println!("starting tui")
}
