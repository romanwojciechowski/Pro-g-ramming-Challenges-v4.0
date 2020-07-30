#include <iostream>
using namespace std;

string text = "";
int choice;
string converted_text = "";

void convertText(string letter)
{
    if (choice == 1)
    {
        for (int i = 0; i < letter.length(); i++)
        {
            char x = letter.at(i) + 3;
            cout << x;
        }
    }
    else
    {
        for (int i = 0; i < letter.length(); i++)
        {
            char x = letter.at(i) - 3;
            cout << x;
        }
    }
}

int main()
{
    cout << "1 - encrypt, 0 - decrypt: ";
    cin >> choice;
    cout << "Text: ";
    cin >> text;

    convertText(text);

    return 0;
}
