# 1️⃣🐝🏎️ The One Billion Row Challenge


## Author - Dan Richards 



## Iterations 


| Iteration | Title                             | Duration    | Improvement |
|-----------|-----------------------------------|-------------| ------------|
| 1         | Simple Single Threaded            | 308.678428  |             |
| 2         | Simple Single Threaded (Integers) | 354.827043  |  




## 1 - Simple Single Threaded

> **Duration** 308.678428s  - Not that good
  
  
> My intention with this was to simply get some code working and see how she went! This very much goes along the lines of red, green refactor. 
> Although in this repository I won't be using test cases... Just simply getting it to work and seeing how I do was important.


### Reaction
- My first reaction was that it was going to be slow, I noticed that some other programs were processing it in 90s vs my 308 seconds. So something must be quite wrong here. Likely to do with the use of the scanner.



### Possible improvements
In order of easier to implement.... 
- Use integers and / 10 at the end 
- Use buffers for injecting the input 
- Use concurrency to use more than one core to


-------


## 2 - Simple Single Threaded - Using integers

### Reaction

- IT WENT BACKWARDS!? 
- Put it back into floats for the moment I guess. I'm still not convinced

### Possible Improvements 
- Use float...
- Consider different file scanning options

-------




