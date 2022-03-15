# bookshop-micro-rpg

Command line application written in Go which reproduces the micro RPG created by Oliver Darkshire of Sotheran's bookshop called BOOKSTORE, [posted on Twitter](https://twitter.com/Sotherans/status/1493279170188693506?s=20&t=2RmmSgLk4ycn0w6V3tpcTQ).

## Installation
``` bash
go get github.com/s0rbus/bookshop-micro-rpg
```

## Usage

### Command line interface
``` bash
go install github.com/s0rbus/bookshop-micro-rpg
```

or download binaries from the releases page.

To see usage message with the various options explained run:
```bash
./bookshop-micro-rpg --help
```
Without specifying any options, the defaults will be used which will run three simulations (attempts) using the basic '10 day' rule and not using any plugins.

## Plots
There is an option to plot the changes in Money during an attempt. This uses the excellent Go package [asciigraph](https://github.com/guptarohit/asciigraph) and is enabled using
```bash
--plot
```

Example output:
````
```bash
Attempt 1 The shop is no longer viable. You survived for 10 days, but now the business is closing for good. Your maximum amount of money was 3
Attempt 1 money: [1 -2 -5 -5 -2 -6 -5 -2 1 3]
  3.00 ┤        ╭
  2.00 ┤        │
  1.00 ┼╮      ╭╯
  0.00 ┤│      │
 -1.00 ┤│      │
 -2.00 ┤╰╮ ╭╮ ╭╯
 -3.00 ┼ │ ││ │
 -4.00 ┤ │ ││ │
 -5.00 ┤ ╰─╯│╭╯
 -6.00 ┤    ╰╯
Attempt 2 The shop is no longer viable. You survived for 10 days, but now the business is closing for good. Your maximum amount of money was 0
Attempt 2 money: [0 -1 -3 -2 -2 -1 0 -2 -3 -3]
  0.00 ┼╮    ╭╮
 -1.00 ┤╰╮  ╭╯│
 -2.00 ┤ │╭─╯ ╰╮
 -3.00 ┼ ╰╯    ╰─
Attempt 3 The shop is no longer viable. You survived for 20 days, but now the business is closing for good. Your maximum amount of money was 9
Attempt 3 money: [3 5 6 7 8 9 9 9 9 0 -1 -1 -1 -1 -1 -1 -1 0 0 0]
  9.00 ┤    ╭───╮
  8.00 ┼   ╭╯   │
  7.00 ┤  ╭╯    │
  6.00 ┤ ╭╯     │
  5.00 ┤╭╯      │
  4.00 ┤│       │
  3.00 ┼╯       │
  2.00 ┤        │
  1.00 ┤        │
  0.00 ┤        ╰╮      ╭──
 -1.00 ┤         ╰──────╯
 ```
````

## JSON output
The output from the applciation can be formatted as JSON using the --json option. This is ignored if either --verbose or --plot are used. The idea is that the JSON output provides structured data which can be further extracted and analyzed. This would be more difficult with the standard, non-structured output.

## Expansions

As of version 2.0.0 and above, the expansion system uses Javascript files which are run using an embedded Javascript engine. This makes it easier for people not familiar with Go to add expansions. The same options are used to find and load expansions except that --expansion should now take a javascript filename, so for example to use the sale expansion:
```bash
--plugin-dir=expansions --expansion=sale.js
```

This tells the application to look for and load the plugin file sale.js which can be found in the folder expansions.

This example is a 'sales' expansion. When a sale is started as a result of a dice throw, money decreases by 5 and patience by 1, but thereafter, while the sale is still 'on', sales of books are more likely to occur (50/50)

There is another simple example included in the repository for a 'coffee shop' within the bookshop.

## Building
make is used for building so the make tool is a requirement. The default target will just build the application. Use ```make all``` to also build expansion libraries in the expansions folder. Using make to build the application will also ensure that build info (version, git hash and built timestamp) are added to the application. These can be displayed by using ```--version``` flag.

## Licence
The default position when a licence is not provided, according to [github documentation](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/customizing-your-repository/licensing-a-repository#choosing-the-right-license) is:
```
You're under no obligation to choose a license. However, without a license, the default copyright laws apply, meaning that you retain all rights to your source code and no one may reproduce, distribute, or create derivative works from your work. 
```
Providing an Open Source licence would allow commercial use of the software, but since the code in this reposistory is an attempt to reproduce an idea by Oliver Darkshire, it does not feel right to include this permission. Therefore a licence is not provided and users of the code in this repo must abide by the default position as stated above.
## Acknowledgement

This project would not exist without the clever insight of Oliver Darkshire to create a micro RPG to give insight into running a bookshop :smile:.


## Contributing

This is a hobby/spare time project so there may be delays in responding to pull requests. Responses are 'best endeavour'. Having said that, contributions in the form of bug fixes, enhancements, suggestions for improvements are welcome, submit a pull request.


