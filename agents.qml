import QtQuick 2.2
import QtQuick.Controls 1.0
import QtQuick.Layouts 1.1
import QtQuick.Dialogs 1.0

ApplicationWindow {
    width: 850
    height: 600
    color: "white"
    ColumnLayout {
        id: screen
        anchors.fill: parent
        spacing: 0
        Rectangle {
            id: statusRect
            Layout.fillWidth: true
            Layout.preferredHeight: statusRow.height + 10
            border.color: "black"
            border.width: 1
            RowLayout {
                id: statusRow
                anchors { verticalCenter: parent.verticalCenter; margins: 5; horizontalCenter: parent.horizontalCenter }
                Text {
                    id: statusText
                    objectName: "statusText"
                    text: "Click the grid cells to make a Nest, food, and walls. Right click to inspect."
                }
                Text {
                    text: "count:"
                    visible: sb1.visible
                }
                SpinBox {
                    id: sb1
                    objectName: "countSpinner"
                    maximumValue: 999
                    minimumValue: 1
                    Layout.minimumWidth: 50
                    value: 1
                    onEditingFinished: {
                        grid.setCount(value)
                    }
                    visible: false
                }
                Text {
                    visible: sb2.visible
                    text: "lifetime:"
                }
                SpinBox {
                    id: sb2
                    objectName: "lifeSpinner"
                    maximumValue: 999
                    minimumValue: 1
                    Layout.minimumWidth: 50
                    value: 25
                    onEditingFinished: {
                        grid.setLife(value)
                    }
                    visible: false
                }
            }
        }
        Rectangle {
            Layout.fillHeight: true
            Layout.fillWidth: true
            Flickable {
                clip: true
                boundsBehavior: Flickable.StopAtBounds
                anchors.centerIn: parent
                width: { return Math.min(parent.width, g.width) }
                height: { return Math.min(parent.height, g.height) }
                contentWidth: g.width; contentHeight: g.height
                flickableDirection: Flickable.HorizontalAndVerticalFlick
                Grid {
                    id: g
                    objectName: "grid"
                    spacing: 1
                }
            }
        }
        Rectangle {
            id: botBorder
            height: 1
            Layout.minimumHeight: 1
            Layout.fillWidth: true
            color: "black"
        }
        Rectangle {
            id: toolBar
            Layout.fillWidth: true
            Layout.preferredHeight: 70
            color: "white"
            TabView {
                id: btmtab
                anchors.fill: parent
                objectName: "tv"
                Tab {
                    title: "editor"
                    Rectangle {
                        color: "white"
                        RowLayout {
                            id: bottomLayout
                            anchors.fill: parent
                            Button {
                                id: newBtn
                                text: "New"
                                onClicked: grid.buildGrid()
                            }
                            Button {
                                text: "Save"
                                onClicked: {
                                    fileDialog.close()
                                    fileDialog.type = 0 //save
                                    fileDialog.open()
                                }
                            }
                            FileDialog {
                                id: fileDialog
                                property int type: 0
                                objectName: "fileDialog"
                                title: "Choose a filename"
                                onAccepted: {
                                    console.log("You chose: " + fileDialog.fileUrls)
                                    if(!type) {
                                        grid.saveGrid(fileDialog.fileUrl.toString())
                                    } else {
                                        grid.loadGrid(fileDialog.fileUrl.toString())
                                    }
                                }
                            }
                            Button {
                                text: "Load"
                                onClicked: {
                                    fileDialog.close()
                                    fileDialog.type = 1 //load
                                    fileDialog.open()
                                }
                            }
                            Text {
                                id: rowText
                                text: "Rows:"
                            }
                            SpinBox {
                                id: colRect
                                objectName: "cols"
                                Layout.minimumWidth: 42
                                maximumValue: 99
                                minimumValue: 1
                                value: 25
                            }
                            Text {
                                id: colText
                                text: "Columns:"
                            }
                            SpinBox {
                                id: rowRect
                                objectName: "rows"
                                Layout.minimumWidth: 42
                                maximumValue: 99
                                minimumValue: 1
                                value: 25
                            }
                            Rectangle {
                                Layout.fillWidth: true
                            }
                            Text {
                                id: foodcnttxt
                                text: "Food Count:"
                            }
                            SpinBox {
                                id: foodcntbox
                                objectName: "defaultFoodCountCombo"
                                Layout.minimumWidth: 50
                                maximumValue: 999
                                minimumValue: 1
                                value: 75
                            }
                            Text {
                                id: foodexptxt
                                text: "Food Life:"
                            }
                            SpinBox {
                                id: foodexpbox
                                objectName: "defaultFoodLifetimeCombo"
                                Layout.minimumWidth: 50
                                maximumValue: 999
                                minimumValue: 1
                                value: 125
                            }
                        }
                    }
                }

                Tab {
                    title: "simulator"
                    Rectangle {
                        color: "white"
                        RowLayout {
                            anchors.fill: parent

                            Button {
                                id: simRunBtn
                                text: "Run"
                                enabled: grid.nest != null
                                Component.onCompleted: grid.assign("Run",this)
                                onClicked: grid.runClicked()
                            }
                            Button {
                                text: "Step"
                                Component.onCompleted: grid.assign("Step",this)
                                onClicked: grid.step()
                                enabled: false
                            }
                            Button {
                                text: "Pause"
                                Component.onCompleted: grid.assign("Pause",this)
                                onClicked: grid.pauseClicked()
                                enabled: false
                            }
                            Text {
                                text: "Delay ms"
                            }
                            SpinBox {
                                Component.onCompleted: grid.assign("delaySpinner", this)
                                Layout.minimumWidth: 50
                                maximumValue: 999
                                minimumValue: 1
                                value: 100
                            }
                            Rectangle {
                                Layout.fillWidth: true
                            }
                            Text {
                                text: "Time:"
                            }
                            Text {
                                Component.onCompleted: grid.assign("simStatus", this)
                                text: "0"
                            }
                            Text {
                                text: "Food Gathered:"
                            }
                            Text {
                                Component.onCompleted: grid.assign("foodStatus", this)
                                text: "0"
                            }
                        }
                    }
                }
            }
        }
    }
}
