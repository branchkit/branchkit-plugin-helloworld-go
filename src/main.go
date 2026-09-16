package main

import "github.com/branchkit/plugin-sdk-go"



func main() {
	plugin := branchkit.NewPlugin()

	// HandleGreet and GreetParams are generated from plugin.json into
	// actions_gen.go, so the action string is never spelled here and the
	// params arrive typed. Edit action_types in plugin.json and re-run
	// branchkit-gen to change them.
	HandleGreet(plugin, func(p GreetParams, _ *branchkit.OnActionRequest) (any, error) {
		name := "BranchKit"
		if p.Name != nil {
			name = *p.Name
		}

		// Generated wrapper — the method name and argument shape are
		// checked at compile time, unlike a raw plugin.Call.
		return nil, plugin.InputTypeText("Hello, " + name + "!")
	})

	// One renderer per tab declared in plugin.json. The SDK owns the
	// render_settings hook: it dispatches on the tab key and re-renders the
	// tab through the settings stream whenever a method returns.
	plugin.SettingsTab("getting_started", func(_ *branchkit.RenderSettingsRequest) (string, error) {
		return `<div style="padding: 16px; font-family: system-ui;">

	<h2 style="margin: 0 0 12px 0;">Helloworld</h2>
	<p style="color: #888; margin: 0 0 16px 0;">A BranchKit plugin</p>

	<h3 style="margin: 0 0 8px 0;">Voice Commands</h3>
	<table style="border-collapse: collapse; width: 100%;">
		<tr>
			<td style="padding: 6px 12px; border-bottom: 1px solid #333;"><em>"hello branchkit"</em></td>
			<td style="padding: 6px 12px; border-bottom: 1px solid #333; color: #888;">Types "Hello, BranchKit!"</td>
		</tr>
		<tr>
			<td style="padding: 6px 12px;"><em>"hello &lt;name&gt;"</em></td>
			<td style="padding: 6px 12px; color: #888;">Types "Hello, &lt;name&gt;!" with any spoken word</td>
		</tr>
	</table>
</div>`, nil
	})


	plugin.Run()
}
