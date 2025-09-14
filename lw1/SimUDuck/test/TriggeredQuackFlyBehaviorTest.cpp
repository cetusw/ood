#include "gtest/gtest.h"
#include "gmock/gmock.h"

#include "lib/Duck/Fly/TriggeredQuackFlyBehavior.h"

#include "Mock/MockFlyBehavior.h"
#include "Mock/MockQuackBehavior.h"

using testing::Exactly;
using testing::Return;

class TriggeredQuackFlyBehaviorTest : public testing::Test
{
protected:
    std::unique_ptr<MockFlyBehavior> mockFly;
    std::unique_ptr<MockQuackBehavior> mockQuack;

    MockFlyBehavior *mockFlyPtr = nullptr;
    MockQuackBehavior *mockQuackPtr = nullptr;

    void SetUp() override
    {
        mockFly = std::make_unique<MockFlyBehavior>();
        mockQuack = std::make_unique<MockQuackBehavior>();
        mockFlyPtr = mockFly.get();
        mockQuackPtr = mockQuack.get();
    }
};

TEST_F(TriggeredQuackFlyBehaviorTest, DoesNotQuackOnFirstSuccessfulFlight)
{
    EXPECT_CALL(*mockFlyPtr, Fly()).WillOnce(Return(true));
    EXPECT_CALL(*mockQuackPtr, Quack()).Times(Exactly(0));

    TriggeredQuackFlyBehavior triggeredFly(
        std::move(mockFly),
        std::move(mockQuack)
    );

    ASSERT_TRUE(triggeredFly.Fly());
}

TEST_F(TriggeredQuackFlyBehaviorTest, QuacksOnlyAfterSecondAndFourthFly)
{
    EXPECT_CALL(*mockFlyPtr, Fly()).Times(Exactly(4)).WillRepeatedly(Return(true));
    EXPECT_CALL(*mockQuackPtr, Quack()).Times(Exactly(2));
    TriggeredQuackFlyBehavior triggeredFly(
        std::move(mockFly),
        std::move(mockQuack)
    );

    triggeredFly.Fly();
    triggeredFly.Fly();
    triggeredFly.Fly();
    triggeredFly.Fly();
}

TEST_F(TriggeredQuackFlyBehaviorTest, DoesNotQuackAndCounterNotIncreasedIfFlyFails)
{
    EXPECT_CALL(*mockFlyPtr, Fly()).WillOnce(Return(false));
    EXPECT_CALL(*mockQuackPtr, Quack()).Times(Exactly(0));
    TriggeredQuackFlyBehavior triggeredFly(
        std::move(mockFly),
        std::move(mockQuack)
    );

    ASSERT_FALSE(triggeredFly.Fly());
}

TEST_F(TriggeredQuackFlyBehaviorTest, DoesNotQuackIfFlyFailsOnAnEvenFlightNumber)
{
    EXPECT_CALL(*mockFlyPtr, Fly())
            .WillOnce(Return(true))
            .WillOnce(Return(false));
    EXPECT_CALL(*mockQuackPtr, Quack()).Times(Exactly(0));
    TriggeredQuackFlyBehavior triggeredFly(
        std::move(mockFly),
        std::move(mockQuack)
    );

    triggeredFly.Fly();
    triggeredFly.Fly();
}

TEST_F(TriggeredQuackFlyBehaviorTest, CorrectlyCountsFlightsAndQuacksAfterAFailedAttempt)
{
    EXPECT_CALL(*mockFlyPtr, Fly())
        .WillOnce(Return(true))
        .WillOnce(Return(false))
        .WillOnce(Return(true));
    EXPECT_CALL(*mockQuackPtr, Quack()).Times(Exactly(1));
    TriggeredQuackFlyBehavior triggeredFly(std::move(mockFly), std::move(mockQuack));

    triggeredFly.Fly();
    triggeredFly.Fly();
    triggeredFly.Fly();
}