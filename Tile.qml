import QtQuick 2.0

Rectangle {
    id: tile
    property int type: 0
    property int index: 0
    property int food: 0
    property bool solution: false
    property bool selected: false
    width: 25
    height: 25
    color: {
        if (type == 0) // open
        return "white"
        else if (type == 1) // wall
        return "black"
        else if (type == 2) // home
        return "brown"
        else if (type == 3) // food
        return "green"
    }
    border.color: {
        if (!grid.Edited) {
            solution = false
            if(solution){
                return "blue"
            } else if(selected){
                return "blue"
            }
        }
        return "black"
    }
    Text {
        anchors.centerIn: parent
        font.pixelSize: 10
        color: "white"
        visible: tile.type == 3
        text: food
    }
    border.width: 5
    MouseArea {
        id: mouseArea
        anchors.fill: parent
        hoverEnabled: true
        acceptedButtons: Qt.LeftButton | Qt.RightButton
        onClicked: {
            if (mouse.button == Qt.RightButton) {
                grid.setSelected(index)
                return
            }
            var oldType = type
            if(!grid.edited) {
                grid.clearGrid()
            }

            if (grid.home != null) {
                if (type == 0) {
                    type = 1 //wall
                } else if (type == 1) {
                    type = 3 //food
                } else {
                    type  = 0 //open
                }
            } else {
                if (grid.home == null) {
                    if (type == 3) {
                        type = 0
                    } else {
                        grid.setHome(index)
                        type = 2
                    }
                }
            }

            if (oldType == 2) {
                grid.clearHome()
            } 
        }
    }
}
