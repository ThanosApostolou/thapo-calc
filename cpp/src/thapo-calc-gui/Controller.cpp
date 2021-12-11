#include "Controller.h"
#include <iostream>
#include "thapo-calc-lib.h"

Controller::Controller(QObject *parent) : QObject(parent) {
}

QString Controller::getInput() {
    return m_input;
}

void Controller::setInput(const QString &input) {
    if (input == m_input) {
        return;
    }
    m_input = input;
    try {
        long double result = calculation(m_input.toStdString().c_str());
        setOutput(QString::number(result, 'd'));
    }
    catch (Myerror myerror) {
        std::cout << myerror.msg << "\n\n";
        setOutput(QString::fromStdString(myerror.msg));
    }
    emit inputChanged();
}

QString Controller::getOutput() {
    return m_output;
}

void Controller::setOutput(const QString &output) {
    if (output == m_output) {
        return;
    }
    m_output = "= "+ output;
    emit outputChanged();
}
