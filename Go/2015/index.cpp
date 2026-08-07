#include <iostream>
#include <fstream>
#include <string>
using namespace std;

int main() {
    ifstream file("input.txt");

    if (!file) {
        cout << "Error: Could not open input.txt" << endl;
        return 1;
    }

    string s;
    file >> s;   // Reads the entire sequence of parentheses
    int floor = 0;
    for (int i = 0; i < s.size(); i++) {
        if (s[i] == '(')
            floor++;
        else if (s[i] == ')')
            floor--;

        if (floor == -1) {
            cout << "Basement reached at position: " << i + 1<<endl;
            break;
        }
    }
    file.close();

    return 0;
}
