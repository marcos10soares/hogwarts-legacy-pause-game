# hogwarts-legacy-pause-game

As a parent pausing a game is a must have for me.

Since Hogwarts Legacy does not allow to pause the game, this is a work around to enable players to pause the game. This is a very minimal hacked approach, windows system tools are used to pause the game.

Example videos:
* [Pause game remotely (from other device)](https://www.youtube.com/watch?v=yHBkuH7SgME)
* [Pause game locally by doing alt+tab and pressing 'p'](https://www.youtube.com/watch?v=2GVgz_MCS-M)

If this helped you in any way and you would like to buy me a cup of ☕, [you can donate here](https://www.paypal.com/donate/?hosted_button_id=M29WRLEENT5MA).

**Notes:**

* This works for script talking, but does not work for cutscenes (although the game gets paused, the game will skip forward the same amount of time you've paused when you resume)
* Use at your own risk (do not start multiple programs)
* Known bug: if you pause the game and close this tool, you might not be able to resume the game

**Hopefully this will become useless when the developers enable the feature to pause the game (lol)**

## Features

* Pause the game by doing `alt+tab` and then pressing `p` in the command line where you are running this tool.
* Pause the game by opening a url from any device in your local network (for example: smartphone).


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