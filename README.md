# hogwarts-legacy-pause-game

As a parent pausing a game is a must have for me.

Since Hogwarts Legacy does not allow to pause the game, this is a work around to enable players to pause the game. This is a very minimal hacked approach, windows system tools are used to pause the game.
This tool can be used in two different ways, start this program using the command line, while in game alt+tab to the command line and press `p`.
Another way to use it, is to access the url the program gives you when you start the program, it can be accessed from any device on your local network, this enables you for example to pause the game using your mobile phone.

If this helped you in any way and you would like to buy me a cup of ☕, [you can donate here](https://www.paypal.com/donate/?hosted_button_id=M29WRLEENT5MA).

**Notes:**

* This works during cutscenes
* Use at your own risk (do not start multiple programs)
* Known bug: if you pause the game and close this tool, you might not be able to resume the game

**Hopefully this will become useless when the developers enable the feature to pause the game (lol)**

## Requirements
* [SysInternals Suite](https://apps.microsoft.com/store/detail/sysinternals-suite/9P7KNL5RWT25?hl=pt-pt&gl=pt&rtc=1)
* [Golang](https://go.dev/dl/) (only if you want to clone and run the project)

## Instructions

1. Install [SysInternals Suite](https://apps.microsoft.com/store/detail/sysinternals-suite/9P7KNL5RWT25?hl=pt-pt&gl=pt&rtc=1) from Microsoft Store
2. Open the command line, go to the folder where you downloaded the program or cloned the project
3. Start Hogwarts Legacy 
4. Run the program by executing the program `.\hl_pause.exe` or by doing `go run main.go` (in case you cloned the project)
5. Open the given url (only when you want to pause/resume the game) or press `p` in the command line


## References

https://learn.microsoft.com/en-us/sysinternals/downloads/pstools