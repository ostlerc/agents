import QtQuick 2.0

Rectangle {
    id: tile
    property int type: 0
    property int index: 0
    property bool solution: false
    width: 30
    height: 30
    color: {
        if (type == 0) //open
        return "white"
        else if (type == 1) //wall
        return "black"
        else if (type == 2) //start
        return "green"
        else //end
        return "red"
    }
    border.color: {
        if (grid.Edited || !solution) {
            solution = false
            return "black"
        }
        return "blue"
    }
    border.width: 5
    MouseArea {
        id: mouseArea
        anchors.fill: parent
        onClicked: {
            var oldType = type
            grid.clearGrid()

            if (grid.start != null && grid.end != null) {
                if (type == 0) {
                    type = 1 //wall
                } else {
                    type  = 0 //open
                }
            } else {
                if (grid.start == null) {
                    if (type == 3) {
                        type = 0
                    } else {
                        grid.setStart(index)
                        type = 2
                    }
                } else if (grid.end == null) {
                    grid.setEnd(index)
                    type = 3
                }
            }

            if (oldType == 2) {
                grid.clearStart()
            } else if (oldType == 3) {
                grid.clearEnd()
            }
        }
    }
}
