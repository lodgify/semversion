# Semversion

## Presentation

This is a small application to manipulate sementic version from the command line

```
 ~/g/s/g/k/semversion echo "1.2.3" | ./semversion -action increment -what minor -value 3
 1.5.3
 ~/g/s/g/k/semversion echo "1.2.3" | ./semversion -action increment -what patch -value 3
 1.2.6
 ~/g/s/g/k/semversion echo "1.2.3" | ./semversion -action set -what patch -value 3
 1.2.3
```

## Documentation

```
Usage of ./semversion:
  -action string
    	Action to do: set, increment, decrement
  -value uint
    	The value to use during the action (default 1)
  -what string
    	Where to run the action: major, minor, patch (default "patch")
```