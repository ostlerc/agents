import QtQuick 2.0
import QtQuick.Particles 2.0

Rectangle {
    id: tile
    property int type: 0
    width: 50
    height: 50
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
            if (type == 0)
            type = 1
            else if (type == 1)
            type = 0
        }
    }
}
