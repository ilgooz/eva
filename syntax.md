# Functions

### Function without paramaters
```ruby
func my
  ...
end
```

### Function with paramaters
```ruby
func my(arg1, arg2) 
  ...
end
```

### Function with default paramater
```ruby
func my(anum: 12) 
  ...
end
```

### Elliptical paramaters
```ruby
func my(arg1, arg2, ...others)
  ...
end
```

```ruby
func my(...others, argN)
  ...
end
```

### Inline function
```ruby
func my: ...
```

```ruby
func my(arg): ...
```

### Functions can return multiple values
```ruby
func my(num)
  num++
  return "my", num
end
```

```ruby
func my(num): num++; return "my", num
```

# Structs

### Simple
```ruby
struct my {
  name,
  Age, // public
  lang "tr" // default val
}
```

```ruby
var me: my{"ilker", 20, "tr"}
```
or
```ruby
var me: my{name: "ilker", Age: 20}
```
you can also add params subsequently
```ruby
var me: my{
  name: "ilker",
  Lastname: "öztürk", // added as public
  Age: 20,
  lang: ("tr", "en")
}
```

### Recursive
```ruby
struct puppy {
  toy {
    color,
    weight: 250
  }
}
```

### They can have rules
```ruby
struct Person {
  name      `json:name`
  lang "en" `json:more.lang`
}
```

### They can have methods
```ruby
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
```ruby
import "json"
import "https://github.com/ilgooz/uniqe-id" as uid
```

# JSON
```ruby
var jstring = `{"name": "ilker", "more": {"lang": "tr"}}`
var me: new Person
err = json.Decode(jstring, &me)
print(err, me)
```