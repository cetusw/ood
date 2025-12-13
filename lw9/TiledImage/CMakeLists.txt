cmake_minimum_required(VERSION 4.0)
project(Proxy)

set(CMAKE_CXX_STANDARD 20)

add_library(tiled-image
        include/cow.h
        include/geom.h
        include/tile.h
        include/image.h
        include/drawer.h
        src/image.cpp
        src/drawer.cpp
        src/tile.cpp
        tests/TileTest.cpp
        tests/ImageTest.cpp
        tests/main.cpp
)

target_include_directories(tiled-image PUBLIC ${CMAKE_CURRENT_SOURCE_DIR}/include)

add_executable(TiledImage src/main.cpp)
target_link_libraries(TiledImage PRIVATE tiled-image)

enable_testing()
add_subdirectory(tests)