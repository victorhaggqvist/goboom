======
goboom
======

Synopsis
========
.. code-block:: sh

    goboom [option]

    goobom_run

Description
===========
goboom is wrapper around the dmenu.
goboom is the successor and rewrite of xboomx in Go.

goboom sorts commands to launch according to their launch frequency.
In other words - if you launch Chromium and KeePassX all the time - they will appear in the list of commands first.

Build
-----
goboom is built using Go 1.5 vendoring and godeps.

Install
-------
goboom is available via AUR as `goboom-bin`. For binary downloads see `releases`_.

.. _releases: https://github.com/victorhaggqvist/goboom/releases/latest

.. code-block:: sh

    sudo cp goboom /usr/bin
    sudo cp goboom_run /usr/bin
    mkdir -p ~/.goboom
    cp config.ini.default ~/.goboom/config.ini

Set your keybinding to `goboom_run`.

Options
=======

--gc         Run garbage collection of the DB
--launcher   Output launcher command
--post       Update ranking DB
--pre        Generate dmenu in
--stats      View DB stats

Migration from xboomx
=====================
goboom will look for a config file in `~/.goboom`.
You will need to will need to convert you old config file to ini-format and name it `config.ini`, see the bundled default ini-file for guidence.

goboom uses a csv-file instead sqlite as datastore.
You will need to export your xboomx db as so.
At this point there is not a provided tool to do so, but you can easely export it with something like `sqlitebrowser`_.

.. _sqlitebrowser: http://sqlitebrowser.org/

The contents of your exported database should look along the lines of this::

    name,count
    chroimum,13
    keepassx,17
    gimp,4

License
=======
::

    goboom - a dmenu wrapper
    Copyright (C) 2016 Victor Häggqvist

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <http://www.gnu.org/licenses/>.
