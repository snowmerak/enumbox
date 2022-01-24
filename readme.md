# enumbox

## install

`go install github.com/snowmerak/enumbox@latest`

## how to use

### init package and make yaml file

`enumbox init <package-path>`

if you write `const/names`, enumbox makes enumbox.yml in `const/names`(if you use windows, `const\names`).

MUST USE SLASH.

### generate package

`enumbox generate <package-path>`

if you write `const/names`, enumbox create enumbox.go in `const/names`.

MUST USE SLASH, too.

## sample

### init

`enum init names`

### replace

```yaml
version: 0.0.1
variables:
- name: jhon
  type: string
  value: '"jhon"'
- name: jerry
  type: string
  value: '"jerry"'
- name: julia
  type: string
  value: '"julia"'
- name: james
  type: string
  value: '"james"'
- name: jill
  type: string
  value: '"jill"'
- name: joshua
  type: string
  value: '"joshua"'
- name: jessica
  type: string
  value: '"jessica"'
```

copy and paste to names/enumbox.yml

### generate

`enumbox generate names`

### result

```go
package names

var (
	jhon    string = "jhon"
	jerry   string = "jerry"
	julia   string = "julia"
	james   string = "james"
	jill    string = "jill"
	joshua  string = "joshua"
	jessica string = "jessica"
)

func JHON() string {
	return jhon
}

func JERRY() string {
	return jerry
}

func JULIA() string {
	return julia
}

func JAMES() string {
	return james
}

func JILL() string {
	return jill
}

func JOSHUA() string {
	return joshua
}

func JESSICA() string {
	return jessica
}

func NameOf(data interface{}) string {
	switch data {
	case jhon:
		return "jhon"
	case jerry:
		return "jerry"
	case julia:
		return "julia"
	case james:
		return "james"
	case jill:
		return "jill"
	case joshua:
		return "joshua"
	case jessica:
		return "jessica"
	}
	return ""
}

func IndexOf(data interface{}) int {
	switch data {
	case jhon:
		return 0
	case jerry:
		return 1
	case julia:
		return 2
	case james:
		return 3
	case jill:
		return 4
	case joshua:
		return 5
	case jessica:
		return 6
	}
	return -1
}

func At(index int) interface{} {
	switch index {
	case 0:
		return jhon
	case 1:
		return jerry
	case 2:
		return julia
	case 3:
		return james
	case 4:
		return jill
	case 5:
		return joshua
	case 6:
		return jessica
	}
	return nil
}

func EqualToJHON(b *string) bool {
	return jhon == *b
}

func EqualToJERRY(b *string) bool {
	return jerry == *b
}

func EqualToJULIA(b *string) bool {
	return julia == *b
}

func EqualToJAMES(b *string) bool {
	return james == *b
}

func EqualToJILL(b *string) bool {
	return jill == *b
}

func EqualToJOSHUA(b *string) bool {
	return joshua == *b
}

func EqualToJESSICA(b *string) bool {
	return jessica == *b
}

func Names() []string {
	return []string{
		"jhon",
		"jerry",
		"julia",
		"james",
		"jill",
		"joshua",
		"jessica",
	}
}

```
