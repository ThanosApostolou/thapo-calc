pub mod calc_cmd;
pub mod root_cmd;
pub mod tui_cmd;

use clap::{Args, Parser, Subcommand};
use serde::{Deserialize, Serialize};
use std::{default, fmt::Debug};

use self::root_cmd::{root_cmd_handler, RootCmd};

pub fn initialize_cli() {
    let root_cmd = RootCmd::parse();
    root_cmd_handler(root_cmd)
}
