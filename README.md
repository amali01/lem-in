




###### vvvveerrryyy immmportant  in check conflict: check if the tunnel is used at the same time 









Any unknown command will be ignored. "in set Colony"
The program must handle errors carefully. In no way can it quit in an unexpected manner.

<!-- needs to fix if the line contains a specific room and also extra string -->

##Some colonies will have many rooms and many links, but no path between ##start and ##end. if len(allPaths) == 0 then error
##Some will have rooms that link to themselves,// i think it is done
 sending your path-search spinning in circles. Some will have too many/too few ants, no ##start or ##end, duplicated rooms, links to unknown rooms, rooms with invalid coordinates and a variety of other invalid or poorly-formatted input. In those cases the program will return an error message ERROR: invalid data format. If you wish, you can elaborate a more specific error message (example: ERROR: invalid data format, invalid number of Ants or ERROR: invalid data format, no start room found).

Each tunnel can only be used once per turn.
<!-- room name must be unique print error message if there is more than one -->