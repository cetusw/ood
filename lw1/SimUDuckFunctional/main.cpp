#include <iostream>
#include "lib/Duck/MallardDuck.h"
#include "lib/Duck/DecoyDuck.h"
#include "lib/Duck/RedheadDuck.h"

void PlayWithDuck(const Duck& duck)
{
    duck.Display();
    duck.Quack();
    duck.Fly();
    duck.Dance();
    std::cout << "--------------------" << std::endl;
}

int main()
{
    MallardDuck mallard1;
    MallardDuck mallard2;
    DecoyDuck decoy;
    RedheadDuck redhead;

    PlayWithDuck(mallard1);
    PlayWithDuck(decoy);
    PlayWithDuck(redhead);

    mallard1.Fly();
    mallard2.Fly();
    mallard1.Fly();

    return 0;
}
