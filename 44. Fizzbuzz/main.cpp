#include <iostream>

int main()
{
    for (int i = 1; i <= 50; i++)
    {
        if (i%3 == 0)
            std::cout << "fizz\n";
        else if (i%5 == 0)
            std::cout << "buzz\n";
        else
            std::cout << i << std::endl;
    }
}
