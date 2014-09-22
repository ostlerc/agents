agents
=====
This project is a qml graphical UI for showing agents in a grid

By clicking on the tiles you can change their color. Each color represents a different grid element.
Right click allows you to inspect the grid element and modify it's properties.

Requirements
------------
* golang 1.3.1

    To install golang visit: https://golang.org/doc/install

* Qt >= 5.2.1

    To install Qt visit: http://qt-project.org/downloads

* go-qml

    To install go-qml visit: http://github.com/go-qml/qml

Building
--------
Once all requirements have been met, you should be able to run 'go build' from the command line.
This will build a binary which you can then execute. Note that you must run the binary in the 
same directory as the qml files.
