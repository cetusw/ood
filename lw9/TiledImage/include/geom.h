#pragma once
#include <cstdint>

struct Point
{
	int x = 0;
	int y = 0;
};

struct Size
{
	int width = 0;
	int height = 0;
};

using Color = uint32_t;

inline bool IsPointInSize(const Point p, const Size size) noexcept
{
	return p.x >= 0 && p.y >= 0 && p.x < size.width && p.y < size.height;
}

constexpr Color ToColor(const uint8_t r, const uint8_t g, const uint8_t b) noexcept
{
	return (static_cast<uint32_t>(r) << 16) | (static_cast<uint32_t>(g) << 8) | static_cast<uint32_t>(b);
}

constexpr uint8_t GetR(const Color c) noexcept { return (c >> 16) & 0xFF; }
constexpr uint8_t GetG(const Color c) noexcept { return (c >> 8) & 0xFF; }
constexpr uint8_t GetB(const Color c) noexcept { return c & 0xFF; }