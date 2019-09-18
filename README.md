# OSRS Hiscore Fetcher

Application written in Go to fetch a player's skills, using the oficial Runescape api.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

What you need installed to run the application:

- Go

### Running

To run the application, simply run the following command:

```
go run src/*.go <player_name>
```

An example would be:


```
go run src/*.go SiMoonLight1
```

## Results

On the current stage, the application just prints on the console the skill level, xp and rank. Example:

```
Player SiMoonLight1 
Overall: level 756291 with 1201xp ranked 6790969 
Attack: level 802980 with 70xp ranked 756944 
Defence: level 801349 with 67xp ranked 561198 
Strength: level 1027420 with 70xp ranked 738505 
Hitpoints: level 868199 with 74xp ranked 1106487 
Ranged: level 810585 with 75xp ranked 1244176 
Prayer: level 1207016 with 44xp ranked 56400 
Magic: level 1072690 with 63xp ranked 386704 
Cooking: level 1080415 with 54xp ranked 159648 
Woodcutting: level 1340549 with 52xp ranked 131932 
Fletching: level 731949 with 60xp ranked 274241 
Fishing: level 1328875 with 50xp ranked 108619 
Firemaking: level 1175991 with 50xp ranked 101380 
Crafting: level 915398 with 54xp ranked 151093 
Smithing: level 824597 with 50xp ranked 108729 
Mining: level 1058556 with 50xp ranked 108417 
Herblore: level 602334 with 50xp ranked 102209 
Agility: level 955671 with 50xp ranked 101337 
Thieving: level 801212 with 50xp ranked 101352 
Slayer: level 656637 with 62xp ranked 355659 
Farming: level 986772 with 17xp ranked 3500 
Runecrafting: level 581255 with 38xp ranked 30970 
Hunter: level -1 with 1xp ranked -1 
Construction: level 603893 with 50xp ranked 101469 
```

