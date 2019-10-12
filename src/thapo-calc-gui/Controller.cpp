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
    std::cout << "input: " << m_input.toStdString() << std::endl;
    try {
        long double result = calculation(m_input.toStdString());
        setOutput(QString::number(result, 'D'));
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
    m_output = output;
    std::cout << m_input.toStdString() << std::endl;
    emit outputChanged();
}
