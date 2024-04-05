# 1️⃣🐝🏎️ The One Billion Row Challenge


## Author - Dan Richards 


## Hardware 
- MacBook Pro 16" M1 2019 


## Runtime table 

| Iteration | Title                                                                                                                                         | Duration   | Improvement |
|-----------|-----------------------------------------------------------------------------------------------------------------------------------------------|------------| ------------|
| 0         | [Reference Code](https://r2p.dev/b/2024-03-18-1brc-go/#:~:text=One%20Billion%20Row%20Challenge%20in%20Golang%20%2D%20From%2095s%20to%201.96s) | 217.976069 |             |
| 1         | Simple Single Threaded                                                                                                                        | 308.678428 |             |
| 2         | Simple Single Threaded (Integers)                                                                                                             | 354.827043 |             | 
| 3         | Accessing the Map once per line                                                                                                               | 200.885222 |             | 
| 4         | One stack frame per line                                                                                                                      | 204.891398 |             | 

------


## 1 - Simple Single Threaded

### **Processing Time**: 308.678428 Seconds

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

### **Processing Time**: 354.827043 Seconds

### Reaction
- **IT WENT BACKWARDS!?**
- Put it back into floats for the moment I guess. I'm still not convinced

### Possible Improvements 
- Use float... (Done)
- Consider different file scanning options


## Comparison to an online article 
I am using [this article](https://r2p.dev/b/2024-03-18-1brc-go/#:~:text=One%20Billion%20Row%20Challenge%20in%20Golang%20%2D%20From%2095s%20to%201.96sthis) as a guide for helping me develop my code and understand the different aspects of the go language. 
I noticed that this person's code was run in 96 seconds (when run on my computer 217), now obviously I cannot account for differences in hardware... So I will run his code on my MacBook Pro M1 2019 to see how my code compares to his.  

I wanted to consider why my code was slower. On first glance the code seemed relatively similar... 
I did note that my code did call getMin and getMax functions on every iteration which could add overhead to the program by needing to push memory items to the stack... Although in saying that I would assume that the compiler would make my code more efficient and simply remove the need for adding items to the stack with a new function all. I notice that simply using pointers when referencing the values in getMin and getMax wouldn't decrease the overhead. This is because as the function call is executed, the program still has to add the pointer to the stack and being 64bits is the same as passing it by reference. (I assume). 


I noticed whilst removing the getMin and getMax functions. I am not certain how the compiler is actually decoding this, but the idea of needing to calculate the hash for the key, and then find the item in memory (considering clashing algorithms for 10k weather stations) this could be quite slow. I will firstly get the value of the struct, store it in the current stack and reference the values from there. 


-------

## 3 - Only accessing the Map once, not multiple times per weather station 

> As per the text above when comparing the reference code to my code.. I saw the first point of call was that the reference code was not calling a separate function for getting the min/Max temperature. However whilst pulling out the code to do this I saw that I accessed the map multiple times per weather station which inherently seems inefficient to me. Lets see what happens if I get the weather station data once per line of the text. 


### **Processing Time**: 200.885222 Seconds

```go
func myCode() {
    if len(os.Args) < 2 {
        panic("No arguments")
    }
    fileName := os.Args[1]

    // Open File
    file, fileError := os.Open(fileName)
    defer file.Close() // This will close the file after the function has been run.
    if fileError != nil {
        panic(fileError)
    }

    // Start processing the records

    stations := make(map[string]*StationData)

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {

    line := scanner.Text()
    parts := strings.Split(line, ";")

    stationName := parts[0]
    // To do, change float to int
    temp, parseError := strconv.ParseFloat(parts[1], 64)

    if parseError != nil {
        fmt.Printf("Error parsing float on line, " + line)
        panic(parseError)
    }

    station := stations[stationName]
    if station == nil {
        stations[stationName] = &StationData{temp, temp, temp, 1}
    } else {
        station.count += 1
        station.total += temp
        station.min = getMin(station.min, temp)
        station.max = getMax(station.min, temp)
    }

    }

    for key, value := range stations {
        average := value.total / float64(value.count)
        fmt.Printf("%s=%f/%f/%f\n", key, value.min, value.max, average)
    }

}
```

## 4 - Removing function calls and keeping it in once stack frame per line.

### Processing Time: 204.891398

> This is such a small impact that quite frankly I think the environment of me doing other tasks likely impacted this. 
This increase in speed is negligible I am willing to ignore it at this stage. But to be considered in the future. 


### Possible Improvements
I wanted to consider how changing the file scanning would work. [Time to start reading about bufio.scaner](https://medium.com/golangspec/in-depth-introduction-to-bufio-scanner-in-golang-55483bb689b4)


# Scanning the file 

```go
	for scanner.Scan() {
    scanner.Text() // or 
    scanner.Bytes()
}
```

1. Bytes 28s
2. Text 31.651927s
Not much difference between the two!


> Note I just spent the last few hours understanding how channels, mutexs and how the go runtime works!
 


 
 