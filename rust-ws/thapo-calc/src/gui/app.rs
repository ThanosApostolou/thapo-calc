use dioxus::prelude::*;

pub fn app(cx: Scope) -> Element {
    let state = use_state(&cx, || "red".to_string());
    cx.render(rsx! {
        div {
            hidden: "false",
            background_color: "blue",
            onclick:  |_| state.set(format!("{}a", state.get())),
            h1 { "hello world" }
            p { "Some body content" }
            p { "{state}"}

        }
    })
}
