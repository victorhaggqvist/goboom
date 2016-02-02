# goboom
goboom is wrapper around the dmenu. goboom is the successor and rewrite of xboomx in Go.

goboom sorts commands to launch according to their launch frequency.
In other words - if you launch Chromium and KeePassX all the time - they will appear in the list of commands first.

## Building
goboom is build using Go 1.5 vendoring and godeps.

    export GO15VENDOREXPERIMENT=1
    godep get github.com/victorhaggqvist/goboom
    go build goboom.go

## Install
```sh
sudo cp goboom /usr/bin
sudo cp goboom_run /usr/bin
mkdir -p ~/.goboom
cp config.ini.default ~/.goboom/config.ini
```

Set your keybinding to `goboom_run`.

### Migration from xboomx
goboom will look for a config file in `~/.goboom`.
You will need to will need to convert you old config file to ini-format and name it `config.ini`, see the bundeled default ini-file for guidence.

goboom uses a csv-file instead sqlite as datastore. You will need to export your xboomx db as so. At this point there is not a provided tool to do so, but you can easely export it with something like [sqlitebrowser](http://sqlitebrowser.org/).

The contents of your exported database should look along the lines of this.

    name,count
    chroimum,13
    keepassx,17
    gimp,4


## License

GPL v3
