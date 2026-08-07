#include <iostream>
#include <fstream>
#include <string>
using namespace std;

int main() {
    ifstream fin("input.txt",ios::in);

    if (!fin) {
        cout << "Error: Cannot open input.txt" << endl;
        return 1;
    }

    string inputData;
    fin >> inputData;
    fin.close();

    int n = inputData.length();

    // Size of the matrix
    int size = 2 * n + 1;
    int offset = n ;

    bool **visited = new bool*[size];

    for (int i = 0; i < size; i++) {
        visited[i] = new bool[size];
        for (int j = 0; j < size; j++) {
            visited[i][j] = false;
        }
    }

    // Start from the center
    int x = offset;
    int y = offset;

    visited[x][y] = true;
    int count = 1;

    // Process all inputData
    for (char ch : inputData) {
        switch (ch) {
            case '^':
                y++;
                break;
            case 'v':
                y--;
                break;
            case '>':
                x++;
                break;
            case '<':
                x--;
                break;
        }

        if (visited[x][y] == false) {
            visited[x][y] = true;
            count++;
        }
    }

    cout << "Houses receiving at least one present: " << count << endl;

    // Free allocated memory
    for (int i = 0; i < size; i++) {
        delete[] visited[i];
    }
    delete[] visited;

    return 0;
}
