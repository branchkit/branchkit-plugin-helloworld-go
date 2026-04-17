package main

import shared "github.com/branchkit/plugin-sdk-go"

func main() {
	plugin := shared.NewPlugin()

	plugin.HandleAction("hello.greet", func(req *shared.OnActionRequest) (any, error) {
		var p GreetParams
		req.UnmarshalParams(&p)

		name := "BranchKit"
		if p.Name != nil {
			name = *p.Name
		}

		plugin.Call("input.type_text", map[string]any{"text": "Hello, " + name + "!"}, nil)
		return nil, nil
	})

	plugin.Run()
}
