# Functions

## Function without paramaters
```
func my
  ...
end
```

## Function with paramaters
```
func my(arg1, arg2) 
  ...
end
```

## Elliptical paramaters
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

## Inline function
```
func my: ...
```

```
func my(arg): ...
```

## Functions can return multiple values
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

## Simple
```
struct my {
  name,
  Age, // public
  lang: "tr" // default val
}
```

```
var me: my{"ilker", 20}
```
// or
```
var me: my{name: "ilker", Age: 20}
```
// you can also add params subsequently
```
var me: my{
  name: "ilker",
  Lastname: "öztürk", // public
  Age: 20,
  lang: ("tr", "en")
}
```

## Extended
```
struct puppy {
  toy {
    color,
    weight
  }
}
```

## structs can have methods
```
func puppy:love
  ...
end

func puppy:say(word)
  print(word)
end

var dog: puppy{toy: {color: "red"}}
dog.say("love you")
```

# JSON