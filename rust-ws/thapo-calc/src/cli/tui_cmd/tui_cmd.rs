use clap::Args;
use serde::{Deserialize, Serialize};
use std::fmt::Debug;

#[derive(Clone, Debug, Serialize, Deserialize, Hash, PartialEq, PartialOrd, Args)]
pub struct TuiCmd {}

pub fn tui_cmd_handler(_tui_cmd: &TuiCmd) {
    println!("starting tui")
}
