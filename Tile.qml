import QtQuick 2.0

Rectangle {
    id: tile
    property int type: 0
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
    border.color: "black"
    border.width: 5

    MouseArea {
        id: mouseArea
        anchors.fill: parent
        onClicked: {
            var oldType = type

            if (grid.hasStart && grid.hasEnd) {
                if (type == 0) {
                    type = 1 //wall
                } else {
                    type  = 0 //start
                }
            } else {
                if (!grid.hasStart) {
                    grid.hasStart = true
                    type = 2
                }
                if (!grid.hasEnd) {
                    grid.hasEnd = true
                    type = 3
                }
            }

            if (oldType == 2) {
                grid.hasStart = false
            } else if (oldType == 3) {
                grid.hasEnd = false
            }
        }
    }
}
