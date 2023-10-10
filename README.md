<h1 align="center">Lem-In Project</h1>


<p align="center" >
    <img style="width:350px ; height:350px ;" src="lm-in3.jpg" alt="Lem-In Logo" />
</p>

<h2 align="center">About The Project</h2>
<h4 align="center">Lem-in is a project that simulates a digital version of an ant farm. The objective of this program is to find the quickest path for a given number of ants to navigate through a colony composed of rooms and tunnels. The program reads input from a file, which describes the layout of the ant farm and specifies the starting and ending points.</h4>

## Getting Started
You can run the Lem-In project with the following command:
```console
git clone https://learn.reboot01.com/git/emahfood/lem-in.git
cd lem-in
```

## Usage
```
go run . [filename]
```
### Directory Structure
```console
─ lem-in/
│
├── examples/
│   ├── example00.txt
│   └── ...
│
├── funcs/
│   ├── myInfo.go
│   ├── myStruct.go
│   ├── mySort.go
│   ├── CleanInput.go  
│   ├── myPrint.go
│   ├── myFlags.go
│   └── myColors.go
│
├── main.go
├── go.mod
├── README.md
└── ...
```

## Example

```
$ go run . test1.txt
3
2 5 0
##start
0 1 2
##end
1 9 2
3 5 4
0-2
0-3
2-1
3-1
2-3
```
- the result
```
L1-2 L2-3
L1-1 L2-1 L3-2
L3-1
```


## Instructions
- An ant farm is constructed with interconnected rooms and tunnels.
-   The ants are initially placed in the room labeled "##start".
-  The goal is to guide the ants to the room labeled "##end" using the shortest possible path.
-  It's important to note that the shortest path may not always be the simplest.
-  The program handles various scenarios, such as colonies with complex room structures, loops, invalid data formats, missing start/end rooms, duplicated rooms, and unknown room connections.
-  In such cases, the program provides appropriate error messages to indicate the issue.


## Additional information
- The project is written in Go.
- Only standard go packages were used.

## Authors
- emahfood
- amali
- akhaled

