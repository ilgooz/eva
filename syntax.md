# Functions

### Function without paramaters
```
func my
  ...
end
```

### Function with paramaters
```
func my(arg1, arg2) 
  ...
end
```

### Elliptical paramaters
```
func my(arg1, arg2, ...others)
  ...
end
```

```
func my(...others, argN)
  ...
end
```

### Inline function
```
func my: ...
```

```
func my(arg): ...
```

### Functions can return multiple values
```
func my(num)
  num++
  return "my", num
end
```

```
func my(num): num++; return "my", num
```

# Structs

### Simple
```
struct my {
  name,
  Age, // public
  lang "tr" // default val
}
```

```
var me: my{"ilker", 20, "tr"}
```
or
```
var me: my{name: "ilker", Age: 20}
```
you can also add params subsequently
```
var me: my{
  name: "ilker",
  Lastname: "öztürk", // added as public
  Age: 20,
  lang: ("tr", "en")
}
```

### Recursive
```
struct puppy {
  toy {
    color,
    weight: 250
  }
}
```

### They can have rules
```
struct Person {
  name      `json:name`
  lang "en" `json:more.lang`
}
```

### They can have methods
```
func puppy:love
  print("heart")
end

func puppy:say(word)
  print(word)
end

var dog: puppy{toy: {color: "red"}}
dog.say("love you")
dog.love()
```

# Packages
```
import "json"
import "https://github.com/ilgooz/uniqe-id" as uid
```

# JSON
```
var jstring = `{"name": "ilker", "more": {"lang": "tr"}}`
var me: new Person
err = json.Decode(jstring, &me)
print(err, me)
```