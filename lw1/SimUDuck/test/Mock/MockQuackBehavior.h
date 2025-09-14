#ifndef SIMUDUCK_MOCKQUACKBEHAVIOR_H
#define SIMUDUCK_MOCKQUACKBEHAVIOR_H

#include "gmock/gmock.h"
#include "lib/Duck/Quack/IQuackBehavior.h"

class MockQuackBehavior : public IQuackBehavior {
public:
    MOCK_METHOD(void, Quack, (), (override));
};


#endif //SIMUDUCK_MOCKQUACKBEHAVIOR_H
