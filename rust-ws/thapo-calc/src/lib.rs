#[cfg(any(target_family = "unix", target_family = "windows"))]
pub mod cli;
pub mod domain;
pub mod gui;

#[cfg(any(target_family = "unix", target_family = "windows"))]
pub fn run() {
    cli::initialize_cli();
}

#[cfg(target_family = "wasm")]
pub fn run() {
    dioxus::web::launch(gui::app::app);
}
