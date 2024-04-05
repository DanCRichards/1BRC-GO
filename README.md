# 1️⃣🐝🏎️ The One Billion Row Challenge


## Author - Dan Richards 


## Hardware 
- MacBook Pro 16" M1 2019 


## Runtime table 

| Iteration | Title                                                                                                                                         | Duration    | Improvement |
|-----------|-----------------------------------------------------------------------------------------------------------------------------------------------|-------------| ------------|
| 0         | [Reference Code](https://r2p.dev/b/2024-03-18-1brc-go/#:~:text=One%20Billion%20Row%20Challenge%20in%20Golang%20%2D%20From%2095s%20to%201.96s) | 308.678428  |             |
| 1         | Simple Single Threaded                                                                                                                        | 308.678428  |             |
| 2         | Simple Single Threaded (Integers)                                                                                                             | 354.827043  |             | 

------


## 1 - Simple Single Threaded

> My intention with this was to simply get some code working and see how she went! This very much goes along the lines of red, green refactor. 
> Although in this repository I won't be using test cases... I think getting something to work and then refactoring is VERY IMPORTANT 


### Reaction
- My first reaction was that it was going to be slow, I noticed that some other programs were processing it in 90 seconds vs my 308 seconds. So something must be quite wrong here. Likely to do with the use of the scanner.

### Possible improvements
In order of easier to implement.... 
- Use integers and / 10 at the end 
- Use buffers for injecting the input 
- Use concurrency to use more than one core to

-------

## 2 - Simple Single Threaded - Using integers

### Reaction
- **IT WENT BACKWARDS!?**
- Put it back into floats for the moment I guess. I'm still not convinced

### Possible Improvements 
- Use float... (Done)
- Consider different file scanning options



### Comparison to an online article 
I am using [this article](https://r2p.dev/b/2024-03-18-1brc-go/#:~:text=One%20Billion%20Row%20Challenge%20in%20Golang%20%2D%20From%2095s%20to%201.96sthis) as a guide for helping me develop my code and understand the different aspects of the go language. 
I noticed that this person's code was run in 96 seconds, now obviously I cannot account for differences in hardware... So I will run his code on my MacBook Pro M1 2019 to see how my code compares to his.  


### Learnings: I started to read [this article](https://medium.com/golangspec/in-depth-introduction-to-bufio-scanner-in-golang-55483bb689b4) about the bufio package in go. 



-------




