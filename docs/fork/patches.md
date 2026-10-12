# notch-lab patches

This fork tracks upstream `egoist/mygo` tag `v0.4.5` on branch `notch-lab`. Every patch is one commit with the subject prefix `notch:`. The requirement ids are those of notch-lab's `docs/spec/14-mygo-fork.md`.

| Patch | Commit subject | API | Requirement | Example |
|---|---|---|---|---|
| screen-flyout | `notch: screen flyout` | `FlyoutOptions.Screen`, `platform.Flyout.Screen` | R-MYGO-010 | `examples/notch-screen-flyout` |
| key-on-click | `notch: flyout key on click` | `FlyoutOptions.KeyOnClick`, `platform.Flyout.KeyOnClick` | R-MYGO-010 | `examples/notch-key-on-click` |
| screen-notch | `notch: Screen.Notch` | `Screen.Notch()`, `platform.Screen.Notch` | R-MYGO-019 | `examples/notch-screen-notch` |
| global-pointer | `notch: global pointer` | `App.OnGlobalPointer`, `platform.AppController.WatchPointer`, `platform.AppHandler.PointerMoved` | R-MYGO-020 | `examples/notch-pointer` |
| flyout-dialogs | `notch: dialogs over flyouts` | none: `Dialog.*` with a flyout `Parent` is application modal | R-MYGO-021 | — |
| webview-hooks | `notch: web view hooks` | `PageOptions.Configure`, `WebView.NativeHandle`, `Page.SetWebViewOpenHandler`, `WebViewOptions.Native`, `WindowOpenRequest.Native` | R-MYGO-030…037 (replaces them) | — |
