import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.VirtualKeyboard 2.3

ApplicationWindow {
    id: main_window
    visible: true
    width: 640
    height: 480
    title: qsTr("Thapo Calc")

    Grid {
        id: grid
        x: 0
        y: 80
        width: 100
        height: 100
    }

    Button {
        id: button
        text: qsTr("Calculate")
        anchors.rightMargin: 0
        anchors.bottomMargin: 0
        anchors.leftMargin: 0
        anchors.topMargin: 358
        width: 200
        height: 100
    }

    Rectangle {
        x: 100
        y: 186
        width: 200
        height: 100
        color: "red"

        Text {
            anchors.centerIn: parent
            text: "Hello, World!"
        }

        TapHandler {
            onTapped: parent.color = "blue"
        }
    }

    TextField {
        id: textField
        x: 274
        y: 212
        width: 200
        height: 100
        text: qsTr("Text Field")
    }
}
