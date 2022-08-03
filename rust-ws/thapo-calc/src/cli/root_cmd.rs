use clap::{Parser, Subcommand};
use serde::{Deserialize, Serialize};
use std::fmt::Debug;

use crate::gui;

use super::{
    calc_cmd::calc_cmd::{calc_cmd_handler, CalcCmd},
    tui_cmd::tui_cmd::{tui_cmd_handler, TuiCmd},
};

/// Simple program to greet a person
#[derive(Clone, Debug, Serialize, Deserialize, Hash, PartialEq, PartialOrd, Parser)]
#[clap(author, version, about, long_about = None)]
pub struct RootCmd {
    /// Name of the person to greet
    #[clap(short, long, value_parser)]
    name: Option<String>,

    /// Number of times to greet
    #[clap(short, long, value_parser, default_value_t = 1)]
    count: u8,

    #[clap(subcommand)]
    commands: Option<Commands>,
}

#[derive(Clone, Debug, Serialize, Deserialize, Hash, PartialEq, PartialOrd, Subcommand)]
enum Commands {
    /// Adds files to myapp
    Calc(CalcCmd),
    Tui(TuiCmd),
}

pub fn root_cmd_handler(root_cmd: RootCmd) {
    match &root_cmd.commands {
        Some(command) => match command {
            Commands::Calc(calc_cmd) => {
                calc_cmd_handler(calc_cmd);
            }
            Commands::Tui(tui_cmd) => {
                tui_cmd_handler(tui_cmd);
            }
        },
        None => {
            println!("root");
            dioxus::desktop::launch(gui::app::app);
        }
    }
}
