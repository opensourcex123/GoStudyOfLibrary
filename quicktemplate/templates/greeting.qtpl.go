// Code generated by qtc from "greeting.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// All text outside function is treated as comments.
//

//line greeting.qtpl:3
package templates

//line greeting.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line greeting.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line greeting.qtpl:3
func StreamGreeting(qw422016 *qt422016.Writer, name string, count int) {
//line greeting.qtpl:3
	qw422016.N().S(`
    `)
//line greeting.qtpl:4
	for i := 0; i < count; i++ {
//line greeting.qtpl:4
		qw422016.N().S(`
        Hello, `)
//line greeting.qtpl:5
		qw422016.E().S(name)
//line greeting.qtpl:5
		qw422016.N().S(`
      `)
//line greeting.qtpl:6
	}
//line greeting.qtpl:6
	qw422016.N().S(`
`)
//line greeting.qtpl:7
}

//line greeting.qtpl:7
func WriteGreeting(qq422016 qtio422016.Writer, name string, count int) {
//line greeting.qtpl:7
	qw422016 := qt422016.AcquireWriter(qq422016)
//line greeting.qtpl:7
	StreamGreeting(qw422016, name, count)
//line greeting.qtpl:7
	qt422016.ReleaseWriter(qw422016)
//line greeting.qtpl:7
}

//line greeting.qtpl:7
func Greeting(name string, count int) string {
//line greeting.qtpl:7
	qb422016 := qt422016.AcquireByteBuffer()
//line greeting.qtpl:7
	WriteGreeting(qb422016, name, count)
//line greeting.qtpl:7
	qs422016 := string(qb422016.B)
//line greeting.qtpl:7
	qt422016.ReleaseByteBuffer(qb422016)
//line greeting.qtpl:7
	return qs422016
//line greeting.qtpl:7
}
