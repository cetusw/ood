#ifndef SIMUDUCK_MOCKDANCEBEHAVIOR_H
#define SIMUDUCK_MOCKDANCEBEHAVIOR_H

#include "gmock/gmock.h"
#include "lib/Duck/Dance/IDanceBehavior.h"

class MockDanceBehavior : public IDanceBehavior {
public:
    MOCK_CONST_METHOD0(Dance, void());
};

#endif //SIMUDUCK_MOCKDANCEBEHAVIOR_H