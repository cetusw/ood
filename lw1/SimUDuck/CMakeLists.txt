cmake_minimum_required(VERSION 4.0)
project(SimUDuck)

set(CMAKE_CXX_STANDARD 20)

find_package(GTest REQUIRED)

add_executable(SimUDuck SimUDuck.cpp
        lib/Duck/Dance/IDanceBehavior.h
        lib/Duck/Dance/DanceWaltz.h
        lib/Duck/Dance/DanceMinuet.h
        lib/Duck/Dance/DanceNoWay.h
)

add_executable(duck-test
        lib/Duck/Fly/TriggeredQuackFlyBehavior.h
        test/TriggeredQuackFlyBehaviorTest.cpp
        test/Mock/MockFlyBehavior.h
        test/Mock/MockDanceBehavior.h
        test/DanceBehaviorTest.cpp
)

target_link_libraries(duck-test PRIVATE GTest::gmock_main)

target_include_directories(duck-test PRIVATE ${GTest_INCLUDE_DIRS})

target_include_directories(duck-test PRIVATE ${CMAKE_CURRENT_SOURCE_DIR})

enable_testing()
add_test(NAME unit_tests COMMAND duck-test)
