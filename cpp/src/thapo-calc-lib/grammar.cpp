#include <math.h>

#include "grammar.h"
#include "error.h"
#include <iostream>

long double expression (std::stringstream& input, bool parentheses) {
    long double returned_expression=0;
    char c=' ', op='+';
    bool primary_next=true;

    c = read_nonspace_char(input);
    if (!isdigit(c) && c != '(') {
        primary_next=false;
    }
    while (! input.eof() && c != '\n') {
        std::cout << "read c: " << c << std::endl;
        if (primary_next) {
            input.putback(c);
            if (op == '+') {
                returned_expression += term(input);
            }
            else if (op == '-') {
                returned_expression -= term(input);
            }
            primary_next=false;
        }
        else {
            if (c == '+') {
                op = '+';
            }
            else if (c == '-') {
                op = '-';
            }
            else if (parentheses && c == ')') {
                return returned_expression;
            }
            else {
                throw Myerror ("Error: expression expected +, - or )");
            }
            primary_next = true;
        }
        c = read_nonspace_char(input);
        if ((input.eof() || c == '\n') && primary_next) {
            throw Myerror ("Error: expression didn't end with number");
        }
    }
    return returned_expression;
}

long double term (std::stringstream& input) {
    long double returned_term=1;
    char c=' ', op='*';
    c = read_nonspace_char(input);
    while (!input.eof() && c != '\n') {
        if (isprimary(c)) {
            input.putback(c);
            if (op == '*') {
                returned_term = returned_term * power(input);
            }
            else if (op == '/') {
                returned_term = returned_term / power(input);
            }
        }
        else {
            throw Myerror ("Term expected power");
        }
        c = read_nonspace_char(input);
        if (c == '*') {
            op = '*';
        }
        else if (c == '/') {
            op = '/';
        }
        else {
            input.putback(c);
            break;
        }
        c = read_nonspace_char(input);
    }
    return returned_term;
}

long double power (std::stringstream& input) {
    long double returned_power=1, exponent=1;
    char c=' ';
    c = read_nonspace_char(input);
    if (isprimary(c)) {
        input.putback(c);
        returned_power = primary(input);
    }
    else {
        throw Myerror ("Power expected primary");
    }
    c = read_nonspace_char(input);
    if (c == '^') {
        returned_power = pow(returned_power, power(input));
    }
    else {
        input.putback(c);
    }
    return returned_power;
}

long double primary (std::stringstream& input) {
    long double returned_primary=0;
    char c=' ';
    c = read_nonspace_char(input);
    if (isdigit(c)) {
        input.putback(c);
        input >> returned_primary;
    }
    else if (c == '(') {
        return expression (input, true);
    }
    else {
        throw Myerror ("Primary expected number or (");
    }
    return returned_primary;
}


char read_nonspace_char (std::stringstream& input) {
    char c='\n';
    input.get(c);
    while (c == ' ') {
        std::cout << "nonspace reading: " << c << std::endl;
        input.get(c);
    }
    return c;
}

bool isprimary(char c) {
    if (isdigit(c) || c == '(') {
        return true;
    }
    else return false;
}

