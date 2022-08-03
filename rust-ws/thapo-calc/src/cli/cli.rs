use clap::Parser;

use super::{root_cmd_handler, RootCmd};

pub fn initialize_cli() {
    let root_cmd = RootCmd::parse();
    root_cmd_handler(root_cmd)
}
