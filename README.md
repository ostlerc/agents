astar
=====
This project is a qml graphical UI for showing the A\* (A STAR) path finding algorithm.

By clicking on the tiles you can change their color. Each color represents a different grid element.
* white: open path
* black: wall
* green: starting point
* red: goal or end point

After the grid has a starting and ending point the Run button will become enabled.
After clicking the 'Run' button solution nodes will be highlighted with a blue border.
You can continue to edit the grid at this point by clicking on more nodes, this
will clear the solution, or you can re-create a new grid by clicking the 'New' button.

Requirements
------------
* golang 1.3.1

    To install golang visit: https://golang.org/doc/install

* Qt >= 5.2.1

    To install Qt visit: http://qt-project.org/downloads

* go-qml

    To install go-qml visit: http://github.com/qml/go-qml

Building
--------
Once all requirements have been met, you should be able to run 'go build' from the command line.
This will build a binary which you can then execute. Note that you must run the binary in the 
same directory as the qml files.
