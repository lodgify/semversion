# Semversion

## Presentation

This is a small application to manipulate semantic version from the command line (Either STDIN or -original flag)

```
 echo "1.2.3" | ./semversion -action increment -what minor -value 3
 1.5.3
 echo "1.2.3" | ./semversion -action increment -what patch -value 3
 1.2.6
 echo "1.2.3" | ./semversion -action set -what patch -value 5
 1.2.5
 ./semversion -action set -what patch -value 3 -original 1.2.3
 1.2.5
 
```

## Documentation

```
Usage of ./semversion:                        
  -action string                              
        Action to do: set, increment, decrement                                             
  -original string                            
        The original value, can also be omitted and passed in STDIN                         
  -value uint                                 
        The value to use during the action (default 1)                                      
  -what string                                
        Where to run the action: major, minor, patch (default "patch") 
```
