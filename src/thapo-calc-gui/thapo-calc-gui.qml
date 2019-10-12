import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQml 2.3
import thapo.calc.Controller 1.0

ApplicationWindow {
    id: main_window
    visible: true
    width: 640
    height: 480
    title: qsTr("Thapo Calc")

    Controller {
        id: controller
    }


    Grid {
        id: grid
        x: 0
        y: 80
        width: 200
        height: 100
    }

    TextField {
        id: output
        x: 0
        y: 0
        width: 400
        height: 100
        readOnly: true
        text: controller.output
    }

    TextField {
        id: input
        x: 0
        y: 100
        width: 400
        height: 100
        text: qsTr("")
        onTextChanged: controller.input = text
    }

    Button {
        id: button
        text: qsTr("Calculate")
        x: 0
        y: 200
        width: 400
        height: 100
    }

}
