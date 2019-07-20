package examples

import "github.com/gascore/gas"

// Example application #6
//
// 'htmlDirective' shows how you can use component.Directive.HTML
func HTMLDirective() *gas.E {
	root := &HTMLRoot{}
	c := &gas.C{NotPointer: true, Root: root}
	root.c = c

	return c.Init()
}

type HTMLRoot struct {
	c               *gas.C
	isArticleActive bool
}

func (root *HTMLRoot) Render() []interface{} {
	return gas.CL(
		gas.NE(
			&gas.E{
				Handlers: map[string]gas.Handler{
					"click": func(e gas.Object) {
						root.isArticleActive = !root.isArticleActive
						root.c.Update()
					},
				},
				Tag: "button",
			},
			func() interface{} {
				if root.isArticleActive {
					return "Show article"
				} else {
					return "Hide article"
				}
			}(),
		),
		gas.NE(
			&gas.E{
				HTML: gas.HTMLDirective{
					Render: func() string {
						if root.isArticleActive {
							return articleText
						} else {
							return helloText
						}
					},
				},
				Tag: "article",
				Attrs: map[string]string{
					"id":    "article",
					"style": `border: 1px solid #dedede;padding: 2px 4px;margin-top:12px;`,
				},
			}))
}

const articleText = `<h1>
Lorem ipsum dolor sit amet, consectetur adipiscing elit.
</h1>
<p>
Vivamus arcu nibh, sodales nec lectus ut, vestibulum porta est. Nunc in odio eu tellus feugiat volutpat vitae a erat.
</p>
<p>
<i>Phasellus sit amet suscipit urna</i>. 
Quisque vitae risus lobortis, aliquam orci at, pulvinar urna. Quisque vitae lobortis libero.
Nullam a faucibus dolor. Ut eu turpis et purus mollis ullamcorper. Vivamus interdum felis quis volutpat volutpat. Mauris id auctor nisi.
</p>
<hr/>
<p>
<strong>Integer aliquam tellus nunc, ac dapibus felis pulvinar viverra</strong>. 
Donec dapibus dolor in massa vehicula ornare. Duis molestie velit vitae purus consectetur pulvinar. Aliquam ac purus placerat, laoreet tortor at, aliquet ex.
</p>
<h3>
Nulla facilisi. Donec mattis auctor finibus.
</h3>`

const helloText = `<h1>To see article click button!</h1>`
