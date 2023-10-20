# pokedexcli
A command line interface for fetching and getting details about Pokemon


## Installation
- get go!
- then clone
- $ go build . 
- $ pokedexcli

## usage:
- type "help" to get a list of available commands
- catch pokemon and check their stats.

## Things that would be kinda interesting to add:
[x] number of poke balls to limit spamming Catch commmand
[] A a shop for buying more items to make catching pokemon easier
    [x] !!!! fix bug where buying a poke-ball only gives it to you the first time and only gets the cost after the first attempt (possibly something to do with cache)
[] Possibly limiting the amount of pokemon in an area before you have to leave to find new pokemon
    [] tracking your current location
    [] tracking of number of pokemon in area
        [] adding a respawn time for escaped/caught pokemon to an area
[] pokemon giving coin when caught
[] a leveling system 
    [x] pokemon give 1/3 XP of their baseExperience to player when caught
    [] pokemon not caught should give some XP to player
    [] an exponential or teired XP neeeded to gain a level
        [] a bonus to the player for catching higher level pokemon
[] getting player Name?
[] messaging for throwing a pokeball
[] choose a team of pokemon from pokedex to go exploring with up to 6
    [] bonus for battles based on party pokemon type vs wild pokemon type
[] a map of literals
    [] cost of items at store
    [] XP multiplyer
[] a save state
[] explore command should limit areas to only a few locations at a time. possibly based on player level

