include(FetchContent)

FetchContent_Declare(
        googletest
        URL https://github.com/google/googletest/archive/refs/tags/v1.14.0.zip
)
FetchContent_MakeAvailable(googletest)


enable_testing()

add_executable(run_tests
        main.cpp
        ImageTest.cpp
        TileTest.cpp
)

target_link_libraries(run_tests PRIVATE tiled-image gtest)

include(GoogleTest)
gtest_discover_tests(run_tests)